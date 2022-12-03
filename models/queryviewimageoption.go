package models

import (
	"fmt"
	"net/url"
	"strings"
)

type QueryViewImageOption struct {
	MaxAge int
	Params map[string]string
}

func (o *QueryViewImageOption) GetMaxAge() int {
	if o.MaxAge >= minMaxAge {
		return o.MaxAge
	}

	return defaultMaxAge
}

func (o *QueryViewImageOption) SetMaxAge(age int) {
	if age > minMaxAge {
		o.MaxAge = age
		return
	}

	o.MaxAge = minMaxAge
}

func (o *QueryViewImageOption) AddParam(key, value string) {
	o.Params[key] = value
}

func (o *QueryViewImageOption) DeleteParam(key string) string {
	value := ""
	if v, ok := o.Params[key]; ok {
		value = v
		delete(o.Params, key)
	}

	return value
}

func (o *QueryViewImageOption) Encode() string {
	params := strings.Builder{}
	for key, value := range o.Params {
		params.WriteString("&")
		params.WriteString(url.PathEscape(fmt.Sprintf("%s=%s", key, value)))
	}

	return fmt.Sprintf("?resolution=%s&maxAge=%d%s",
		ImageResolutionHigh,
		o.GetMaxAge(),
		params.String(),
	)
}
