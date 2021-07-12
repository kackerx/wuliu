package model

type Article struct {
    *Model
    Title   string
    Content string
}

func (a Article) TableName() string {
    return "wuliu_article"
}

type ArticleVO struct {
    *Article
    PreArticle  *Article
    NextArticle *Article
    CreatedOn   string
}
