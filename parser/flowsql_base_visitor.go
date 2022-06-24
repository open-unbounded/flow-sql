// Code generated from D:/nextgen/golang/flow-sql\FlowSql.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FlowSql

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseFlowSqlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFlowSqlVisitor) VisitSql(ctx *SqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowSqlVisitor) VisitWhere(ctx *WhereContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowSqlVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowSqlVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowSqlVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowSqlVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowSqlVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowSqlVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowSqlVisitor) VisitFields(ctx *FieldsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowSqlVisitor) VisitAlias(ctx *AliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowSqlVisitor) VisitOrderby(ctx *OrderbyContext) interface{} {
	return v.VisitChildren(ctx)
}
