# 設定
環境変数AuthEnvにproduction or developと設定することで、設定ファイルを読み替えています。  
developの場合設定する必要はありませんが、productionの場合は必ず環境変数を設定してください。  

# 認証用のcurl

    curl -H 'Content-Type:application/json' -H 'User-Agent:iPhone' -H 'Accept-Encoding:gzip,deflate' -d '{"ID":"gaku","pass":"gakugaku"}' http://localhost:8080/users/auth

# JWT認証の追加
まずはkong上でユーザの作成を行う。

    curl -X POST http://kong:8001/consumers     --data "username=gaku"     --data "custom_id=gaku"

このユーザを使用し、JWTのトークンを生成する元となるKeyを以下コマンドで問い合わせする。

    curl -X POST http://localhost:8001/consumers/gaku/jwt -H "Content-Type: application/x-www-form-urlencoded"

gaku部分は上記で作成したユーザ名を指定している。  
これで問い合わせを行うと以下のような形で返却される。

    {
      "created_at":1510615420000,
      "id":"79051f7f-ff24-4d69-8e1d-d83146bc9ec7",
      "algorithm":"HS256",
      "key":"RmzcPktBjNbnsGdZPwLioOmdThCjFGIO",
      "secret":"wKtru3BuCiT9vFFki77cg5DE2rt6a4if",
      "consumer_id":"0a086d40-dafd-43e2-94dc-835d1b96c92b"
    }

これを元にgolang側でJsonTokenを発行する。  
トークンの発行が完了したら、以下コマンドで対象のAPIをJWT認証必須にする。

    curl -X POST http://kong:8001/apis/{api}/plugins     --data "name=jwt"

これで、{api}で指定したものが、JWTトークンをヘッダーに設定しないと取得できなくなる。  
JWTトークンの発行まではbasic認証で行い、発行が完了したらJWTトークンのみで通信するような感じにすれば認証機能を追加できる。
