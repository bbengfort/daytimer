package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"time"

	calendar "google.golang.org/api/calendar/v3"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// getClient uses the oauth config to retrieve a token and generate a client.
func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile, err := tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}

	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(cacheFile, tok)
	}

	return config.Client(ctx, tok)
}

// tokenCacheFile generates credential file path/filename.
func tokenCacheFile() (string, error) {
	// Get the user to look up the user's home directory
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	// Get the hidden credentials directory, making sure it's created
	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)

	// Determine the path to the token cache file
	tokenCachePath := filepath.Join(tokenCacheDir, url.QueryEscape("daytimer.json"))
	return tokenCachePath, nil
}

// tokenFromFile loads a Token from the given path
func tokenFromFile(path string) (*oauth2.Token, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	t := new(oauth2.Token)
	err = json.NewDecoder(f).Decode(t)
	return t, err
}

// saveToken uses a path to create a file and store the token in it
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()

	json.NewEncoder(f).Encode(token)
}

// getTokenFromWeb uses Config to request a Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}

	return tok
}

func main() {
	ctx := context.Background()

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := getClient(ctx, config)

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve calendar client. %v", err)
	}

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events. %v", err)
	}

	fmt.Println("Upcoming events:")
	if len(events.Items) > 0 {
		for _, i := range events.Items {
			var when string
			// If the DateTime is an empty string, the event is an all day event
			if i.Start.DateTime != "" {
				when = i.Start.DateTime
			} else {
				when = i.Start.Date
			}
			fmt.Printf("%s (%s)\n", i.Summary, when)
		}
	} else {
		fmt.Println("No upcoming events found.")
	}
}
