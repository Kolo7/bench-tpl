package varmanager

import (
	"fmt"
	"math/rand"

	"github.com/samber/lo"
)

// 定义一个生成随机数的函数
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// 生成一个随机的数字字母组合，长度32位
func RandomLetters(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	for i := 0; i < length; i++ {
		result = append(result, bytes[RandomInt(0, len(bytes))])
	}
	return string(result)
}

// 生成固定长度的随机数字组合
func RandomNumbers(length int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	for i := 0; i < length; i++ {
		result = append(result, bytes[RandomInt(0, len(bytes))])
	}
	return string(result)
}

// 生成随机中文字符
func RandomChinese(length int) string {
	// 中文unicode编码范围
	min := 0x4E00
	max := 0x9FA5
	str := ""
	for i := 0; i < length; i++ {
		str += string(rune(RandomInt(min, max)))
	}
	return str
}

// tag 生成
func ToTag(colName string) string {
	return fmt.Sprintf(`json:"%s" db:"%s" gorm:"column:%s"`, colName, colName, colName)
}

// 判断数组中是否存在某个值
func InExcludedFields(val string) bool {
	return lo.Contains([]string{"Id", "CreatedAt", "CreateTime", "UpdateTime", "UpdatedAt"}, val)
}

func Sub(a, b int) int {
	if a < b {
		return 0
	}
	return a - b
}
