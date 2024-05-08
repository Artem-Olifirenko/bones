package main

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFatalJsonLogger(t *testing.T) {
	tests := []struct {
		name  string
		error error
		msg   string
	}{
		{
			name: "empty",
		},
		{
			name:  "normal strings",
			error: errors.New("test error"),
			msg:   "test msg",
		},
		{
			name:  "strings with quotes and slashes",
			error: errors.New(`some \\ and \ and \" and "test"`),
			msg:   `some \\ and \ and \" and "test"`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			str := fatalJsonLog(test.msg, test.error)

			var data map[string]string
			err := json.Unmarshal([]byte(str), &data)

			assert.NoError(t, err)
			assert.Equal(t, test.msg, data["msg"])

			if test.error == nil {
				assert.Equal(t, "", data["error"])
			} else {
				assert.Equal(t, test.error.Error(), data["error"])
			}
		})
	}
}
