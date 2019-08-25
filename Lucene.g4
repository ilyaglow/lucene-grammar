grammar Lucene;

//
//  Based on https://github.com/lrowe/lucenequery/blob/master/lucenequery/StandardLuceneGrammar.g4
//
//  This is a re-implementation of the standard lucene syntax with ANTLR4
//   http://lucene.apache.org/core/4_3_0/queryparser/index.html
//
//   The query syntax is complete and supports the same features as the
//   original parser written in JCC. The advantage of this grammar is that it
//   is 'pluggable' into Lucene's modern flexible parser, so that you can
//   add custom logic on top of the 'rigid' query syntax. Besides...the grammar
//   is not that 'rigid' - you can modify the grammar and easily recompile.
//
//   # run this commad inside antlrqueryparser
//
//      $ ant generate-antlr -Dgrammar=MyNewGrammar
//
//   # or if you want to test things, do:
//
//      $ ant try-view -Dgrammar=MyNewGrammar -Dquery="foo AND bar"
//
//
//   Implementation note: I have tried hard to avoid putting language specific details
//   into the grammar, unfortunately this was not always possible. But it is kept
//   at minimum. You can generate parser code in your language of choice
//   if you change the following:
//
//   options  :
//     language=<your-target-language-supported-by-antlr3>
//     superClass= the default is to subclass 'UnforgivingParser', this java class
//                         lives in the package oal.queryparser.flixible.aqp.parser
//                         and its purpose is to bark everytime when an exception
//                         happens (otherwise, ANTLR tries to recover from some situations
//                         -- you may want to remove this definition, or add your own
//                         error recovery logic there)
//
//   @header:  this adds the java declaration to the generated parser file,
//                    feel free to remove (if you want to test the grammar using
//                    ANTLRWorks, you want to remove it)
//   @lexer::header: dtto but for lexer
//   @lexer::members: again, necessary for being strict and prevent error
//                                recovery, but this applies only to lexer errors.
//
//  One last note - if you want to implement your own error recovery, have a look
//  at the generated java class
//
//      oal.queryparser.flixible.aqp.parser.<GrammarName>SyntaxParser.java
//
//  There we are raising parse exception as well
//

options {
  language = Python2;
}

tokens {
  OPERATOR,
  ATOM,
  MODIFIER,
  TMODIFIER,
  CLAUSE,
  FIELD,
  FUZZY,
  BOOST,
  QNORMAL,
  QPHRASE,
  QPHRASETRUNC,
  QTRUNCATED,
  QRANGEIN,
  QRANGEEX,
  QANYTHING,
  QDATE
}

mainQ :
  // sep? clause=clauseDefault sep? EOF
  sep? clause=clauseDefault sep?
  ;

clauseDefault
  :
  //m:(a b AND c OR d OR e)
  // without duplicating the rules (but it allows recursion)
  clauseOr (sep? clauseOr)*
  ;

clauseOr
  : clauseAnd (or_ clauseAnd)*
  ;

clauseAnd
  : clauseNot (and_ clauseNot)*
  ;

clauseNot
  : clauseBasic (not_ clauseBasic)*
  ;

clauseBasic
  :
  sep? modifier? LPAREN clauseDefault sep? RPAREN term_modifier?
  | sep? atom
  ;

atom
  :
  modifier? field multi_value term_modifier?
  | modifier? field? value term_modifier?
  ;

field
  :
  TERM_NORMAL COLON sep?
  ;

value
  :
  range_term
  | normal
  | truncated
  | quoted
  | quoted_truncated
  | QMARK
  | anything
  | STAR
  ;

anything
  :
  STAR COLON STAR
  ;

two_sided_range_term
  :
  start_type=(LBRACK|LCURLY)
  sep?
  (a=range_value)
  sep?
  ( TO? sep? b=range_value sep? )?
  end_type=(RBRACK|RCURLY)
  ;

range_term
  :
  two_sided_range_term
  ;

range_value
  :
  truncated
  | quoted
  | quoted_truncated
  | date
  | normal
  | STAR
  ;

