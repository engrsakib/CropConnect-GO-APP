package models

import (
	"time"

	"gorm.io/gorm"
)

const (
    UserTypeSuperAdmin = 0 // super Admin: has all permissions, can manage all aspects of the system
    UserTypeFarmer    = 1 // Farmer: user who can manage their own farm and products
    UserTypeAdmin      = 2 // Admin (Subscriber): user who can manage their own farm and products, and also has some administrative permissions
    UserTypeCustomer   = 3 // Regular User: user who can browse products, place orders, and manage their own profile
)

type User struct {
    ID       int64     `gorm:"primaryKey;autoIncrement" json:"id"`
    Name     string    `gorm:"type:varchar(100);index;not null" json:"name"`
    Mobile   string    `gorm:"type:varchar(20);uniqueIndex;not null" json:"mobile"`
    Email    string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
    Password string    `gorm:"type:varchar(255);not null" json:"-"`
    Photo    string    `gorm:"type:varchar(255)" json:"photo"`

    
    UserType int8 `gorm:"type:smallint;default:3;not null" json:"user_type"` 

   Permissions []UserPermission `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"permissions"`

   DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
    
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}