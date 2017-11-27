FROM golang:1.9.2-alpine3.6
#各種パッケージのインストール
RUN apk add --update alpine-sdk
RUN go get bitbucket.org/liamstask/goose/cmd/goose

ADD ./build.sh /build.sh
RUN chmod +x /build.sh

CMD /build.sh
