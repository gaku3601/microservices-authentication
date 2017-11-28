# ログイン仕様
emailとpasswordをjsonでusers/loginに送信された場合、jwtトークンを生成し返却する。  
ただし、jwtトークンを生成する元となるkeyは別サーバとして稼働しているkongから発行される。  
また、emailとpasswordは事前にUser登録としてDB上に格納されていること。  
登録されていないものが渡された場合、エラーを返却する。  
passwordはdb上にハッシュ化されている。  
  
# TODO
- users/loginで待ち受けるサーバを作成する。
- emailとpasswordをjsonで受け取る
	- emailとpasswordがDBにUser登録されていることを確認する。
		- passwordをhash化する
		- DB接続を行う。
		- DB上にemailとpasswordに合致するUserが存在するか確認する。
			- 存在しない場合エラーを返却する。
			- 存在している場合、userIDのみ取得する。
- jwtトークンをjsonで返却する。
	- kongと接続する。
	- jwtkeyを取得する。
	- jwtトークンを作成する。
		- jwtトークンにはUserIDを格納する。
