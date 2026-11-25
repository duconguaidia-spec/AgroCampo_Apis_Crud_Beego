package controllers

import (
	"crypto/subtle"
	"database/sql"
	"encoding/json"
	"errors"
	"net/mail"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
)

type CuentaController struct{ beego.Controller }

type cuentaDatos struct {
	ID         int      `json:"id"`
	Nombre     string   `json:"nombre"`
	Apellido   string   `json:"apellido"`
	Correo     string   `json:"correo"`
	Telefono   string   `json:"telefono"`
	Rol        string   `json:"rol"`
	Ubicacion  string   `json:"ubicacion"`
	Foto       string   `json:"foto"`
	Intereses  []string `json:"intereses"`
	Contrasena string   `json:"contrasena,omitempty"`
	Token      string   `json:"token,omitempty"`
}

func (c *CuentaController) resultado(status int, datos interface{}, mensaje string) {
	c.Ctx.Output.SetStatus(status)
	c.Data["json"] = map[string]interface{}{"success": status < 400, "data": datos, "message": mensaje}
	c.ServeJSON()
}

func (c *CuentaController) Operar() {
	clave := beego.AppConfig.DefaultString("clave_interna", "")
	if len(clave) < 32 || subtle.ConstantTimeCompare([]byte(clave), []byte(c.Ctx.Input.Header("X-Agrocampo-Interno"))) != 1 {
		c.resultado(403, nil, "El CRUD de usuarios requiere la clave interna de la MID")
		return
	}
	var d cuentaDatos
	if json.Unmarshal(c.Ctx.Input.RequestBody, &d) != nil {
		c.resultado(400, nil, "Datos inválidos")
		return
	}
	db, err := orm.GetDB()
	if err != nil {
		c.resultado(500, nil, "Base de usuarios no disponible")
		return
	}
	ctx := c.Ctx.Request.Context()
	d.Correo = strings.ToLower(strings.TrimSpace(d.Correo))
	switch c.Ctx.Input.Param(":accion") {
	case "contar-productores":
		var total int
		err := db.QueryRowContext(ctx, `SELECT count(*) FROM usuarios.usuarios u JOIN usuarios.roles r ON r.id=u.rol_id WHERE u.activo AND r.activo AND r.nombre='usuario'`).Scan(&total)
		if err != nil {
			c.resultado(500, nil, "No se pudo consultar el total de productores")
			return
		}
		c.resultado(200, total, "")
	case "registro":
		correo, err := mail.ParseAddress(d.Correo)
		if err != nil || correo.Address != d.Correo || strings.TrimSpace(d.Nombre) == "" || strings.TrimSpace(d.Apellido) == "" || len(d.Contrasena) < 8 || len(d.Contrasena) > 72 || (d.Rol != "usuario" && d.Rol != "veterinario") {
			c.resultado(400, nil, "Revisa nombre, correo, rol y contraseña")
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(d.Contrasena), bcrypt.DefaultCost)
		if err != nil {
			c.resultado(500, nil, "No se pudo preparar la contraseña")
			return
		}
		err = db.QueryRowContext(ctx, `INSERT INTO usuarios.usuarios (nombre,apellido,correo,telefono,password_hash,rol_id) SELECT $1,$2,$3,$4,$5,id FROM usuarios.roles WHERE nombre=$6 AND activo ON CONFLICT(correo) DO NOTHING RETURNING id`, d.Nombre, d.Apellido, d.Correo, d.Telefono, string(hash), d.Rol).Scan(&d.ID)
		if errors.Is(err, sql.ErrNoRows) {
			c.resultado(409, nil, "El correo ya está registrado o el rol no está disponible")
			return
		}
		if err != nil {
			c.resultado(500, nil, "No se pudo crear la cuenta")
			return
		}
		c.resultado(201, map[string]int{"id": d.ID}, "")
	case "login":
		var hash string
		err := db.QueryRowContext(ctx, `SELECT u.id,u.nombre,u.apellido,r.nombre,u.password_hash FROM usuarios.usuarios u JOIN usuarios.roles r ON r.id=u.rol_id WHERE u.correo=$1 AND u.activo AND r.activo`, d.Correo).Scan(&d.ID, &d.Nombre, &d.Apellido, &d.Rol, &hash)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			c.resultado(500, nil, "No se pudo consultar la cuenta")
			return
		}
		if err != nil || bcrypt.CompareHashAndPassword([]byte(hash), []byte(d.Contrasena)) != nil {
			c.resultado(401, nil, "Correo o contraseña incorrectos")
			return
		}
		d.Contrasena = ""
		c.resultado(200, d, "")
	case "perfil":
		var lista []byte
		err := db.QueryRowContext(ctx, `SELECT u.nombre,u.apellido,u.correo,coalesce(u.telefono,''),r.nombre,coalesce(p.ubicacion,''),coalesce(u.avatar_url,''),coalesce((SELECT json_agg(i.nombre ORDER BY i.nombre) FROM usuarios.usuario_intereses ui JOIN usuarios.intereses i ON i.id=ui.interes_id WHERE ui.usuario_id=u.id AND ui.activo AND i.activo),'[]'::json) FROM usuarios.usuarios u JOIN usuarios.roles r ON r.id=u.rol_id LEFT JOIN usuarios.perfiles p ON p.usuario_id=u.id WHERE u.id=$1 AND u.activo`, d.ID).Scan(&d.Nombre, &d.Apellido, &d.Correo, &d.Telefono, &d.Rol, &d.Ubicacion, &d.Foto, &lista)
		if errors.Is(err, sql.ErrNoRows) {
			c.resultado(404, nil, "Cuenta no encontrada")
			return
		}
		if err != nil || json.Unmarshal(lista, &d.Intereses) != nil {
			c.resultado(500, nil, "No se pudo consultar el perfil")
			return
		}
		c.resultado(200, d, "")
	case "guardar-perfil":
		if strings.TrimSpace(d.Nombre) == "" || strings.TrimSpace(d.Apellido) == "" || len(d.Intereses) > 30 {
			c.resultado(400, nil, "Revisa los datos del perfil")
			return
		}
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			c.resultado(500, nil, "No se pudo guardar el perfil")
			return
		}
		defer tx.Rollback()
		_, err = tx.ExecContext(ctx, `UPDATE usuarios.usuarios SET nombre=$2,apellido=$3,telefono=$4,avatar_url=$5 WHERE id=$1 AND activo`, d.ID, d.Nombre, d.Apellido, d.Telefono, d.Foto)
		if err == nil {
			_, err = tx.ExecContext(ctx, `INSERT INTO usuarios.perfiles(usuario_id,ubicacion) VALUES($1,$2) ON CONFLICT(usuario_id) DO UPDATE SET ubicacion=excluded.ubicacion`, d.ID, d.Ubicacion)
		}
		if err == nil {
			_, err = tx.ExecContext(ctx, `UPDATE usuarios.usuario_intereses SET activo=false WHERE usuario_id=$1`, d.ID)
		}
		for _, interes := range d.Intereses {
			if err != nil {
				break
			}
			_, err = tx.ExecContext(ctx, `WITH dato AS (INSERT INTO usuarios.intereses(nombre) VALUES($2) ON CONFLICT(nombre) DO UPDATE SET nombre=excluded.nombre RETURNING id) INSERT INTO usuarios.usuario_intereses(usuario_id,interes_id) SELECT $1,id FROM dato ON CONFLICT(usuario_id,interes_id) DO UPDATE SET activo=true`, d.ID, interes)
		}
		if err == nil {
			err = tx.Commit()
		}
		if err != nil {
			c.resultado(500, nil, "No se pudo guardar el perfil")
			return
		}
		c.resultado(200, nil, "")
	case "crear-token":
		res, err := db.ExecContext(ctx, `INSERT INTO usuarios.tokens_recuperacion(usuario_id,token_hash,expira_en) SELECT id,$2,now()+interval '15 minutes' FROM usuarios.usuarios WHERE correo=$1 AND activo`, d.Correo, d.Token)
		if err != nil {
			c.resultado(500, nil, "No se pudo preparar la recuperación")
			return
		}
		n, _ := res.RowsAffected()
		c.resultado(200, map[string]bool{"encontrada": n == 1}, "")
	case "restablecer":
		if len(d.Contrasena) < 8 || len(d.Contrasena) > 72 {
			c.resultado(400, nil, "La contraseña debe tener entre 8 y 72 bytes")
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(d.Contrasena), bcrypt.DefaultCost)
		if err != nil {
			c.resultado(500, nil, "No se pudo preparar la contraseña")
			return
		}
		err = db.QueryRowContext(ctx, `WITH token AS (UPDATE usuarios.tokens_recuperacion SET usado_en=now() WHERE token_hash=$1 AND usado_en IS NULL AND expira_en>now() RETURNING usuario_id) UPDATE usuarios.usuarios SET password_hash=$2 WHERE activo AND id=(SELECT usuario_id FROM token) RETURNING id`, d.Token, string(hash)).Scan(&d.ID)
		if errors.Is(err, sql.ErrNoRows) {
			c.resultado(400, nil, "Enlace inválido o expirado")
			return
		}
		if err != nil {
			c.resultado(500, nil, "No se pudo actualizar la contraseña")
			return
		}
		c.resultado(200, map[string]int{"id": d.ID}, "")
	default:
		c.resultado(404, nil, "Operación no disponible")
	}
}
