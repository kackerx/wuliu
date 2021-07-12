package controller

import (
    "github.com/gin-gonic/gin"
    "wuliu/service"
)

type test struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func Home(c *gin.Context) {
    c.HTML(200, "home.html", gin.H{
        "title": "Home",
    })
}

func List(c *gin.Context) {
    c.HTML(200, "list.html", gin.H{
        "title":    "List",
        "articles": service.GetArticles(4, 0),
    })
}

func Content(c *gin.Context) {
    c.HTML(200, "content.html", gin.H{
        "title": "content",
        "resp":  service.GetArticleByID(1),
    })
}

