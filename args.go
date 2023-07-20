package sqlz

import (
	"fmt"
	"strings"
)

type Args []any

func (a *Args) Add(value any, cond ...bool) string {
	if len(cond) > 0 && !cond[0] {
		return ""
	}
	*a = append(*a, value)
	return fmt.Sprintf("$%d", len(*a))
}

func (a *Args) Addf(s string, value any, cond ...bool) string {
	if len(cond) > 0 && !cond[0] {
		return ""
	}
	return strings.ReplaceAll(s, "$?", a.Add(value))
}
