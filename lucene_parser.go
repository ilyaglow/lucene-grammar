// Code generated from Lucene.g4 by ANTLR 4.7.1. DO NOT EDIT.

package lucene // Lucene
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 44, 259,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 3,
	2, 5, 2, 58, 10, 2, 3, 2, 3, 2, 5, 2, 62, 10, 2, 3, 3, 3, 3, 5, 3, 66,
	10, 3, 3, 3, 7, 3, 69, 10, 3, 12, 3, 14, 3, 72, 11, 3, 3, 4, 3, 4, 3, 4,
	3, 4, 7, 4, 78, 10, 4, 12, 4, 14, 4, 81, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5,
	7, 5, 87, 10, 5, 12, 5, 14, 5, 90, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6,
	96, 10, 6, 12, 6, 14, 6, 99, 11, 6, 3, 7, 5, 7, 102, 10, 7, 3, 7, 5, 7,
	105, 10, 7, 3, 7, 3, 7, 3, 7, 5, 7, 110, 10, 7, 3, 7, 3, 7, 5, 7, 114,
	10, 7, 3, 7, 5, 7, 117, 10, 7, 3, 7, 5, 7, 120, 10, 7, 3, 8, 5, 8, 123,
	10, 8, 3, 8, 3, 8, 3, 8, 5, 8, 128, 10, 8, 3, 8, 5, 8, 131, 10, 8, 3, 8,
	5, 8, 134, 10, 8, 3, 8, 3, 8, 5, 8, 138, 10, 8, 5, 8, 140, 10, 8, 3, 9,
	3, 9, 3, 9, 5, 9, 145, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 5, 10, 155, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3,
	12, 5, 12, 163, 10, 12, 3, 12, 3, 12, 5, 12, 167, 10, 12, 3, 12, 5, 12,
	170, 10, 12, 3, 12, 5, 12, 173, 10, 12, 3, 12, 3, 12, 5, 12, 177, 10, 12,
	5, 12, 179, 10, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 5, 14, 191, 10, 14, 3, 15, 3, 15, 3, 15, 5, 15, 196,
	10, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19,
	3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 5, 21, 212, 10, 21, 3, 21, 3, 21, 5,
	21, 216, 10, 21, 5, 21, 218, 10, 21, 3, 22, 3, 22, 5, 22, 222, 10, 22,
	3, 23, 3, 23, 5, 23, 226, 10, 23, 3, 24, 5, 24, 229, 10, 24, 3, 24, 3,
	24, 5, 24, 233, 10, 24, 3, 24, 3, 24, 5, 24, 237, 10, 24, 3, 24, 5, 24,
	240, 10, 24, 3, 25, 5, 25, 243, 10, 25, 3, 25, 3, 25, 3, 26, 5, 26, 248,
	10, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 6, 28, 255, 10, 28, 13, 28,
	14, 28, 256, 3, 28, 2, 2, 29, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 2, 6, 4, 2,
	5, 5, 12, 12, 4, 2, 6, 6, 13, 13, 4, 2, 23, 23, 25, 25, 3, 2, 8, 9, 2,
	282, 2, 57, 3, 2, 2, 2, 4, 63, 3, 2, 2, 2, 6, 73, 3, 2, 2, 2, 8, 82, 3,
	2, 2, 2, 10, 91, 3, 2, 2, 2, 12, 119, 3, 2, 2, 2, 14, 139, 3, 2, 2, 2,
	16, 141, 3, 2, 2, 2, 18, 154, 3, 2, 2, 2, 20, 156, 3, 2, 2, 2, 22, 160,
	3, 2, 2, 2, 24, 182, 3, 2, 2, 2, 26, 190, 3, 2, 2, 2, 28, 192, 3, 2, 2,
	2, 30, 199, 3, 2, 2, 2, 32, 201, 3, 2, 2, 2, 34, 203, 3, 2, 2, 2, 36, 205,
	3, 2, 2, 2, 38, 207, 3, 2, 2, 2, 40, 217, 3, 2, 2, 2, 42, 219, 3, 2, 2,
	2, 44, 223, 3, 2, 2, 2, 46, 239, 3, 2, 2, 2, 48, 242, 3, 2, 2, 2, 50, 247,
	3, 2, 2, 2, 52, 251, 3, 2, 2, 2, 54, 254, 3, 2, 2, 2, 56, 58, 5, 54, 28,
	2, 57, 56, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 61,
	5, 4, 3, 2, 60, 62, 5, 54, 28, 2, 61, 60, 3, 2, 2, 2, 61, 62, 3, 2, 2,
	2, 62, 3, 3, 2, 2, 2, 63, 70, 5, 6, 4, 2, 64, 66, 5, 54, 28, 2, 65, 64,
	3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 69, 5, 6, 4, 2,
	68, 65, 3, 2, 2, 2, 69, 72, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 70, 71, 3,
	2, 2, 2, 71, 5, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 73, 79, 5, 8, 5, 2, 74,
	75, 5, 50, 26, 2, 75, 76, 5, 8, 5, 2, 76, 78, 3, 2, 2, 2, 77, 74, 3, 2,
	2, 2, 78, 81, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 7,
	3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 82, 88, 5, 10, 6, 2, 83, 84, 5, 48, 25,
	2, 84, 85, 5, 10, 6, 2, 85, 87, 3, 2, 2, 2, 86, 83, 3, 2, 2, 2, 87, 90,
	3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 9, 3, 2, 2, 2,
	90, 88, 3, 2, 2, 2, 91, 97, 5, 12, 7, 2, 92, 93, 5, 46, 24, 2, 93, 94,
	5, 12, 7, 2, 94, 96, 3, 2, 2, 2, 95, 92, 3, 2, 2, 2, 96, 99, 3, 2, 2, 2,
	97, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 11, 3, 2, 2, 2, 99, 97, 3,
	2, 2, 2, 100, 102, 5, 54, 28, 2, 101, 100, 3, 2, 2, 2, 101, 102, 3, 2,
	2, 2, 102, 104, 3, 2, 2, 2, 103, 105, 5, 38, 20, 2, 104, 103, 3, 2, 2,
	2, 104, 105, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 107, 7, 3, 2, 2, 107,
	109, 5, 4, 3, 2, 108, 110, 5, 54, 28, 2, 109, 108, 3, 2, 2, 2, 109, 110,
	3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 113, 7, 4, 2, 2, 112, 114, 5, 40,
	21, 2, 113, 112, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 120, 3, 2, 2, 2,
	115, 117, 5, 54, 28, 2, 116, 115, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117,
	118, 3, 2, 2, 2, 118, 120, 5, 14, 8, 2, 119, 101, 3, 2, 2, 2, 119, 116,
	3, 2, 2, 2, 120, 13, 3, 2, 2, 2, 121, 123, 5, 38, 20, 2, 122, 121, 3, 2,
	2, 2, 122, 123, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 125, 5, 16, 9, 2,
	125, 127, 5, 28, 15, 2, 126, 128, 5, 40, 21, 2, 127, 126, 3, 2, 2, 2, 127,
	128, 3, 2, 2, 2, 128, 140, 3, 2, 2, 2, 129, 131, 5, 38, 20, 2, 130, 129,
	3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 133, 3, 2, 2, 2, 132, 134, 5, 16,
	9, 2, 133, 132, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2,
	135, 137, 5, 18, 10, 2, 136, 138, 5, 40, 21, 2, 137, 136, 3, 2, 2, 2, 137,
	138, 3, 2, 2, 2, 138, 140, 3, 2, 2, 2, 139, 122, 3, 2, 2, 2, 139, 130,
	3, 2, 2, 2, 140, 15, 3, 2, 2, 2, 141, 142, 7, 25, 2, 2, 142, 144, 7, 7,
	2, 2, 143, 145, 5, 54, 28, 2, 144, 143, 3, 2, 2, 2, 144, 145, 3, 2, 2,
	2, 145, 17, 3, 2, 2, 2, 146, 155, 5, 24, 13, 2, 147, 155, 5, 30, 16, 2,
	148, 155, 5, 32, 17, 2, 149, 155, 5, 36, 19, 2, 150, 155, 5, 34, 18, 2,
	151, 155, 7, 11, 2, 2, 152, 155, 5, 20, 11, 2, 153, 155, 7, 10, 2, 2, 154,
	146, 3, 2, 2, 2, 154, 147, 3, 2, 2, 2, 154, 148, 3, 2, 2, 2, 154, 149,
	3, 2, 2, 2, 154, 150, 3, 2, 2, 2, 154, 151, 3, 2, 2, 2, 154, 152, 3, 2,
	2, 2, 154, 153, 3, 2, 2, 2, 155, 19, 3, 2, 2, 2, 156, 157, 7, 10, 2, 2,
	157, 158, 7, 7, 2, 2, 158, 159, 7, 10, 2, 2, 159, 21, 3, 2, 2, 2, 160,
	162, 9, 2, 2, 2, 161, 163, 5, 54, 28, 2, 162, 161, 3, 2, 2, 2, 162, 163,
	3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 166, 5, 26, 14, 2, 165, 167, 5,
	54, 28, 2, 166, 165, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 178, 3, 2,
	2, 2, 168, 170, 7, 18, 2, 2, 169, 168, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2,
	170, 172, 3, 2, 2, 2, 171, 173, 5, 54, 28, 2, 172, 171, 3, 2, 2, 2, 172,
	173, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 176, 5, 26, 14, 2, 175, 177,
	5, 54, 28, 2, 176, 175, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 179, 3,
	2, 2, 2, 178, 169, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 180, 3, 2, 2,
	2, 180, 181, 9, 3, 2, 2, 181, 23, 3, 2, 2, 2, 182, 183, 5, 22, 12, 2, 183,
	25, 3, 2, 2, 2, 184, 191, 5, 32, 17, 2, 185, 191, 5, 36, 19, 2, 186, 191,
	5, 34, 18, 2, 187, 191, 5, 52, 27, 2, 188, 191, 5, 30, 16, 2, 189, 191,
	7, 10, 2, 2, 190, 184, 3, 2, 2, 2, 190, 185, 3, 2, 2, 2, 190, 186, 3, 2,
	2, 2, 190, 187, 3, 2, 2, 2, 190, 188, 3, 2, 2, 2, 190, 189, 3, 2, 2, 2,
	191, 27, 3, 2, 2, 2, 192, 193, 7, 3, 2, 2, 193, 195, 5, 4, 3, 2, 194, 196,
	5, 54, 28, 2, 195, 194, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 197, 3,
	2, 2, 2, 197, 198, 7, 4, 2, 2, 198, 29, 3, 2, 2, 2, 199, 200, 9, 4, 2,
	2, 200, 31, 3, 2, 2, 2, 201, 202, 7, 26, 2, 2, 202, 33, 3, 2, 2, 2, 203,
	204, 7, 28, 2, 2, 204, 35, 3, 2, 2, 2, 205, 206, 7, 27, 2, 2, 206, 37,
	3, 2, 2, 2, 207, 208, 9, 5, 2, 2, 208, 39, 3, 2, 2, 2, 209, 211, 5, 42,
	22, 2, 210, 212, 5, 44, 23, 2, 211, 210, 3, 2, 2, 2, 211, 212, 3, 2, 2,
	2, 212, 218, 3, 2, 2, 2, 213, 215, 5, 44, 23, 2, 214, 216, 5, 42, 22, 2,
	215, 214, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 218, 3, 2, 2, 2, 217,
	209, 3, 2, 2, 2, 217, 213, 3, 2, 2, 2, 218, 41, 3, 2, 2, 2, 219, 221, 7,
	14, 2, 2, 220, 222, 7, 23, 2, 2, 221, 220, 3, 2, 2, 2, 221, 222, 3, 2,
	2, 2, 222, 43, 3, 2, 2, 2, 223, 225, 7, 15, 2, 2, 224, 226, 7, 23, 2, 2,
	225, 224, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2, 226, 45, 3, 2, 2, 2, 227, 229,
	5, 54, 28, 2, 228, 227, 3, 2, 2, 2, 228, 229, 3, 2, 2, 2, 229, 230, 3,
	2, 2, 2, 230, 232, 7, 19, 2, 2, 231, 233, 5, 54, 28, 2, 232, 231, 3, 2,
	2, 2, 232, 233, 3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234, 240, 7, 21, 2, 2,
	235, 237, 5, 54, 28, 2, 236, 235, 3, 2, 2, 2, 236, 237, 3, 2, 2, 2, 237,
	238, 3, 2, 2, 2, 238, 240, 7, 21, 2, 2, 239, 228, 3, 2, 2, 2, 239, 236,
	3, 2, 2, 2, 240, 47, 3, 2, 2, 2, 241, 243, 5, 54, 28, 2, 242, 241, 3, 2,
	2, 2, 242, 243, 3, 2, 2, 2, 243, 244, 3, 2, 2, 2, 244, 245, 7, 19, 2, 2,
	245, 49, 3, 2, 2, 2, 246, 248, 5, 54, 28, 2, 247, 246, 3, 2, 2, 2, 247,
	248, 3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 249, 250, 7, 20, 2, 2, 250, 51,
	3, 2, 2, 2, 251, 252, 7, 24, 2, 2, 252, 53, 3, 2, 2, 2, 253, 255, 7, 22,
	2, 2, 254, 253, 3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256, 254, 3, 2, 2, 2,
	256, 257, 3, 2, 2, 2, 257, 55, 3, 2, 2, 2, 43, 57, 61, 65, 70, 79, 88,
	97, 101, 104, 109, 113, 116, 119, 122, 127, 130, 133, 137, 139, 144, 154,
	162, 166, 169, 172, 176, 178, 190, 195, 211, 215, 217, 221, 225, 228, 232,
	236, 239, 242, 247, 256,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'['", "']'", "':'", "'+'", "", "'*'", "", "'{'", "'}'",
	"", "", "'\"'", "'''", "'TO'",
}
var symbolicNames = []string{
	"", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "COLON", "PLUS", "MINUS", "STAR",
	"QMARK", "LCURLY", "RCURLY", "CARAT", "TILDE", "DQUOTE", "SQUOTE", "TO",
	"AND", "OR", "NOT", "WS", "NUMBER", "DATE_TOKEN", "TERM_NORMAL", "TERM_TRUNCATED",
	"PHRASE", "PHRASE_ANYTHING", "OPERATOR", "ATOM", "MODIFIER", "TMODIFIER",
	"CLAUSE", "FIELD", "FUZZY", "BOOST", "QNORMAL", "QPHRASE", "QPHRASETRUNC",
	"QTRUNCATED", "QRANGEIN", "QRANGEEX", "QANYTHING", "QDATE",
}

