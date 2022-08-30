# MM News APIs

This project sets up RESTful APIs for managing news `articles` under `channels`.

Features:
- Able to manage channels.
- Able to manage articles for a channel.
- When adding articles, only the URL is required.
    The Application will fetch the URL and calculate the word count (HTML tags stripped).
- Fetching the URLs and counting the words is to be done in the background after the article URL has been received.
    Can search articles within word count ranges e.g. 0-100, 100-500, 0-501.

A SQLite database is set up for the data persisting, for both channels and articles.

## RESTful endpoints

`/api/channel`
- `GET`: Get a list of all channels, return as JSON
- `POST`: Add a new channel as requested in JSON

Manage articles of a channel

`/api/channel/:id`
- `GET`: Get a list of articles in a channel by its id return as JSON, can be optionally filterd by word count as requested in JSON
- `POST`: Add a news article into a channel, url in JSON as requested, process in background 

### Models

- Channels: `id`, `title`
- Articles: `id`, `channel_id`, `url`, `wc`

## Usage

### Prerequisites

SQLite3 is assumed for the database. 

### Launch of application

```bash
sqlite3 news.db < setup.sql
```

### Launch server

Add dependencies and 

```bash
go get .
go run .
```

### Swagger UI

Access the Swagger UI at `http://localhost:8080/docs/index.html#/` when the app is launched.

Update docs when changes are made

```
export PATH=$(go env GOPATH)/bin:$PATH
swag init
```

## Progress

- [x] Set up database
- [x] Lauch server
- [x] Add endpoint for channel manipulation
- [x] Add endpoint for adding an article
- [x] Add endpoint for search articles of a channel
- [x] Process the word count of an article
- [x] Filter articles by word count
- [x] Handle invalid and errors
- [x] Integrate UI for APIs
- [x] Wrap up and documentation
