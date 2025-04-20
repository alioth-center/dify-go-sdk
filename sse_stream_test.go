package dify

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func Test_sseStream_readOneLineRaw(t *testing.T) {
	type testData struct {
		Content string `json:"content"`
	}
	t.Run("normal", func(t *testing.T) {
		testDataBytes := []byte(`data: {"content":"test1"}
data: {"content":"test2"}
`)
		s := newSseStream[testData](bytes.NewReader(testDataBytes))
		var actual []byte
		for {
			got, err := s.readOneLineRaw()
			if err != nil {
				if !errors.Is(err, io.EOF) {
					assert.NoError(t, err)
					return
				}
				break
			}
			actual = append(actual, got...)
			actual = append(actual, '\n')
		}
		assert.Equal(t, `{"content":"test1"}
{"content":"test2"}
`, string(actual))
	})

	t.Run("error header", func(t *testing.T) {
		testDataBytes := []byte(`ddd: {"content":"test1"}
`)
		s := newSseStream[testData](bytes.NewReader(testDataBytes))
		_, err := s.readOneLineRaw()
		assert.ErrorContains(t, err, "data format error")
	})

	t.Run("no blank line", func(t *testing.T) {
		testDataBytes := []byte(`data: {"content":"test1"}`)
		s := newSseStream[testData](bytes.NewReader(testDataBytes))
		actual, err := s.readOneLineRaw()
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, `{"content":"test1"}`, string(actual))
	})
}

func Test_sseStream_readOneLine(t *testing.T) {
	type testData struct {
		Content string `json:"content"`
	}
	t.Run("normal", func(t *testing.T) {
		testDataBytes := []byte(`data: {"content":"test1"}
data: {"content":"test2"}
`)
		s := newSseStream[testData](bytes.NewReader(testDataBytes))
		var actual []testData
		for {
			got, err := s.readOneLine()
			if err != nil {
				if !errors.Is(err, io.EOF) {
					assert.NoError(t, err)
					return
				}
				break
			}
			actual = append(actual, got)
		}
		assert.Equal(t, []testData{{Content: "test1"}, {Content: "test2"}}, actual)
	})
}