var ruleNames = []string{
	"mainQ", "clauseDefault", "clauseOr", "clauseAnd", "clauseNot", "clauseBasic",
	"atom", "field", "value", "anything", "two_sided_range_term", "range_term",
	"range_value", "multi_value", "normal", "truncated", "quoted_truncated",
	"quoted", "modifier", "term_modifier", "boost", "fuzzy", "not_", "and_",
	"or_", "date", "sep",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type LuceneParser struct {
	*antlr.BaseParser
}

func NewLuceneParser(input antlr.TokenStream) *LuceneParser {
	this := new(LuceneParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Lucene.g4"

	return this
}

// LuceneParser tokens.
const (
	LuceneParserEOF             = antlr.TokenEOF
	LuceneParserLPAREN          = 1
	LuceneParserRPAREN          = 2
	LuceneParserLBRACK          = 3
	LuceneParserRBRACK          = 4
	LuceneParserCOLON           = 5
	LuceneParserPLUS            = 6
	LuceneParserMINUS           = 7
	LuceneParserSTAR            = 8
	LuceneParserQMARK           = 9
	LuceneParserLCURLY          = 10
	LuceneParserRCURLY          = 11
	LuceneParserCARAT           = 12
	LuceneParserTILDE           = 13
	LuceneParserDQUOTE          = 14
	LuceneParserSQUOTE          = 15
	LuceneParserTO              = 16
	LuceneParserAND             = 17
	LuceneParserOR              = 18
	LuceneParserNOT             = 19
	LuceneParserWS              = 20
	LuceneParserNUMBER          = 21
	LuceneParserDATE_TOKEN      = 22
	LuceneParserTERM_NORMAL     = 23
	LuceneParserTERM_TRUNCATED  = 24
	LuceneParserPHRASE          = 25
	LuceneParserPHRASE_ANYTHING = 26
	LuceneParserOPERATOR        = 27
	LuceneParserATOM            = 28
	LuceneParserMODIFIER        = 29
	LuceneParserTMODIFIER       = 30
	LuceneParserCLAUSE          = 31
	LuceneParserFIELD           = 32
	LuceneParserFUZZY           = 33
	LuceneParserBOOST           = 34
	LuceneParserQNORMAL         = 35
	LuceneParserQPHRASE         = 36
	LuceneParserQPHRASETRUNC    = 37
	LuceneParserQTRUNCATED      = 38
	LuceneParserQRANGEIN        = 39
	LuceneParserQRANGEEX        = 40
	LuceneParserQANYTHING       = 41
	LuceneParserQDATE           = 42
)

// LuceneParser rules.
const (
	LuceneParserRULE_mainQ                = 0
	LuceneParserRULE_clauseDefault        = 1
	LuceneParserRULE_clauseOr             = 2
	LuceneParserRULE_clauseAnd            = 3
	LuceneParserRULE_clauseNot            = 4
	LuceneParserRULE_clauseBasic          = 5
	LuceneParserRULE_atom                 = 6
	LuceneParserRULE_field                = 7
	LuceneParserRULE_value                = 8
	LuceneParserRULE_anything             = 9
	LuceneParserRULE_two_sided_range_term = 10
	LuceneParserRULE_range_term           = 11
	LuceneParserRULE_range_value          = 12
	LuceneParserRULE_multi_value          = 13
	LuceneParserRULE_normal               = 14
	LuceneParserRULE_truncated            = 15
	LuceneParserRULE_quoted_truncated     = 16
	LuceneParserRULE_quoted               = 17
	LuceneParserRULE_modifier             = 18
	LuceneParserRULE_term_modifier        = 19
	LuceneParserRULE_boost                = 20
	LuceneParserRULE_fuzzy                = 21
	LuceneParserRULE_not_                 = 22
	LuceneParserRULE_and_                 = 23
	LuceneParserRULE_or_                  = 24
	LuceneParserRULE_date                 = 25
	LuceneParserRULE_sep                  = 26
)

// IMainQContext is an interface to support dynamic dispatch.
type IMainQContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetClause returns the clause rule contexts.
	GetClause() IClauseDefaultContext

	// SetClause sets the clause rule contexts.
	SetClause(IClauseDefaultContext)

	// IsMainQContext differentiates from other interfaces.
	IsMainQContext()
}

type MainQContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	clause IClauseDefaultContext
}

