package main

import (
	"github.com/gofast-bbs/model"
	"github.com/gofast-bbs/routers"
)

func main() {
	model.InitDb()
	routers.InitRouter()
}
