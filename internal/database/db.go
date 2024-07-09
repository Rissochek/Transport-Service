package database

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	
	"LastProject/internal/model"
)

func InitDataBase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=1717 dbname=postgres port=5000 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed init db")
	}

	db.AutoMigrate(&model.User{}, &model.Order{})

	return db, nil
}

func AddUserToDataBase(db *gorm.DB, user *model.User) error{
	result := db.Create(user)
	if result != nil {
		return errors.New("failed add user to db")
	}
	return nil
}

func AddOrderToDataBase(db *gorm.DB, order *model.Order) error{
	result := db.Create(order)
	if result != nil {
		return errors.New("failed add order to db")
	}
	return nil
}