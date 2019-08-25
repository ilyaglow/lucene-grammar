// Code generated from Lucene.g4 by ANTLR 4.7.1. DO NOT EDIT.

package lucene // Lucene
import "github.com/antlr/antlr4/runtime/Go/antlr"

// LuceneListener is a complete listener for a parse tree produced by LuceneParser.
type LuceneListener interface {
	antlr.ParseTreeListener

	// EnterMainQ is called when entering the mainQ production.
	EnterMainQ(c *MainQContext)

	// EnterClauseDefault is called when entering the clauseDefault production.
	EnterClauseDefault(c *ClauseDefaultContext)

	// EnterClauseOr is called when entering the clauseOr production.
	EnterClauseOr(c *ClauseOrContext)

	// EnterClauseAnd is called when entering the clauseAnd production.
	EnterClauseAnd(c *ClauseAndContext)

	// EnterClauseNot is called when entering the clauseNot production.
	EnterClauseNot(c *ClauseNotContext)

	// EnterClauseBasic is called when entering the clauseBasic production.
	EnterClauseBasic(c *ClauseBasicContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterAnything is called when entering the anything production.
	EnterAnything(c *AnythingContext)

	// EnterTwo_sided_range_term is called when entering the two_sided_range_term production.
	EnterTwo_sided_range_term(c *Two_sided_range_termContext)

	// EnterRange_term is called when entering the range_term production.
	EnterRange_term(c *Range_termContext)

	// EnterRange_value is called when entering the range_value production.
	EnterRange_value(c *Range_valueContext)

	// EnterMulti_value is called when entering the multi_value production.
	EnterMulti_value(c *Multi_valueContext)

	// EnterNormal is called when entering the normal production.
	EnterNormal(c *NormalContext)

	// EnterTruncated is called when entering the truncated production.
	EnterTruncated(c *TruncatedContext)

	// EnterQuoted_truncated is called when entering the quoted_truncated production.
	EnterQuoted_truncated(c *Quoted_truncatedContext)

	// EnterQuoted is called when entering the quoted production.
	EnterQuoted(c *QuotedContext)

	// EnterModifier is called when entering the modifier production.
	EnterModifier(c *ModifierContext)

	// EnterTerm_modifier is called when entering the term_modifier production.
	EnterTerm_modifier(c *Term_modifierContext)

	// EnterBoost is called when entering the boost production.
	EnterBoost(c *BoostContext)

	// EnterFuzzy is called when entering the fuzzy production.
	EnterFuzzy(c *FuzzyContext)

	// EnterNot_ is called when entering the not_ production.
	EnterNot_(c *Not_Context)

	// EnterAnd_ is called when entering the and_ production.
	EnterAnd_(c *And_Context)

	// EnterOr_ is called when entering the or_ production.
	EnterOr_(c *Or_Context)

	// EnterDate is called when entering the date production.
	EnterDate(c *DateContext)

	// EnterSep is called when entering the sep production.
	EnterSep(c *SepContext)

	// ExitMainQ is called when exiting the mainQ production.
	ExitMainQ(c *MainQContext)

	// ExitClauseDefault is called when exiting the clauseDefault production.
	ExitClauseDefault(c *ClauseDefaultContext)

	// ExitClauseOr is called when exiting the clauseOr production.
	ExitClauseOr(c *ClauseOrContext)

	// ExitClauseAnd is called when exiting the clauseAnd production.
	ExitClauseAnd(c *ClauseAndContext)

	// ExitClauseNot is called when exiting the clauseNot production.
	ExitClauseNot(c *ClauseNotContext)

	// ExitClauseBasic is called when exiting the clauseBasic production.
	ExitClauseBasic(c *ClauseBasicContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitAnything is called when exiting the anything production.
	ExitAnything(c *AnythingContext)

	// ExitTwo_sided_range_term is called when exiting the two_sided_range_term production.
	ExitTwo_sided_range_term(c *Two_sided_range_termContext)

	// ExitRange_term is called when exiting the range_term production.
	ExitRange_term(c *Range_termContext)

	// ExitRange_value is called when exiting the range_value production.
	ExitRange_value(c *Range_valueContext)

	// ExitMulti_value is called when exiting the multi_value production.
	ExitMulti_value(c *Multi_valueContext)

	// ExitNormal is called when exiting the normal production.
	ExitNormal(c *NormalContext)

	// ExitTruncated is called when exiting the truncated production.
	ExitTruncated(c *TruncatedContext)

	// ExitQuoted_truncated is called when exiting the quoted_truncated production.
	ExitQuoted_truncated(c *Quoted_truncatedContext)

	// ExitQuoted is called when exiting the quoted production.
	ExitQuoted(c *QuotedContext)

	// ExitModifier is called when exiting the modifier production.
	ExitModifier(c *ModifierContext)

	// ExitTerm_modifier is called when exiting the term_modifier production.
	ExitTerm_modifier(c *Term_modifierContext)

	// ExitBoost is called when exiting the boost production.
	ExitBoost(c *BoostContext)

	// ExitFuzzy is called when exiting the fuzzy production.
	ExitFuzzy(c *FuzzyContext)

	// ExitNot_ is called when exiting the not_ production.
	ExitNot_(c *Not_Context)

	// ExitAnd_ is called when exiting the and_ production.
	ExitAnd_(c *And_Context)

	// ExitOr_ is called when exiting the or_ production.
	ExitOr_(c *Or_Context)

	// ExitDate is called when exiting the date production.
	ExitDate(c *DateContext)

	// ExitSep is called when exiting the sep production.
	ExitSep(c *SepContext)
}
