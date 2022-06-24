// Code generated from D:/nextgen/golang/flow-sql\FlowSql.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FlowSql

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FlowSqlListener is a complete listener for a parse tree produced by FlowSqlParser.
type FlowSqlListener interface {
	antlr.ParseTreeListener

	// EnterSql is called when entering the sql production.
	EnterSql(c *SqlContext)

	// EnterWhere is called when entering the where production.
	EnterWhere(c *WhereContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterParens is called when entering the parens production.
	EnterParens(c *ParensContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterFields is called when entering the fields production.
	EnterFields(c *FieldsContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterOrderby is called when entering the orderby production.
	EnterOrderby(c *OrderbyContext)

	// ExitSql is called when exiting the sql production.
	ExitSql(c *SqlContext)

	// ExitWhere is called when exiting the where production.
	ExitWhere(c *WhereContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitParens is called when exiting the parens production.
	ExitParens(c *ParensContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitFields is called when exiting the fields production.
	ExitFields(c *FieldsContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitOrderby is called when exiting the orderby production.
	ExitOrderby(c *OrderbyContext)
}
