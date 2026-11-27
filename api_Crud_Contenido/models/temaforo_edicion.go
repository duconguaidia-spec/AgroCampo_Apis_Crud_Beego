package models

import (
	"context"

	"github.com/beego/beego/v2/client/orm"
)

type EdicionTemaForo struct {
	Titulo      string `json:"titulo"`
	Descripcion string `json:"descripcion"`
	CategoriaID int    `json:"categoria_id"`
}

func EditarTemaForo(ctx context.Context, id int, datos EdicionTemaForo) (EdicionTemaForo, error) {
	db, err := orm.GetDB()
	if err != nil {
		return EdicionTemaForo{}, err
	}
	var editado EdicionTemaForo
	err = db.QueryRowContext(ctx, `
		UPDATE contenido.temas_foro AS tema
		SET titulo = $2, descripcion = $3, categoria_id = categoria.id, fecha_modificacion = now()
		FROM contenido.categorias_contenido AS categoria
		WHERE tema.id = $1 AND tema.activo AND categoria.id = $4 AND categoria.activo AND categoria.tipo = 'foro'
		RETURNING tema.titulo, tema.descripcion, tema.categoria_id`,
		id, datos.Titulo, datos.Descripcion, datos.CategoriaID,
	).Scan(&editado.Titulo, &editado.Descripcion, &editado.CategoriaID)
	return editado, err
}
