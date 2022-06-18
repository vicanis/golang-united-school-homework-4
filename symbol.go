package string_sum

import (
	"fmt"
	"strings"
)

type Symbol struct {
	c    byte
	Type int
}

const (
	SYM_DIGIT = iota
	SYM_PLUS
	SYM_MINUS
	SYM_WHITESPACE
	SYM_UNKNOWN
)

func includes(item string, array []string) bool {
	for _, s := range array {
		if item == s {
			return true
		}
	}

	return false
}

func parse(c byte) int {
	s := string(c)

	if includes(s, strings.Split("0123456789", "")) {
		return SYM_DIGIT
	}

	if s == "+" {
		return SYM_PLUS
	}

	if s == "-" {
		return SYM_MINUS
	}

	if includes(s, strings.Split(" \r\n\t", "")) {
		return SYM_WHITESPACE
	}

	return SYM_UNKNOWN
}

func NewSymbol(c byte) Symbol {
	sym := Symbol{}
	sym.Type = parse(c)
	return sym
}

func (sym Symbol) String() string {
	var t string

	switch sym.Type {
	case SYM_DIGIT:
		t = "digit"
	case SYM_PLUS:
		t = "plus"
	case SYM_MINUS:
		t = "minus"
	case SYM_WHITESPACE:
		t = "whitespace"
	default:
		t = "unknown"
	}

	return fmt.Sprintf("[%s]: %s", t, string(sym.c))
}
