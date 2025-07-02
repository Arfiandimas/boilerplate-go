package entity

import "time"

//go:generate easytags $GOFILE db,json

/*
Contoh database entity
Untuk membuat semua komponen table database maka wajib menggunakan tag db:column_name
nama kolom harus sama dengan nama kolom pada tabel database

Agar dapat menggunakan Entity langsung pada response dapat menambahkan tag json:column_name
nama tag dapat disesuaikan.

Anda dapat menambahkan entity sesuai dengan kehendak
*/
type Example struct {
	ID        uint64     `db:"id" json:"id"`
	Name      string     `db:"name" json:"name"`
	Address   string     `db:"address" json:"address"`
	Email     string     `db:"email" json:"email"`
	Phone     string     `db:"phone" json:"phone"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}
