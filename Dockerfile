FROM golang:1.16-buster as builder

RUN echo "deb http://deb.debian.org/debian buster-backports main" > /etc/apt/sources.list.d/backports.list \
    && apt-get update -qq \
    && apt-get install -y zsh git/buster-backports tig vim less bash \
    && curl -fsSL https://raw.github.com/git/git/master/contrib/completion/git-completion.bash -o $HOME/.git-completion.bash \
    && curl -fsSL https://raw.github.com/git/git/master/contrib/completion/git-prompt.sh -o $HOME/.git-prompt.sh \
    && chmod a+x $HOME/.git* \
    && echo "source ~/.git-completion.bash" >> $HOME/.bashrc \
    && echo "source ~/.git-prompt.sh" >> $HOME/.bashrc \
    && echo "export PS1=\"\u@\w\$(__git_ps1)\\$ \"" >> $HOME/.bashrc \
    ;

ENV GO111MODULE=on EDITOR=vim

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
