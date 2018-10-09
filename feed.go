package main

import (
	"github.com/gorilla/feeds"
	"time"
)

func generateFeed(domain string) (*feeds.Feed, error) {
	client := NewClient(domain)
	options := &ListOption{
		PerPage: 10,
		Page:    1,
	}
	posts, err := client.GetPosts(options)
	if err != nil {
		return nil, err
	}

	site, err := client.GetSiteInfo()
	if err != nil {
		return nil, err
	}

	feed := &feeds.Feed{
		Title: site.Name,
		Link: &feeds.Link{
			Href: site.URL,
		},
		Description: site.Description,
		Author: &feeds.Author{
			Name: site.Name,
		},
		Items: []*feeds.Item{},
	}

	for _, post := range posts {
		item := &feeds.Item{
			Title:       post.Title.Rendered,
			Link:        &feeds.Link{Href: post.GUID.Rendered},
			Description: post.Excerpt.Rendered,
			Content:     post.Content.Rendered,
		}
		layout := "2006-01-02T15:04:05"
		createdTime, err := time.Parse(layout, post.Date)
		if err != nil {
			return nil, err
		}
		item.Created = createdTime
		feed.Items = append(feed.Items, item)
	}

	return feed, nil
}
