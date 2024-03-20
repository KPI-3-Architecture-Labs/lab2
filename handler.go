package lab2

import (
	"bytes"
	"io"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	buffer, readerErr := io.ReadAll(ch.Input)
	if readerErr != nil {
		return readerErr
	}
	buffer = bytes.Trim(buffer, "\x00")

	text := string(buffer)
	trimmed := strings.Trim(text, " \n")
	res, calcError := CalculatePostfix(trimmed)
	if calcError != nil {
		return calcError
	}

	_, writerErr := ch.Output.Write([]byte(res))
	if writerErr != nil {
		return writerErr
	}

	return nil
}
