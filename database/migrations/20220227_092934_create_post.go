package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreatePost_20220227_092934 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreatePost_20220227_092934{}
	m.Created = "20220227_092934"

	migration.Register("CreatePost_20220227_092934", m)
}

// Run the migrations
func (m *CreatePost_20220227_092934) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE posts(id SERIAL,title TEXT NOT NULL,content TEXT NOT NULL,user_id integer DEFAULT NULL,created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL)")
}

// Reverse the migrations
func (m *CreatePost_20220227_092934) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE create_post")
}