func NewEmptyMainQContext() *MainQContext {
	var p = new(MainQContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_mainQ
	return p
}

func (*MainQContext) IsMainQContext() {}

func NewMainQContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainQContext {
	var p = new(MainQContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_mainQ

	return p
}

func (s *MainQContext) GetParser() antlr.Parser { return s.parser }

func (s *MainQContext) GetClause() IClauseDefaultContext { return s.clause }

func (s *MainQContext) SetClause(v IClauseDefaultContext) { s.clause = v }

func (s *MainQContext) ClauseDefault() IClauseDefaultContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClauseDefaultContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClauseDefaultContext)
}

func (s *MainQContext) AllSep() []ISepContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISepContext)(nil)).Elem())
	var tst = make([]ISepContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISepContext)
		}
	}

	return tst
}

func (s *MainQContext) Sep(i int) ISepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISepContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISepContext)
}

func (s *MainQContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainQContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainQContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterMainQ(s)
	}
}

func (s *MainQContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitMainQ(s)
	}
}

func (p *LuceneParser) MainQ() (localctx IMainQContext) {
	localctx = NewMainQContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LuceneParserRULE_mainQ)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(55)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(54)
			p.Sep()
		}

	}
	{
		p.SetState(57)

		var _x = p.ClauseDefault()

		localctx.(*MainQContext).clause = _x
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuceneParserWS {
		{
			p.SetState(58)
			p.Sep()
		}

	}

	return localctx
}

// IClauseDefaultContext is an interface to support dynamic dispatch.
type IClauseDefaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClauseDefaultContext differentiates from other interfaces.
	IsClauseDefaultContext()
}

type ClauseDefaultContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClauseDefaultContext() *ClauseDefaultContext {
	var p = new(ClauseDefaultContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_clauseDefault
	return p
}

func (*ClauseDefaultContext) IsClauseDefaultContext() {}

func NewClauseDefaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClauseDefaultContext {
	var p = new(ClauseDefaultContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_clauseDefault

	return p
}

func (s *ClauseDefaultContext) GetParser() antlr.Parser { return s.parser }

func (s *ClauseDefaultContext) AllClauseOr() []IClauseOrContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClauseOrContext)(nil)).Elem())
	var tst = make([]IClauseOrContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClauseOrContext)
		}
	}

	return tst
}

func (s *ClauseDefaultContext) ClauseOr(i int) IClauseOrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClauseOrContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClauseOrContext)
}

func (s *ClauseDefaultContext) AllSep() []ISepContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISepContext)(nil)).Elem())
	var tst = make([]ISepContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISepContext)
		}
	}

	return tst
}

func (s *ClauseDefaultContext) Sep(i int) ISepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISepContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISepContext)
}

