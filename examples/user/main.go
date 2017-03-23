package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lyft/lyft-go-sdk/lyft"

	"log"

	"github.com/davecgh/go-spew/spew"
	"golang.org/x/oauth2"
)

var (
	config = &oauth2.Config{
		ClientID:     "YOUR-CLIENT-ID",
		ClientSecret: "YOUR-CLIENT-SECRET",
		Scopes:       []string{"public", "profile", "rides.read", "rides.request", "offline"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.lyft.com/oauth/authorize",
			TokenURL: "https://api.lyft.com/oauth/token",
		},
	}
	oauthStateString = "random"
)

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/redirect", handleOAuth2Callback)

	fmt.Println("Started running on http://127.0.0.1:7001")
	fmt.Println(http.ListenAndServe(":7001", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`<html><body>Login using <a href="/login">Lyft</a></body></html>`))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := config.AuthCodeURL(oauthStateString, oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleOAuth2Callback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	callAPI(token)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func callAPI(token *oauth2.Token) {
	httpClient := config.Client(context.Background(), token)
	client := lyft.NewAPIClient(httpClient, "sample-app")

	// Make sure to prefix your Client Secret above with SANDBOX-, otherwise this will request a real ride.
	// See https://developer.lyft.com/docs/sandbox
	opt := map[string]interface{}{
		"lat": 37.7763,
		"lng": -122.3918,
	}
	result, _, err := client.UserApi.NewRide(lyft.Ride{
		RideType: lyft.RideTypeLyft,
		Origin:   opt,
	})

	if err != nil {
		log.Fatal(err)
	} else {
		spew.Dump(result)
	}
}
