package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddPasswordDigestToUser_20220515_143909 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddPasswordDigestToUser_20220515_143909{}
	m.Created = "20220515_143909"

	migration.Register("AddPasswordDigestToUser_20220515_143909", m)
}

// Run the migrations
func (m *AddPasswordDigestToUser_20220515_143909) Up() {
	m.SQL("ALTER TABLE users ADD password_digest varchar(255) NOT NULL DEFAULT ''")

}

// Reverse the migrations
func (m *AddPasswordDigestToUser_20220515_143909) Down() {
	m.SQL("ALTER TABLE users DROP password_digest")

}
