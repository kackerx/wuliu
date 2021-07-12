package controller

import "github.com/gin-gonic/gin"

func AddNavigation(c *gin.Context) {
    c.HTML(200, "add.html", gin.H{
        "title": "admin",
    })
}