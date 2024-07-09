package model

import (
	"time"
)

type User struct{
	UserId    uint 	   `gorm:"primaryKey"`
	Username string    
	Password string    `json:"-"`
	Orders	[]Order    `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
}

type Order struct {
	UserId    	uint 	     
	OrderId	   	uint         `gorm:"primaryKey"`	
	Status    	string       
	Name        string		 
    CreatedAt 	time.Time    
	FinishedAt 	time.Time    
}
