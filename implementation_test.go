package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestCalculatePostfix(cs *C) {
	result, _ := CalculatePostfix("2 2 +")
	cs.Assert(result, Equals, "4")
}

func (s *MySuite) TestCalculatePostfixPower(cs *C) {
	result, _ := CalculatePostfix("3 5 ^")
	cs.Assert(result, Equals, "243")
}

func (s *MySuite) TestCalculatePostfixMedium(cs *C) {
	result, _ := CalculatePostfix("4 5 + 5 *")
	cs.Assert(result, Equals, "45")
}

func (s *MySuite) TestCalculatePostfixComplex(cs *C) {
	result, _ := CalculatePostfix("4 2 - 3 * 5 +")
	cs.Assert(result, Equals, "11")
}

func (s *MySuite) TestCalculatePostfixComplexPower(cs *C) {
	result, _ := CalculatePostfix("2 3 ^ 2 + 2 *")
	cs.Assert(result, Equals, "20")
}

func (s *MySuite) TestCalculatePostfixBig(cs *C) {
	result, _ := CalculatePostfix("5 2 + 3 * 4 - 2 / 6 + 3 +")
	cs.Assert(result, Equals, "17")
}

func (s *MySuite) TestCalculatePostfixEmptyError(cs *C) {
	_, err := CalculatePostfix("")
	cs.Assert(err, ErrorMatches, "invalid expression: empty input")
}

func (s *MySuite) TestCalculatePostfixOperandsError(cs *C) {
	_, err := CalculatePostfix("2 +")
	cs.Assert(err, ErrorMatches, "invalid expression: not enough operands")
}

func (s *MySuite) TestCalculatePostfixSymbolError1(cs *C) {
	_, err := CalculatePostfix("--4 2 - 3bd * 5 +")
	cs.Assert(err, ErrorMatches, "invalid expression: unsupported ssymbol")
}

func (s *MySuite) TestCalculatePostfixSymbolError2(cs *C) {
	_, err := CalculatePostfix("2 a +")
	cs.Assert(err, ErrorMatches, "invalid expression: unsupported symbol")
}

func (s *MySuite) TestCalculatePostfixExtraOperandsError(cs *C) {
	_, err := CalculatePostfix("2 3 5 +")
	cs.Assert(err, ErrorMatches, "invalid expression: incorrect number of operands")
}

func ExampleCalculatePostfix() {
	res, _ := CalculatePostfix("2 2 +")
	fmt.Println(res)
}
