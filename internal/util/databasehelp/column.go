package databasehelp

import (
	"reflect"
	"strings"

	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// m: type A struct {
// 	T1 string `gorm:"column:t1"`
// 	T2 string
// }
// return to ...
// []clause.Column{
// 	{
// 		Alias:  "A__T1",
// 		Name:  "t1",
// 		Table: "a",
// 	},
// 	{
// 		Alias:  "A__T2",
// 		Name:  "t2",
// 		Table: "a",
// 	},
// }
func (dh *DatabaseHelp) ModelToColumns(m interface{}, tableNames ...string) []clause.Column {
	fields := make([]clause.Column, 0)
	mType := reflect.TypeOf(m)
	mValue := reflect.ValueOf(m)
	namingStrategy := new(schema.NamingStrategy)
	var tableName string
	if len(tableNames) > 0 {
		tableName = tableNames[0]
	} else {
		mMethod := mValue.MethodByName("TableName")
		if mMethod.IsValid() && mMethod.Type().NumIn() == 0 && mMethod.Type().NumOut() == 1 && mMethod.Type().Out(0).Kind() == reflect.String {
			tableName = mMethod.Call([]reflect.Value{})[0].String()
		} else {
			tableName = namingStrategy.TableName(mType.Name())
		}
	}
	for i := 0; i < mType.NumField(); i++ {
		elementStructField := mType.Field(i)
		elementType := elementStructField.Type
		if elementType.Kind() == reflect.Ptr {
			elementType = elementType.Elem()
		}
		if elementType.Kind() == reflect.Array || elementType.Kind() == reflect.Slice || elementType.Kind() == reflect.Struct {
			continue
		}
		column, ok := schema.ParseTagSetting(elementStructField.Tag.Get("gorm"), ";")["COLUMN"]
		if !ok {
			column = namingStrategy.ColumnName(tableName, elementStructField.Name)
		}

		fields = append(fields, clause.Column{
			Alias: mType.Name() + "__" + elementStructField.Name,
			Name:  column,
			Table: tableName,
		})
	}
	return fields
}

// m: type A struct {
// 	T1 string `gorm:"column:t1"`
// 	T2 string
// }
// return to ...
// []clause.Column{
// 	{
// 		Alias:  "t1",
// 		Name:  "t1",
// 		Table: "a",
// 	},
// 	{
// 		Alias:  "t2",
// 		Name:  "t2",
// 		Table: "a",
// 	},
// }
func (dh *DatabaseHelp) ModelToColumns2(m interface{}, tableNames ...string) []clause.Column {
	fields := make([]clause.Column, 0)
	mType := reflect.TypeOf(m)
	mValue := reflect.ValueOf(m)
	namingStrategy := new(schema.NamingStrategy)
	var tableName string
	if len(tableNames) > 0 {
		tableName = tableNames[0]
	} else {
		mMethod := mValue.MethodByName("TableName")
		if mMethod.IsValid() && mMethod.Type().NumIn() == 0 && mMethod.Type().NumOut() == 1 && mMethod.Type().Out(0).Kind() == reflect.String {
			tableName = mMethod.Call([]reflect.Value{})[0].String()
		} else {
			tableName = namingStrategy.TableName(mType.Name())
		}
	}
	for i := 0; i < mType.NumField(); i++ {
		elementStructField := mType.Field(i)
		elementType := elementStructField.Type
		if elementType.Kind() == reflect.Ptr {
			elementType = elementType.Elem()
		}
		if elementType.Kind() == reflect.Array || elementType.Kind() == reflect.Slice || elementType.Kind() == reflect.Struct {
			continue
		}
		column, ok := schema.ParseTagSetting(elementStructField.Tag.Get("gorm"), ";")["COLUMN"]
		if !ok {
			column = namingStrategy.ColumnName(tableName, elementStructField.Name)
		}

		fields = append(fields, clause.Column{
			Alias: column,
			Name:  column,
			Table: tableName,
		})
	}
	return fields
}

// source: []clause.Column{
// 	{
// 		Alias:  "A__T1",
// 		Name:  "t1",
// 		Table: "a",
// 	},
// 	{
// 		Alias:  "A__T2",
// 		Name:  "t2",
// 		Table: "a",
// 	},
// }
// interestFields: []string{"X", "Y"}
// return to ...
// []clause.Column{
// 	{
// 		Alias:  "X",
// 		Name:  "t1",
// 		Table: "a",
// 	},
// 	{
// 		Alias:  "Y",
// 		Name:  "t2",
// 		Table: "a",
// 	},
// }
func (dh *DatabaseHelp) RenameColumnAlias(source []clause.Column, interestFields []string) []clause.Column {
	var results []clause.Column
	for _, v := range source {
		str := strings.Split(v.Alias, "__")[1]
		for _, v2 := range interestFields {
			if str == v2 {
				results = append(results, clause.Column{
					Alias: v2,
					Name:  v.Name,
					Table: v.Table,
				})
			}
		}
	}
	return results
}

// source: []clause.Column{
// 	{
// 		Alias:  "A__T1",
// 		Name:  "t1",
// 		Table: "a",
// 	},
// 	{
// 		Alias:  "A__T2",
// 		Name:  "t2",
// 		Table: "a",
// 	},
// }
// interestFields: []string{"X", "Y"}
// return to ...
// []clause.Column{
// 	{
// 		Alias:  "A__T1",
// 		Name:  "t1",
// 		Table: "a",
// 	},
// 	{
// 		Alias:  "A__T2",
// 		Name:  "t2",
// 		Table: "a",
// 	},
// }
func (dh *DatabaseHelp) RenameColumnAlias2(source []clause.Column, interestFields []string) []clause.Column {
	var results []clause.Column
	for _, v := range source {
		str := strings.Split(v.Alias, "__")[1]
		for _, v2 := range interestFields {
			if str == v2 {
				results = append(results, clause.Column{
					Alias: v.Alias,
					Name:  v.Name,
					Table: v.Table,
				})
			}
		}
	}
	return results
}

// source: []clause.Column{
// 	{
// 		Alias:  "A__T1",
// 		Name:  "t1",
// 		Table: "a",
// 	},
// 	{
// 		Alias:  "A__T2",
// 		Name:  "t2",
// 		Table: "a",
// 	},
// }
// title: "H"
// tableName: "j"
// interestFields: []string{"X", "Y"}
// return to ...
// []clause.Column{
// 	{
// 		Alias:  "H__X",
// 		Name:  "A__T1",
// 		Table: "j",
// 	},
// 	{
// 		Alias:  "H__Y",
// 		Name:  "A__T2",
// 		Table: "j",
// 	},
// }
func (dh *DatabaseHelp) RenameColumnAlias3(source []clause.Column, title string, tableName string, interestFields []string) []clause.Column {
	var results []clause.Column
	for _, v := range source {
		str := strings.Split(v.Alias, "__")[1]
		for _, v2 := range interestFields {
			if str == v2 {
				results = append(results, clause.Column{
					Alias: title + "__" + str,
					Name:  v.Alias,
					Table: tableName,
				})
			}
		}
	}
	return results
}

// source: []clause.Column{
// 	{
// 		Alias:  "A__T1",
// 		Name:  "t1",
// 		Table: "a",
// 	},
// 	{
// 		Alias:  "A__T2",
// 		Name:  "t2",
// 		Table: "a",
// 	},
// }
// title: "H"
// tableName: "j"
// interestFields: []string{"X", "Y"}
// return to ...
// []clause.Column{
// 	{
// 		Alias:  "H__X",
// 		Name:  "t1",
// 		Table: "j",
// 	},
// 	{
// 		Alias:  "H__Y",
// 		Name:  "t2",
// 		Table: "j",
// 	},
// }
func (dh *DatabaseHelp) RenameColumnAlias4(source []clause.Column, title string, tableName string, interestFields []string) []clause.Column {
	var results []clause.Column
	for _, v := range source {
		str := strings.Split(v.Alias, "__")[1]
		for _, v2 := range interestFields {
			if str == v2 {
				results = append(results, clause.Column{
					Alias: title + "__" + str,
					Name:  v.Name,
					Table: tableName,
				})
			}
		}
	}
	return results
}

func (dh *DatabaseHelp) AppendColumn(a []clause.Column, b ...clause.Column) []clause.Column {
	unique := make([]clause.Column, 0)
	type key struct{ Alias string }
	m := make(map[key]int)
	d := append(a, b...)
	for _, v := range d {
		k := key{v.Alias}
		if _, ok := m[k]; ok {
			// enable to delete previous duplicate values
			// disable to ignore duplicate values
			// unique[i] = v
		} else {
			m[k] = len(unique)
			unique = append(unique, v)
		}
	}

	return unique
}

func (dh *DatabaseHelp) MakeInterestFields(m interface{}) []string {
	interestFields := make([]string, 0)
	resultType := reflect.TypeOf(m)
	for i := 0; i < resultType.NumField(); i++ {
		interestFields = append(interestFields, resultType.Field(i).Name)
	}
	return interestFields
}

func (dh *DatabaseHelp) ComposeColumn(m interface{}, tableName string, interestFields []string) []clause.Column {
	if tableName != "" {
		return dh.RenameColumnAlias(dh.ModelToColumns(m, tableName), interestFields)
	}
	return dh.RenameColumnAlias(dh.ModelToColumns(m), interestFields)
}

// m: type A struct {
// 	T1 string `gorm:"column:t1"`
// 	T2 string
// }
// return to ...
// []clause.Column{
// 	{
// 		Alias:  "A__T1",
// 		Name:  "A__T1",
// 		Table: "a",
// 	},
// 	{
// 		Alias:  "A__T2",
// 		Name:  "A__T2",
// 		Table: "a",
// 	},
// }
func (dh *DatabaseHelp) ModelToColumnsIgnoreDB(m interface{}, tableName string) []clause.Column {
	fields := make([]clause.Column, 0)
	mType := reflect.TypeOf(m)
	for i := 0; i < mType.NumField(); i++ {
		elementStructField := mType.Field(i)
		elementType := elementStructField.Type
		if elementType.Kind() == reflect.Ptr {
			elementType = elementType.Elem()
		}
		if elementType.Kind() == reflect.Array || elementType.Kind() == reflect.Slice || elementType.Kind() == reflect.Struct {
			continue
		}

		fieldName := mType.Name() + "__" + elementStructField.Name
		fields = append(fields, clause.Column{
			Alias: fieldName,
			Name:  fieldName,
			Table: tableName,
		})
	}
	return fields
}

func (dh *DatabaseHelp) ComposeColumnIgnoreDB(m interface{}, tableName string, interestFields []string) []clause.Column {
	return dh.RenameColumnAlias(dh.ModelToColumnsIgnoreDB(m, tableName), interestFields)
}
