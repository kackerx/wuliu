package model

type Model struct {
	ID       uint  `json:"id" gorm:"primary_key"`
	CreateOn int32  `json:"create_on"`
	CreateBy string `json:"create_by"`
}

