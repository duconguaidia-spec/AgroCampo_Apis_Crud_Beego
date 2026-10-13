package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/beego/beego/v2/client/orm/migration"
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
	file, err := readSQLFile("20260521_143218_creacion_db_up.sql")
	if err != nil {
		fmt.Println(err)
		return
	}
	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		request = strings.TrimSpace(request)
		if request != "" {
			fmt.Println(request)
			m.SQL(request)
		}
	}
}

// Reverse the migrations
func (m *CreacionDb_20260521_143218) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := readSQLFile("20260521_143218_creacion_db_down.sql")
	if err != nil {
		fmt.Println(err)
		return
	}
	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		request = strings.TrimSpace(request)
		if request != "" {
			fmt.Println(request)
			m.SQL(request)
		}
	}

}

func readSQLFile(name string) ([]byte, error) {
	paths := []string{
		filepath.Join("database", "scripts", name),
		filepath.Join("..", "scripts", name),
		filepath.Join("..", "database", "scripts", name),
	}

	for _, path := range paths {
		file, err := os.ReadFile(path)
		if err == nil {
			return file, nil
		}
	}

	return nil, fmt.Errorf("no se pudo leer el script SQL %s", name)
}
