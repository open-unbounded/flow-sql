// Code generated from D:/nextgen/golang/flow-sql\FlowSql.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FlowSql

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by FlowSqlParser.
type FlowSqlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FlowSqlParser#sql.
	VisitSql(ctx *SqlContext) interface{}

	// Visit a parse tree produced by FlowSqlParser#where.
	VisitWhere(ctx *WhereContext) interface{}

	// Visit a parse tree produced by FlowSqlParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by FlowSqlParser#parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by FlowSqlParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by FlowSqlParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by FlowSqlParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by FlowSqlParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by FlowSqlParser#fields.
	VisitFields(ctx *FieldsContext) interface{}

	// Visit a parse tree produced by FlowSqlParser#alias.
	VisitAlias(ctx *AliasContext) interface{}
}
