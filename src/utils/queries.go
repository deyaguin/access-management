package utils

import "strconv"

func LimitQuery(tableName, selectorQuery string, limit, offset int) string {
	query := "SELECT * FROM ( SELECT *, ROW_NUMBER() OVER (ORDER BY name) as row FROM " +
		tableName + " WHERE " + selectorQuery + ") a WHERE a.row > " + strconv.Itoa(limit) +
		" AND a.row <= " + strconv.Itoa(offset)

	return query
}

func LikeQuery(tableName, columnName, columnValue string) string {
	query := `"` + tableName + `".deleted_at IS NULL AND ((` + columnName +
		" LIKE '%" + columnValue + "%'))"

	return query
}
