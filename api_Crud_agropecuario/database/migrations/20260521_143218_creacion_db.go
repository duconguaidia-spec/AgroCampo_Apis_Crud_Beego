package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
	"io/ioutil"
	"fmt"
	"strings"
)

// DO NOT MODIFY
type CreacionDb_20260521_143218 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionDb_20260521_143218{}
	m.Created = "20260521_143218"

	migration.Register("CreacionDb_20260521_143218", m)
}

// Run the migrations
func (m *CreacionDb_20260521_143218) Up() {
	file, err := ioutil.ReadFile("../database/migrations/20260521_143218_creacion_db.up.sql")
	if err != nil {
		fmt.Println(err)
	}
	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}

// Reverse the migrations
func (m *CreacionDb_20260521_143218) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../database/migrations/20260521_143218_creacion_db.down.sql")
	if err != nil {
		fmt.Println(err)
	}
	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}

}
