# Channel Newss APIs

## Backend APIs

Endpoints:

`/api/channel`
- `GET`: Get a list of all channels, return as JSON
- `POST`: Add a new channel as requested in JSON

Manage articles of a channel

`/api/channel/:id`
- `GET`: Get a list of articles in a channel by its id return as JSON, can be optionally filterd by word count as requested in JSON
- `POST`: Add a news article into a channel, url in JSON as requested, process in background 

## Schemas

- Channels: `id`, `name`
- Articles: `id`, `channel`, `url`, `wc`

## Usage

## Progress

- [ ] Set up database
- [ ] Lauch server
- [ ] Add endpoints for channel
- [ ] Integrate UI for APIs
- [ ] Add endpoints for articles
- [ ] Wrap up and documentation
