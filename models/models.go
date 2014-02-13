package models

import (
	"database/sql"
	"strconv"

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
	//db := GetLink()
	//db.FindAll(&blogs)
	db, derr := sql.Open("sqlite3", "db/data.sqlite")
	if derr != nil {
		panic(derr)
	}
	rows, err := db.Query("select * from blog")
	checkErr(err)
	for rows.Next() {
		var blog Blog
		err = rows.Scan(&blog.Id, &blog.Title, &blog.Content, &blog.Created)
		blogs = append(blogs, blog)
	}
	return
}

func GetBlog(id int) (blog Blog) {
	//db := GetLink()
	//db.Where("id=?", id).Find(&blog)
	db, derr := sql.Open("sqlite3", "db/data.sqlite")
	if derr != nil {
		panic(derr)
	}
	rows, err := db.Query("select * from blog where id=?", strconv.Itoa(id))
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&blog.Id, &blog.Title, &blog.Content, &blog.Created)
	}
	return
}

func SaveBlog(blog Blog) (bg Blog) {
	db := GetLink()
	db.Save(&blog)
	return bg
}

func DelBlog(id int) {
	//db := GetLink()
	//db.Delete(&blog)
	db, derr := sql.Open("sqlite3", "db/data.sqlite")
	checkErr(derr)
	stmt, err := db.Prepare("delete from blog where id=?")
	checkErr(err)
	_, err = stmt.Exec(id)
	checkErr(err)

	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
