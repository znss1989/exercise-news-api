package main

import (
	"database/sql"
	"fmt"
	"strings"

	strip "github.com/grokify/html-strip-tags-go"
)

func queryChannels() ([]Channel, error) {
	var channels []Channel

	rows, err := db.Query("SELECT id, title FROM channels")
	if err != nil {
		return nil, fmt.Errorf("queryChannels: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var chn Channel
		if err := rows.Scan(&chn.ID, &chn.Title); err != nil {
			return nil, fmt.Errorf("queryChannels: %v", err)
		}
		channels = append(channels, chn)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("queryChannels: %v", err)
	}
	return channels, nil
}

func insertChannel(chn Channel) (int64, error) {
	result, err := db.Exec("INSERT INTO channels (title) VALUES (?)", chn.Title)
	if err != nil {
		return 0, fmt.Errorf("insertChannel: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("insertChannel: %v", err)
	}
	return id, nil
}

func queryArticles(channelID int64) ([]Article, error) {
	var articles []Article

	rows, err := db.Query("SELECT id, channel_id, url, wc FROM articles WHERE channel_id=?", channelID)
	if err != nil {
		return nil, fmt.Errorf("queryArticles: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var atc Article
		if err := rows.Scan(&atc.ID, &atc.ChannelID, &atc.Url, &atc.WordCount); err != nil {
			return nil, fmt.Errorf("queryArticles: %v", err)
		}
		articles = append(articles, atc)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("queryArticles: %v", err)
	}
	return articles, nil
}

func queryArticlesFiltering(channelID int64, ftr Filter) ([]Article, error) {
	var articles []Article

	rows, err := db.Query(
		`SELECT id, channel_id, url, wc 
		 FROM articles 
		 WHERE channel_id=? 
		 AND (wc BETWEEN ? AND ?)`,
		channelID, ftr.Lo, ftr.Hi)
	if err != nil {
		return nil, fmt.Errorf("queryArticles: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var atc Article
		if err := rows.Scan(&atc.ID, &atc.ChannelID, &atc.Url, &atc.WordCount); err != nil {
			return nil, fmt.Errorf("queryArticles: %v", err)
		}
		articles = append(articles, atc)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("queryArticles: %v", err)
	}
	return articles, nil
}

func insertArticle(channelID int64, atc Article) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO articles (channel_id, url, wc) VALUES (?, ?, ?)",
		channelID, atc.Url, 0)
	if err != nil {
		return 0, fmt.Errorf("insertArticle: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("insertArticle: %v", err)
	}
	return id, nil
}

func queryArticleUrl(articleID int64) (string, error) {
	var url string

	row := db.QueryRow("SELECT url FROM articles WHERE id = ?", articleID)
	if err := row.Scan(&url); err != nil {
		if err == sql.ErrNoRows {
			return url, fmt.Errorf("urlByID %d: no such article", articleID)
		}
		return url, fmt.Errorf("urlById %d: %v", articleID, err)
	}
	return url, nil
}

func updateArticleWordCount(articleID int64) {
	url, err := queryArticleUrl(articleID)
	if err != nil {
		return
	}

	// parse and process hmtl
	htmlRaw, err := getRawHTML(url)
	if err != nil {
		return
	}
	html := p.Sanitize(htmlRaw)
	content := strip.StripTags(html)
	wc := int64(len(strings.Fields(content)))

	fmt.Printf("Async update for article for %v", url)
	fmt.Printf(" --> word count: %v", wc)

	// update word count
	_, err = db.Exec("UPDATE articles SET wc = ? WHERE id=?", wc, articleID)
	if err != nil {
		return
	}
}
