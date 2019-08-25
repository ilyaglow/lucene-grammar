// Code generated from Lucene.g4 by ANTLR 4.7.1. DO NOT EDIT.

package lucene // Lucene
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLuceneListener is a complete listener for a parse tree produced by LuceneParser.
type BaseLuceneListener struct{}

var _ LuceneListener = &BaseLuceneListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLuceneListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLuceneListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLuceneListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLuceneListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMainQ is called when production mainQ is entered.
func (s *BaseLuceneListener) EnterMainQ(ctx *MainQContext) {}

// ExitMainQ is called when production mainQ is exited.
func (s *BaseLuceneListener) ExitMainQ(ctx *MainQContext) {}

// EnterClauseDefault is called when production clauseDefault is entered.
func (s *BaseLuceneListener) EnterClauseDefault(ctx *ClauseDefaultContext) {}

// ExitClauseDefault is called when production clauseDefault is exited.
func (s *BaseLuceneListener) ExitClauseDefault(ctx *ClauseDefaultContext) {}

// EnterClauseOr is called when production clauseOr is entered.
func (s *BaseLuceneListener) EnterClauseOr(ctx *ClauseOrContext) {}

// ExitClauseOr is called when production clauseOr is exited.
func (s *BaseLuceneListener) ExitClauseOr(ctx *ClauseOrContext) {}

// EnterClauseAnd is called when production clauseAnd is entered.
func (s *BaseLuceneListener) EnterClauseAnd(ctx *ClauseAndContext) {}

// ExitClauseAnd is called when production clauseAnd is exited.
func (s *BaseLuceneListener) ExitClauseAnd(ctx *ClauseAndContext) {}

// EnterClauseNot is called when production clauseNot is entered.
func (s *BaseLuceneListener) EnterClauseNot(ctx *ClauseNotContext) {}

// ExitClauseNot is called when production clauseNot is exited.
func (s *BaseLuceneListener) ExitClauseNot(ctx *ClauseNotContext) {}

// EnterClauseBasic is called when production clauseBasic is entered.
func (s *BaseLuceneListener) EnterClauseBasic(ctx *ClauseBasicContext) {}

// ExitClauseBasic is called when production clauseBasic is exited.
func (s *BaseLuceneListener) ExitClauseBasic(ctx *ClauseBasicContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseLuceneListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseLuceneListener) ExitAtom(ctx *AtomContext) {}

// EnterField is called when production field is entered.
func (s *BaseLuceneListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseLuceneListener) ExitField(ctx *FieldContext) {}

// EnterValue is called when production value is entered.
func (s *BaseLuceneListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseLuceneListener) ExitValue(ctx *ValueContext) {}

// EnterAnything is called when production anything is entered.
func (s *BaseLuceneListener) EnterAnything(ctx *AnythingContext) {}

// ExitAnything is called when production anything is exited.
func (s *BaseLuceneListener) ExitAnything(ctx *AnythingContext) {}

// EnterTwo_sided_range_term is called when production two_sided_range_term is entered.
func (s *BaseLuceneListener) EnterTwo_sided_range_term(ctx *Two_sided_range_termContext) {}

// ExitTwo_sided_range_term is called when production two_sided_range_term is exited.
func (s *BaseLuceneListener) ExitTwo_sided_range_term(ctx *Two_sided_range_termContext) {}

// EnterRange_term is called when production range_term is entered.
func (s *BaseLuceneListener) EnterRange_term(ctx *Range_termContext) {}

// ExitRange_term is called when production range_term is exited.
func (s *BaseLuceneListener) ExitRange_term(ctx *Range_termContext) {}

// EnterRange_value is called when production range_value is entered.
func (s *BaseLuceneListener) EnterRange_value(ctx *Range_valueContext) {}

// ExitRange_value is called when production range_value is exited.
func (s *BaseLuceneListener) ExitRange_value(ctx *Range_valueContext) {}

