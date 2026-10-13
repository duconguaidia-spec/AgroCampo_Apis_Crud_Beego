package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Reporteactividad struct {
	Id                   int        `orm:"column(id);pk;auto" json:"id"`
	TipoReporte          string     `orm:"column(tipo)" json:"tipo"`
	Formato              string     `orm:"column(formato)" json:"formato"`
	Filtros              string     `orm:"column(filtros);type(jsonb)" json:"filtros"`
	IdUsuarioSolicitante *int       `orm:"column(solicitado_por);null" json:"solicitado_por,omitempty"`
	Resultado            string     `orm:"column(estado)" json:"estado"`
	ArchivoURL           string     `orm:"column(archivo_url);null" json:"archivo_url,omitempty"`
	Error                string     `orm:"column(error);null" json:"error,omitempty"`
	Activo               bool       `orm:"column(activo)" json:"activo"`
	FechaGeneracion      time.Time  `orm:"column(fecha_creacion);type(timestamp with time zone);auto_now_add" json:"fecha_creacion"`
	FechaModificacion    time.Time  `orm:"column(fecha_modificacion);type(timestamp with time zone);auto_now" json:"fecha_modificacion"`
	FechaCompletado      *time.Time `orm:"column(completado_en);type(timestamp with time zone);null" json:"completado_en,omitempty"`
}

func (t *Reporteactividad) TableName() string {
	return "solicitudes_reporte"
}

func init() {
	orm.RegisterModel(new(Reporteactividad))
}

// AddReporteactividad insert a new Reporteactividad into database and returns
// last inserted Id on success.
func AddReporteactividad(m *Reporteactividad) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetReporteactividadById retrieves Reporteactividad by Id. Returns error if
// Id doesn't exist
func GetReporteactividadById(id int) (v *Reporteactividad, err error) {
	o := orm.NewOrm()
	v = &Reporteactividad{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllReporteactividad retrieves all Reporteactividad matches certain condition. Returns empty list if
// no records exist
func GetAllReporteactividad(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Reporteactividad))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Reporteactividad
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateReporteactividad updates Reporteactividad by Id and returns error if
// the record to be updated doesn't exist
func UpdateReporteactividadById(m *Reporteactividad) (err error) {
	o := orm.NewOrm()
	v := Reporteactividad{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteReporteactividad deletes Reporteactividad by Id and returns error if
// the record to be deleted doesn't exist
func DeleteReporteactividad(id int) (err error) {
	o := orm.NewOrm()
	v := Reporteactividad{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Reporteactividad{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
