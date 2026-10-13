package main

import (
	_ "API_CRUD_NOTICIAS/routers"
	"fmt"
	"github.com/beego/beego/v2/server/web/filter/cors"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error cargando archivo .env")
	}

	// Reacargar app.conf despues de cargar el archivo .env
	err = beego.LoadAppConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}

	pgUser, _ := beego.AppConfig.String("PGuser")
	pgPass, _ := beego.AppConfig.String("PGpass")
	pgHost, _ := beego.AppConfig.String("PGhost")
	pgPort, _ := beego.AppConfig.String("PGport")
	pgDB, _ := beego.AppConfig.String("PGdb")
	pgSSchema, _ := beego.AppConfig.String("PGsschema")

	fmt.Printf("PostgreSQL configured database=%s search_path=%s\n", pgDB, pgSSchema)

	orm.RegisterDataBase(
		"default",
		"postgres",
		"postgres://"+
			pgUser+":"+
			pgPass+"@"+
			pgHost+":"+
			pgPort+"/"+
			pgDB+
			"?sslmode=disable&search_path="+
			pgSSchema,
	)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"Origin",
			"x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"content-Length"},
		AllowCredentials: false,
	}))
	beego.Run()
}