// EnterMulti_value is called when production multi_value is entered.
func (s *BaseLuceneListener) EnterMulti_value(ctx *Multi_valueContext) {}

// ExitMulti_value is called when production multi_value is exited.
func (s *BaseLuceneListener) ExitMulti_value(ctx *Multi_valueContext) {}

// EnterNormal is called when production normal is entered.
func (s *BaseLuceneListener) EnterNormal(ctx *NormalContext) {}

// ExitNormal is called when production normal is exited.
func (s *BaseLuceneListener) ExitNormal(ctx *NormalContext) {}

// EnterTruncated is called when production truncated is entered.
func (s *BaseLuceneListener) EnterTruncated(ctx *TruncatedContext) {}

// ExitTruncated is called when production truncated is exited.
func (s *BaseLuceneListener) ExitTruncated(ctx *TruncatedContext) {}

// EnterQuoted_truncated is called when production quoted_truncated is entered.
func (s *BaseLuceneListener) EnterQuoted_truncated(ctx *Quoted_truncatedContext) {}

// ExitQuoted_truncated is called when production quoted_truncated is exited.
func (s *BaseLuceneListener) ExitQuoted_truncated(ctx *Quoted_truncatedContext) {}

// EnterQuoted is called when production quoted is entered.
func (s *BaseLuceneListener) EnterQuoted(ctx *QuotedContext) {}

// ExitQuoted is called when production quoted is exited.
func (s *BaseLuceneListener) ExitQuoted(ctx *QuotedContext) {}

// EnterModifier is called when production modifier is entered.
func (s *BaseLuceneListener) EnterModifier(ctx *ModifierContext) {}

// ExitModifier is called when production modifier is exited.
func (s *BaseLuceneListener) ExitModifier(ctx *ModifierContext) {}

// EnterTerm_modifier is called when production term_modifier is entered.
func (s *BaseLuceneListener) EnterTerm_modifier(ctx *Term_modifierContext) {}

// ExitTerm_modifier is called when production term_modifier is exited.
func (s *BaseLuceneListener) ExitTerm_modifier(ctx *Term_modifierContext) {}

// EnterBoost is called when production boost is entered.
func (s *BaseLuceneListener) EnterBoost(ctx *BoostContext) {}

// ExitBoost is called when production boost is exited.
func (s *BaseLuceneListener) ExitBoost(ctx *BoostContext) {}

// EnterFuzzy is called when production fuzzy is entered.
func (s *BaseLuceneListener) EnterFuzzy(ctx *FuzzyContext) {}

// ExitFuzzy is called when production fuzzy is exited.
func (s *BaseLuceneListener) ExitFuzzy(ctx *FuzzyContext) {}

// EnterNot_ is called when production not_ is entered.
func (s *BaseLuceneListener) EnterNot_(ctx *Not_Context) {}

// ExitNot_ is called when production not_ is exited.
func (s *BaseLuceneListener) ExitNot_(ctx *Not_Context) {}

// EnterAnd_ is called when production and_ is entered.
func (s *BaseLuceneListener) EnterAnd_(ctx *And_Context) {}

// ExitAnd_ is called when production and_ is exited.
func (s *BaseLuceneListener) ExitAnd_(ctx *And_Context) {}

// EnterOr_ is called when production or_ is entered.
func (s *BaseLuceneListener) EnterOr_(ctx *Or_Context) {}

// ExitOr_ is called when production or_ is exited.
func (s *BaseLuceneListener) ExitOr_(ctx *Or_Context) {}

// EnterDate is called when production date is entered.
func (s *BaseLuceneListener) EnterDate(ctx *DateContext) {}

// ExitDate is called when production date is exited.
func (s *BaseLuceneListener) ExitDate(ctx *DateContext) {}

// EnterSep is called when production sep is entered.
func (s *BaseLuceneListener) EnterSep(ctx *SepContext) {}

// ExitSep is called when production sep is exited.
func (s *BaseLuceneListener) ExitSep(ctx *SepContext) {}
