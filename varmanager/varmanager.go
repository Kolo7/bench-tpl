package varmanager

import (
	"html/template"
	"strings"

	"github.com/Kolo7/bench-tpl/db"
	"github.com/Kolo7/bench-tpl/utils"
	"github.com/samber/lo"
)

type VarManager struct {
	globalVars       map[string]interface{}
	packageVars      map[string]interface{}
	tableExampleVars map[string]map[string]interface{}
	tables           map[string]*db.Table
	funcs            template.FuncMap
}

func NewVarManager() *VarManager {
	return &VarManager{
		globalVars:       make(map[string]interface{}),
		tables:           make(map[string]*db.Table),
		tableExampleVars: make(map[string]map[string]interface{}),
		packageVars:      make(map[string]interface{}),
		funcs: map[string]interface{}{
			"rand":             RandomInt,
			"randomLetters":    RandomLetters,
			"randomNumbers":    RandomNumbers,
			"randomChinese":    RandomChinese,
			"toTag":            ToTag,
			"toUpperCamelCase": utils.ToUpperCamelCase,
			"toLowerCamelCase": utils.ToLowerCamelCase,
			"toSnakeCase":      utils.ToSnakeCase,
			"inExcludedFields": InExcludedFields,
			"sub":              Sub,
		},
	}
}

func (vm *VarManager) SetTableExampleVar(table *db.Table) {
	if table == nil {
		return
	}
	vm.tables[table.Name] = table

	rows := lo.Shuffle[db.Row](table.Rows)
	if len(rows) == 0 {
		return
	}
	rowMap := make(map[string]interface{})
	for _, ele := range rows[0] {
		rowMap[ele.Column.Field] = ele.Val
	}
	vm.tableExampleVars[table.Name] = rowMap
}

func (vm *VarManager) GetTablesExampleVar() map[string]map[string]interface{} {
	return vm.tableExampleVars
}

func (vm *VarManager) GetFuncMap() template.FuncMap {
	return vm.funcs
}

func (vm *VarManager) SetGlobalVar(fqdn string) {
	if fqdn != "" {
		vm.globalVars["fqdn"] = fqdn
	}
}

func (vm *VarManager) GetGlobalVar() map[string]interface{} {
	return vm.globalVars
}

// 设置包级别的变量
// 入参：包路径，包名
func (vm *VarManager) SetPackageVar(pkgPath ...string) {
	if len(pkgPath) == 0 {
		return
	}
	// 从全局变量获取fqdn
	fqdn := vm.globalVars["fqdn"].(string)
	vm.packageVars = make(map[string]interface{})
	// 把全局变量也放入包级别变量中
	loadMap2Map(vm.packageVars, vm.globalVars)
	// 包全名
	pkgFullName := strings.Join([]string{fqdn, strings.Join(pkgPath, "/")}, "/")
	upperPkgName := utils.ToUpperCamelCase(pkgPath[len(pkgPath)-1])
	lowerPkgName := utils.ToLowerCamelCase(pkgPath[len(pkgPath)-1])
	vm.packageVars["pkgFullName"] = pkgFullName
	vm.packageVars["upperPkgName"] = upperPkgName
	vm.packageVars["lowerPkgName"] = lowerPkgName

	// 特殊包名设置为全局变量
	// model
	if strings.HasSuffix(pkgFullName, "/model") {
		vm.globalVars["modelPackageName"] = pkgFullName
	}
	// dao
	if strings.HasSuffix(pkgFullName, "/dao") {
		vm.globalVars["daoPackageName"] = pkgFullName
	}
	// api
	if strings.HasSuffix(pkgFullName, "/api") {
		vm.globalVars["apiPackageName"] = pkgFullName
	}
}

// 获取包级别的变量
func (vm *VarManager) GetPackageVar() map[string]interface{} {
	return vm.packageVars
}

// 获取表级别的变量
func (vm *VarManager) GetTableVar(tableName string) map[string]interface{} {
	table := vm.tables[tableName]
	if table == nil {
		return nil
	}
	// 表级别的变量
	varMap := make(map[string]interface{})
	// 拷贝全局变量
	loadMap2Map(varMap, vm.globalVars)
	// 拷贝包级别变量
	loadMap2Map(varMap, vm.packageVars)
	// 将常用的表相关变量放入变量管理器
	varMap["tableName"] = table.Name
	// 大写开头的表名
	varMap["upperTableName"] = utils.ToUpperCamelCase(table.Name)
	// 小写开头的表名
	varMap["lowerTableName"] = utils.ToLowerCamelCase(table.Name)
	// 字段名列表
	varMap["tableColumnFields"] = lo.Map(table.Columns, func(col *db.Column, _ int) interface{} { return col.Field })
	// 大驼峰字段名列表
	varMap["tableColumnUpperFields"] = lo.Map(table.Columns, func(col *db.Column, _ int) interface{} { return utils.ToUpperCamelCase(col.Field) })
	// 主键字段
	varMap["tablePrimaryKey"] = lo.Reduce(table.Columns, func(agg *db.Column, col *db.Column, _ int) *db.Column {
		if col.Key == "PRI" {
			// 主键字段名
			varMap["tableUpperPrimaryKeyField"] = utils.ToUpperCamelCase(col.Field)
			return col
		}
		return agg
	}, nil)
	varMap["tableColumns"] = table.Columns
	return varMap
}

func loadMap2Map(m1, m2 map[string]interface{}) {
	for k, v := range m2 {
		m1[k] = v
	}
}
