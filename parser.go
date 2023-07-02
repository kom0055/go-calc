package main

import "strings"

func Evaluate(d string, debug bool) (interface{}, error) {
	e := NewEvaluator(strings.NewReader(d), debug)
	calcParse(e)
	if err := e.Err(); err != nil {
		return nil, err
	}
	return e.Data(), nil
}
