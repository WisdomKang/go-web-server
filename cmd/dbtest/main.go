package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	MemberNumber string
	CreditCards  []CreditCard `gorm:"foreignKey:UserNumber;references:MemberNumber"`
}

type CreditCard struct {
	gorm.Model
	Number     string
	UserNumber string
}

func main() {
	dsn := "user=postgres password=test1234 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Seoul"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return
	}
	// Migrate the schema
	db.AutoMigrate(&CreditCard{})
	db.AutoMigrate(&User{})

	var newUser User = User{}
	db.Create(&newUser)

	fmt.Println(newUser)

}
