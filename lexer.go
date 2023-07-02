package main

import (
	"bufio"
	"fmt"
	"io"
	"log"

	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Integer | constraints.Float
}

type AAA interface {
	comparable
}

type Evaluator interface {
	calcLexer
	Err() error
	Data() interface{}
}

func NewEvaluator(r io.Reader, debug bool) Evaluator {
	return &evaluator{
		buf:   bufio.NewReader(r),
		debug: debug,
	}
}

type evaluator struct {
	buf   *bufio.Reader
	data  interface{}
	err   error
	debug bool
}

func (e *evaluator) Lex(lval *calcSymType) int {
	return e.lex(lval)
}

func (e *evaluator) Error(s string) {
	log.Println(fmt.Sprintf("syntax error: %s", s))
}

func (e *evaluator) Err() error {
	return e.err
}

func (e *evaluator) Data() interface{} {
	return e.data
}

func (e *evaluator) read() rune {
	ch, _, _ := e.buf.ReadRune()
	return ch
}
func (e *evaluator) lex(lval *calcSymType) int {
	for {
		r := e.read()
		if r == 0 {
			return 0
		}
		switch r {
		case '(':
			return LEFT_PAREN
		case ')':
			return RIGHT_PAREN
		default:
			return
		}

	}
}
