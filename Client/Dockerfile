FROM golang:1.22-alpine

WORKDIR /${appName}


RUN go get -u github.com/wailsapp/wails/cmd/wails


CMD ["wails init myapp"]