package lab2

import (
	"bytes"
	. "gopkg.in/check.v1"
	"strings"
)

type ComputeHandlerSuite struct{}

var _ = Suite(&ComputeHandlerSuite{})

func (s *ComputeHandlerSuite) TestCompute(cs *C) {
	input := strings.NewReader("2 3 + 5 *")
	output := new(bytes.Buffer)
	handler := ComputeHandler{
		Input:  input,
		Output: output,
	}
	err := handler.Compute()
	cs.Assert(output.String(), Equals, "25")
	cs.Assert(err, IsNil)
}

func (s *ComputeHandlerSuite) TestComputeError(c *C) {
	input := strings.NewReader("2 3 + invalid input")
	output := new(bytes.Buffer)
	handler := ComputeHandler{
		Input:  input,
		Output: output,
	}
	err := handler.Compute()
	c.Assert(err, NotNil)
}
