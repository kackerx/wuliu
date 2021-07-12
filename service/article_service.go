package service

import (
	"time"
	"wuliu/model"
	"wuliu/repository"
)

func GetArticles(pageIndex, pageSize int) []model.Article {
	return repository.GetArticleList(pageIndex, pageSize)
}

func GetArticleByID(id uint) model.ArticleVO {
	preArticle := repository.GetPreByID(id)
	if preArticle.Model == nil {
		preArticle.Title = "暂无"
	}
	nextArticle := repository.GetNextByID(id)
	if nextArticle.Model == nil {
		nextArticle.Title = "暂无"
	}
	currentArticle := repository.GetArticleByID(id)
	respVO := model.ArticleVO{
		Article:     &currentArticle,
		PreArticle:  &preArticle,
		NextArticle: &nextArticle,
		CreatedOn: time.Unix(int64(currentArticle.CreateOn), 0).Format("2006-01-02 03:04:05"),
	}
	return respVO
}
