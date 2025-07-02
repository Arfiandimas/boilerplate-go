// Package presentations
package presentations

//go:generate easytags $GOFILE json,form

/*
ExampleRequest Presentation Request Model HTTP Example

Pada bagian presentation anda dapat membuat sebuat entity secara general yaitu Response dan Request
tetapi tidak tertutup kemungkinan anda dapat menggunakannya untuk membuat sebuah entity lainnya

Pada request terdapat validasi sederhana yang dapat digunakan dengan menyisipkan pada sebuah tag
contoh sederhana yaitu tag validate dan binding untuk penjelasan lebih lanjut
dapat membaca document dari paket go-validation
*/
type ExampleRequest struct {
	ID      uint64 `json:"id,omitempty" form:"id"`
	Name    string `json:"name" validate:"required" binding:"required" form:"name"`
	Address string `json:"address" validate:"required" binding:"required" form:"address"`
	Email   string `json:"email" validate:"required" binding:"required" form:"email"`
	Phone   string `json:"phone" validate:"required" binding:"required" form:"phone"`
}

// PublishMessageRequest Presentation Request Model MQ Example
type PublishMessageRequest struct {
	Topic     string `json:"topic" form:"topic" validate:"required" binding:"required"`
	Partition int32  `json:"partition" form:"partition"`
	Messages  string `json:"messages" form:"messages" validate:"required" binding:"required"`
}

type PublishMessagePubsub struct {
	Topic    string `json:"topic" form:"topic" validate:"required" binding:"required"`
	Messages string `json:"messages" form:"messages" validate:"required" binding:"required"`
}
