package common

import "errors"

var(
	HttpPostOnly = errors.New("wrong http request, post only")
	JsonFormatError = errors.New("wrong json format")
	ParamError = errors.New("parameter error")
)