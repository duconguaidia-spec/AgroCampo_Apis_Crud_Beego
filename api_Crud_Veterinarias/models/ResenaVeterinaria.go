package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type ResenaVeterinaria struct {
	Id            int          `orm:"column(id_resena);pk"`
	IdVeterinaria *Veterinaria `orm:"column(id_veterinaria);rel(fk)"`
	IdUsuario     int          `orm:"column(id_usuario)"`
	Calificacion  int          `orm:"column(calificacion)"`
	Comentario    string       `orm:"column(comentario);null"`
	FechaCreacion time.Time    `orm:"column(fecha_creacion);type(timestamp without time zone)"`
}

func (t *ResenaVeterinaria) TableName() string {
	return "ResenaVeterinaria"
}

func init() {
	orm.RegisterModel(new(ResenaVeterinaria))
}

// AddResenaVeterinaria insert a new ResenaVeterinaria into database and returns
// last inserted Id on success.
func AddResenaVeterinaria(m *ResenaVeterinaria) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetResenaVeterinariaById retrieves ResenaVeterinaria by Id. Returns error if
// Id doesn't exist
func GetResenaVeterinariaById(id int) (v *ResenaVeterinaria, err error) {
	o := orm.NewOrm()
	v = &ResenaVeterinaria{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
		o.LoadRelated(v, "IdVeterinaria")
	}
	return nil, err
}

// GetAllResenaVeterinaria retrieves all ResenaVeterinaria matches certain condition. Returns empty list if
// no records exist
func GetAllResenaVeterinaria(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ResenaVeterinaria)).RelatedSel()
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

	var l []ResenaVeterinaria
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

// UpdateResenaVeterinaria updates ResenaVeterinaria by Id and returns error if
// the record to be updated doesn't exist
func UpdateResenaVeterinariaById(m *ResenaVeterinaria) (err error) {
	o := orm.NewOrm()
	v := ResenaVeterinaria{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteResenaVeterinaria deletes ResenaVeterinaria by Id and returns error if
// the record to be deleted doesn't exist
func DeleteResenaVeterinaria(id int) (err error) {
	o := orm.NewOrm()
	v := ResenaVeterinaria{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ResenaVeterinaria{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
