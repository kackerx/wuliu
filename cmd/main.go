package main

import (
	"github.com/gin-gonic/gin"
	"wuliu/router"
)

func main() {
	r := gin.Default()
	router.NewRouter(r)
	router.NewRouter2(r)
	r.Run()
}
