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
	result, err := CalculatePostfix("2 2 +")
	cs.Assert(result, Equals, "4")

	result, err = CalculatePostfix("4 5 + 5 *")
	cs.Assert(result, Equals, "45")

	result, err = CalculatePostfix("4 2 - 3 * 5 +")
	cs.Assert(result, Equals, "11")

	result, err = CalculatePostfix("2 3 ^ 2 + 2 *")
	cs.Assert(result, Equals, "20")

	result, err = CalculatePostfix("--4 2 - 3bd * 5 +")
	cs.Assert(err, ErrorMatches, "invalid expression: not enough operands")
}

func ExampleCalculatePostfix() {
	res, _ := CalculatePostfix("2 2 +")
	fmt.Println(res)
}
