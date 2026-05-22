package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableServicioveterinaria_20260522_130523 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableServicioveterinaria_20260522_130523{}
	m.Created = "20260522_130523"

	migration.Register("InsertTableServicioveterinaria_20260522_130523", m)
}

// Run the migrations
func (m *InsertTableServicioveterinaria_20260522_130523) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *InsertTableServicioveterinaria_20260522_130523) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
