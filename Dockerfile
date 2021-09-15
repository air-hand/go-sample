FROM golang:1.17-bullseye as builder

RUN echo "deb http://deb.debian.org/debian bullseye-backports main" > /etc/apt/sources.list.d/backports.list \
    && apt-get update -qq \
    && apt-get install -y software-properties-common dirmngr apt-file \
    && apt-key adv --fetch-keys 'https://mariadb.org/mariadb_release_signing_key.asc' \
    && add-apt-repository 'deb [arch=amd64,arm64,ppc64el] https://tw1.mirror.blendbyte.net/mariadb/repo/10.6/debian bullseye main' \
    && apt-get update -qq \
    && apt-get install -y git tig vim less bash sudo mariadb-client \
    && apt-get install -y fontconfig fonts-noto-cjk \
    && fc-cache -fv \
    ;

#ENV USER=app \
ENV USER=root \
    GO111MODULE=on \
    EDITOR=vim \
    LANG=C.UTF-8

#RUN addgroup wheel \
#    && echo "auth sufficient pam_wheel.so trust group=wheel" >> /etc/pam.d/su \
#    && echo "%wheel ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers \
#    && useradd -s /bin/bash -m -G wheel $USER \
#    ;

USER $USER

#ENV HOME=/home/$USER
ENV HOME=/$USER

COPY --chown=$USER:$USER .bashrc .vimrc $HOME/

RUN curl -fsSL https://raw.github.com/git/git/master/contrib/completion/git-completion.bash -o ~/.git-completion.bash \
    && curl -fsSL https://raw.github.com/git/git/master/contrib/completion/git-prompt.sh -o ~/.git-prompt.sh \
    && chmod a+x ~/.git-* \
    ;

WORKDIR /opt/app/src

RUN go get golang.org/x/tools/gopls@latest \
    && go get -u github.com/ramya-rao-a/go-outline \
    && go get github.com/clipperhouse/gen \
    && go install github.com/go-delve/delve/cmd/dlv@master \
    && mv $GOPATH/bin/dlv $GOPATH/bin/dlv-dap \
    ;

COPY --chown=$USER:$USER src ./

RUN go mod download

RUN mkdir -p /go/out && go build -o /go/out/app

WORKDIR /opt/app

COPY --chown=$USER:$USER Makefile .editorconfig ./

# multi stage build for slim
FROM gcr.io/distroless/base-debian10:latest as prod

COPY --from=builder /go/out/app /go/out/app

ENTRYPOINT ["/go/out/app"]
