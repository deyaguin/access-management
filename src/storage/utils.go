package storage

import "strconv"

func mssqlLimit(tableName string, limit, offset int) string {
	query := "SELECT * FROM ( SELECT *, ROW_NUMBER() OVER (ORDER BY name) as row FROM " +
		tableName + ") a WHERE a.row > " + strconv.Itoa(limit) + " and a.row <= " +
		strconv.Itoa(offset)
	return query
}
