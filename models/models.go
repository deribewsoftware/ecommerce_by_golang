package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.OrderID `json:"_id" bson:"_id"`
	First_Name      *string           `json:"first_name"       validate:"required,min=2 ,max=30" `
	Last_Name       *string           `json:"last_name"        validate:"required,min=2,max=30"`
	Phone           *string           `json:"phone"            validate:"required"`
	Email           *string           `json:"email"            validate:"email,required"`
	Passwords       *string           `json:"passwords"        validate:"required min=6"`
	Token           *string           `json:"token"`
	Refresh_Token   *string           `json:"refresh_token"`
	Created_At      time.Time         `json:"created_at"`
	Updated_At      time.Time         `json:"updated_at"`
	User_ID         string            `json:"user_id"`
	UserCart        []ProductUser     `json:"user_cart" bson:"user_cart"`
	Address_Details []Address         `json:"address_details" bson:"address_details"`
	Order_Status    []Order           `json:"order_status" bson:"order_status"`
}
type Product struct {
	ProductID    primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name"`
	Price        *uint64            `json:"price"`
	Rating       *uint8             `json:"rating"`
	Image        *string            `json:"image"`
}

type ProductUser struct {
	ProductID    primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name" bson:"rating"`
	Price        int                `json:"price" bson:"_id"`
	Rating       *uint              `json:"rating" bson:"rating"`
	Image        *string            `json:"image" bson:"image"`
}

type Address struct {
	AddressID primitive.ObjectID `bson:"_id"`
	House     *string            `json:"house" bson:"house"`
	Street    *string            `json:"street" bson:"street"`
	City      *string            `json:"city" bson:"city"`
	Pincode   *string            `json:"pincode" bson:"pincode"`
}
type Order struct {
	OrderID       primitive.ObjectID `bson:"_id"`
	Order_Cart    []ProductUser      `json:"order_cart" bson:"order_cart"`
	Order_At      time.Time          `json:"order_at" bson:"order_at"`
	Price         int64              `json:"price" bson:"price"`
	Discount      *int               `json:"discount" bson:"discount"`
	PaymentMethod Payment            `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool
	COD     bool
}
