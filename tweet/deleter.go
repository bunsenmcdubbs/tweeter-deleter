package tweet

import (
	"context"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/sync/errgroup"
	"sync/atomic"
	"time"
)

type Deleter struct {
	client *twitter.Client
}

func NewDeleter(client *twitter.Client) *Deleter {
	return &Deleter{client}
}

// Delete tweeters for the authenticated user which were posted before the oldest allowed time (exclusive).
// "force" must be true in order to delete tweets. Otherwise, a dry run is executed.
// Returns the number of deleted (or potentially deleted if force=false) tweets.
func (d Deleter) Delete(oldest time.Time, force bool, verbose bool) (int, error) {
	var oldestID int64
	var numDeleted int64
	for {
		tweets, _, err := d.client.Timelines.UserTimeline(&twitter.UserTimelineParams{
			TrimUser:        twitter.Bool(true),
			IncludeRetweets: twitter.Bool(true),
			MaxID:           oldestID,
			Count:           200,
		})
		if len(tweets) == 0 || err != nil {
			return int(numDeleted), nil
		}

		g, innerCtx := errgroup.WithContext(context.Background())
		g.SetLimit(20)
		for _, tweet := range tweets {
			t, err := tweet.CreatedAtTime()
			if err != nil {
				return int(numDeleted), fmt.Errorf("unable to parse timestamp for tweet %d: %w", tweet.ID, err)
			}
			if t.Before(oldest) {
				if force {
					func(tweetID int64) {
						g.Go(func() error {
							if innerCtx.Err() != nil { // skip/no-op if any error has already been encountered
								return nil
							}

							deletedTweet, _, err := d.client.Statuses.Destroy(tweetID, &twitter.StatusDestroyParams{TrimUser: twitter.Bool(false)})
							if err != nil || deletedTweet.ID != tweetID {
								return fmt.Errorf("unable to delete tweet #%d: %w", tweetID, err)
							}
							if verbose {
								fmt.Printf("Deleted tweet %d from %s: %v\n", deletedTweet.ID, deletedTweet.CreatedAt, deletedTweet.Text)
							}
							atomic.AddInt64(&numDeleted, 1)
							return nil
						})
					}(tweet.ID)
				} else {
					if verbose {
						fmt.Printf("Would have deleted tweet %d from %s: %v\n", tweet.ID, tweet.CreatedAt, tweet.Text)
					}
					numDeleted += 1
				}
			}

			if oldestID == 0 || tweet.ID < oldestID {
				oldestID = tweet.ID
			}
		}
		if err := g.Wait(); err != nil {
			return int(numDeleted), err
		}

		oldestID -= 1
	}
}
