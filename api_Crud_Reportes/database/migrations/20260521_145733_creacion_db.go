package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreacionDb_20260521_145733 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionDb_20260521_145733{}
	m.Created = "20260521_145733"

	migration.Register("CreacionDb_20260521_145733", m)
}

// Run the migrations
func (m *CreacionDb_20260521_145733) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *CreacionDb_20260521_145733) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
