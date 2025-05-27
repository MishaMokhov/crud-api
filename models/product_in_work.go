package models

import "gorm.io/gorm"

type ProductInWork struct {
    gorm.Model

    UserID uint
    ProductID uint

    User    User
    Product    Product
}
