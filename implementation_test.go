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

	// simple tests

	result, err := CalculatePostfix("2 2 +")
	cs.Assert(result, Equals, "4")

	result, err = CalculatePostfix("3 5 ^")
	cs.Assert(result, Equals, "243")

	// more operands

	result, err = CalculatePostfix("4 5 + 5 *")
	cs.Assert(result, Equals, "45")

	result, err = CalculatePostfix("4 2 - 3 * 5 +")
	cs.Assert(result, Equals, "11")

	result, err = CalculatePostfix("2 3 ^ 2 + 2 *")
	cs.Assert(result, Equals, "20")

	// complex tests

	result, err = CalculatePostfix("5 2 + 3 * 4 - 2 / 6 + 3 +")
	cs.Assert(result, Equals, "17")

	// validation tests

	result, err = CalculatePostfix("")
	cs.Assert(err, ErrorMatches, "invalid expression: empty input")

	result, err = CalculatePostfix("2 +")
	cs.Assert(err, ErrorMatches, "invalid expression: not enough operands")

	result, err = CalculatePostfix("--4 2 - 3bd * 5 +")
	cs.Assert(err, ErrorMatches, "invalid expression: unsupported symbol")

	result, err = CalculatePostfix("2 a +")
	cs.Assert(err, ErrorMatches, "invalid expression: unsupported symbol")

	result, err = CalculatePostfix("2 3 5 +")
	cs.Assert(err, ErrorMatches, "invalid expression: incorrect number of operands")
}

func ExampleCalculatePostfix() {
	res, _ := CalculatePostfix("2 2 +")
	fmt.Println(res)
}
