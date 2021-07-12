package router

import (
    "github.com/gin-contrib/multitemplate"
    "github.com/gin-gonic/gin"
    "net/http"
    "path/filepath"
    "wuliu/controller"
)

func NewRouter(r *gin.Engine) {
    //r := gin.New()
    //r.Use(gin.Logger())
    //r.Use(gin.Recovery())
    //r := gin.Default()
    
    //r.LoadHTMLGlob("template/html/*")
    r.StaticFS("/template", http.Dir("./template"))
    //r.HTMLRender = loadTemplates("./template")
    r.LoadHTMLGlob("template/**/*.html")
    {
        r.GET("/", controller.Home)
        r.GET("/list", controller.List)
        r.GET("/content", controller.Content)
        r.GET("/add", controller.AddNavigation)
    }
    
    //admin := r.Group("/admin123")
    //{
    //    admin.GET("/addNavigation", controller.AddNavigation)
    //}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
    r := multitemplate.NewRenderer()
    
    layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
    if err != nil {
        panic(err.Error())
    }
    
    includes, err := filepath.Glob(templatesDir + "/includes/*.html")
    if err != nil {
        panic(err.Error())
    }
    
    // Generate our templates map from our layouts/ and includes/ directories
    for _, include := range includes {
        layoutCopy := make([]string, len(layouts))
        copy(layoutCopy, layouts)
        files := append(layoutCopy, include)
        r.AddFromFiles(filepath.Base(include), files...)
    }
    return r
}
