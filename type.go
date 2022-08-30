package main

type Channel struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type ChannelRequest struct {
	Title string `json:"title"`
}

type Article struct {
	ID        int64  `json:"id"`
	ChannelID int64  `json:"channelID"`
	Url       string `json:"url"`
	WordCount int64  `json:"wordCount"`
}

type ArticleRequest struct {
	Url string `json:"url"`
}

type Filter struct {
	Lo int64 `json:"lo"`
	Hi int64 `json:"hi"`
}
