FROM golang:1.16-buster as builder

RUN apt-get update -qq && \
    apt-get install -y zsh

ENV GO111MODULE=on

RUN go get golang.org/x/tools/gopls@latest \
    && go get -u github.com/lukehoban/go-outline

WORKDIR /opt/app

COPY go.* /opt/app/

RUN go mod download

COPY * /opt/app/

RUN go build -o /go/bin/app

# multi stage build for slim
FROM gcr.io/distroless/base-debian10:latest as prod

COPY --from=builder /go/bin/app /go-app

CMD ["/go-app"]
