package models

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// CategoriaGanado es un tipo de ganado, por ejemplo "Novillo gordo" o "Vaca de cria"
// (tabla agropecuario.categorias_ganado)
type CategoriaGanado struct {
	Id                int       `orm:"column(id);pk;auto" json:"id"`
	Nombre            string    `orm:"column(nombre)" json:"nombre"`
	Descripcion       string    `orm:"column(descripcion);null" json:"descripcion,omitempty"`
	Activo            bool      `orm:"column(activo)" json:"activo"`
	FechaCreacion     time.Time `orm:"column(fecha_creacion);type(datetime);auto_now_add" json:"fecha_creacion"`
	FechaModificacion time.Time `orm:"column(fecha_modificacion);type(datetime);auto_now" json:"fecha_modificacion"`
}

// nombre de la tabla (el esquema agropecuario sale del search_path en app.conf)
func (t *CategoriaGanado) TableName() string {
	return "categorias_ganado"
}

func init() {
	orm.RegisterModel(new(CategoriaGanado))
}

// AddCategoriaGanado inserta una categoria nueva
func AddCategoriaGanado(m *CategoriaGanado) (id int64, err error) {
	m.Nombre = strings.TrimSpace(m.Nombre)
	if m.Nombre == "" {
		return 0, errors.New("falta el nombre")
	}

	o := orm.NewOrm()

	// en la base no puede haber dos categorias con el mismo nombre
	existente := CategoriaGanado{}
	err = o.QueryTable(new(CategoriaGanado)).Filter("nombre", m.Nombre).One(&existente)

	if err == nil {
		if existente.Activo {
			return 0, errors.New("ya existe una categoria con ese nombre")
		}
		// estaba eliminada: se vuelve a activar con los datos nuevos
		m.Id = existente.Id
		m.Activo = true
		m.FechaCreacion = existente.FechaCreacion
		_, err = o.Update(m, "Descripcion", "Activo", "FechaModificacion")
		return int64(m.Id), err
	}
	if err != orm.ErrNoRows {
		return 0, err
	}

	m.Id = 0
	m.Activo = true
	id, err = o.Insert(m)
	return id, err
}

// GetCategoriaGanadoById busca una categoria por su id (si esta eliminada no la devuelve)
func GetCategoriaGanadoById(id int) (v *CategoriaGanado, err error) {
	o := orm.NewOrm()
	v = &CategoriaGanado{Id: id}
	if err = o.Read(v); err != nil {
		return nil, err
	}
	if !v.Activo {
		return nil, errors.New("la categoria esta eliminada")
	}
	return v, nil
}

// GetAllCategoriaGanado trae la lista de categorias activas
func GetAllCategoriaGanado(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {

	o := orm.NewOrm()
	qs := o.QueryTable(new(CategoriaGanado)).Filter("activo", true)

	// filtros: query=nombre:Novillo
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}

	// orden: sortby=nombre&order=asc
	if len(sortby) > 0 {
		var campos []string
		for i, campo := range sortby {
			direccion := "asc"
			if len(order) == 1 {
				direccion = order[0]
			} else if i < len(order) {
				direccion = order[i]
			}
			if direccion == "desc" {
				campo = "-" + campo
			} else if direccion != "asc" {
				return nil, errors.New("order debe ser asc o desc")
			}
			campos = append(campos, campo)
		}
		qs = qs.OrderBy(campos...)
	} else {
		qs = qs.OrderBy("id")
	}

	var lista []CategoriaGanado
	_, err = qs.Limit(limit, offset).All(&lista, fields...)
	if err != nil {
		return nil, err
	}

	for _, cat := range lista {
		if len(fields) == 0 {
			ml = append(ml, cat)
			continue
		}
		// si pidieron solo algunos campos (fields=Id,Nombre)
		fila := make(map[string]interface{})
		valor := reflect.ValueOf(cat)
		for _, campo := range fields {
			f := valor.FieldByName(campo)
			if !f.IsValid() {
				return nil, errors.New("el campo " + campo + " no existe")
			}
			fila[campo] = f.Interface()
		}
		ml = append(ml, fila)
	}
	return ml, nil
}

// UpdateCategoriaGanadoById modifica una categoria
func UpdateCategoriaGanadoById(m *CategoriaGanado) (err error) {
	o := orm.NewOrm()
	v := CategoriaGanado{Id: m.Id}
	if err = o.Read(&v); err != nil {
		return err
	}
	if !v.Activo {
		return errors.New("la categoria esta eliminada")
	}

	// el nombre es obligatorio: si no lo mandaron se queda como estaba
	m.Nombre = strings.TrimSpace(m.Nombre)
	if m.Nombre == "" {
		m.Nombre = v.Nombre
	}
	m.Activo = true
	m.FechaCreacion = v.FechaCreacion

	_, err = o.Update(m, "Nombre", "Descripcion", "FechaModificacion")
	return err
}

// DeleteCategoriaGanado no borra la fila, solo la desactiva (activo = false).
// Asi los precios que ya usan esta categoria no se danan.
func DeleteCategoriaGanado(id int) (err error) {
	o := orm.NewOrm()
	v := CategoriaGanado{Id: id}
	if err = o.Read(&v); err != nil {
		return err
	}
	if !v.Activo {
		return errors.New("la categoria ya estaba eliminada")
	}

	v.Activo = false
	_, err = o.Update(&v, "Activo", "FechaModificacion")
	return err
}