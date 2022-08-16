package models

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

type Filter map[string]any

func (f Filter) String() string {
	const comma = ","
	var pairs []string

	for k, v := range f {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			var data []string

			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				data = append(data, url.QueryEscape(fmt.Sprint(s.Index(i))))
			}

			val := strings.Join(data, comma)
			pairs = append(pairs,
				fmt.Sprintf("%s:in:[%s]", k, val))

		default:
			val := url.QueryEscape(fmt.Sprint(v))
			pairs = append(pairs,
				fmt.Sprintf("%s:eq:%s", k, val))
		}
	}

	return strings.Join(pairs, comma)
}
