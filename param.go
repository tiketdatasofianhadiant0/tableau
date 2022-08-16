package tableau

import (
	"fmt"
	"strings"
)

type Param map[string]any

func (p Param) String() string {
	const ampersand = `&`

	var lines []string
	for k, v := range p {
		line := strings.TrimSpace(k) + "=" + fmt.Sprint(v)
		lines = append(lines, line)
	}

	return "?" + strings.Join(lines, ampersand)
}
