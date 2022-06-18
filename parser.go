package string_sum

import (
	"fmt"
	"strconv"
)

type Parser struct {
	stage         int
	hasFirstMinus bool
	firstDigits   []byte
	secondDigits  []byte
	operation     byte
}

const (
	STAGE_INIT = iota
	STAGE_WS_HEAD
	STAGE_FIRST_MINUS
	STAGE_FIRST_DIGIT
	STAGE_WS_OPERATION_BEFORE
	STAGE_OPERATION
	STAGE_WS_OPERATION_AFTER
	STAGE_SECOND_DIGIT
	STAGE_WS_TAIL
)

func NewParser() *Parser {
	parser := Parser{}
	parser.stage = STAGE_INIT
	return &parser
}

func (p *Parser) NextSymbol(c byte) error {
	s := NewSymbol(c)

	errorUnexpectedSymbol := fmt.Errorf("parse error, unexpected symbol %c", c)

	nextStage := p.stage

	switch p.stage {

	case STAGE_INIT:
		fallthrough
	case STAGE_WS_HEAD:
		switch s.Type {
		case SYM_WHITESPACE:
			nextStage = STAGE_WS_HEAD
		case SYM_DIGIT:
			nextStage = STAGE_FIRST_DIGIT
			p.firstDigits = append(p.firstDigits, c)
		case SYM_MINUS:
			nextStage = STAGE_FIRST_MINUS
			p.hasFirstMinus = true
		default:
			return errorUnexpectedSymbol
		}

	case STAGE_FIRST_MINUS:
		switch s.Type {
		case SYM_WHITESPACE:
			// pass
		case SYM_DIGIT:
			nextStage = STAGE_FIRST_DIGIT
			p.firstDigits = append(p.firstDigits, c)
		default:
			return errorUnexpectedSymbol
		}

	case STAGE_FIRST_DIGIT:
		switch s.Type {
		case SYM_WHITESPACE:
			nextStage = STAGE_WS_OPERATION_BEFORE
		case SYM_DIGIT:
			p.firstDigits = append(p.firstDigits, c)
		case SYM_MINUS:
			fallthrough
		case SYM_PLUS:
			nextStage = STAGE_OPERATION
			p.operation = c
		default:
			return errorUnexpectedSymbol
		}

	case STAGE_WS_OPERATION_BEFORE:
		switch s.Type {
		case SYM_WHITESPACE:
			// pass
		case SYM_MINUS:
			fallthrough
		case SYM_PLUS:
			nextStage = STAGE_OPERATION
			p.operation = c
		default:
			return errorUnexpectedSymbol
		}

	case STAGE_OPERATION:
		switch s.Type {
		case SYM_WHITESPACE:
			nextStage = STAGE_WS_OPERATION_AFTER
		case SYM_MINUS:
			fallthrough
		case SYM_PLUS:
			p.operation = c
		case SYM_DIGIT:
			nextStage = STAGE_SECOND_DIGIT
			p.secondDigits = append(p.secondDigits, c)
		default:
			return errorUnexpectedSymbol
		}

	case STAGE_WS_OPERATION_AFTER:
		switch s.Type {
		case SYM_WHITESPACE:
			// pass
		case SYM_DIGIT:
			nextStage = STAGE_SECOND_DIGIT
			p.secondDigits = append(p.secondDigits, c)
		default:
			return errorUnexpectedSymbol
		}

	case STAGE_SECOND_DIGIT:
		switch s.Type {
		case SYM_WHITESPACE:
			nextStage = STAGE_WS_TAIL
		case SYM_DIGIT:
			p.secondDigits = append(p.secondDigits, c)
		default:
			return errorUnexpectedSymbol
		}

	case STAGE_WS_TAIL:
		switch s.Type {
		case SYM_WHITESPACE:
			// pass
		default:
			return errorUnexpectedSymbol
		}

	}

	p.stage = nextStage

	return nil
}

func (p *Parser) Done() error {
	if len(p.firstDigits) == 0 {
		return fmt.Errorf("%w: empty first number", errorNotTwoOperands)
	}

	if p.stage != STAGE_SECOND_DIGIT {
		return errorEmptyInput
	}

	if len(p.secondDigits) == 0 {
		return fmt.Errorf("%w: empty second number", errorNotTwoOperands)
	}

	return nil
}

func (p *Parser) Calc() (string, error) {
	s1 := string(p.firstDigits)

	if p.hasFirstMinus {
		s1 = "-" + s1
	}

	i1, err := strconv.Atoi(s1)
	if err != nil {
		return "", fmt.Errorf("convert first number: %w", err)
	}

	i2, err := strconv.Atoi(string(p.secondDigits))
	if err != nil {
		return "", fmt.Errorf("convert second number: %w", err)
	}

	var sum int

	switch p.operation {
	case '+':
		sum = i1 + i2
	case '-':
		sum = i1 - i2
	}

	return fmt.Sprintf("%d", sum), nil
}
