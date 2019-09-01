package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spaolacci/murmur3"
	"log"
	"strings"
)

func init() {
	var db *gorm.DB
	var err error
	if db, err = gorm.Open("mysql", "root:ps@(192.168.1.70:3307)/gotest?charset=utf8&parseTime=True&loc=Local"); err != nil {
		log.Fatal(err)
		return
	}

	if ! db.HasTable(&Like{}) {
		if err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Like{}).Error; err != nil {
			panic(err)
		}
	}
	G_db = db
}

func ADDLike(like *Like) error{
	like.Hash = murmur3.Sum64([]byte(strings.Join([]string{string(like.ID), like.Ua, like.Title}, "-"))) >> 1
	if err := G_db.Create(like).Error; err != nil {
		log.Fatal(err)
		return err
	}
	log.Print("insert record successully.")
	return nil
}

func QueryById(id int) *Like {
	var like Like
	if err := G_db.Find(&like, id).Error; err != nil {
		//log.Fatal(err)
		return nil
	}
	return &like;
}

