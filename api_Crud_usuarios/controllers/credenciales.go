package controllers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
)

var (
	contrasenaEnUsuario bool
	consultaRegistro    string
	consultaLogin       string
)

func ConfigurarAlmacenContrasenas() error {
	db, err := orm.GetDB()
	if err != nil {
		return err
	}
	err = db.QueryRow(`
		SELECT EXISTS (
			SELECT 1 FROM information_schema.columns
			WHERE table_schema = 'usuarios' AND table_name = 'usuarios' AND column_name = 'password_hash'
		)
	`).Scan(&contrasenaEnUsuario)
	if err != nil {
		return fmt.Errorf("consultar esquema de credenciales de Usuarios: %w", err)
	}
	if contrasenaEnUsuario {
		consultaRegistro = `
			INSERT INTO usuarios.usuarios (nombre, apellido, correo, telefono, password_hash, rol_id)
			SELECT $1, $2, $3, $4, $5, id FROM usuarios.roles WHERE nombre = $6 AND activo
			ON CONFLICT (correo) DO NOTHING RETURNING id`
		consultaLogin = `
			SELECT u.id, u.nombre, u.apellido, r.nombre, u.password_hash
			FROM usuarios.usuarios u JOIN usuarios.roles r ON r.id = u.rol_id
			WHERE u.correo = $1 AND u.activo AND r.activo`
		return nil
	}

	consultaRegistro = `
		WITH nuevo AS (
			INSERT INTO usuarios.usuarios (nombre, apellido, correo, telefono, rol_id)
			SELECT $1, $2, $3, $4, id FROM usuarios.roles WHERE nombre = $6 AND activo
			ON CONFLICT (correo) DO NOTHING RETURNING id
		)
		INSERT INTO usuarios."Contrasena" (id_usuario, contrasena_hash)
		SELECT id, $5 FROM nuevo RETURNING id_usuario`
	consultaLogin = `
		SELECT u.id, u.nombre, u.apellido, r.nombre, c.contrasena_hash
		FROM usuarios.usuarios u JOIN usuarios.roles r ON r.id = u.rol_id
		JOIN LATERAL (
			SELECT contrasena_hash FROM usuarios."Contrasena"
			WHERE id_usuario = u.id AND activa
			ORDER BY id_contrasena DESC LIMIT 1
		) c ON true
		WHERE u.correo = $1 AND u.activo AND r.activo`
	return nil
}

func restablecerContrasena(ctx context.Context, db *sql.DB, tokenHash, hash string) (int, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var id int
	err = tx.QueryRowContext(ctx, `
		UPDATE usuarios.tokens_recuperacion SET usado_en = now()
		WHERE token_hash = $1 AND usado_en IS NULL AND expira_en > now()
		RETURNING usuario_id`, tokenHash).Scan(&id)
	if err != nil {
		return 0, err
	}
	if contrasenaEnUsuario {
		err = tx.QueryRowContext(ctx, `
			UPDATE usuarios.usuarios SET password_hash = $2
			WHERE id = $1 AND activo RETURNING id`, id, hash).Scan(&id)
	} else {
		// El bloqueo serializa los restablecimientos del mismo usuario.
		err = tx.QueryRowContext(ctx, `
			SELECT id FROM usuarios.usuarios WHERE id = $1 AND activo FOR UPDATE`, id).Scan(&id)
		if err == nil {
			_, err = tx.ExecContext(ctx, `UPDATE usuarios."Contrasena"
				SET activa = false, fecha_modificacion = now() WHERE id_usuario = $1 AND activa`, id)
		}
		if err == nil {
			_, err = tx.ExecContext(ctx, `INSERT INTO usuarios."Contrasena" (id_usuario, contrasena_hash)
				VALUES ($1, $2)`, id, hash)
		}
	}
	if err != nil {
		return 0, err
	}
	return id, tx.Commit()
}