multi_value
  :
  LPAREN clauseDefault sep? RPAREN
  ;

normal
  :
  TERM_NORMAL
  | NUMBER
  ;

truncated
  :
  TERM_TRUNCATED
  ;

quoted_truncated
  :
  PHRASE_ANYTHING
  ;

quoted  :
  PHRASE
  ;

modifier:
  PLUS
  | MINUS;


term_modifier :
  boost fuzzy?
  | fuzzy boost?
  ;

boost :
  (CARAT) // set the default value
  (NUMBER)? //replace the default with user input
  ;

fuzzy :
  (TILDE) // set the default value
  (NUMBER)? //replace the default with user input
  ;

not_  :
  sep? AND sep? NOT
  | sep? NOT
  ;

and_  :
  sep? AND
  ;

or_   :
  sep? OR
  ;

date  :
  DATE_TOKEN
  ;

/* ================================================================
 * =                     LEXER                                    =
 * ================================================================
 */

LPAREN  : '(';

RPAREN  : ')';

LBRACK  : '[';

RBRACK  : ']';

COLON   : ':' ;  //this must NOT be fragment

PLUS  : '+' ;

MINUS : ('-'|'!');

STAR  : '*' ;

QMARK  : '?'+ ;

fragment VBAR  : '|' ;

fragment AMPER : '&' ;

LCURLY  : '{' ;

RCURLY  : '}' ;

CARAT : '^' (INT+ ('.' INT+)?)?;

TILDE : '~' (INT+ ('.' INT+)?)?;

DQUOTE
  : '"';

SQUOTE
  : '\'';

TO  : 'TO';

/* We want to be case insensitive */
AND   : (('a' | 'A') ('n' | 'N') ('d' | 'D') | (AMPER AMPER?)) ;
OR  : (('o' | 'O') ('r' | 'R') | (VBAR VBAR?));
NOT   : ('n' | 'N') ('o' | 'O') ('t' | 'T');

sep : WS+;

WS  :   ( ' '
        | '\t'
        | '\r'
        | '\n'
        | '\u3000'
        ) // -> skip
    ;

/*
fragment TERM_CHAR  :
     ~(' ' | '\t' | '\n' | '\r' | '\u3000'
      | '\\' | '\'' | '\"'
      | '(' | ')' | '[' | ']' | '{' | '}'
      | '+' | '-' | '!' | ':' | '~' | '^'
      | '*' | '|' | '&' | '?' | '\\\"' | '/'  //this line is not present in lucene StandardParser.jj
      );
*/

fragment INT: '0' .. '9';

fragment ESC_CHAR:  '\\' .;

fragment TERM_START_CHAR
  :
  (~(' ' | '\t' | '\n' | '\r' | '\u3000'
        | '\'' | '"'
        | '(' | ')' | '[' | ']' | '{' | '}'
        | '+' | '-' | '!' | ':' | '~' | '^'
        | '?' | '*' | '\\'
        )
   | ESC_CHAR );

fragment TERM_CHAR
  :
  (TERM_START_CHAR | '-' | '+')
  ;

NUMBER
  :
  INT+ ('.' INT+)?
  ;

DATE_TOKEN
  :
  INT INT? ('/'|'-'|'.') INT INT? ('/'|'-'|'.') INT INT (INT INT)?
  ;

TERM_NORMAL
  :
  TERM_START_CHAR ( TERM_CHAR )*
  ;

TERM_TRUNCATED:
  (STAR|QMARK) (TERM_CHAR+ (QMARK|STAR))+ (TERM_CHAR)*
  | TERM_START_CHAR (TERM_CHAR* (QMARK|STAR))+ (TERM_CHAR)*
  | (STAR|QMARK) TERM_CHAR+
  ;

PHRASE
  :
  DQUOTE (ESC_CHAR|~('"'|'\\'|'?'|'*'))+ DQUOTE
  ;

PHRASE_ANYTHING :
  DQUOTE (ESC_CHAR|~('"'|'\\'))+ DQUOTE
  ;
