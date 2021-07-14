FROM golang:1.16-buster as builder

RUN apt-get update -qq \
    && apt-get install -y zsh git tig \
    && chsh -s /bin/zsh \
    && curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh | zsh \
    ;

ENV GO111MODULE=on

RUN go get golang.org/x/tools/gopls@latest \
    && go get -u github.com/lukehoban/go-outline

WORKDIR /opt/app/src

COPY src ./
COPY Makefile .editorconfig ../

RUN go mod download

RUN go build -o /go/bin/app

# multi stage build for slim
FROM gcr.io/distroless/base-debian10:latest as prod

COPY --from=builder /go/bin/app /go/bin/app

ENTRYPOINT ["/go/bin/app"]
