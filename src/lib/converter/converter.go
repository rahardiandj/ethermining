package converter

import (
	"context"
	"time"

	"github.com/rahardiandj/ethermining/src/constant"
)

//TimeToUnixSecond : Convert formatted datetime to unix second
func TimeToUnixSecond(ctx context.Context, date, format, location string) (int64, error) {
	if format == "" {
		format = constant.DateTimeFormat
	}

	if location == "" {
		location = constant.DefaultLocation
	}

	loc, err := time.LoadLocation(location)

	if err != nil {
		return -1, err
	}

	t, err := time.ParseInLocation(format, date, loc)
	if err != nil {
		return -1, err
	}

	return t.UTC().UnixNano() / int64(time.Second), err
}
