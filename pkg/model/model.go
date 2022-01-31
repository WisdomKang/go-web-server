package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string
	Password  string
	LastLogin time.Time

	NoticeSubscription string
	EmailSubscription  string
}

type Notice struct {
	gorm.Model
	Title   string
	Content string
}

type NoticeRead struct {
	gorm.Model
	UserId   uint
	NoticeId uint
}

type Content struct {
	gorm.Model
	Title    string
	Thumnail uint
	Type     string
}

type Movie struct {
	gorm.Model
	Title       string
	Thumnail    uint
	Video       Video
	Description string
}

type Drama struct {
	gorm.Model
	Title       string
	Thumnail    uint
	Part        []DramaPart
	Description string
}

type DramaPart struct {
	gorm.Model
	PartNumber  int
	Video       Video
	Description string
}

type Video struct {
	gorm.Model
	File ContentsFile
}

type ContentsFile struct {
	gorm.Model
	OriginName string
	UUIDName   string
	Path       string
}
