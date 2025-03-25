package main

// gorm gen configure

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/haierkeys/obsidian-image-api-gateway/pkg/fileurl"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	dbType string
	dbDsn  string
	step   int
)

func init() {

	dType := flag.String("type", "", "输入类型")
	dsn := flag.String("dsn", "", "输入DB dsn地址")
	dStep := flag.Int("step", 0, "输入执行步骤")

	flag.Parse()
	dbType = *dType
	dbDsn = *dsn
	step = *dStep
}

// SQLColumnToHumpStyle sql转换成驼峰模式
func SQLColumnToHumpStyle(in string) (ret string) {
	for i := 0; i < len(in); i++ {
		if i > 0 && in[i-1] == '_' && in[i] != '_' {
			s := strings.ToUpper(string(in[i]))
			ret += s
		} else if in[i] == '_' {
			continue
		} else {
			ret += string(in[i])
		}
	}
	return
}

func Db(dsn string, dbType string) *gorm.DB {

	db, err := gorm.Open(useDia(dsn, dbType), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	return db
}

func useDia(dsn string, dbType string) gorm.Dialector {
	if dbType == "mysql" {
		return mysql.Open(dsn)
	} else if dbType == "sqlite" {

		if !fileurl.IsExist(dsn) {
			fileurl.CreatePath(dsn, os.ModePerm)
		}
		return sqlite.Open(dsn)
	}
	return nil
}

func main() {

	g := gen.NewGenerator(gen.Config{
		// 默认会在 OutPath 目录生成CRUD代码，并且同目录下生成 model 包
		// 所以OutPath最终package不能设置为model，在有数据库表同步的情况下会产生冲突
		// 若一定要使用可以通过ModelPkgPath单独指定model package的名称
		OutPath: "./internal/query",
		/* ModelPkgPath: "dal/model"*/

		// gen.WithoutContext：禁用WithContext模式
		// gen.WithDefaultQuery：生成一个全局Query对象Q
		// gen.WithQueryInterface：生成Query接口
		Mode:              gen.WithQueryInterface,
		WithUnitTest:      false,
		FieldWithTypeTag:  false,
		FieldWithIndexTag: true,
	})

	db := Db(dbDsn, dbType)
	g.UseDB(db)

	var dataMap = map[string]func(gorm.ColumnType) (dataType string){
		// int mapping
		"integer": func(columnType gorm.ColumnType) (dataType string) {
			if n, ok := columnType.Nullable(); ok && n {
				return "int64"
			}
			return "int64"
		},
		"int": func(columnType gorm.ColumnType) (dataType string) {
			if n, ok := columnType.Nullable(); ok && n {
				return "int64"
			}
			return "int64"
		},
	}
	g.WithDataTypeMap(dataMap)

	opts := []gen.ModelOpt{
		//gen.FieldType("uid", "int64"),
		gen.FieldType("created_at", "timex.Time"),
		gen.FieldType("updated_at", "timex.Time"),
		gen.FieldType("deleted_at", "timex.Time"),
		gen.FieldGORMTag("created_at", func(tag field.GormTag) field.GormTag {
			tag.Set("autoCreateTime", "")
			tag.Set("type", "datetime")

			return tag
		}),
		gen.FieldGORMTag("updated_at", func(tag field.GormTag) field.GormTag {
			tag.Set("autoUpdateTime", "")
			tag.Set("type", "datetime")

			return tag
		}),
		gen.FieldGORMTag("deleted_at", func(tag field.GormTag) field.GormTag {
			tag.Set("type", "datetime")
			tag.Set("default", "NULL")
			return tag
		}),
		gen.FieldJSONTagWithNS(func(columnName string) string {
			return SQLColumnToHumpStyle(columnName)
		}),

		gen.FieldNewTagWithNS("form", func(columnName string) string {
			return SQLColumnToHumpStyle(columnName)
		}),
	}

	tableList, _ := db.Migrator().GetTables()

	for _, table := range tableList {
		if table == "sqlite_sequence" {
			continue
		}
		if strings.HasPrefix(table, "sqlite_") {
			continue
		}

		g.ApplyBasic(g.GenerateModel(table, opts...))
	}
	g.Execute()

}
