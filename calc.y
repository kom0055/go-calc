%{
package main

import (
	_ "bufio"
	_ "bytes"
	_ "fmt"
	_ "io"
	_ "log"
	_ "math/big"
	_ "os"
	_ "unicode/utf8"
)

%}

%union {
	valI64	*int64
	valF64	*float64
	param	*string
	op	int
}


%type	<expr>	value
%type	<expr>	paren

// quoted string
%token LEFT_QUOTE RIGHT_QUOTE

// sub expr
%token LEFT_PAREN RIGHT_PAREN

// comparator symbols
%token EQ NEQ GT GTE LT LTE

// logical symbols
%token AND OR NOT

// bitwise symbols
%token BIT_XOR BIT_AND BIR_OR BIT_NOT

// bitwise shift symbols
%token LSHIFT RSHIFT

// additive symbols
%token PLUS MINUS

// multiplicative symbols
%token MULTIPLY DIVIDE MODULUS

// exponential symbols
%token EXPONENT

// unary symbols
%token UNARY_POSITIVE UNARY_NEGATIVE UNARY_NOT

%%

calc:
	value
	{
	 //setScannerData(yylex, $1)
	}

value:
paren
{
    $$ = $1
}

paren:
LEFT_PAREN value RIGHT_PAREN
{
	$$ = $2
}
| LEFT_PAREN RIGHT_PAREN
{

}

%%
