package pkg

import "text/template"

// Make sure that the template compiles during package initialization
func parseTemplateOrPanic(t string) *template.Template {
	tpl, err := template.New("output_template").Parse(t)
	if err != nil {
		panic(err)
	}
	return tpl
}

var outputTemplate = parseTemplateOrPanic(`
///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gorm_gen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package {{.PkgName}}

import (
    "fmt"
    "time"

    "github.com/pkg/errors"
    "gorm.io/gorm"
    "gorm.io/gorm/schema"

    "github.com/haierkeys/golang-image-upload-service/global"
    "github.com/haierkeys/golang-image-upload-service/internal/model"
    "github.com/haierkeys/golang-image-upload-service/pkg/timef"
)

func Connection() *gorm.DB {
	db_driver := global.DBEngine
	db_driver.Config.NamingStrategy = schema.NamingStrategy{
		TablePrefix:   "{{.Prefix}}",   // 表名前缀
		SingularTable: true, // 使用单数表名
	}
	return db_driver
}


func NewModel() *{{.StructName}} {
	return new({{.StructName}})
}

type {{.QueryBuilderName}} struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	whereRaw []struct {
		query  string
		values []interface{}
	}
	limit  int
	offset int
}

func NewQueryBuilder() *{{.QueryBuilderName}} {
	return new({{.QueryBuilderName}})
}



func (qb *{{.QueryBuilderName}}) buildQuery() *gorm.DB {
	ret := Connection()
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, where2 := range qb.whereRaw {
		ret = ret.Where(where2.query, where2.values...)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}


func (t *{{.StructName}}) Create() (id int64, err error) {
	t.CreatedAt = timef.Now()
	db_driver := Connection()
	if err = db_driver.Model(t).Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.{{.PrimaryIdName}}, nil
}

func (t *{{.StructName}}) Save() (err error) {
	t.UpdatedAt = timef.Now()

	db_driver := Connection()
	if err = db_driver.Model(t).Save(t).Error; err != nil {
		return errors.Wrap(err, "update err")
	}
	return nil
}


func (qb *{{.QueryBuilderName}}) Updates( m map[string]interface{}) (err error) {

	db_driver := Connection()
	db_driver = db_driver.Model(&{{.StructName}}{})

	for _, where := range qb.where {
		db_driver.Where(where.prefix, where.value)
	}

	if err = db_driver.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}


//自减
func (qb *{{.QueryBuilderName}}) Increment(column string, value int64) (err error) {

	db_driver := Connection()
	db_driver = db_driver.Model(&{{.StructName}}{})

	for _, where := range qb.where {
		db_driver.Where(where.prefix, where.value)
	}

	if err = db_driver.Update(column, gorm.Expr(column+" + ?", value)).Error; err != nil {
		return errors.Wrap(err, "increment err")
	}
	return nil
}

//自增
func (qb *{{.QueryBuilderName}}) Decrement(column string, value int64) (err error) {

	db_driver := Connection()
	db_driver = db_driver.Model(&{{.StructName}}{})

	for _, where := range qb.where {
		db_driver.Where(where.prefix, where.value)
	}

	if err = db_driver.Update(column, gorm.Expr(column+" - ?", value)).Error; err != nil {
		return errors.Wrap(err, "decrement err")
	}
	return nil
}

func (qb *{{.QueryBuilderName}}) Delete() (err error) {

	db_driver := Connection()
	for _, where := range qb.where {
		db_driver = db_driver.Where(where.prefix, where.value)
	}

	if err = db_driver.Delete(&{{.StructName}}{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *{{.QueryBuilderName}}) Count() (int64, error) {
	var c int64
	res := qb.buildQuery().Model(&{{.StructName}}{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *{{.QueryBuilderName}}) First() (*{{.StructName}}, error) {
	ret := &{{.StructName}}{}
	res := qb.buildQuery().First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *{{.QueryBuilderName}}) Get() ([]*{{.StructName}}, error) {
	return qb.QueryAll()
}

func (qb *{{.QueryBuilderName}}) QueryOne() (*{{.StructName}}, error) {
	qb.limit = 1
	ret, err := qb.QueryAll()
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *{{.QueryBuilderName}}) QueryAll() ([]*{{.StructName}}, error) {
	var ret []*{{.StructName}}
	err := qb.buildQuery().Find(&ret).Error
	return ret, err
}

func (qb *{{.QueryBuilderName}}) Limit(limit int) *{{.QueryBuilderName}} {
	qb.limit = limit
	return qb
}

func (qb *{{.QueryBuilderName}}) Offset(offset int) *{{.QueryBuilderName}} {
	qb.offset = offset
	return qb
}

func (qb *{{.QueryBuilderName}}) WhereRaw(query string, values ...interface{})  *{{.QueryBuilderName}} {
	vals := make([]interface{}, len(values))
	for i, v := range values {
		vals[i] = v
	}
	qb.whereRaw = append(qb.whereRaw, struct {
		query  string
		values []interface{}
	}{
		query,
		vals,
	})
	return qb
}

// ----------



{{$queryBuilderName := .QueryBuilderName}}
{{range .OptionFields}}
func (qb *{{$queryBuilderName}}) Where{{call $.Helpers.Titelize .FieldName}}(p model.Predicate, value {{.FieldType}}) *{{$queryBuilderName}} {
	 qb.where = append(qb.where, struct {
		prefix string
		value interface{}
	}{
		fmt.Sprintf("%v %v ?", "{{.ColumnName}}", p),
		value,
	})
	return qb
}

func (qb *{{$queryBuilderName}}) Where{{call $.Helpers.Titelize .FieldName}}In(value []{{.FieldType}}) *{{$queryBuilderName}} {
	 qb.where = append(qb.where, struct {
		prefix string
		value interface{}
	}{
		fmt.Sprintf("%v %v ?", "{{.ColumnName}}", "IN"),
		value,
	})
	return qb
}

func (qb *{{$queryBuilderName}}) Where{{call $.Helpers.Titelize .FieldName}}NotIn(value []{{.FieldType}}) *{{$queryBuilderName}} {
	 qb.where = append(qb.where, struct {
		prefix string
		value interface{}
	}{
		fmt.Sprintf("%v %v ?", "{{.ColumnName}}", "NOT IN"),
		value,
	})
	return qb
}

func (qb *{{$queryBuilderName}}) OrderBy{{call $.Helpers.Titelize .FieldName}}(asc bool) *{{$queryBuilderName}} {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "` + "`{{.ColumnName}}`" + ` " + order)
	return qb
}
{{end}}
`)
