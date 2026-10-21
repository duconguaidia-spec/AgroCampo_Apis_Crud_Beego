package models

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// Subasta es una subasta de ganado (tabla agropecuario.subastas)
type Subasta struct {
	Id                int       `orm:"column(id);pk;auto" json:"id"`
	RegistradaPor     int       `orm:"column(registrada_por);null" json:"registrada_por"`
	FechaSubasta      time.Time `orm:"column(fecha_subasta);type(date)" json:"fecha_subasta"`
	Ubicacion         string    `orm:"column(ubicacion);null" json:"ubicacion"`
	Nombre            string    `orm:"column(nombre);null" json:"nombre"`
	Activo            bool      `orm:"column(activo)" json:"activo"`
	FechaCreacion     time.Time `orm:"column(fecha_creacion);type(datetime);auto_now_add" json:"fecha_creacion"`
	FechaModificacion time.Time `orm:"column(fecha_modificacion);type(datetime);auto_now" json:"fecha_modificacion"`
}

// nombre de la tabla (el esquema agropecuario sale del search_path en app.conf)
func (t *Subasta) TableName() string {
	return "agropecuario\".\"subastas"
}

func init() {
	orm.RegisterModel(new(Subasta))
}

// AddSubasta inserta una subasta nueva
func AddSubasta(m *Subasta) (id int64, err error) {
	if m.RegistradaPor == 0 {
		return 0, errors.New("falta registrada_por (id de un usuario que exista)")
	}
	if m.FechaSubasta.IsZero() {
		return 0, errors.New("falta fecha_subasta")
	}

	m.Id = 0
	m.Activo = true

	o := orm.NewOrm()
	id, err = o.Insert(m)
	return id, err
}

// GetSubastaById busca una subasta por su id (si esta eliminada no la devuelve)
func GetSubastaById(id int) (v *Subasta, err error) {
	o := orm.NewOrm()
	v = &Subasta{Id: id}
	if err = o.Read(v); err != nil {
		return nil, err
	}
	if !v.Activo {
		return nil, errors.New("la subasta esta eliminada")
	}
	return v, nil
}

// GetAllSubasta trae la lista de subastas activas
func GetAllSubasta(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {

	o := orm.NewOrm()
	qs := o.QueryTable(new(Subasta)).Filter("activo", true)

	// filtros: query=nombre:Feria,ubicacion:Bogota
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}

	// orden: sortby=fecha_subasta&order=desc
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

	var lista []Subasta
	_, err = qs.Limit(limit, offset).All(&lista, fields...)
	if err != nil {
		return nil, err
	}

	for _, s := range lista {
		if len(fields) == 0 {
			ml = append(ml, s)
			continue
		}
		// si pidieron solo algunos campos (fields=Id,Nombre)
		fila := make(map[string]interface{})
		valor := reflect.ValueOf(s)
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

// UpdateSubastaById modifica una subasta
func UpdateSubastaById(m *Subasta) (err error) {
	o := orm.NewOrm()
	v := Subasta{Id: m.Id}
	if err = o.Read(&v); err != nil {
		return err
	}
	if !v.Activo {
		return errors.New("la subasta esta eliminada")
	}

	// lo que no se envio se queda como estaba
	if m.RegistradaPor == 0 {
		m.RegistradaPor = v.RegistradaPor
	}
	if m.FechaSubasta.IsZero() {
		m.FechaSubasta = v.FechaSubasta
	}
	m.Activo = true
	m.FechaCreacion = v.FechaCreacion

	_, err = o.Update(m, "RegistradaPor", "FechaSubasta", "Ubicacion", "Nombre", "FechaModificacion")
	return err
}

// DeleteSubasta no borra la fila, solo la desactiva (activo = false)
func DeleteSubasta(id int) (err error) {
	o := orm.NewOrm()
	v := Subasta{Id: id}
	if err = o.Read(&v); err != nil {
		return err
	}
	if !v.Activo {
		return errors.New("la subasta ya estaba eliminada")
	}

	v.Activo = false
	_, err = o.Update(&v, "Activo", "FechaModificacion")
	return err
}