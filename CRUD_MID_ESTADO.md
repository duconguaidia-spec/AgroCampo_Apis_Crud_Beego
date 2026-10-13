# Estado de las CRUD de AgroCampo para el MID

## Conexión a la base

Las siete aplicaciones CRUD apuntan a la base `AGROCAMPO`. Cada `conf/app.conf` define `PGsschema`, que es el esquema principal del servicio. La CRUD de usuarios usa `usuarios,reportes` para leer/escribir auditoría.

| Servicio | Esquema(s) |
|---|---|
| `api_Crud_agropecuario` | `agropecuario` |
| `api_Crud_Contenido` | `contenido` |
| `api_Crud_Noticias` | `noticias` |
| `api_Crud_Reportes` | `reportes` |
| `api_crud_Soporte` | `soporte` |
| `api_Crud_usuarios` | `usuarios,reportes` |
| `api_Crud_Veterinarias` | `veterinarias` |

El `main.go` de cada servicio lee `PGsschema`. En `.env` se deben conservar las variables de conexión existentes (`PGuser`, `PGpass`, `PGhost`, `PGport`) para cada módulo. No se debe imprimir la URL de conexión porque contiene credenciales.

Las siete aplicaciones aceptan el encabezado CORS `Authorization`, necesario para que el MID envíe el token guardado por el frontend. CORS no permite credenciales de cookies junto con origen `*`.

## Recursos existentes alineados

| CRUD | Tablas a las que apuntan los modelos actuales |
|---|---|
| Agropecuario | `categorias_ganado`, `subastas`, `precios_subasta_ganado` |
| Contenido | `categorias_contenido`, `temas_foro`, `respuestas_foro`, `videos_educativos` |
| Noticias | `noticias` |
| Reportes | `solicitudes_reporte` |
| Soporte | `categorias_faq`, `preguntas_frecuentes` |
| Usuarios | `usuarios`, `roles`, `perfiles`, `actividad_usuario` |
| Veterinarias | `veterinarias`, `especialidades_veterinarias`, `veterinaria_especialidades`, `servicios_generales`, `veterinaria_servicios`, `servicios_veterinaria`, `productos_veterinaria`, `resenas_veterinaria` |

## Entidades sin endpoint CRUD todavía

Estas tablas existen en el SQL pero aún no tienen modelo/controlador/ruta completos en los módulos actuales: `noticias.categorias_noticia`, `contenido.temas_guardados`, `agropecuario.indicadores_sector`, `agropecuario.mediciones_sector`, `soporte.tickets_soporte` y `reportes.actividad_usuario` en el servicio de reportes. Deben agregarse si el MID las necesita como recurso independiente.

## Operaciones reservadas al MID

Se retiraron del router CRUD de usuarios las rutas genéricas de contraseñas, tokens de recuperación y verificación de dos pasos. El MID debe implementar registro/inicio de sesión, hash y verificación de contraseña, emisión/validación de tokens, recuperación y 2FA. La columna `password_hash` se excluye de la serialización JSON del modelo de usuario.

El MID debe ser el único punto público. No expongas estas CRUD directamente al navegador: los controladores generados permiten operaciones amplias y no implementan autorización por rol/propietario. Configura el MID para validar identidad y permisos antes de llamar a cada CRUD. La auditoría también debe ser de solo lectura para usuarios normales y sus escrituras deben generarse desde el backend.

## Estado de verificación

La alineación de nombres/campos y configuración se hizo contra `API_MID_AGROCAMPO/database_v2/01_esquema.sql`. Las siete aplicaciones compilaron con `go build .`. El SQL no fue ejecutado contra PostgreSQL desde esta tarea; aún falta probar las peticiones reales en una base `AGROCAMPO` creada con ese archivo. La cobertura de las tablas que siguen sin endpoint está enumerada arriba, así que el proyecto aún no cubre todo el SQL nuevo.
