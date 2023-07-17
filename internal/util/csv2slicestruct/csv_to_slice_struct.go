package csv2slicestruct

import (
	"campaign-service/internal/util/covstr2any"
	"encoding/csv"
	"errors"
	"io"
	"reflect"
	"strings"
)

var (
	ErrDstType = errors.New("destination wrong type")
)

// const cursorKey = "Date"
// const cursorValue = "20210811"

func CSV2SliceStruct(file io.Reader, to interface{}) error {
	var structIsPtr bool
	toStructType := reflect.TypeOf(to) // expect ptr
	if toStructType.Kind() != reflect.Ptr {
		return ErrDstType
	}
	toStructType = toStructType.Elem() // expect slice
	if toStructType.Kind() != reflect.Slice {
		return ErrDstType
	}
	if toStructType.Elem().Kind() == reflect.Ptr {
		toStructType = toStructType.Elem()
		structIsPtr = true
	}
	toStructType = toStructType.Elem() // expect struct
	if toStructType.Kind() != reflect.Struct {
		return ErrDstType
	}
	toValue := reflect.ValueOf(to).Elem()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.ReuseRecord = true
	var csvRowIndex int64 = -1
	mappingKeyToIndex := make(map[string]int)
	for {
		csvRowIndex++

		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if csvRowIndex == 0 {
			for i := 0; i < toStructType.NumField(); i++ {
				elt := toStructType.Field(i)
				tag := elt.Tag
				csvTag := tag.Get("csv")
				for j, v := range record {
					if v == csvTag && v != "" && csvTag != "" {
						mappingKeyToIndex[elt.Name] = j
					}
				}
			}
			continue
		}

		// if key, ok := mappingKeyToIndex[cursorKey]; ok {
		// 	value := record[key]
		// 	t1 := strings.Split(value, "/")
		// 	day, month, year := t1[0], t1[1], t1[2]
		// 	if year+month+fmt.Sprintf("%02s", day) < cursorValue {
		// 		continue
		// 	}
		// }

		tempToFieldValue := reflect.Indirect(reflect.New(toStructType))
		for i := 0; i < toStructType.NumField(); i++ {
			elt := toStructType.Field(i)
			name := elt.Name
			key, ok := mappingKeyToIndex[name]
			if !ok {
				continue
			}
			value := record[key]
			if strings.TrimSpace(value) == "" {
				continue
			}

			tag := elt.Tag
			mutationTag := tag.Get("mutation")
			elv := tempToFieldValue.Field(i)
			rv, err := covstr2any.ConvertStringToAnyType(value, elt.Type)
			if err != nil {
				return err
			}

			if mutationTag != "" {
				mutationMethod := tempToFieldValue.MethodByName(mutationTag)
				if mutationMethod.IsValid() && mutationMethod.Type().NumIn() == 1 && mutationMethod.Type().NumOut() == 1 {
					rv = mutationMethod.Call([]reflect.Value{rv})[0]
				}
			}
			elv.Set(rv)
		}

		if structIsPtr {
			tempToFieldValue = tempToFieldValue.Addr()
		}
		toValue.Set(reflect.Append(toValue, tempToFieldValue))
	}

	return nil
}

// func CSV2SliceStruct(file io.Reader, to interface{}) error {
// 	reader := csv.NewReader(file)
// 	reader.FieldsPerRecord = -1
// 	reader.ReuseRecord = true
// 	csvLines, err := reader.ReadAll()
// 	if err != nil {
// 		return err
// 	}
// 	var structIsPtr bool
// 	toType := reflect.TypeOf(to) // expect ptr
// 	if toType.Kind() != reflect.Ptr {
// 		return ErrDstType
// 	}
// 	toType = toType.Elem() // expect slice
// 	if toType.Kind() != reflect.Slice {
// 		return ErrDstType
// 	}
// 	if toType.Elem().Kind() == reflect.Ptr {
// 		toType = toType.Elem()
// 		structIsPtr = true
// 	}
// 	toType = toType.Elem() // expect struct
// 	if toType.Kind() != reflect.Struct {
// 		return ErrDstType
// 	}

// 	toValue := reflect.ValueOf(to).Elem()
// 	mappingKeyToIndex := make(map[string]int)
// 	for i, fields := range csvLines {
// 		if i == 0 {
// 			for i := 0; i < toType.NumField(); i++ {
// 				elt := toType.Field(i)
// 				tag := elt.Tag
// 				csvTag := tag.Get("csv")
// 				for i2, v := range fields {
// 					if v == csvTag && v != "" && csvTag != "" {
// 						mappingKeyToIndex[elt.Name] = i2
// 					}
// 				}
// 			}
// 			continue
// 		}

// 		tempToFieldValue := reflect.Indirect(reflect.New(toType))
// 		for j := 0; j < toType.NumField(); j++ {
// 			elt := toType.Field(j)
// 			tag := elt.Tag
// 			mutationTag := tag.Get("mutation")
// 			name := elt.Name
// 			key, ok := mappingKeyToIndex[name]
// 			if !ok {
// 				continue
// 			}
// 			value := fields[key]
// 			if strings.TrimSpace(value) == "" {
// 				continue
// 			}

// 			elv := tempToFieldValue.Field(j)
// 			rv, err := covstr2any.ConvertStringToAnyType(value, elt.Type)
// 			if err != nil {
// 				return err
// 			}

// 			if mutationTag != "" {
// 				mutationMethod := tempToFieldValue.MethodByName(mutationTag)
// 				if mutationMethod.IsValid() && mutationMethod.Type().NumIn() == 1 && mutationMethod.Type().NumOut() == 1 {
// 					rv = mutationMethod.Call([]reflect.Value{rv})[0]
// 				}
// 			}
// 			elv.Set(rv)
// 		}

// 		if structIsPtr {
// 			tempToFieldValue = tempToFieldValue.Addr()
// 		}
// 		toValue.Set(reflect.Append(toValue, tempToFieldValue))
// 	}

// 	return nil
// }
