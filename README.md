Create a new Twitter developer account, get "Consumer Keys".

```shell
$ TWITTER_CONSUMER_KEY=<key> TWITTER_CONSUMER_SECRET=<secret> go run ./cmd/head
envdir env go run ./cmd/head
authorizationURL:  https://api.twitter.com/oauth/authorize?oauth_token=xxxxxx
<user input>
Raw Response Body:
```