package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"regexp"

	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func starToString(star *github.StarredRepository) string {
	repo := star.GetRepository()
	starredAt := star.GetStarredAt()

	var text string
	text += fmt.Sprintf("Name: %v", repo.GetName())
	text += "\n"
	text += fmt.Sprintf("Starred at: %v", starredAt)
	text += "\n"
	text += fmt.Sprintf("Description: %v", repo.GetDescription())
	text += "\n"
	text += fmt.Sprintf("Language: %v", repo.GetLanguage())
	text += "\n"
	text += fmt.Sprintf("Stars: %v", repo.GetStargazersCount())
	text += "\n"
	text += fmt.Sprintf("URL: %v", repo.GetURL())

	return text
}

func containsSubstr(subj, substr string) bool {
	subj = strings.ToLower(subj)
	substr = strings.ToLower(substr)
	return strings.Contains(subj, substr)
}

const entrySeparator = "//-----------------------"

var (
	username = flag.String("username", "petuhovskiy", "your github username")
	token    = flag.String("token", "", "github oauth token")
	filter   = flag.String("filter", "github", "filter by substring inclusion (ignoring case) in combined star description")
	regex    = flag.String("regex", "", "filter by regular expression substring")
)

func main() {
	flag.Parse()

	ctx := context.Background()

	var client *github.Client
	if *token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: *token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
	} else {
		client = github.NewClient(nil)
	}

	// Example of listing user's starred repos.
	// https://gist.github.com/derhuerst/19e0844796fa3b62e1e9567a1dc0b5a3

	var stars []*github.StarredRepository

	const (
		perPage = 100
	)
	for page := 1; ; page++ {
		fmt.Printf("Fetching page %v...\n", page)

		list, resp, err := client.Activity.ListStarred(ctx, *username, &github.ActivityListStarredOptions{
			ListOptions: github.ListOptions{
				Page:    page,
				PerPage: perPage,
			},
		})
		if err != nil {
			log.Fatal(err)
		}

		stars = append(stars, list...)
		fmt.Printf("Fetched %v stars so far\n", len(stars))

		spew.Dump(resp)
		if len(list) == 0 {
			break
		}
	}
	fmt.Println(entrySeparator)

	var repos []string
	for _, star := range stars {
		repos = append(repos, starToString(star))
	}

	found := 0
	for _, repo := range repos {
		if !containsSubstr(repo, *filter) {
			continue
		}
		if *regex != "" {
			if ok, _ := regexp.MatchString(*regex, repo); !ok {
				continue
			}
		}

		fmt.Println(repo)
		fmt.Println(entrySeparator)
		found++
	}

	fmt.Printf("Found %v entries\n", found)
}
