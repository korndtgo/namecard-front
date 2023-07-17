package util

import (
	"errors"
	"reflect"
	"regexp"
	"strings"
)

var (
	ErrStcCovDestType = errors.New("destination wrong type")
	stcCovTitle       = "json="
	stcCovTitle2      = "name="
	stcCovRegex, _    = regexp.Compile(stcCovTitle + `\w+`)
	stcCovRegex2, _   = regexp.Compile(stcCovTitle2 + `\w+`)
)

func Struct2ProtoModel(from interface{}, to interface{}) error {
	fromType := reflect.TypeOf(from)
	if fromType.Kind() == reflect.Ptr {
		fromType = fromType.Elem()
	}

	toType := reflect.TypeOf(to)
	if toType.Kind() != reflect.Ptr {
		return ErrStcCovDestType
	}
	toType = toType.Elem()
	if toType.Kind() != reflect.Struct {
		return ErrStcCovDestType
	}

	fromValue := reflect.Indirect(reflect.ValueOf(from))
	toValue := reflect.Indirect(reflect.ValueOf(to))
	for i := 0; i < toType.NumField(); i++ {
		toStructFieldElement := toType.Field(i)
		fromStructFieldName := toStructFieldElement.Name
		_, found := fromType.FieldByName(toStructFieldElement.Name)
		if !found {
			toTagProtobuf := toStructFieldElement.Tag.Get("protobuf")
			title := stcCovTitle
			str := stcCovRegex.FindString(toTagProtobuf)
			if str == "" {
				title = stcCovTitle2
				str = stcCovRegex2.FindString(toTagProtobuf)
				if str == "" {
					continue
				}
			}
			toTagProtobufTitle := str[strings.Index(str, title)+len(title):]
			for j := 0; j < fromType.NumField(); j++ {
				fromStructFieldElement := fromType.Field(j)
				fromTagJson := fromStructFieldElement.Tag.Get("json")
				if toTagProtobufTitle == fromTagJson {
					found = true // change to be found
					fromStructFieldName = fromStructFieldElement.Name
				}
			}
		}

		if found {
			fromValueElement := fromValue.FieldByName(fromStructFieldName)
			toValueElement := toValue.Field(i)
			if fromValueElement.Kind() == reflect.Ptr {
				if fromValueElement.IsNil() {
					toValueElement.Set(reflect.Zero(toStructFieldElement.Type))
					continue
				}
				fromValueElement = fromValueElement.Elem()
			}
			if toStructFieldElement.Type.Kind() == reflect.Ptr {
				if toValueElement.IsNil() {
					toValueElement.Set(reflect.New(toStructFieldElement.Type.Elem())) // set ptr
					if toValueElement.Elem().CanSet() && fromValueElement.Type().ConvertibleTo(toStructFieldElement.Type.Elem()) {
						toValueElement.Elem().Set(fromValueElement.Convert(toStructFieldElement.Type.Elem()))
					}
				}
			} else {
				if toValueElement.CanSet() && fromValueElement.Type().ConvertibleTo(toStructFieldElement.Type) {
					toValueElement.Set(fromValueElement.Convert(toStructFieldElement.Type))
				}
			}

		}
	}

	return nil
}

func ProtoModel2Struct(from interface{}, to interface{}) error {
	fromType := reflect.TypeOf(from)
	if fromType.Kind() == reflect.Ptr {
		fromType = fromType.Elem()
	}

	toType := reflect.TypeOf(to)
	if toType.Kind() != reflect.Ptr {
		return ErrStcCovDestType
	}
	toType = toType.Elem()
	if toType.Kind() != reflect.Struct {
		return ErrStcCovDestType
	}

	fromValue := reflect.Indirect(reflect.ValueOf(from))
	toValue := reflect.Indirect(reflect.ValueOf(to))
	for i := 0; i < toType.NumField(); i++ {
		toStructFieldElement := toType.Field(i)
		fromStructFieldName := toStructFieldElement.Name
		_, found := fromType.FieldByName(toStructFieldElement.Name)
		if !found {
			toTagJson := toStructFieldElement.Tag.Get("json")
			for j := 0; j < fromType.NumField(); j++ {
				fromStructFieldElement := fromType.Field(j)
				fromTagProtobuf := fromStructFieldElement.Tag.Get("protobuf")
				title := stcCovTitle
				str := stcCovRegex.FindString(fromTagProtobuf)
				if str == "" {
					title = stcCovTitle2
					str = stcCovRegex2.FindString(fromTagProtobuf)
					if str == "" {
						continue
					}
				}
				fromTagProtobufTitle := str[strings.Index(str, title)+len(title):]
				if toTagJson == fromTagProtobufTitle {
					found = true // change to be found
					fromStructFieldName = fromStructFieldElement.Name
				}
			}
		}

		if found {
			fromValueElement := fromValue.FieldByName(fromStructFieldName)
			toValueElement := toValue.Field(i)
			if fromValueElement.Kind() == reflect.Ptr {
				if fromValueElement.IsNil() {
					toValueElement.Set(reflect.Zero(toStructFieldElement.Type))
					continue
				}
				fromValueElement = fromValueElement.Elem()
			}
			if toStructFieldElement.Type.Kind() == reflect.Ptr {
				if toValueElement.IsNil() {
					toValueElement.Set(reflect.New(toStructFieldElement.Type.Elem())) // set ptr
					if toValueElement.Elem().CanSet() && fromValueElement.Type().ConvertibleTo(toStructFieldElement.Type.Elem()) {
						toValueElement.Elem().Set(fromValueElement.Convert(toStructFieldElement.Type.Elem()))
					}
				}
			} else {
				if toValueElement.CanSet() && fromValueElement.Type().ConvertibleTo(toStructFieldElement.Type) {
					toValueElement.Set(fromValueElement.Convert(toStructFieldElement.Type))
				}
			}

		}
	}

	return nil
}
