package models

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// IndicadorSector es un dato que se mide en el sector, por ejemplo "Precio promedio de leche"
// (tabla agropecuario.indicadores_sector)
type IndicadorSector struct {
	Id                int       `orm:"column(id);pk;auto" json:"id"`
	Nombre            string    `orm:"column(nombre)" json:"nombre"`
	Descripcion       string    `orm:"column(descripcion);null" json:"descripcion"`
	Unidad            string    `orm:"column(unidad)" json:"unidad"`
	Fuente            string    `orm:"column(fuente);null" json:"fuente"`
	Activo            bool      `orm:"column(activo)" json:"activo"`
	FechaCreacion     time.Time `orm:"column(fecha_creacion);type(datetime);auto_now_add" json:"fecha_creacion"`
	FechaModificacion time.Time `orm:"column(fecha_modificacion);type(datetime);auto_now" json:"fecha_modificacion"`
}

// nombre de la tabla (el esquema agropecuario sale del search_path en app.conf)
func (t *IndicadorSector) TableName() string {
	return "indicadores_sector"
}

func init() {
	orm.RegisterModel(new(IndicadorSector))
}

// AddIndicadorSector inserta un indicador nuevo
func AddIndicadorSector(m *IndicadorSector) (id int64, err error) {
	m.Nombre = strings.TrimSpace(m.Nombre)
	if m.Nombre == "" {
		return 0, errors.New("falta el nombre")
	}
	if strings.TrimSpace(m.Unidad) == "" {
		return 0, errors.New("falta la unidad")
	}

	o := orm.NewOrm()

	// en la base no puede haber dos indicadores con el mismo nombre
	existente := IndicadorSector{}
	err = o.QueryTable(new(IndicadorSector)).Filter("nombre", m.Nombre).One(&existente)

	if err == nil {
		if existente.Activo {
			return 0, errors.New("ya existe un indicador con ese nombre")
		}
		// estaba eliminado: se vuelve a activar con los datos nuevos
		m.Id = existente.Id
		m.Activo = true
		m.FechaCreacion = existente.FechaCreacion
		_, err = o.Update(m, "Descripcion", "Unidad", "Fuente", "Activo", "FechaModificacion")
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

// GetIndicadorSectorById busca un indicador por su id (si esta eliminado no lo devuelve)
func GetIndicadorSectorById(id int) (v *IndicadorSector, err error) {
	o := orm.NewOrm()
	v = &IndicadorSector{Id: id}
	if err = o.Read(v); err != nil {
		return nil, err
	}
	if !v.Activo {
		return nil, errors.New("el indicador esta eliminado")
	}
	return v, nil
}

// GetAllIndicadorSector trae la lista de indicadores activos
func GetAllIndicadorSector(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {

	o := orm.NewOrm()
	qs := o.QueryTable(new(IndicadorSector)).Filter("activo", true)

	// filtros: query=unidad:COP/litro
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

	var lista []IndicadorSector
	_, err = qs.Limit(limit, offset).All(&lista, fields...)
	if err != nil {
		return nil, err
	}

	for _, ind := range lista {
		if len(fields) == 0 {
			ml = append(ml, ind)
			continue
		}
		// si pidieron solo algunos campos (fields=Id,Nombre)
		fila := make(map[string]interface{})
		valor := reflect.ValueOf(ind)
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

// UpdateIndicadorSectorById modifica un indicador
func UpdateIndicadorSectorById(m *IndicadorSector) (err error) {
	o := orm.NewOrm()
	v := IndicadorSector{Id: m.Id}
	if err = o.Read(&v); err != nil {
		return err
	}
	if !v.Activo {
		return errors.New("el indicador esta eliminado")
	}

	// el nombre y la unidad son obligatorios: si no los mandaron se quedan como estaban
	m.Nombre = strings.TrimSpace(m.Nombre)
	if m.Nombre == "" {
		m.Nombre = v.Nombre
	}
	if strings.TrimSpace(m.Unidad) == "" {
		m.Unidad = v.Unidad
	}
	m.Activo = true
	m.FechaCreacion = v.FechaCreacion

	_, err = o.Update(m, "Nombre", "Descripcion", "Unidad", "Fuente", "FechaModificacion")
	return err
}

// DeleteIndicadorSector no borra la fila, solo la desactiva (activo = false).
// Tambien desactiva las mediciones de ese indicador para que no queden sueltas.
func DeleteIndicadorSector(id int) (err error) {
	o := orm.NewOrm()
	v := IndicadorSector{Id: id}
	if err = o.Read(&v); err != nil {
		return err
	}
	if !v.Activo {
		return errors.New("el indicador ya estaba eliminado")
	}

	v.Activo = false
	_, err = o.Update(&v, "Activo", "FechaModificacion")
	if err != nil {
		return err
	}

	_, err = o.QueryTable(new(MedicionSector)).
		Filter("indicador_id", id).
		Update(orm.Params{"activo": false})
	return err
}
