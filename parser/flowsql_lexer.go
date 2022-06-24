// Code generated from D:/nextgen/golang/flow-sql\FlowSql.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type FlowSqlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var flowsqllexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func flowsqllexerLexerInit() {
	staticData := &flowsqllexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "';'", "'*'", "'/'", "'-'", "'+'", "'('", "')'", "','",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "SELECT", "FROM", "WHERE", "ORDER_BY",
		"AND", "AS", "ID", "NUMBER", "CP", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "SELECT",
		"FROM", "WHERE", "ORDER_BY", "AND", "AS", "LETTER", "ID", "INT", "NUMBER",
		"CP", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 18, 142, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 4, 11, 82, 8, 11, 11, 11, 12, 11, 83, 1, 11,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 15, 1, 15, 5, 15, 100, 8, 15, 10, 15, 12, 15, 103, 9, 15, 1, 16,
		1, 16, 1, 16, 5, 16, 108, 8, 16, 10, 16, 12, 16, 111, 9, 16, 3, 16, 113,
		8, 16, 1, 17, 3, 17, 116, 8, 17, 1, 17, 1, 17, 1, 17, 4, 17, 121, 8, 17,
		11, 17, 12, 17, 122, 3, 17, 125, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 3, 18, 134, 8, 18, 1, 19, 4, 19, 137, 8, 19, 11, 19,
		12, 19, 138, 1, 19, 1, 19, 0, 0, 20, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 0,
		31, 15, 33, 0, 35, 16, 37, 17, 39, 18, 1, 0, 22, 2, 0, 83, 83, 115, 115,
		2, 0, 69, 69, 101, 101, 2, 0, 76, 76, 108, 108, 2, 0, 67, 67, 99, 99, 2,
		0, 84, 84, 116, 116, 2, 0, 70, 70, 102, 102, 2, 0, 82, 82, 114, 114, 2,
		0, 79, 79, 111, 111, 2, 0, 77, 77, 109, 109, 2, 0, 87, 87, 119, 119, 2,
		0, 72, 72, 104, 104, 2, 0, 68, 68, 100, 100, 2, 0, 66, 66, 98, 98, 2, 0,
		89, 89, 121, 121, 2, 0, 65, 65, 97, 97, 2, 0, 78, 78, 110, 110, 3, 0, 65,
		90, 95, 95, 97, 122, 2, 0, 65, 90, 97, 122, 3, 0, 48, 57, 65, 90, 97, 122,
		1, 0, 49, 57, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 150, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1,
		0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17,
		1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0,
		25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0,
		0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 1, 41, 1, 0, 0, 0, 3, 43, 1, 0, 0,
		0, 5, 45, 1, 0, 0, 0, 7, 47, 1, 0, 0, 0, 9, 49, 1, 0, 0, 0, 11, 51, 1,
		0, 0, 0, 13, 53, 1, 0, 0, 0, 15, 55, 1, 0, 0, 0, 17, 57, 1, 0, 0, 0, 19,
		64, 1, 0, 0, 0, 21, 69, 1, 0, 0, 0, 23, 75, 1, 0, 0, 0, 25, 88, 1, 0, 0,
		0, 27, 92, 1, 0, 0, 0, 29, 95, 1, 0, 0, 0, 31, 97, 1, 0, 0, 0, 33, 112,
		1, 0, 0, 0, 35, 115, 1, 0, 0, 0, 37, 133, 1, 0, 0, 0, 39, 136, 1, 0, 0,
		0, 41, 42, 5, 59, 0, 0, 42, 2, 1, 0, 0, 0, 43, 44, 5, 42, 0, 0, 44, 4,
		1, 0, 0, 0, 45, 46, 5, 47, 0, 0, 46, 6, 1, 0, 0, 0, 47, 48, 5, 45, 0, 0,
		48, 8, 1, 0, 0, 0, 49, 50, 5, 43, 0, 0, 50, 10, 1, 0, 0, 0, 51, 52, 5,
		40, 0, 0, 52, 12, 1, 0, 0, 0, 53, 54, 5, 41, 0, 0, 54, 14, 1, 0, 0, 0,
		55, 56, 5, 44, 0, 0, 56, 16, 1, 0, 0, 0, 57, 58, 7, 0, 0, 0, 58, 59, 7,
		1, 0, 0, 59, 60, 7, 2, 0, 0, 60, 61, 7, 1, 0, 0, 61, 62, 7, 3, 0, 0, 62,
		63, 7, 4, 0, 0, 63, 18, 1, 0, 0, 0, 64, 65, 7, 5, 0, 0, 65, 66, 7, 6, 0,
		0, 66, 67, 7, 7, 0, 0, 67, 68, 7, 8, 0, 0, 68, 20, 1, 0, 0, 0, 69, 70,
		7, 9, 0, 0, 70, 71, 7, 10, 0, 0, 71, 72, 7, 1, 0, 0, 72, 73, 7, 6, 0, 0,
		73, 74, 7, 1, 0, 0, 74, 22, 1, 0, 0, 0, 75, 76, 7, 7, 0, 0, 76, 77, 7,
		6, 0, 0, 77, 78, 7, 11, 0, 0, 78, 79, 7, 1, 0, 0, 79, 81, 7, 6, 0, 0, 80,
		82, 5, 32, 0, 0, 81, 80, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 81, 1, 0,
		0, 0, 83, 84, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 86, 7, 12, 0, 0, 86,
		87, 7, 13, 0, 0, 87, 24, 1, 0, 0, 0, 88, 89, 7, 14, 0, 0, 89, 90, 7, 15,
		0, 0, 90, 91, 7, 11, 0, 0, 91, 26, 1, 0, 0, 0, 92, 93, 7, 14, 0, 0, 93,
		94, 7, 0, 0, 0, 94, 28, 1, 0, 0, 0, 95, 96, 7, 16, 0, 0, 96, 30, 1, 0,
		0, 0, 97, 101, 7, 17, 0, 0, 98, 100, 7, 18, 0, 0, 99, 98, 1, 0, 0, 0, 100,
		103, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 32, 1,
		0, 0, 0, 103, 101, 1, 0, 0, 0, 104, 113, 5, 48, 0, 0, 105, 109, 7, 19,
		0, 0, 106, 108, 7, 20, 0, 0, 107, 106, 1, 0, 0, 0, 108, 111, 1, 0, 0, 0,
		109, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111,
		109, 1, 0, 0, 0, 112, 104, 1, 0, 0, 0, 112, 105, 1, 0, 0, 0, 113, 34, 1,
		0, 0, 0, 114, 116, 5, 45, 0, 0, 115, 114, 1, 0, 0, 0, 115, 116, 1, 0, 0,
		0, 116, 117, 1, 0, 0, 0, 117, 124, 3, 33, 16, 0, 118, 120, 5, 46, 0, 0,
		119, 121, 3, 33, 16, 0, 120, 119, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122,
		120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 125, 1, 0, 0, 0, 124, 118,
		1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 36, 1, 0, 0, 0, 126, 134, 2, 60,
		62, 0, 127, 128, 5, 62, 0, 0, 128, 134, 5, 61, 0, 0, 129, 130, 5, 60, 0,
		0, 130, 134, 5, 61, 0, 0, 131, 132, 5, 33, 0, 0, 132, 134, 5, 61, 0, 0,
		133, 126, 1, 0, 0, 0, 133, 127, 1, 0, 0, 0, 133, 129, 1, 0, 0, 0, 133,
		131, 1, 0, 0, 0, 134, 38, 1, 0, 0, 0, 135, 137, 7, 21, 0, 0, 136, 135,
		1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0,
		0, 0, 139, 140, 1, 0, 0, 0, 140, 141, 6, 19, 0, 0, 141, 40, 1, 0, 0, 0,
		10, 0, 83, 101, 109, 112, 115, 122, 124, 133, 138, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// FlowSqlLexerInit initializes any static state used to implement FlowSqlLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewFlowSqlLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func FlowSqlLexerInit() {
	staticData := &flowsqllexerLexerStaticData
	staticData.once.Do(flowsqllexerLexerInit)
}

// NewFlowSqlLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewFlowSqlLexer(input antlr.CharStream) *FlowSqlLexer {
	FlowSqlLexerInit()
	l := new(FlowSqlLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &flowsqllexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "FlowSql.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// FlowSqlLexer tokens.
const (
	FlowSqlLexerT__0     = 1
	FlowSqlLexerT__1     = 2
	FlowSqlLexerT__2     = 3
	FlowSqlLexerT__3     = 4
	FlowSqlLexerT__4     = 5
	FlowSqlLexerT__5     = 6
	FlowSqlLexerT__6     = 7
	FlowSqlLexerT__7     = 8
	FlowSqlLexerSELECT   = 9
	FlowSqlLexerFROM     = 10
	FlowSqlLexerWHERE    = 11
	FlowSqlLexerORDER_BY = 12
	FlowSqlLexerAND      = 13
	FlowSqlLexerAS       = 14
	FlowSqlLexerID       = 15
	FlowSqlLexerNUMBER   = 16
	FlowSqlLexerCP       = 17
	FlowSqlLexerWS       = 18
)
