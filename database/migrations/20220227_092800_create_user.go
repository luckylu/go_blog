package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateUser_20220227_092800 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUser_20220227_092800{}
	m.Created = "20220227_092800"

	migration.Register("CreateUser_20220227_092800", m)
}

// Run the migrations
func (m *CreateUser_20220227_092800) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE users(id SERIAL,name TEXT NOT NULL,created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL)")
}

// Reverse the migrations
func (m *CreateUser_20220227_092800) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE create_user")
}
