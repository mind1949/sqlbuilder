package sqlbuilder

type Sqlizer interface {
	ToSql() (string, []interface{}, error)
}

type rawSqlizer interface {
	toSqlRaw() (string, []interface{}, error)
}
