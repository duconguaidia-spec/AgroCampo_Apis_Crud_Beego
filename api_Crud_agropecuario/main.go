package main

import (
	_ "API_CRUD_AGROPECUARIO/routers"
	"errors"
	"os"
	"path/filepath"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	// Cargar variables del .env
	err := loadProjectEnv()
	if err != nil {
		panic(err)
	}

	// Recargar app.conf después del .env
	err = beego.LoadAppConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}

	// Obtener variables de PostgreSQL
	pgUser, _ := beego.AppConfig.String("PGuser")
	pgPass, _ := beego.AppConfig.String("PGpass")
	pgHost, _ := beego.AppConfig.String("PGhost")
	pgPort, _ := beego.AppConfig.String("PGport")
	pgDb, _ := beego.AppConfig.String("PGdb")
	pgsSchema, _ := beego.AppConfig.String("PGsschema")

	// Registrar base de datos
	err = orm.RegisterDataBase(
		"default",
		"postgres",
		"postgres://"+
			pgUser+":"+
			pgPass+"@"+
			pgHost+":"+
			pgPort+"/"+
			pgDb+
			"?sslmode=disable&search_path="+
			pgsSchema,
	)

	if err != nil {
		panic(err)
	}

	// Swagger en modo dev
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// Configuración CORS
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			"PUT",
			"PATCH",
			"GET",
			"POST",
			"OPTIONS",
			"DELETE",
		},
		AllowHeaders: []string{
			"Origin",
			"x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken",
		},
		ExposeHeaders:    []string{"content-length"},
		AllowCredentials: false,
	}))

	
	// Ejecutar servidor
	beego.Run()
}

// loadProjectEnv permite ejecutar este CRUD desde su carpeta y reutilizar el .env de la raíz.
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
