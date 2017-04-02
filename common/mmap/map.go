package mmap

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

//FillStruct 用map填充结构体
func FillStruct(data map[string]interface{}, object interface{}) error {
	for k, v := range data {
		err := SetFiled(object, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

//SetFiled 用map的值替换struct的值
func SetFiled(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFiledValue := structValue.FieldByName(name)
	if !structFiledValue.IsValid() {
		return fmt.Errorf("no such filed %s in obj", name)
	}
	if !structFiledValue.CanSet() {
		return fmt.Errorf("cannot set %s filed value", name)
	}
	structFiledType := structFiledValue.Type()
	val := reflect.ValueOf(value)

	var err error
	if structFiledType != val.Type() {
		val, err = TypeConversion(fmt.Sprintf("%v", value), structFiledValue.Type().Name())
		if err != nil {
			return err
		}
	}
	structFiledValue.Set(val)
	return nil
}

//TypeConversion 类型转换
func TypeConversion(value string, ntype string) (reflect.Value, error) {
	if ntype == "string" {
		return reflect.ValueOf(value), nil
	} else if ntype == "time.Time" {
		t, err := time.ParseInLocation("", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "Time" {
		t, err := time.ParseInLocation("", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "int" {
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err
	} else if ntype == "int8" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int8(i)), err
	} else if ntype == "int16" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int16(i)), err
	} else if ntype == "int32" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int32(i)), err
	} else if ntype == "int64" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	} else if ntype == "float32" {
		f, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(float32(f)), err
	} else if ntype == "float64" {
		f, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(f), err
	}
	return reflect.ValueOf(value), errors.New("unknow type：" + ntype)
}
