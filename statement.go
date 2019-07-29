package sqlbuilder

import "github.com/lann/builder"

type StatementBuilderType builder.Builder

func (b StatementBuilderType) Select(columns ...string) SelectBuilder {
	return SelectBuilder(b).Columns(columns...)
}

func Select(columns ...string) SelectBuilder {
	return StatementBuilderType{}.Select(columns...)
}
