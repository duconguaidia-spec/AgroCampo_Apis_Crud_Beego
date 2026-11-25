# AgroCampo_Apis_Crud_Beego
Este repositorio es para lograr manejo de Beego y Apis en nuestro proyecto

## Ejecutar un CRUD por separado

La configuración está centralizada en el archivo `.env` de la raíz. Desde la
carpeta del CRUD que quieras iniciar, ejecuta:

```powershell
go run .
```

Por ejemplo, para iniciar el CRUD de noticias:

```powershell
cd api_Crud_Noticias
go run .
```

Si Go muestra `cannot find package` al compilar Noticias dentro de OneDrive, ejecuta esto **desde `api_Crud_Noticias`**:

```powershell
.\Iniciar-Noticias.ps1
```

El script compila las fuentes en una carpeta temporal local y ejecuta Noticias desde el proyecto original para usar su configuración, Swagger y `.env`. Para verificar solo la compilación, agrega `-SoloCompilar`. Si editas el código, detén el servicio con `Ctrl+C` y vuelve a ejecutar el script.

Cada servicio toma su puerto y sus credenciales de PostgreSQL del `.env` raíz.
Si cambias el usuario, contraseña o dirección de PostgreSQL, actualiza las
variables `API_CRUD_*_PGUSER`, `API_CRUD_*_PGPASS` y `API_CRUD_*_PGHOST` allí.
