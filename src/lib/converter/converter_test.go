package converter

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeToUnixSecond(t *testing.T) {

	ctx := context.Background()
	var testTimeConversion = map[string]struct {
		inputdate   string
		inputloc    string
		inputformat string
		output      interface{}
		err         interface{}
	}{
		"correctConversion": {
			"2018-07-21 10:17:00",
			"Asia/Jakarta",
			"2006-01-02 15:04:05",
			1532143020,
			nil,
		},
		"nilLocationConversion": {
			"2018-07-21 10:17:00",
			"",
			"2006-01-02 15:04:05",
			1532143020,
			nil,
		},
		"nilFormatConversion": {
			"2018-07-21 10:17:00",
			"Asia/Jakarta",
			"",
			1532143020,
			nil,
		},
		"errorLocationConversion": {
			"",
			"FalseLocation",
			"2006-01-02 15:04:05",
			-1,
			errors.New("Error Location"),
		},
		"errorParseConversion": {
			"2006-01-02 1",
			"",
			"2006-01-02 15:04:05",
			-1,
			errors.New("Incorrect Parsing"),
		},
	}

	for testName, test := range testTimeConversion {
		t.Log("Running ", testName)
		result, err := TimeToUnixSecond(ctx, test.inputdate, test.inputformat, test.inputloc)

		assert.Equal(t, test.err, err)
		assert.Equal(t, test.output, result)
	}
}