func (s *ClauseDefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClauseDefaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClauseDefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterClauseDefault(s)
	}
}

func (s *ClauseDefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitClauseDefault(s)
	}
}

func (p *LuceneParser) ClauseDefault() (localctx IClauseDefaultContext) {
	localctx = NewClauseDefaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LuceneParserRULE_clauseDefault)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.ClauseOr()
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(63)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(62)
					p.Sep()
				}

			}
			{
				p.SetState(65)
				p.ClauseOr()
			}

		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IClauseOrContext is an interface to support dynamic dispatch.
type IClauseOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClauseOrContext differentiates from other interfaces.
	IsClauseOrContext()
}

type ClauseOrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClauseOrContext() *ClauseOrContext {
	var p = new(ClauseOrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_clauseOr
	return p
}

func (*ClauseOrContext) IsClauseOrContext() {}

func NewClauseOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClauseOrContext {
	var p = new(ClauseOrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_clauseOr

	return p
}

func (s *ClauseOrContext) GetParser() antlr.Parser { return s.parser }

func (s *ClauseOrContext) AllClauseAnd() []IClauseAndContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClauseAndContext)(nil)).Elem())
	var tst = make([]IClauseAndContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClauseAndContext)
		}
	}

	return tst
}

func (s *ClauseOrContext) ClauseAnd(i int) IClauseAndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClauseAndContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClauseAndContext)
}

func (s *ClauseOrContext) AllOr_() []IOr_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOr_Context)(nil)).Elem())
	var tst = make([]IOr_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOr_Context)
		}
	}

	return tst
}

func (s *ClauseOrContext) Or_(i int) IOr_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOr_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOr_Context)
}

func (s *ClauseOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClauseOrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClauseOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterClauseOr(s)
	}
}

func (s *ClauseOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitClauseOr(s)
	}
}

func (p *LuceneParser) ClauseOr() (localctx IClauseOrContext) {
	localctx = NewClauseOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LuceneParserRULE_clauseOr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.ClauseAnd()
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(72)
				p.Or_()
			}
			{
				p.SetState(73)
				p.ClauseAnd()
			}

		}
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IClauseAndContext is an interface to support dynamic dispatch.
type IClauseAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClauseAndContext differentiates from other interfaces.
	IsClauseAndContext()
}

type ClauseAndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClauseAndContext() *ClauseAndContext {
	var p = new(ClauseAndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_clauseAnd
	return p
}

func (*ClauseAndContext) IsClauseAndContext() {}

func NewClauseAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClauseAndContext {
	var p = new(ClauseAndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_clauseAnd

	return p
}

func (s *ClauseAndContext) GetParser() antlr.Parser { return s.parser }

func (s *ClauseAndContext) AllClauseNot() []IClauseNotContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClauseNotContext)(nil)).Elem())
	var tst = make([]IClauseNotContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClauseNotContext)
		}
	}

	return tst
}

func (s *ClauseAndContext) ClauseNot(i int) IClauseNotContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClauseNotContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClauseNotContext)
}

func (s *ClauseAndContext) AllAnd_() []IAnd_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnd_Context)(nil)).Elem())
	var tst = make([]IAnd_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnd_Context)
		}
	}

	return tst
}

func (s *ClauseAndContext) And_(i int) IAnd_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnd_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnd_Context)
}

func (s *ClauseAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClauseAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClauseAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterClauseAnd(s)
	}
}

func (s *ClauseAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitClauseAnd(s)
	}
}

func (p *LuceneParser) ClauseAnd() (localctx IClauseAndContext) {
	localctx = NewClauseAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LuceneParserRULE_clauseAnd)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.ClauseNot()
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(81)
				p.And_()
			}
			{
				p.SetState(82)
				p.ClauseNot()
			}

		}
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IClauseNotContext is an interface to support dynamic dispatch.
type IClauseNotContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClauseNotContext differentiates from other interfaces.
	IsClauseNotContext()
}

type ClauseNotContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClauseNotContext() *ClauseNotContext {
	var p = new(ClauseNotContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_clauseNot
	return p
}

func (*ClauseNotContext) IsClauseNotContext() {}

func NewClauseNotContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClauseNotContext {
	var p = new(ClauseNotContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_clauseNot

	return p
}

func (s *ClauseNotContext) GetParser() antlr.Parser { return s.parser }

func (s *ClauseNotContext) AllClauseBasic() []IClauseBasicContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClauseBasicContext)(nil)).Elem())
	var tst = make([]IClauseBasicContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClauseBasicContext)
		}
	}

	return tst
}

func (s *ClauseNotContext) ClauseBasic(i int) IClauseBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClauseBasicContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClauseBasicContext)
}

func (s *ClauseNotContext) AllNot_() []INot_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INot_Context)(nil)).Elem())
	var tst = make([]INot_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INot_Context)
		}
	}

	return tst
}

func (s *ClauseNotContext) Not_(i int) INot_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INot_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INot_Context)
}

func (s *ClauseNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClauseNotContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClauseNotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterClauseNot(s)
	}
}

func (s *ClauseNotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitClauseNot(s)
	}
}

func (p *LuceneParser) ClauseNot() (localctx IClauseNotContext) {
	localctx = NewClauseNotContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LuceneParserRULE_clauseNot)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.ClauseBasic()
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(90)
				p.Not_()
			}
			{
				p.SetState(91)
				p.ClauseBasic()
			}

		}
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IClauseBasicContext is an interface to support dynamic dispatch.
type IClauseBasicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClauseBasicContext differentiates from other interfaces.
	IsClauseBasicContext()
}

type ClauseBasicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClauseBasicContext() *ClauseBasicContext {
	var p = new(ClauseBasicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_clauseBasic
	return p
}

func (*ClauseBasicContext) IsClauseBasicContext() {}

func NewClauseBasicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClauseBasicContext {
	var p = new(ClauseBasicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_clauseBasic

	return p
}

func (s *ClauseBasicContext) GetParser() antlr.Parser { return s.parser }

func (s *ClauseBasicContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LuceneParserLPAREN, 0)
}

func (s *ClauseBasicContext) ClauseDefault() IClauseDefaultContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClauseDefaultContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClauseDefaultContext)
}

func (s *ClauseBasicContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LuceneParserRPAREN, 0)
}

func (s *ClauseBasicContext) AllSep() []ISepContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISepContext)(nil)).Elem())
	var tst = make([]ISepContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISepContext)
		}
	}

	return tst
}

func (s *ClauseBasicContext) Sep(i int) ISepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISepContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISepContext)
}

func (s *ClauseBasicContext) Modifier() IModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModifierContext)
}

func (s *ClauseBasicContext) Term_modifier() ITerm_modifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITerm_modifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITerm_modifierContext)
}

func (s *ClauseBasicContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ClauseBasicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClauseBasicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClauseBasicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterClauseBasic(s)
	}
}

func (s *ClauseBasicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitClauseBasic(s)
	}
}

