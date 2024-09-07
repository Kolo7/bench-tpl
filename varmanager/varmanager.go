package varmanager

import (
	"reflect"
	"strconv"
	"strings"
)

type VarManager struct {
	vars map[string]interface{}
}

func NewVarManager() *VarManager {
	return &VarManager{
		vars: make(map[string]interface{}),
	}
}

func (vm *VarManager) SetVar(name string, value interface{}) {
	// 反射判断value是不是结构体
	typ := reflect.TypeOf(value)
	if typ.Kind() == reflect.Struct {
		// 如果是结构体，则将结构体的字段名和字段值都存入变量管理器
		for i := 0; i < typ.NumField(); i++ {
			subName := name + "_" + strings.ToLower(typ.Field(i).Name)
			vm.SetVar(subName, value.(reflect.Value).Field(i).Interface())
		}
	} else if typ.Kind() == reflect.Ptr {
		// 如果是指针，则递归调用SetVar函数
		vm.SetVar(name, value.(*reflect.Value).Interface())
	} else if typ.Kind() == reflect.Slice {
		// 如果是切片，则将切片的元素都存入变量管理器
		for i := 0; i < reflect.ValueOf(value).Len(); i++ {
			subName := name + "_" + strconv.Itoa(i)
			vm.SetVar(subName, reflect.ValueOf(value).Index(i).Interface())
		}
	} else if typ.Kind() == reflect.Map {
		// 如果是map，则将map的键值对都存入变量管理器
		for _, key := range reflect.ValueOf(value).MapKeys() {
			subName := name + "_" + key.String()
			vm.SetVar(subName, reflect.ValueOf(value).MapIndex(key).Interface())
		}
	}

	// 其他类型直接存入变量管理器
	vm.vars[name] = value

}

func (vm *VarManager) Get() map[string]interface{} {
	return vm.vars
}
