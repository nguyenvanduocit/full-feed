package main

import "github.com/gorilla/feeds"

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
		feed.Items = append(feed.Items, &feeds.Item{
			Title:       post.Title.Rendered,
			Link:        &feeds.Link{Href: post.GUID.Rendered},
			Description: post.Excerpt.Rendered,
		})
	}

	return feed, nil
}
