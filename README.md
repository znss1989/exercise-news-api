# MM News APIs

## RESTful endpoints

`/api/channel`
- `GET`: Get a list of all channels, return as JSON
- `POST`: Add a new channel as requested in JSON

Manage articles of a channel

`/api/channel/:id`
- `GET`: Get a list of articles in a channel by its id return as JSON, can be optionally filterd by word count as requested in JSON
- `POST`: Add a news article into a channel, url in JSON as requested, process in background 

## Schemas

- Channels: `id`, `title`
- Articles: `id`, `channel_id`, `url`, `wc`

## Usage

### Prerequisites

SQLite3 is assumed for the database. 

### Launch of application

```bash
sqlite3 news.db < setup.sql
```

## Progress

- [x] Set up database
- [ ] Lauch server
- [ ] Add endpoints for channel
- [ ] Integrate UI for APIs
- [ ] Add endpoints for articles
- [ ] Filter articles by word count
- [ ] Wrap up and documentation
