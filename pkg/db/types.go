package db

import (
	"github.com/jinzhu/gorm"
	"time"
)

var (
	G_db *gorm.DB
)

type Like struct {
	gorm.Model
	Ip        string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua        string `gorm:"type:varchar(256);not null;"`
	Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
	Hash      uint64 `gorm:"unique_index:hash_idx;"`
	Hates []Hate `gorm:"FOREIGNKEY:LikeId;ASSOCIATION_FOREIGNKEY:ID"`
	//Hates []Hate
}

type Hate struct {
	gorm.Model
	Name string
	LikeId uint
}

func (l *Like) ToDto() LikeDto {
	dto := LikeDto {
		Id: l.ID,
		Ua: l.Ua,
		Title: l.Title,
		Ip: l.Ip,
		CreatedAt: l.CreatedAt,
		UpdatedAt: l.UpdatedAt,
	}
	return dto;
}

type LikeDto struct {
	Id uint	`json:"id"`
	Ua string `json:"ua"`
	Title string `json:"title"`
	Ip string `json:"ip"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (dto *LikeDto) ToLike() Like{
	like := Like{
		Ip: dto.Ip,
		Ua: dto.Ua,
		Title: dto.Title,
	}
	return like
}


