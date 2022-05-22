package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddAdminToUser_20220522_082031 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddAdminToUser_20220522_082031{}
	m.Created = "20220522_082031"

	migration.Register("AddAdminToUser_20220522_082031", m)
}

// Run the migrations
func (m *AddAdminToUser_20220522_082031) Up() {
	m.SQL("ALTER TABLE users ADD admin BOOLEAN DEFAULT FALSE")

}

// Reverse the migrations
func (m *AddAdminToUser_20220522_082031) Down() {
	m.SQL("ALTER TABLE users DROP admin")

}
