package main

import (
	_ "api_crud_soporte/routers"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {

	err := loadProjectEnv()
	if err != nil {
		panic(err)
	}

	err = beego.LoadAppConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}

	pgUser, _ := web.AppConfig.String("PGuser")
	pgPass, _ := web.AppConfig.String("PGpass")
	pgHost, _ := web.AppConfig.String("PGhost")
	pgPort, _ := web.AppConfig.String("PGport")
	pgDb, _ := web.AppConfig.String("PGdb")
	pgSchema, _ := web.AppConfig.String("PGsschema")

	fmt.Printf("PostgreSQL configured database=%s search_path=%s\n", pgDb, pgSchema)

	orm.RegisterDataBase(
		"default",
		"postgres",
		"postgres://"+
			pgUser+":"+
			pgPass+"@"+
			pgHost+":"+
			pgPort+"/"+
			pgDb+
			"?sslmode=disable&search_path="+
			pgSchema,
	)

	// Permite leer correctamente el body en JSON
	beego.BConfig.CopyRequestBody = true

	// Evita errores de directorios
	beego.BConfig.WebConfig.DirectoryIndex = false

	// Configuración swagger
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{
			"Origin",
			"x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))

	// Ejecutar servidor
	beego.Run()
}

func loadProjectEnv() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	for {
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			return godotenv.Load(envPath)
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return errors.New("no se encontró el archivo .env en la carpeta del proyecto ni en sus carpetas superiores")
		}
		dir = parent
	}
}
