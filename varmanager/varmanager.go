package varmanager

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/kolo7/bench-tpl/db"
)

type VarManager struct {
	vars map[string]interface{}
}

func NewVarManager() *VarManager {
	return &VarManager{
		vars: make(map[string]interface{}),
	}
}

func (vm *VarManager) SetGlobalVar(name string, value interface{}) {
	if value == nil {
		return
	}
	vm.vars[name] = value
}

func (vm *VarManager) SetTableVar(Table *db.Table) {
	vm.vars["table_name"] = Table.Name
}

// 展平变量
func (vm *VarManager) SetFlattenVar(name string, value interface{}) {
	if value == nil {
		return
	}
	// 反射判断value是不是结构体
	typ := reflect.TypeOf(value)
	val := reflect.ValueOf(value)
	if typ.Kind() == reflect.Struct {
		// 如果是结构体，则将结构体的字段名和字段值都存入变量管理器
		for i := 0; i < typ.NumField(); i++ {
			subName := name + "_" + strings.ToLower(typ.Field(i).Name)
			vm.SetGlobalVar(subName, val.Field(i).Interface())
		}
	} else if typ.Kind() == reflect.Ptr {
		// 如果是指针，则递归调用SetVar函数
		vm.SetGlobalVar(name, val.Elem().Interface())
	} else if typ.Kind() == reflect.Slice {
		// 如果是切片，则将切片的元素都存入变量管理器
		for i := 0; i < val.Len(); i++ {
			subName := name + "_" + strconv.Itoa(i)
			vm.SetGlobalVar(subName, val.Index(i).Interface())
		}
	} else if typ.Kind() == reflect.Map {
		// 如果是map，则将map的键值对都存入变量管理器
		for _, key := range val.MapKeys() {
			subName := name + "_" + key.String()
			vm.SetGlobalVar(subName, val.MapIndex(key).Interface())
		}
	}

	// 其他类型直接存入变量管理器
	vm.vars[name] = value
}

func (vm *VarManager) Get() map[string]interface{} {
	return vm.vars
}
