package common

import "errors"

var(
	HttpGetOnly = errors.New("wrong http request, get only")
	HttpPostOnly = errors.New("wrong http request, post only")
	JsonFormatError = errors.New("wrong json format")
	ParamError = errors.New("parameter error")
)