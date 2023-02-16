package util

import (
	"strings"
	"time"
)

func CollectInsert(table string, fields []string, values []string) string {
	currentTime := time.Now()
	t := currentTime.Format("2006.01.02 15:04:05")
	sqltxt := "INSERT INTO `" + table + "` "
	sqltxt += " (`" + strings.Join(fields, "`,`") + "`,`created_at`, `updated_at`)"
	sqltxt += " VALUES "
	sqltxt += "('" + strings.Join(values, "','") + "','" + t + "', '" + t + "' );"
	return sqltxt
}
