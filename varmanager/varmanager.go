package varmanager

import (
	"html/template"
	"reflect"
	"strconv"
	"strings"

	"github.com/kolo7/bench-tpl/db"
	"github.com/samber/lo"
)

type VarManager struct {
	vars   map[string]interface{}
	tables map[string]*db.Table
	funcs  template.FuncMap
}

func NewVarManager() *VarManager {
	return &VarManager{
		vars:   make(map[string]interface{}),
		tables: make(map[string]*db.Table),
		funcs: map[string]interface{}{
			"rand":          RandomInt,
			"randomLetters": RandomLetters,
			"randomNumbers": RandomNumbers,
			"randomChinese": RandomChinese,
		},
	}
}

func (vm *VarManager) SetGlobalVar(name string, value interface{}) {
	if value == nil {
		return
	}
	vm.vars[name] = value
}

func (vm *VarManager) SetTableVar(table *db.Table) {
	if table == nil {
		return
	}
	vm.tables[table.Name] = table
}

func (vm *VarManager) SetFuncVar(name string, value interface{}) {
	if value == nil {
		return
	}
	vm.funcs[name] = value
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

func (vm *VarManager) GetTables() map[string]interface{} {
	tableMap := make(map[string]interface{})
	for name, table := range vm.tables {
		rowMap := make(map[string]interface{})
		// 打乱rows顺序，取第一个元素作为基准
		rows := lo.Shuffle[db.Row](table.Rows)
		if len(rows) == 0 {
			return rowMap
		}
		for _, ele := range rows[0] {
			rowMap[ele.Column.Field] = ele.Val
		}
		tableMap[name] = rowMap
	}
	return tableMap
}

func (vm *VarManager) GetFuncMap() template.FuncMap {
	return vm.funcs
}