func (p *LuceneParser) ClauseBasic() (localctx IClauseBasicContext) {
	localctx = NewClauseBasicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LuceneParserRULE_clauseBasic)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuceneParserWS {
			{
				p.SetState(98)
				p.Sep()
			}

		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuceneParserPLUS || _la == LuceneParserMINUS {
			{
				p.SetState(101)
				p.Modifier()
			}

		}
		{
			p.SetState(104)
			p.Match(LuceneParserLPAREN)
		}
		{
			p.SetState(105)
			p.ClauseDefault()
		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuceneParserWS {
			{
				p.SetState(106)
				p.Sep()
			}

		}
		{
			p.SetState(109)
			p.Match(LuceneParserRPAREN)
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuceneParserCARAT || _la == LuceneParserTILDE {
			{
				p.SetState(110)
				p.Term_modifier()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuceneParserWS {
			{
				p.SetState(113)
				p.Sep()
			}

		}
		{
			p.SetState(116)
			p.Atom()
		}

	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *AtomContext) Multi_value() IMulti_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMulti_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMulti_valueContext)
}

func (s *AtomContext) Modifier() IModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModifierContext)
}

func (s *AtomContext) Term_modifier() ITerm_modifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITerm_modifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITerm_modifierContext)
}

func (s *AtomContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (p *LuceneParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LuceneParserRULE_atom)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuceneParserPLUS || _la == LuceneParserMINUS {
			{
				p.SetState(119)
				p.Modifier()
			}

		}
		{
			p.SetState(122)
			p.Field()
		}
		{
			p.SetState(123)
			p.Multi_value()
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuceneParserCARAT || _la == LuceneParserTILDE {
			{
				p.SetState(124)
				p.Term_modifier()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuceneParserPLUS || _la == LuceneParserMINUS {
			{
				p.SetState(127)
				p.Modifier()
			}

		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(130)
				p.Field()
			}

		}
		{
			p.SetState(133)
			p.Value()
		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuceneParserCARAT || _la == LuceneParserTILDE {
			{
				p.SetState(134)
				p.Term_modifier()
			}

		}

	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) TERM_NORMAL() antlr.TerminalNode {
	return s.GetToken(LuceneParserTERM_NORMAL, 0)
}

func (s *FieldContext) COLON() antlr.TerminalNode {
	return s.GetToken(LuceneParserCOLON, 0)
}

func (s *FieldContext) Sep() ISepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISepContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISepContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *LuceneParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LuceneParserRULE_field)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(LuceneParserTERM_NORMAL)
	}
	{
		p.SetState(140)
		p.Match(LuceneParserCOLON)
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuceneParserWS {
		{
			p.SetState(141)
			p.Sep()
		}

	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Range_term() IRange_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRange_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRange_termContext)
}

func (s *ValueContext) Normal() INormalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INormalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INormalContext)
}

func (s *ValueContext) Truncated() ITruncatedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITruncatedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITruncatedContext)
}

func (s *ValueContext) Quoted() IQuotedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuotedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuotedContext)
}

func (s *ValueContext) Quoted_truncated() IQuoted_truncatedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuoted_truncatedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuoted_truncatedContext)
}

func (s *ValueContext) QMARK() antlr.TerminalNode {
	return s.GetToken(LuceneParserQMARK, 0)
}

func (s *ValueContext) Anything() IAnythingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnythingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnythingContext)
}

func (s *ValueContext) STAR() antlr.TerminalNode {
	return s.GetToken(LuceneParserSTAR, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *LuceneParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LuceneParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(144)
			p.Range_term()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(145)
			p.Normal()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(146)
			p.Truncated()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(147)
			p.Quoted()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(148)
			p.Quoted_truncated()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(149)
			p.Match(LuceneParserQMARK)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(150)
			p.Anything()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(151)
			p.Match(LuceneParserSTAR)
		}

	}

	return localctx
}

// IAnythingContext is an interface to support dynamic dispatch.
type IAnythingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnythingContext differentiates from other interfaces.
	IsAnythingContext()
}

type AnythingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnythingContext() *AnythingContext {
	var p = new(AnythingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_anything
	return p
}

func (*AnythingContext) IsAnythingContext() {}

func NewAnythingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnythingContext {
	var p = new(AnythingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_anything

	return p
}

func (s *AnythingContext) GetParser() antlr.Parser { return s.parser }

func (s *AnythingContext) AllSTAR() []antlr.TerminalNode {
	return s.GetTokens(LuceneParserSTAR)
}

func (s *AnythingContext) STAR(i int) antlr.TerminalNode {
	return s.GetToken(LuceneParserSTAR, i)
}

func (s *AnythingContext) COLON() antlr.TerminalNode {
	return s.GetToken(LuceneParserCOLON, 0)
}

func (s *AnythingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnythingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnythingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterAnything(s)
	}
}

func (s *AnythingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitAnything(s)
	}
}

func (p *LuceneParser) Anything() (localctx IAnythingContext) {
	localctx = NewAnythingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LuceneParserRULE_anything)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Match(LuceneParserSTAR)
	}
	{
		p.SetState(155)
		p.Match(LuceneParserCOLON)
	}
	{
		p.SetState(156)
		p.Match(LuceneParserSTAR)
	}

	return localctx
}

// ITwo_sided_range_termContext is an interface to support dynamic dispatch.
type ITwo_sided_range_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStart_type returns the start_type token.
	GetStart_type() antlr.Token

	// GetEnd_type returns the end_type token.
	GetEnd_type() antlr.Token

	// SetStart_type sets the start_type token.
	SetStart_type(antlr.Token)

	// SetEnd_type sets the end_type token.
	SetEnd_type(antlr.Token)

	// GetA returns the a rule contexts.
	GetA() IRange_valueContext

	// GetB returns the b rule contexts.
	GetB() IRange_valueContext

	// SetA sets the a rule contexts.
	SetA(IRange_valueContext)

	// SetB sets the b rule contexts.
	SetB(IRange_valueContext)

	// IsTwo_sided_range_termContext differentiates from other interfaces.
	IsTwo_sided_range_termContext()
}

type Two_sided_range_termContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	start_type antlr.Token
	a          IRange_valueContext
	b          IRange_valueContext
	end_type   antlr.Token
}

func NewEmptyTwo_sided_range_termContext() *Two_sided_range_termContext {
	var p = new(Two_sided_range_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_two_sided_range_term
	return p
}

func (*Two_sided_range_termContext) IsTwo_sided_range_termContext() {}

func NewTwo_sided_range_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Two_sided_range_termContext {
	var p = new(Two_sided_range_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_two_sided_range_term

	return p
}

func (s *Two_sided_range_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Two_sided_range_termContext) GetStart_type() antlr.Token { return s.start_type }

func (s *Two_sided_range_termContext) GetEnd_type() antlr.Token { return s.end_type }

func (s *Two_sided_range_termContext) SetStart_type(v antlr.Token) { s.start_type = v }

func (s *Two_sided_range_termContext) SetEnd_type(v antlr.Token) { s.end_type = v }

func (s *Two_sided_range_termContext) GetA() IRange_valueContext { return s.a }

func (s *Two_sided_range_termContext) GetB() IRange_valueContext { return s.b }

func (s *Two_sided_range_termContext) SetA(v IRange_valueContext) { s.a = v }

func (s *Two_sided_range_termContext) SetB(v IRange_valueContext) { s.b = v }

func (s *Two_sided_range_termContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(LuceneParserLBRACK, 0)
}

func (s *Two_sided_range_termContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(LuceneParserLCURLY, 0)
}

func (s *Two_sided_range_termContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(LuceneParserRBRACK, 0)
}

func (s *Two_sided_range_termContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(LuceneParserRCURLY, 0)
}

func (s *Two_sided_range_termContext) AllSep() []ISepContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISepContext)(nil)).Elem())
	var tst = make([]ISepContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISepContext)
		}
	}

	return tst
}

