package helper_funcs

import (
	"fmt"
	"reflect"
)

// Column 通用字段提取函数
// src: 源切片（[]struct / []map[string]interface{}）
// field: 要提取的字段名（结构体字段名/映射的key）
// 返回: 提取后的字段切片
func Column(src interface{}, field string) ([]interface{}, error) {
	// 解析源数据的反射类型
	srcVal := reflect.ValueOf(src)
	if srcVal.Kind() != reflect.Slice {
		return nil, fmt.Errorf("src must be a slice")
	}

	result := make([]interface{}, 0, srcVal.Len())

	// 遍历切片元素
	for i := 0; i < srcVal.Len(); i++ {
		elem := srcVal.Index(i)
		switch elem.Kind() {
		// 处理结构体切片（最常用）
		case reflect.Struct:
			fieldVal := elem.FieldByName(field)
			if !fieldVal.IsValid() {
				return nil, fmt.Errorf("struct has no field: %s", field)
			}
			result = append(result, fieldVal.Interface())

		// 处理映射切片（如 []map[string]interface{}）
		case reflect.Map:
			key := reflect.ValueOf(field)
			fieldVal := elem.MapIndex(key)
			if !fieldVal.IsValid() {
				return nil, fmt.Errorf("map has no key: %s", field)
			}
			result = append(result, fieldVal.Interface())

		default:
			return nil, fmt.Errorf("unsupported element type: %s", elem.Kind())
		}
	}

	return result, nil
}

// SliceColumn 增强版字段提取函数（对齐 PHP array_column）
// src: 源切片（[]struct / []map[string]interface{}）
// valueField: 要提取的数值字段名（结构体字段名/映射的key）
// keyField: 可选的键名字段名（为空则返回 []interface{}，不为空则返回 map[interface{}]interface{}）
// 返回: 提取结果（[]interface{} 或 map[interface{}]interface{}）、错误信息
func SliceColumn(src interface{}, valueField string, keyField ...string) (interface{}, error) {
	// 解析源数据的反射类型
	srcVal := reflect.ValueOf(src)
	if srcVal.Kind() != reflect.Slice {
		return nil, fmt.Errorf("src must be a slice, got %s", srcVal.Kind())
	}

	// 判断是否需要按key返回（第三个参数是否存在）
	useKey := len(keyField) > 0 && keyField[0] != ""

	var (
		resultSlice []interface{}               // 纯值切片结果
		resultMap   map[interface{}]interface{} // 键值映射结果
	)

	// 初始化结果容器
	if useKey {
		resultMap = make(map[interface{}]interface{}, srcVal.Len())
	} else {
		resultSlice = make([]interface{}, 0, srcVal.Len())
	}

	// 遍历切片元素
	for i := 0; i < srcVal.Len(); i++ {
		elem := srcVal.Index(i)
		// 确保能获取到元素的实际值（处理指针类型）
		for elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}

		// 提取数值字段值
		val, err := getFieldValue(elem, valueField)
		if err != nil {
			return nil, fmt.Errorf("get value field %s failed: %w", valueField, err)
		}

		// 如果指定了key字段，则提取key并组装map
		if useKey {
			key, err := getFieldValue(elem, keyField[0])
			if err != nil {
				return nil, fmt.Errorf("get key field %s failed: %w", keyField[0], err)
			}
			resultMap[key] = val
		} else {
			// 否则直接追加到切片
			resultSlice = append(resultSlice, val)
		}
	}

	// 根据是否使用key返回不同类型
	if useKey {
		return resultMap, nil
	}
	return resultSlice, nil
}

// getFieldValue 通用字段值提取（适配结构体/映射）
func getFieldValue(elem reflect.Value, field string) (interface{}, error) {
	switch elem.Kind() {
	case reflect.Struct:
		fieldVal := elem.FieldByName(field)
		if !fieldVal.IsValid() {
			return nil, fmt.Errorf("struct has no field: %s", field)
		}
		return fieldVal.Interface(), nil

	case reflect.Map:
		key := reflect.ValueOf(field)
		fieldVal := elem.MapIndex(key)
		if !fieldVal.IsValid() {
			return nil, fmt.Errorf("map has no key: %s", field)
		}
		return fieldVal.Interface(), nil

	default:
		return nil, fmt.Errorf("unsupported element type: %s", elem.Kind())
	}
}
