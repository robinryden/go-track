# go-track

URL monitor written in Go.

@TODO
> 1. Write tests
> 2. Get urls from a txt-file etc.

## Settings are set inside the .env file
```golang
SLACK_WEBHOOK_URL="https://hooks.slack.com/services/id"
REDIS_HOST="0.0.0.0"
REDIS_PASSWORD=""
```

## URL's to ping
For the moment it's located in main.go

```golang
urls = []string{"https://www.google.com/urldoesntexist", "https://www.google.com/"}
```

## Run
```golang
go run main.go
```
