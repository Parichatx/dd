package entity

import (
	"gorm.io/gorm"
)

type Mem struct{
	gorm.Model
	Star string `valid:"required~Star is required,matches(^[A-Z]{3}\\d{2}$)~Star is invalid"`
//	IDE string `valid: ",required~ID is required"`
	Email string `valid:"email~Email is invalid,required~Email is required"`
}