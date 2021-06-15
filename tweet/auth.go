package tweet

import (
	"fmt"
	"net/http"

	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"
)

const (
	oobCallbackURL = "oob"
)

func AuthorizeOOB(consumerKey, consumerSecret string) (*http.Client, error) {
	config := oauth1.Config{
		ConsumerKey: consumerKey,
		ConsumerSecret: consumerSecret,
		CallbackURL: oobCallbackURL,
		Endpoint: twitter.AuthorizeEndpoint,
	}

	reqToken, reqSecret, err := config.RequestToken()
	if err != nil {
		return nil, err
	}
	authorizationURL, err := config.AuthorizationURL(reqToken)
	if err != nil {
		return nil, err
	}
	fmt.Println("authorizationURL: ", authorizationURL)

	var verifierCode string
	fmt.Scanln(&verifierCode)

	accessToken, accessSecret, err := config.AccessToken(reqToken, reqSecret, verifierCode)
	if err != nil {
		return nil, err
	}

	token := oauth1.NewToken(accessToken, accessSecret)
	return config.Client(oauth1.NoContext, token), nil
}