func (s *Two_sided_range_termContext) Sep(i int) ISepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISepContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISepContext)
}

func (s *Two_sided_range_termContext) AllRange_value() []IRange_valueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRange_valueContext)(nil)).Elem())
	var tst = make([]IRange_valueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRange_valueContext)
		}
	}

	return tst
}

func (s *Two_sided_range_termContext) Range_value(i int) IRange_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRange_valueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRange_valueContext)
}

func (s *Two_sided_range_termContext) TO() antlr.TerminalNode {
	return s.GetToken(LuceneParserTO, 0)
}

func (s *Two_sided_range_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Two_sided_range_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Two_sided_range_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterTwo_sided_range_term(s)
	}
}

func (s *Two_sided_range_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitTwo_sided_range_term(s)
	}
}

func (p *LuceneParser) Two_sided_range_term() (localctx ITwo_sided_range_termContext) {
	localctx = NewTwo_sided_range_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LuceneParserRULE_two_sided_range_term)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Two_sided_range_termContext).start_type = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == LuceneParserLBRACK || _la == LuceneParserLCURLY) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Two_sided_range_termContext).start_type = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuceneParserWS {
		{
			p.SetState(159)
			p.Sep()
		}

	}

	{
		p.SetState(162)

		var _x = p.Range_value()

		localctx.(*Two_sided_range_termContext).a = _x
	}

	p.SetState(164)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(163)
			p.Sep()
		}

	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LuceneParserSTAR)|(1<<LuceneParserTO)|(1<<LuceneParserWS)|(1<<LuceneParserNUMBER)|(1<<LuceneParserDATE_TOKEN)|(1<<LuceneParserTERM_NORMAL)|(1<<LuceneParserTERM_TRUNCATED)|(1<<LuceneParserPHRASE)|(1<<LuceneParserPHRASE_ANYTHING))) != 0 {
		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuceneParserTO {
			{
				p.SetState(166)
				p.Match(LuceneParserTO)
			}

		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuceneParserWS {
			{
				p.SetState(169)
				p.Sep()
			}

		}
		{
			p.SetState(172)

			var _x = p.Range_value()

			localctx.(*Two_sided_range_termContext).b = _x
		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuceneParserWS {
			{
				p.SetState(173)
				p.Sep()
			}

		}

	}
	{
		p.SetState(178)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Two_sided_range_termContext).end_type = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == LuceneParserRBRACK || _la == LuceneParserRCURLY) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Two_sided_range_termContext).end_type = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRange_termContext is an interface to support dynamic dispatch.
type IRange_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRange_termContext differentiates from other interfaces.
	IsRange_termContext()
}

type Range_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRange_termContext() *Range_termContext {
	var p = new(Range_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_range_term
	return p
}

func (*Range_termContext) IsRange_termContext() {}

func NewRange_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Range_termContext {
	var p = new(Range_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_range_term

	return p
}

func (s *Range_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Range_termContext) Two_sided_range_term() ITwo_sided_range_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITwo_sided_range_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITwo_sided_range_termContext)
}

func (s *Range_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Range_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Range_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterRange_term(s)
	}
}

func (s *Range_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitRange_term(s)
	}
}

func (p *LuceneParser) Range_term() (localctx IRange_termContext) {
	localctx = NewRange_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LuceneParserRULE_range_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Two_sided_range_term()
	}

	return localctx
}

// IRange_valueContext is an interface to support dynamic dispatch.
type IRange_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRange_valueContext differentiates from other interfaces.
	IsRange_valueContext()
}

type Range_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRange_valueContext() *Range_valueContext {
	var p = new(Range_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_range_value
	return p
}

func (*Range_valueContext) IsRange_valueContext() {}

func NewRange_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Range_valueContext {
	var p = new(Range_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_range_value

	return p
}

func (s *Range_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Range_valueContext) Truncated() ITruncatedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITruncatedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITruncatedContext)
}

func (s *Range_valueContext) Quoted() IQuotedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuotedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuotedContext)
}

func (s *Range_valueContext) Quoted_truncated() IQuoted_truncatedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuoted_truncatedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuoted_truncatedContext)
}

func (s *Range_valueContext) Date() IDateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateContext)
}

func (s *Range_valueContext) Normal() INormalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INormalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INormalContext)
}

func (s *Range_valueContext) STAR() antlr.TerminalNode {
	return s.GetToken(LuceneParserSTAR, 0)
}

func (s *Range_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Range_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Range_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterRange_value(s)
	}
}

func (s *Range_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitRange_value(s)
	}
}

func (p *LuceneParser) Range_value() (localctx IRange_valueContext) {
	localctx = NewRange_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LuceneParserRULE_range_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(188)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuceneParserTERM_TRUNCATED:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(182)
			p.Truncated()
		}

	case LuceneParserPHRASE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(183)
			p.Quoted()
		}

	case LuceneParserPHRASE_ANYTHING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(184)
			p.Quoted_truncated()
		}

	case LuceneParserDATE_TOKEN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(185)
			p.Date()
		}

	case LuceneParserNUMBER, LuceneParserTERM_NORMAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(186)
			p.Normal()
		}

	case LuceneParserSTAR:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(187)
			p.Match(LuceneParserSTAR)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMulti_valueContext is an interface to support dynamic dispatch.
type IMulti_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMulti_valueContext differentiates from other interfaces.
	IsMulti_valueContext()
}

type Multi_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMulti_valueContext() *Multi_valueContext {
	var p = new(Multi_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_multi_value
	return p
}

func (*Multi_valueContext) IsMulti_valueContext() {}

func NewMulti_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Multi_valueContext {
	var p = new(Multi_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_multi_value

	return p
}

func (s *Multi_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Multi_valueContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LuceneParserLPAREN, 0)
}

func (s *Multi_valueContext) ClauseDefault() IClauseDefaultContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClauseDefaultContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClauseDefaultContext)
}

func (s *Multi_valueContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LuceneParserRPAREN, 0)
}

func (s *Multi_valueContext) Sep() ISepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISepContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISepContext)
}

func (s *Multi_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Multi_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Multi_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterMulti_value(s)
	}
}

func (s *Multi_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitMulti_value(s)
	}
}

func (p *LuceneParser) Multi_value() (localctx IMulti_valueContext) {
	localctx = NewMulti_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LuceneParserRULE_multi_value)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Match(LuceneParserLPAREN)
	}
	{
		p.SetState(191)
		p.ClauseDefault()
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuceneParserWS {
		{
			p.SetState(192)
			p.Sep()
		}

	}
	{
		p.SetState(195)
		p.Match(LuceneParserRPAREN)
	}

	return localctx
}

// INormalContext is an interface to support dynamic dispatch.
type INormalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNormalContext differentiates from other interfaces.
	IsNormalContext()
}

type NormalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNormalContext() *NormalContext {
	var p = new(NormalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_normal
	return p
}

func (*NormalContext) IsNormalContext() {}

func NewNormalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NormalContext {
	var p = new(NormalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_normal

	return p
}

func (s *NormalContext) GetParser() antlr.Parser { return s.parser }

func (s *NormalContext) TERM_NORMAL() antlr.TerminalNode {
	return s.GetToken(LuceneParserTERM_NORMAL, 0)
}

func (s *NormalContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(LuceneParserNUMBER, 0)
}

func (s *NormalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NormalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NormalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterNormal(s)
	}
}

func (s *NormalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitNormal(s)
	}
}

func (p *LuceneParser) Normal() (localctx INormalContext) {
	localctx = NewNormalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LuceneParserRULE_normal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LuceneParserNUMBER || _la == LuceneParserTERM_NORMAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITruncatedContext is an interface to support dynamic dispatch.
type ITruncatedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTruncatedContext differentiates from other interfaces.
	IsTruncatedContext()
}

type TruncatedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTruncatedContext() *TruncatedContext {
	var p = new(TruncatedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_truncated
	return p
}

func (*TruncatedContext) IsTruncatedContext() {}

func NewTruncatedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TruncatedContext {
	var p = new(TruncatedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_truncated

	return p
}

func (s *TruncatedContext) GetParser() antlr.Parser { return s.parser }

func (s *TruncatedContext) TERM_TRUNCATED() antlr.TerminalNode {
	return s.GetToken(LuceneParserTERM_TRUNCATED, 0)
}

func (s *TruncatedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TruncatedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TruncatedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterTruncated(s)
	}
}

func (s *TruncatedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitTruncated(s)
	}
}

func (p *LuceneParser) Truncated() (localctx ITruncatedContext) {
	localctx = NewTruncatedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LuceneParserRULE_truncated)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Match(LuceneParserTERM_TRUNCATED)
	}

	return localctx
}

// IQuoted_truncatedContext is an interface to support dynamic dispatch.
type IQuoted_truncatedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuoted_truncatedContext differentiates from other interfaces.
	IsQuoted_truncatedContext()
}

type Quoted_truncatedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuoted_truncatedContext() *Quoted_truncatedContext {
	var p = new(Quoted_truncatedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_quoted_truncated
	return p
}

func (*Quoted_truncatedContext) IsQuoted_truncatedContext() {}

func NewQuoted_truncatedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Quoted_truncatedContext {
	var p = new(Quoted_truncatedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_quoted_truncated

	return p
}

func (s *Quoted_truncatedContext) GetParser() antlr.Parser { return s.parser }

func (s *Quoted_truncatedContext) PHRASE_ANYTHING() antlr.TerminalNode {
	return s.GetToken(LuceneParserPHRASE_ANYTHING, 0)
}

func (s *Quoted_truncatedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Quoted_truncatedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Quoted_truncatedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterQuoted_truncated(s)
	}
}

func (s *Quoted_truncatedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitQuoted_truncated(s)
	}
}

func (p *LuceneParser) Quoted_truncated() (localctx IQuoted_truncatedContext) {
	localctx = NewQuoted_truncatedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, LuceneParserRULE_quoted_truncated)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(LuceneParserPHRASE_ANYTHING)
	}

	return localctx
}

// IQuotedContext is an interface to support dynamic dispatch.
type IQuotedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuotedContext differentiates from other interfaces.
	IsQuotedContext()
}

type QuotedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuotedContext() *QuotedContext {
	var p = new(QuotedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_quoted
	return p
}

func (*QuotedContext) IsQuotedContext() {}

func NewQuotedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuotedContext {
	var p = new(QuotedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_quoted

	return p
}

func (s *QuotedContext) GetParser() antlr.Parser { return s.parser }

func (s *QuotedContext) PHRASE() antlr.TerminalNode {
	return s.GetToken(LuceneParserPHRASE, 0)
}

func (s *QuotedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuotedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterQuoted(s)
	}
}

func (s *QuotedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitQuoted(s)
	}
}

func (p *LuceneParser) Quoted() (localctx IQuotedContext) {
	localctx = NewQuotedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, LuceneParserRULE_quoted)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Match(LuceneParserPHRASE)
	}

	return localctx
}

// IModifierContext is an interface to support dynamic dispatch.
type IModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModifierContext differentiates from other interfaces.
	IsModifierContext()
}

type ModifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModifierContext() *ModifierContext {
	var p = new(ModifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_modifier
	return p
}

func (*ModifierContext) IsModifierContext() {}

func NewModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModifierContext {
	var p = new(ModifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_modifier

	return p
}

func (s *ModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ModifierContext) PLUS() antlr.TerminalNode {
	return s.GetToken(LuceneParserPLUS, 0)
}

func (s *ModifierContext) MINUS() antlr.TerminalNode {
	return s.GetToken(LuceneParserMINUS, 0)
}

func (s *ModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterModifier(s)
	}
}

func (s *ModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitModifier(s)
	}
}

func (p *LuceneParser) Modifier() (localctx IModifierContext) {
	localctx = NewModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, LuceneParserRULE_modifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LuceneParserPLUS || _la == LuceneParserMINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITerm_modifierContext is an interface to support dynamic dispatch.
type ITerm_modifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTerm_modifierContext differentiates from other interfaces.
	IsTerm_modifierContext()
}

type Term_modifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerm_modifierContext() *Term_modifierContext {
	var p = new(Term_modifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_term_modifier
	return p
}

func (*Term_modifierContext) IsTerm_modifierContext() {}

func NewTerm_modifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Term_modifierContext {
	var p = new(Term_modifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_term_modifier

	return p
}

func (s *Term_modifierContext) GetParser() antlr.Parser { return s.parser }

func (s *Term_modifierContext) Boost() IBoostContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoostContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoostContext)
}

func (s *Term_modifierContext) Fuzzy() IFuzzyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuzzyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuzzyContext)
}

func (s *Term_modifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Term_modifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Term_modifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterTerm_modifier(s)
	}
}

func (s *Term_modifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitTerm_modifier(s)
	}
}

func (p *LuceneParser) Term_modifier() (localctx ITerm_modifierContext) {
	localctx = NewTerm_modifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, LuceneParserRULE_term_modifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(215)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuceneParserCARAT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(207)
			p.Boost()
		}
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuceneParserTILDE {
			{
				p.SetState(208)
				p.Fuzzy()
			}

		}

	case LuceneParserTILDE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(211)
			p.Fuzzy()
		}
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuceneParserCARAT {
			{
				p.SetState(212)
				p.Boost()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBoostContext is an interface to support dynamic dispatch.
type IBoostContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoostContext differentiates from other interfaces.
	IsBoostContext()
}

type BoostContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoostContext() *BoostContext {
	var p = new(BoostContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_boost
	return p
}

func (*BoostContext) IsBoostContext() {}

func NewBoostContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoostContext {
	var p = new(BoostContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_boost

	return p
}

func (s *BoostContext) GetParser() antlr.Parser { return s.parser }

func (s *BoostContext) CARAT() antlr.TerminalNode {
	return s.GetToken(LuceneParserCARAT, 0)
}

func (s *BoostContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(LuceneParserNUMBER, 0)
}

func (s *BoostContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoostContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoostContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterBoost(s)
	}
}

func (s *BoostContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitBoost(s)
	}
}

func (p *LuceneParser) Boost() (localctx IBoostContext) {
	localctx = NewBoostContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, LuceneParserRULE_boost)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
		p.Match(LuceneParserCARAT)
	}

	p.SetState(219)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(218)
			p.Match(LuceneParserNUMBER)
		}

	}

	return localctx
}

