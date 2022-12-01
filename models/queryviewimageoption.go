package models

import (
	"fmt"
	"net/url"
	"strings"
)

type QueryViewImageOption struct {
	maxAge     int
	resolution string
	params     map[string]string
}

func (o *QueryViewImageOption) GetMaxAge() int {
	if o.maxAge >= minMaxAge {
		return o.maxAge
	}

	return defaultMaxAge
}

func (o *QueryViewImageOption) SetMaxAge(age int) {
	if age > minMaxAge {
		o.maxAge = age
		return
	}

	o.maxAge = minMaxAge
}

func (o *QueryViewImageOption) GetImageResolution() string {
	switch o.resolution {
	case ImageResolutionHigh, ImageResolutionLow:
		return o.resolution
	default:
		return ImageResolutionHigh
	}
}

func (o *QueryViewImageOption) SetImageResolution(resolution string) {
	switch resolution {
	case ImageResolutionHigh, ImageResolutionLow:
		o.resolution = resolution
	default:
		o.resolution = ImageResolutionHigh
	}
}

func (o *QueryViewImageOption) AddParam(key, value string) {
	o.params[key] = value
}

func (o *QueryViewImageOption) DeleteParam(key string) string {
	value := ""
	if v, ok := o.params[key]; ok {
		value = v
		delete(o.params, key)
	}

	return value
}

func (o *QueryViewImageOption) Encode() string {
	params := strings.Builder{}
	for key, value := range o.params {
		params.WriteString("&")
		params.WriteString(url.PathEscape(fmt.Sprintf("%s=%s", key, value)))
	}

	return fmt.Sprintf("?resolution=%s&maxAge=%d%s",
		o.GetImageResolution(),
		o.GetMaxAge(),
		params.String(),
	)
}
