package main

import (
	"flag"
	"fmt"
	"github.com/bunsenmcdubbs/tweeter-deleter/cmd"
	"github.com/bunsenmcdubbs/tweeter-deleter/tweet"
	"github.com/dghubble/go-twitter/twitter"
	"time"
)

func main() {
	cfg, err := cmd.ParseEnvConfig()
	if err != nil {
		panic(fmt.Errorf("unable to parse appConfig: %w", err))
	}

	yearsFlag := flag.Int("year", 10, "Maximum allowable tweet age in years")
	forceFlag := flag.Bool("force", false, "Actually delete tweets. If unset, command only prints number of potentially affected tweets without deleting tweets.")
	verboseFlag := flag.Bool("verbose", false, "Print out all (potentially) deleted tweets")
	flag.Parse()

	client, err := tweet.AuthorizeOOBInteractive(cfg.ConsumerKey, cfg.ConsumerSecret)
	if err != nil {
		panic(err)
	}

	d := tweet.NewDeleter(twitter.NewClient(client))
	oldestDate := time.Now().AddDate(-*yearsFlag, 0, 0)
	fmt.Printf("Deleting tweets older than %d years ago (before %s). Dry run: %t\n", *yearsFlag, oldestDate, !*forceFlag)
	numDeleted, err := d.Delete(oldestDate, *forceFlag, *verboseFlag)
	fmt.Printf("Deleted %d tweets. Dry run: %t. Err: %v\n", numDeleted, !*forceFlag, err)
}
