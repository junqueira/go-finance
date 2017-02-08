package finance

import (
	"reflect"

	"github.com/shopspring/decimal"
)

func mapFields(vals []string, v interface{}) {

	typ := reflect.TypeOf(v)
	val := reflect.ValueOf(v).Elem()

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	for i := 0; i < typ.NumField(); i++ {

		f := val.Field(i)

		switch f.Interface().(type) {
		case string:
			f.Set(reflect.ValueOf(vals[i]))
		case int:
			f.Set(reflect.ValueOf(toInt(vals[i])))
		case Timestamp:
			f.Set(reflect.ValueOf(newStamp(vals[i])))
		case decimal.Decimal:
			f.Set(reflect.ValueOf(toDecimal(vals[i])))
		}
	}

}

func constructFields(in interface{}) (fields string) {

	typ := reflect.TypeOf(in)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		tag := f.Tag.Get("yfin")
		fields = fields + tag
	}

	return
}
