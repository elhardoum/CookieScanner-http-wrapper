# CookieScanner HTTP Wrapper

This project wraps the GDPR cookie scanner [`CovenantSQL/CookieScanner`](https://github.com/CovenantSQL/CookieScanner) in HTTP to make it easy for your app to call-in the service with website URLs and get the scan results in JSON format.

## Installation

### Using Docker

It's as simple as running the containers in a detached state:

```sh
docker-compose up -d
```

It'll build the containers on first runs.

### Without Docker

In a golang-ready environment, where you also have `chromium-browser` installed, add these two dependencies that this project requires:

```sh
go get gopkg.in/alessio/shellescape.v1
go get github.com/CovenantSQL/CookieScanner
```

Next, run the HTTP server

```sh
go run server.go
```

## Usage

```sh
curl http://127.0.0.1:8000/scan-website?url=https://www.wikipedia.org
# or POST
curl http://127.0.0.1:8000/scan-website -d 'url=https://www.wikipedia.org'
```