package {{.lowerPkgName}}

import (
	"context"
	"fmt"
	"reflect"
	"strings"
    "github.com/jinzhu/gorm"
	"github.com/samber/lo"
	"git.imgo.tv/ft/go-ceres/pkg/db/mysql"
)

const (
	dbTag          = "json"
)

var (
	ErrNotFound = gorm.ErrRecordNotFound
	ErrUniqueConflict = fmt.Errorf("unique conflict")
	
	ormDB *mysql.OrmDb
)

type count struct {
	Count int64 `json:"count"`
}

// RawFieldNames converts golang struct field into slice string.
func RawFieldNames(in any, postgreSql ...bool) []string {
	out := make([]string, 0)
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	var pg bool
	if len(postgreSql) > 0 {
		pg = postgreSql[0]
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("ToMap only accepts structs; got %T", v))
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		tagv := fi.Tag.Get(dbTag)
		switch tagv {
		case "-":
			continue
		case "":
			if pg {
				out = append(out, fi.Name)
			} else {
				out = append(out, fmt.Sprintf("`%s`", fi.Name))
			}
		default:
			// get tag name with the tag option, e.g.:
			// `db:"id"`
			// `db:"id,type=char,length=16"`
			// `db:",type=char,length=16"`
			// `db:"-,type=char,length=16"`
			if strings.Contains(tagv, ",") {
				tagv = strings.TrimSpace(strings.Split(tagv, ",")[0])
			}
			if tagv == "-" {
				continue
			}
			if len(tagv) == 0 {
				tagv = fi.Name
			}
			if pg {
				out = append(out, tagv)
			} else {
				out = append(out, fmt.Sprintf("`%s`", tagv))
			}
		}
	}

	return out
}

// Remove removes given strs from strings.
func Remove(strings []string, strs ...string) []string {
	out := append([]string(nil), strings...)

	for _, str := range strs {
		var n int
		for _, v := range out {
			if v != str {
				out[n] = v
				n++
			}
		}
		out = out[:n]
	}

	return out
}

func Count(ctx context.Context, db *gorm.DB, searchSql string, args ...any) (int64, error) {
	var (
		count    count
		countSql = fmt.Sprintf("SELECT COUNT(*) as count FROM (%s) AS t", searchSql)
	)
	if searchSql == "" {
		return 0, fmt.Errorf("dao Count searchSql should not be empty")
	}

	if strings.Contains(strings.ToLower(searchSql), "limit") {
		return 0, fmt.Errorf("dao Count searchSql should not contain limit")
	}

	err := db.Raw(countSql, args...).Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return count.Count, nil
}

func placeHolder(num int) string {
	if num == 0 {
		return ""
	}
	return strings.Repeat("?,", num-1) + "?"
}

func buildInsertSql(rows []string) string {
	if len(rows) == 0 {
		return ""
	}
	m := []string{}
	for _, row := range rows {
		m = append(m, "record."+row)
	}
	return strings.Join(m, ",")
}

func parseOrder(order string, fieldNames []string) string {
	order = strings.ToLower(order)
	orders := strings.Split(order, ",")
	result := make([]string, 0)
	for _, order := range orders {
		result = append(result, parseOrderSingle(strings.TrimSpace(order), fieldNames))
	}
	result = lo.Filter(result, func(s string, _ int) bool { return s != "" })
	if len(result) == 0 {
		return "id desc"
	}
	return strings.Join(result, ",")
}

func parseOrderSingle(order string, fieldNames []string) string {
	result := ""
	for _, col := range fieldNames {
		if strings.Contains(order, col) {
			result = col
			break
		}
	}
	if result == "" {
		return ""
	}
	if strings.Contains(order, "desc") {
		result = result + " desc"
	} else if strings.Contains(order, "asc") {
		result = result + " asc"
	}

	return result
}


type Dao struct {
	db *gorm.DB
}
func (d *Dao) DB(ctx context.Context) *gorm.DB {
	return d.db
}
