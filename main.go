package main

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	url, ok := request.QueryStringParameters["url"]
	if !ok {
		return nil, errors.New("param url is required")
	}
	feed, err := generateFeed(url)
	if err != nil {
		return nil, err
	}

	rss, err := feed.ToRss()
	if err != nil {
		return nil, err
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       rss,
	}, nil
}

func main() {
	lambda.Start(handler)
}
