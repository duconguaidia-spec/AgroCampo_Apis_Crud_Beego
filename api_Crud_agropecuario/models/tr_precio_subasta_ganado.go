package models

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// TrPrecioSubastaGanado es el precio por kilo de una categoria de ganado
// dentro de una subasta (tabla agropecuario.precios_subasta_ganado)
type TrPrecioSubastaGanado struct {
	Id                int       `orm:"column(id);pk;auto" json:"id"`
	SubastaId         int       `orm:"column(subasta_id)" json:"subasta_id"`
	CategoriaGanadoId int       `orm:"column(categoria_ganado_id)" json:"categoria_ganado_id"`
	PrecioPorKg       float64   `orm:"column(precio_por_kg)" json:"precio_por_kg"`
	Unidad            string    `orm:"column(unidad)" json:"unidad"`
	Activo            bool      `orm:"column(activo)" json:"activo"`
	FechaCreacion     time.Time `orm:"column(fecha_creacion);type(datetime);auto_now_add" json:"fecha_creacion"`
	FechaModificacion time.Time `orm:"column(fecha_modificacion);type(datetime);auto_now" json:"fecha_modificacion"`
}

func (t *TrPrecioSubastaGanado) TableName() string {
	return "agropecuario\".\"precios_subasta_ganado"
}

func init() {
	orm.RegisterModel(new(TrPrecioSubastaGanado))
}

// AddTrPrecioSubastaGanado inserta un precio nuevo
func AddTrPrecioSubastaGanado(m *TrPrecioSubastaGanado) (id int64, err error) {
	if m.SubastaId == 0 {
		return 0, errors.New("falta subasta_id")
	}
	if m.CategoriaGanadoId == 0 {
		return 0, errors.New("falta categoria_ganado_id")
	}
	if m.PrecioPorKg < 0 {
		return 0, errors.New("el precio no puede ser negativo")
	}
	if m.Unidad == "" {
		m.Unidad = "COP/kg"
	}

	o := orm.NewOrm()

	// en la base solo puede haber un precio por categoria en cada subasta
	existente := TrPrecioSubastaGanado{}
	err = o.QueryTable(new(TrPrecioSubastaGanado)).
		Filter("subasta_id", m.SubastaId).
		Filter("categoria_ganado_id", m.CategoriaGanadoId).
		One(&existente)

	if err == nil {
		if existente.Activo {
			return 0, errors.New("ya existe un precio para esa categoria en esa subasta")
		}
		// estaba eliminado: se vuelve a activar con el precio nuevo
		m.Id = existente.Id
		m.Activo = true
		m.FechaCreacion = existente.FechaCreacion
		_, err = o.Update(m, "PrecioPorKg", "Unidad", "Activo", "FechaModificacion")
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

// GetTrPrecioSubastaGanadoById busca un precio por su id
func GetTrPrecioSubastaGanadoById(id int) (v *TrPrecioSubastaGanado, err error) {
	o := orm.NewOrm()
	v = &TrPrecioSubastaGanado{Id: id}
	if err = o.Read(v); err != nil {
		return nil, err
	}
	if !v.Activo {
		return nil, errors.New("el precio esta eliminado")
	}
	return v, nil
}

// GetAllTrPrecioSubastaGanado trae la lista de precios activos
func GetAllTrPrecioSubastaGanado(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {

	o := orm.NewOrm()
	qs := o.QueryTable(new(TrPrecioSubastaGanado)).Filter("activo", true)

	// filtros: query=subasta_id:1
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}

	// orden: sortby=precio_por_kg&order=desc
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

	var lista []TrPrecioSubastaGanado
	_, err = qs.Limit(limit, offset).All(&lista, fields...)
	if err != nil {
		return nil, err
	}

	for _, p := range lista {
		if len(fields) == 0 {
			ml = append(ml, p)
			continue
		}
		fila := make(map[string]interface{})
		valor := reflect.ValueOf(p)
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

// UpdateTrPrecioSubastaGanadoById modifica el precio y la unidad
// (la subasta y la categoria no se cambian: si te equivocaste, elimina y crea otro)
func UpdateTrPrecioSubastaGanadoById(m *TrPrecioSubastaGanado) (err error) {
	o := orm.NewOrm()
	v := TrPrecioSubastaGanado{Id: m.Id}
	if err = o.Read(&v); err != nil {
		return err
	}
	if !v.Activo {
		return errors.New("el precio esta eliminado")
	}
	if m.PrecioPorKg < 0 {
		return errors.New("el precio no puede ser negativo")
	}

	m.SubastaId = v.SubastaId
	m.CategoriaGanadoId = v.CategoriaGanadoId
	m.Activo = true
	m.FechaCreacion = v.FechaCreacion
	if m.Unidad == "" {
		m.Unidad = v.Unidad
	}

	_, err = o.Update(m, "PrecioPorKg", "Unidad", "FechaModificacion")
	return err
}

// DeleteTrPrecioSubastaGanado no borra la fila, solo la desactiva
func DeleteTrPrecioSubastaGanado(id int) (err error) {
	o := orm.NewOrm()
	v := TrPrecioSubastaGanado{Id: id}
	if err = o.Read(&v); err != nil {
		return err
	}
	if !v.Activo {
		return errors.New("el precio ya estaba eliminado")
	}

	v.Activo = false
	_, err = o.Update(&v, "Activo", "FechaModificacion")
	return err
}