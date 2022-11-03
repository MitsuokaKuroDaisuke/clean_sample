# Sample

VSCodeとRemote Containerインストールして、Reopen in Container
基本コンテナの立ち上げまでは↓と一緒
https://zenn.dev/mitsu5/articles/a0755adbf4b55f

コンテナ立ち上がってコンテナ内のターミナル開いたら、

$ go run batch/db/refresh.go

を実行。DBに初期データが入る。

サーバーの起動↓

$ air

http://localhost:8080/users で全てのユーザがjson形式で取得。こっちはAPI。

http://localhost:8080/user/1 とかで指定されたIDのユーザデータが、テンプレートエンジンを使ってhtml表示