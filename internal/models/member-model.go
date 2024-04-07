package models

import (
	"github.com/jinzhu/gorm"
	"github.com/oyevamos/jun-test-0.git/internal/repository"
)

var db *gorm.DB

type Member struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email"`
	Country string `json:"country"`
	Color   string `json:"color"`
}

func init() {
	repository.Connect()
	db = repository.GetDB()
	db.AutoMigrate(&Member{})
}

func (m *Member) CreateMember() *Member {
	db.NewRecord(m)
	db.Create(&m)
	return m
}

func GetAllMembers() []Member {
	var Members []Member
	db.Find(&Members)
	return Members
}

func GetMemberById(Id int64) (*Member, *gorm.DB) {
	var getMember Member
	db := db.Where("ID=?", Id).Find(&getMember)
	return &getMember, db
}

func DeleteMember(ID int64) Member {
	var member Member
	db.Where("ID=?", ID).Delete(member)
	return member
}
