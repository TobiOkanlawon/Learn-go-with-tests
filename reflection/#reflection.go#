package main

import (
	"reflect"
)

func Walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.Slice:
			for i := 0; i < val.Len(); i++ {
				Walk(val.Index(i).Interface(), fn)
			}

		case reflect.String:
			fn(field.String())

		case reflect.Struct:
			Walk(field.Interface(), fn)

		}
	}
}
