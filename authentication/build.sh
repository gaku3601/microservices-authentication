cd /go/src/github.com/gaku3601/study-microservices/authentication

#DBセットアップ
goose -env production up

#ビルド&起動
go build .
./authentication
#while sleep 3600; do :; done
