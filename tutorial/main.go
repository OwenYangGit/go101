package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User struct
type User struct {
	ID        int64     `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"id,omitempty"`
	Username  string    `gorm:"type:varchar(20) NOT NULL;" json:"username,omitempty"`
	Password  string    `gorm:"type:varchar(100) NOT NULL;" json:"password,omitempty"`
	Status    int32     `gorm:"type:int(5);" json:"status,omitempty"`
	CreatedAt time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db2, err1 := db.DB()
	if err1 != nil {
		fmt.Println("get db failed", err)
		return
	}
	db2.SetConnMaxLifetime(time.Duration(10) * time.Second)
	db2.SetMaxIdleConns(10)
	db2.SetMaxOpenConns(10)

	//產生table
	db.Debug().AutoMigrate(&User{})

	migrator := db.Migrator()
	has := migrator.HasTable(&User{})
	//has := migrator.HasTable("GG")
	if !has {
		fmt.Println("table not exist")
	}

	user := User{Username: "tester", Password: "12333", Status: 1}
	result := db.Debug().Create(&user)
	if result.Error != nil {
		fmt.Println("Create failed")
	}
	if result.RowsAffected != 1 {
		fmt.Println("RowsAffected Number failt")
	}
}
