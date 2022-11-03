package main

import (
	"www/infra/router"
)

// ArticleJob 記事取得スケジューラ
type ArticleJob struct {
}

func main() {

	rt := router.Init()

	rt.Logger.Fatal(rt.Start(":8080"))
}
