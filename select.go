package sqlbuilder

import (
	"bytes"
	"fmt"
	"github.com/lann/builder"
)

type selectData struct {
	Columns    []Sqlizer
	From       Sqlizer
	WhereParts []Sqlizer
}

func (d *selectData) ToSql() (sqlStr string, args []interface{}, err error) {
	sqlStr, args, err = d.toSql()
	if err != nil {
		return
	}

	// sqlStr, err = d.PlaceholderFormat.ReplacePlaceholder(sqlStr)
	return
}

func (d *selectData) toSqlRaw() (sqlStr string, args []interface{}, err error) {
	return d.toSql()
}

func (d *selectData) toSql() (sqlStr string, args []interface{}, err error) {
	if len(d.Columns) == 0 {
		err = fmt.Errorf("select statements must have at least one result column")
		return
	}

	sql := &bytes.Buffer{}

	sql.WriteString("SELECT ")

	if len(d.Columns) > 0 {
		args, err = appendToSql(d.Columns, sql, ", ", args)
		if err != nil {
			return
		}
	}

	if d.From != nil {
		sql.WriteString(" FROM ")
		args, err = appendToSql([]Sqlizer{d.From}, sql, "", args)
		if err != nil {
			return
		}
	}

	if len(d.WhereParts) > 0 {
		sql.WriteString(" WHERE ")
		args, err = appendToSql(d.WhereParts, sql, " AND ", args)
		if err != nil {
			return
		}
	}

	sqlStr = sql.String()
	return
}

type SelectBuilder builder.Builder

func init() {
	builder.Register(SelectBuilder{}, selectData{})
}

func (b SelectBuilder) ToSql() (string, []interface{}, error) {
	data := builder.GetStruct(b).(selectData)
	return data.ToSql()
}

func (b SelectBuilder) Columns(columns ...string) SelectBuilder {
	parts := make([]interface{}, 0, len(columns))
	for _, str := range columns {
		parts = append(parts, newPart(str))
	}
	return builder.Extend(b, "Columns", parts).(SelectBuilder)
}

func (b SelectBuilder) From(from string) SelectBuilder {
	return builder.Set(b, "From", newPart(from)).(SelectBuilder)
}

func (b SelectBuilder) Where(pred interface{}, args ...interface{}) SelectBuilder {
	if pred == nil || pred == "" {
		return b
	}
	return builder.Append(b, "WhereParts", newWherePart(pred, args...)).(SelectBuilder)
}
