package models

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// MedicionSector es el valor de un indicador en una fecha y una region
// por ejemplo: Precio promedio de leche, 2026-10-01, Cundinamarca, 2100
// (tabla agropecuario.mediciones_sector)
type MedicionSector struct {
	Id                int       `orm:"column(id);pk;auto" json:"id"`
	IndicadorId       int       `orm:"column(indicador_id)" json:"indicador_id"`
	Fecha             time.Time `orm:"column(fecha);type(date)" json:"fecha"`
	Region            string    `orm:"column(region)" json:"region"`
	Valor             float64   `orm:"column(valor)" json:"valor"`
	Fuente            string    `orm:"column(fuente);null" json:"fuente"`
	Activo            bool      `orm:"column(activo)" json:"activo"`
	FechaCreacion     time.Time `orm:"column(fecha_creacion);type(datetime);auto_now_add" json:"fecha_creacion"`
	FechaModificacion time.Time `orm:"column(fecha_modificacion);type(datetime);auto_now" json:"fecha_modificacion"`
}

func (t *MedicionSector) TableName() string {
	return "mediciones_sector"
}

func init() {
	orm.RegisterModel(new(MedicionSector))
}

// AddMedicionSector inserta una medicion nueva
func AddMedicionSector(m *MedicionSector) (id int64, err error) {
	if m.IndicadorId == 0 {
		return 0, errors.New("falta indicador_id")
	}
	if m.Fecha.IsZero() {
		return 0, errors.New("falta la fecha")
	}
	m.Region = strings.TrimSpace(m.Region)
	if m.Region == "" {
		m.Region = "Colombia"
	}

	o := orm.NewOrm()

	// el indicador tiene que existir y no estar eliminado
	indicador := IndicadorSector{Id: m.IndicadorId}
	err = o.Read(&indicador)
	if err == orm.ErrNoRows {
		return 0, errors.New("el indicador no existe")
	}
	if err != nil {
		return 0, err
	}
	if !indicador.Activo {
		return 0, errors.New("el indicador esta eliminado")
	}

	// en la base solo puede haber una medicion por indicador, fecha y region
	existente := MedicionSector{}
	err = o.QueryTable(new(MedicionSector)).
		Filter("indicador_id", m.IndicadorId).
		Filter("fecha", m.Fecha).
		Filter("region", m.Region).
		One(&existente)

	if err == nil {
		if existente.Activo {
			return 0, errors.New("ya existe una medicion de ese indicador en esa fecha y region")
		}
		// estaba eliminada: se vuelve a activar con el valor nuevo
		m.Id = existente.Id
		m.Activo = true
		m.FechaCreacion = existente.FechaCreacion
		_, err = o.Update(m, "Valor", "Fuente", "Activo", "FechaModificacion")
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

// GetMedicionSectorById busca una medicion por su id
func GetMedicionSectorById(id int) (v *MedicionSector, err error) {
	o := orm.NewOrm()
	v = &MedicionSector{Id: id}
	if err = o.Read(v); err != nil {
		return nil, err
	}
	if !v.Activo {
		return nil, errors.New("la medicion esta eliminada")
	}
	return v, nil
}

// GetAllMedicionSector trae la lista de mediciones activas
func GetAllMedicionSector(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {

	o := orm.NewOrm()
	qs := o.QueryTable(new(MedicionSector)).Filter("activo", true)

	// filtros: query=indicador_id:1,region:Cundinamarca
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}

	// orden: sortby=fecha&order=desc
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

	var lista []MedicionSector
	_, err = qs.Limit(limit, offset).All(&lista, fields...)
	if err != nil {
		return nil, err
	}

	for _, med := range lista {
		if len(fields) == 0 {
			ml = append(ml, med)
			continue
		}
		fila := make(map[string]interface{})
		valor := reflect.ValueOf(med)
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

// UpdateMedicionSectorById modifica el valor y la fuente de una medicion
// (el indicador, la fecha y la region no se cambian: si te equivocaste, elimina y crea otra)
func UpdateMedicionSectorById(m *MedicionSector) (err error) {
	o := orm.NewOrm()
	v := MedicionSector{Id: m.Id}
	if err = o.Read(&v); err != nil {
		return err
	}
	if !v.Activo {
		return errors.New("la medicion esta eliminada")
	}

	m.IndicadorId = v.IndicadorId
	m.Fecha = v.Fecha
	m.Region = v.Region
	m.Activo = true
	m.FechaCreacion = v.FechaCreacion

	_, err = o.Update(m, "Valor", "Fuente", "FechaModificacion")
	return err
}

// DeleteMedicionSector no borra la fila, solo la desactiva
func DeleteMedicionSector(id int) (err error) {
	o := orm.NewOrm()
	v := MedicionSector{Id: id}
	if err = o.Read(&v); err != nil {
		return err
	}
	if !v.Activo {
		return errors.New("la medicion ya estaba eliminada")
	}

	v.Activo = false
	_, err = o.Update(&v, "Activo", "FechaModificacion")
	return err
}
