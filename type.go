package bemongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role,moitempty" bson:"role,omitempty"`
	Token    string `json:"token,omitempty" bson:"token,omitempty"`
	Private  string `json:"private,omitempty" bson:"private,omitempty"`
	Public   string `json:"public,omitempty" bson:"public,omitempty"`
}

type Credetial struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	NomorID     int                `json:"nomorid" bson:"nomorid"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Price       int                `json:"price" bson:"price"`
	Stock       int                `json:"stock" bson:"stock"`
	Size        string             `json:"size" bson:"size"`
	Image       string             `json:"image" bson:"image"`
}

type Response struct {
	Status  bool        `json:"status" bson:"status"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data" bson:"data"`
}

type Content struct {
	ID          int    `json:"id" bson:"id"`
	Content     string `json:"content" bson:"content"`
	Image       string `json:"image" bson:"image"`
	Description string `json:"description" bson:"description"`
}

type Blog struct {
	ID          int       `json:"id" bson:"id"`
	Title       string    `json:"title" bson:"title"`
	Tanggal     string    `json:"tanggal" bson:"tanggal"`
	Description string    `json:"description" bson:"description"`
	Content     []Content `json:"content" bson:"content"`
}

type Tags struct {
	Tags []string `json:"tags" bson:"tags"`
}

type Category struct {
	Category []string `json:"category" bson:"category"`
}

type Comment struct {
	ID        int    `json:"id" bson:"id"`
	Username  string `json:"username" bson:"username"`
	Answer    string `json:"comment" bson:"comment"`
	Questions string `json:"questions" bson:"questions"`
	Tanggal   string `json:"tanggal" bson:"tanggal"`
}

type Share struct {
	Share []string `json:"share" bson:"share"`
}

type EventGlobal struct {
	ID          int    `json:"id" bson:"id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Tanggal     string `json:"tanggal" bson:"tanggal"`
	Image       string `json:"image" bson:"image"`
	Harga       int    `json:"harga" bson:"harga"`
}

type Event struct {
	ID          int    `json:"id" bson:"id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Tanggal     string `json:"tanggal" bson:"tanggal"`
	Image       string `json:"image" bson:"image"`
	Harga       int    `json:"harga" bson:"harga"`
}

type About struct {
	ID          int    `json:"id" bson:"id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Image       string `json:"image" bson:"image"`
}

type Gallery struct {
	ID          int    `json:"id" bson:"id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Image       string `json:"image" bson:"image"`
}

type Contact struct {
	ID      int    `json:"id" bson:"id"`
	Name    string `json:"name" bson:"name"`
	Subject string `json:"subject" bson:"subject"`
	Image   string `json:"image" bson:"image"`
	Email   string `json:"email" bson:"email"`
	Phone   string `json:"phone" bson:"phone"`
}

type Iklan struct {
	ID          int    `json:"id" bson:"id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Image       string `json:"image" bson:"image"`
}
