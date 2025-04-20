package dify

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
)

type sseStream[T any] struct {
	reader *bufio.Reader
	closed bool
}

func newSseStream[T any](reader io.Reader) *sseStream[T] {
	return &sseStream[T]{
		reader: bufio.NewReader(reader),
	}
}

func (s *sseStream[T]) Close() error {
	s.closed = true
	return nil
}

func (s *sseStream[T]) Closed() bool {
	return s.closed
}

var sseEventDataPrefix = []byte("data: ")

func (s *sseStream[T]) readOneLineRaw() ([]byte, error) {
	if s.Closed() {
		return nil, io.EOF
	}

	lineBytes, _, err := s.reader.ReadLine()
	if err != nil {
		return nil, err
	}
	if !bytes.HasPrefix(lineBytes, sseEventDataPrefix) || len(lineBytes) <= len(sseEventDataPrefix) {
		var errMessage []byte
		if len(lineBytes) <= len(sseEventDataPrefix) {
			errMessage = lineBytes
		} else {
			errMessage = lineBytes[len(sseEventDataPrefix):]
		}
		return nil, errors.Errorf("data format error, start with `%s`, length %d", errMessage, len(lineBytes))
	}

	return lineBytes[len(sseEventDataPrefix):], nil
}

func (s *sseStream[T]) readOneLine() (T, error) {
	var response T
	lineBytes, err := s.readOneLineRaw()
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(lineBytes, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
