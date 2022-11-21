# Tweeter Deleter

Create a new Twitter developer account, get "Consumer Keys". Ensure that "App permissions" are set to "Read and Write" if you want deletes (`-force`) to succeed.

```shell
$ TWITTER_CONSUMER_KEY=<key> TWITTER_CONSUMER_SECRET=<secret> go run ./cmd/deleter -year 5 -force
Please authorize this app: <URL>
Enter PIN: XXXXXXX
Deleting tweets older than 5 years ago (before 2017-11-20 22:04:59.268364 -0800 PST). Dry run: false
Deleted 68 tweets. Dry run: false. Err: <nil>
```

View help with `go run ./cmd/deleter -h`
```
  -force
    	Actually delete tweets. If unset, command only prints number of potentially affected tweets without deleting tweets.
  -verbose
    	Print out all (potentially) deleted tweets
  -year int
    	Maximum allowable tweet age in years (default 10)
```