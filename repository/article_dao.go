package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wuliu/model"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	
	dsn := "root:wasd4044516520@tcp(127.0.0.1:3306)/wuliu?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetArticleList(pageIndex, pageSize int) []model.Article {
	var articles []model.Article
	db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Find(&articles)
	return articles
}

func GetArticleByID(id uint) (article model.Article) {
	db.First(&article, id)
	return
}

func GetPreByID(id uint) model.Article {
	//db.Raw("select * from wuliu_article where id = (select id  from wuliu_article  where id < ? order by id desc limit 1);", id).Scan(&article)
	//return
	var pre model.Article
	db.Where("id < ?", id).Last(&pre)
	return pre
}

func GetNextByID(id uint) model.Article {
	var next model.Article
	db.Where("id > ?", id).First(&next)
	return next
}
