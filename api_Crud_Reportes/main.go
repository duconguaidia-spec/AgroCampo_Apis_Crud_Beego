package main

import (
	_ "api_crud_reportes/routers"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load()
	if err != nil{
		panic("Error cargando archivo, env")
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
	pgSchema, _ := web.AppConfig.String("PGschema")

	fmt.Printf("PostgreSQL connection string: postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s\n", pgUser, pgPass, pgHost, pgPort, pgDb, pgSchema)

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


	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{
		 "Origin", 
		 "x-requested-with",
		 "content-type",
		 "accept",
		 "origin",
		 "autorization",
		 "x-csrftoken"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials:true,
	}))

	beego.Run()
}