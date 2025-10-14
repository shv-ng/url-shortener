
# URL Shortener

A simple URL shortener written in Go, using PostgreSQL for storage.

## Features

* Shorten URLs via a simple API
* Store URLs in PostgreSQL
* Retrieve original URLs by accessing the short URL
* Easy to run with Docker Compose

## Getting Started

### Prerequisites

* Docker and Docker Compose installed on your machine

### Running the Project

```bash
git clone https://github.com/shv-ng/url-shortener
cd url-shortener
docker-compose up -d
```

This will start the application and the PostgreSQL database in the background.

## How to Use

### Shorten a URL

Make a `GET` request to:

```
http://127.0.0.1:8000/shorturl
```

With a JSON body containing the URL you want to shorten:

```json
{
  "url": "https://github.com"
}
```

Response example:

```json
{
  "short_url": "3097fca9"
}
```

### Access the Original URL

Open your browser or make a `GET` request to:

```
http://127.0.0.1:8000/3097fca9
```

This will redirect you to the original URL, e.g., `https://github.com`.

## Notes

* The project currently does **not** use Go routines or concurrency for simplicity.
* Storage is handled by PostgreSQL.
* Feel free to contribute or suggest improvements!

