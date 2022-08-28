package main

type Channel struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type Article struct {
	ID        int64  `json:"id"`
	ChannelID int64  `json:"channelID"`
	Url       string `json:"url"`
	WordCount int    `json:"wordCount"`
}
