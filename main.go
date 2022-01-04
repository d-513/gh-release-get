package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/google/go-github/v41/github"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage:", os.Args[0], "<owner> <repo> [content type]")
		os.Exit(1)
	}
	if len(os.Args) > 4 {
		log.Fatal("Too many arguments; check if your shell isn't doing the path matching itself")
		os.Exit(1)
	}
	client := github.NewClient(nil)
	opt := github.ListOptions{}
	releases, _, err := client.Repositories.ListReleases(context.Background(), os.Args[1], os.Args[2], &opt)
	if err != nil {
		log.Fatal("error:", err)
		os.Exit(1)
	}

	release := releases[0]
	var dl string

	if len(os.Args) > 3 {
		for _, asset := range release.Assets {
			match, errA := filepath.Match(os.Args[3], asset.GetName())
			if errA != nil {
				log.Fatal("error:", err)
				os.Exit(1)
			}
			if match {
				dl = asset.GetBrowserDownloadURL()
				break
			}
		}
	} else {
		dl = release.Assets[0].GetBrowserDownloadURL()
	}
	fmt.Println(dl)
}