// IFuzzyContext is an interface to support dynamic dispatch.
type IFuzzyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuzzyContext differentiates from other interfaces.
	IsFuzzyContext()
}

type FuzzyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuzzyContext() *FuzzyContext {
	var p = new(FuzzyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_fuzzy
	return p
}

func (*FuzzyContext) IsFuzzyContext() {}

func NewFuzzyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuzzyContext {
	var p = new(FuzzyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_fuzzy

	return p
}

func (s *FuzzyContext) GetParser() antlr.Parser { return s.parser }

func (s *FuzzyContext) TILDE() antlr.TerminalNode {
	return s.GetToken(LuceneParserTILDE, 0)
}

func (s *FuzzyContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(LuceneParserNUMBER, 0)
}

func (s *FuzzyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuzzyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuzzyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterFuzzy(s)
	}
}

func (s *FuzzyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitFuzzy(s)
	}
}

func (p *LuceneParser) Fuzzy() (localctx IFuzzyContext) {
	localctx = NewFuzzyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, LuceneParserRULE_fuzzy)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.Match(LuceneParserTILDE)
	}

	p.SetState(223)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(222)
			p.Match(LuceneParserNUMBER)
		}

	}

	return localctx
}

// INot_Context is an interface to support dynamic dispatch.
type INot_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNot_Context differentiates from other interfaces.
	IsNot_Context()
}

type Not_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNot_Context() *Not_Context {
	var p = new(Not_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_not_
	return p
}

func (*Not_Context) IsNot_Context() {}

func NewNot_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Not_Context {
	var p = new(Not_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_not_

	return p
}

func (s *Not_Context) GetParser() antlr.Parser { return s.parser }

func (s *Not_Context) AND() antlr.TerminalNode {
	return s.GetToken(LuceneParserAND, 0)
}

func (s *Not_Context) NOT() antlr.TerminalNode {
	return s.GetToken(LuceneParserNOT, 0)
}

func (s *Not_Context) AllSep() []ISepContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISepContext)(nil)).Elem())
	var tst = make([]ISepContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISepContext)
		}
	}

	return tst
}

func (s *Not_Context) Sep(i int) ISepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISepContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISepContext)
}

func (s *Not_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Not_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Not_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterNot_(s)
	}
}

func (s *Not_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitNot_(s)
	}
}

func (p *LuceneParser) Not_() (localctx INot_Context) {
	localctx = NewNot_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, LuceneParserRULE_not_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuceneParserWS {
			{
				p.SetState(225)
				p.Sep()
			}

		}
		{
			p.SetState(228)
			p.Match(LuceneParserAND)
		}
		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuceneParserWS {
			{
				p.SetState(229)
				p.Sep()
			}

		}
		{
			p.SetState(232)
			p.Match(LuceneParserNOT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuceneParserWS {
			{
				p.SetState(233)
				p.Sep()
			}

		}
		{
			p.SetState(236)
			p.Match(LuceneParserNOT)
		}

	}

	return localctx
}

// IAnd_Context is an interface to support dynamic dispatch.
type IAnd_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnd_Context differentiates from other interfaces.
	IsAnd_Context()
}

type And_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnd_Context() *And_Context {
	var p = new(And_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_and_
	return p
}

func (*And_Context) IsAnd_Context() {}

func NewAnd_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *And_Context {
	var p = new(And_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_and_

	return p
}

func (s *And_Context) GetParser() antlr.Parser { return s.parser }

func (s *And_Context) AND() antlr.TerminalNode {
	return s.GetToken(LuceneParserAND, 0)
}

func (s *And_Context) Sep() ISepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISepContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISepContext)
}

func (s *And_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *And_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *And_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterAnd_(s)
	}
}

func (s *And_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitAnd_(s)
	}
}

func (p *LuceneParser) And_() (localctx IAnd_Context) {
	localctx = NewAnd_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, LuceneParserRULE_and_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuceneParserWS {
		{
			p.SetState(239)
			p.Sep()
		}

	}
	{
		p.SetState(242)
		p.Match(LuceneParserAND)
	}

	return localctx
}

// IOr_Context is an interface to support dynamic dispatch.
type IOr_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOr_Context differentiates from other interfaces.
	IsOr_Context()
}

type Or_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOr_Context() *Or_Context {
	var p = new(Or_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_or_
	return p
}

func (*Or_Context) IsOr_Context() {}

func NewOr_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Or_Context {
	var p = new(Or_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_or_

	return p
}

func (s *Or_Context) GetParser() antlr.Parser { return s.parser }

func (s *Or_Context) OR() antlr.TerminalNode {
	return s.GetToken(LuceneParserOR, 0)
}

func (s *Or_Context) Sep() ISepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISepContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISepContext)
}

func (s *Or_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Or_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Or_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterOr_(s)
	}
}

func (s *Or_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitOr_(s)
	}
}

func (p *LuceneParser) Or_() (localctx IOr_Context) {
	localctx = NewOr_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, LuceneParserRULE_or_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuceneParserWS {
		{
			p.SetState(244)
			p.Sep()
		}

	}
	{
		p.SetState(247)
		p.Match(LuceneParserOR)
	}

	return localctx
}

// IDateContext is an interface to support dynamic dispatch.
type IDateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateContext differentiates from other interfaces.
	IsDateContext()
}

type DateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateContext() *DateContext {
	var p = new(DateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_date
	return p
}

func (*DateContext) IsDateContext() {}

func NewDateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateContext {
	var p = new(DateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_date

	return p
}

func (s *DateContext) GetParser() antlr.Parser { return s.parser }

func (s *DateContext) DATE_TOKEN() antlr.TerminalNode {
	return s.GetToken(LuceneParserDATE_TOKEN, 0)
}

func (s *DateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterDate(s)
	}
}

func (s *DateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitDate(s)
	}
}

func (p *LuceneParser) Date() (localctx IDateContext) {
	localctx = NewDateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, LuceneParserRULE_date)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.Match(LuceneParserDATE_TOKEN)
	}

	return localctx
}

// ISepContext is an interface to support dynamic dispatch.
type ISepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSepContext differentiates from other interfaces.
	IsSepContext()
}

type SepContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySepContext() *SepContext {
	var p = new(SepContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuceneParserRULE_sep
	return p
}

func (*SepContext) IsSepContext() {}

func NewSepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SepContext {
	var p = new(SepContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuceneParserRULE_sep

	return p
}

func (s *SepContext) GetParser() antlr.Parser { return s.parser }

func (s *SepContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(LuceneParserWS)
}

func (s *SepContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(LuceneParserWS, i)
}

func (s *SepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.EnterSep(s)
	}
}

func (s *SepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuceneListener); ok {
		listenerT.ExitSep(s)
	}
}

func (p *LuceneParser) Sep() (localctx ISepContext) {
	localctx = NewSepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, LuceneParserRULE_sep)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(251)
				p.Match(LuceneParserWS)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())
	}

	return localctx
}
