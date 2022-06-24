// Code generated from D:/nextgen/golang/flow-sql\FlowSql.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FlowSql

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFlowSqlListener is a complete listener for a parse tree produced by FlowSqlParser.
type BaseFlowSqlListener struct{}

var _ FlowSqlListener = &BaseFlowSqlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFlowSqlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFlowSqlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFlowSqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFlowSqlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSql is called when production sql is entered.
func (s *BaseFlowSqlListener) EnterSql(ctx *SqlContext) {}

// ExitSql is called when production sql is exited.
func (s *BaseFlowSqlListener) ExitSql(ctx *SqlContext) {}

// EnterWhere is called when production where is entered.
func (s *BaseFlowSqlListener) EnterWhere(ctx *WhereContext) {}

// ExitWhere is called when production where is exited.
func (s *BaseFlowSqlListener) ExitWhere(ctx *WhereContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseFlowSqlListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseFlowSqlListener) ExitNumber(ctx *NumberContext) {}

// EnterParens is called when production parens is entered.
func (s *BaseFlowSqlListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production parens is exited.
func (s *BaseFlowSqlListener) ExitParens(ctx *ParensContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseFlowSqlListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseFlowSqlListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseFlowSqlListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseFlowSqlListener) ExitAddSub(ctx *AddSubContext) {}

// EnterId is called when production id is entered.
func (s *BaseFlowSqlListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseFlowSqlListener) ExitId(ctx *IdContext) {}

// EnterField is called when production field is entered.
func (s *BaseFlowSqlListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseFlowSqlListener) ExitField(ctx *FieldContext) {}

// EnterAlias is called when production alias is entered.
func (s *BaseFlowSqlListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BaseFlowSqlListener) ExitAlias(ctx *AliasContext) {}

// EnterOrderby is called when production orderby is entered.
func (s *BaseFlowSqlListener) EnterOrderby(ctx *OrderbyContext) {}

// ExitOrderby is called when production orderby is exited.
func (s *BaseFlowSqlListener) ExitOrderby(ctx *OrderbyContext) {}
