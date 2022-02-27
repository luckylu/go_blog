package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateComments_20220227_092423 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateComments_20220227_092423{}
	m.Created = "20220227_092423"

	migration.Register("CreateComments_20220227_092423", m)
}

// Run the migrations
func (m *CreateComments_20220227_092423) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE comments(id SERIAL,content TEXT NOT NULL,post_id integer DEFAULT NULL,user_id integer DEFAULT NULL,created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL)")
}

// Reverse the migrations
func (m *CreateComments_20220227_092423) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE create_comments")
}
