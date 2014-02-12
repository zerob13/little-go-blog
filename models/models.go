package models

import (
	"database/sql"

	"github.com/astaxie/beedb"
	_ "github.com/mattn/go-sqlite3"
)

/**
**/

type Blog struct {
	Id      int `PK`
	Title   string
	Content string
	Created string
}

func GetLink() beedb.Model {
	db, err := sql.Open("sqlite3", "db/data.sqlite")
	if err != nil {
		panic(err)
	}
	orm := beedb.New(db)
	return orm
}

func GetAll() (blogs []Blog) {
	db := GetLink()
	db.FindAll(&blogs)
	return
}

func GetBlog(id int) (blog Blog) {
	db := GetLink()
	db.Where("id=?", id).Find(&blog)
	return
}

func SaveBlog(blog Blog) (bg Blog) {
	db := GetLink()
	db.Save(&blog)
	return bg
}

func DelBlog(blog Blog) {
	db := GetLink()
	db.Delete(&blog)
	return
}
