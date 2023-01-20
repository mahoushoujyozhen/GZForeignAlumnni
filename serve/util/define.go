package util

import (
	"errors"
)

var ERR_SNOWFLAKE_PARAM = errors.New("the snowflake params are not valid")
var ERR_CLOCK_BACKWARDS = errors.New("the clock is moved backwards")

