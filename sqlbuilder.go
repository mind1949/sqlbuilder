package sqlbuilder

type Sqlizer interface {
	ToSql() (string, []interface{}, error)
}
