package substitutions

import (
	"fmt"
	"strings"
)

func SubVariables(s string, variables map[string]string) string {
	for varName, varVal := range variables {
		searchTarget := fmt.Sprintf("{%s}", strings.ToLower(varName))

		idx := strings.Index(strings.ToLower(s), searchTarget)
		for idx != -1 {
			s = s[:idx] + varVal + s[idx+len(searchTarget):]
			idx = strings.Index(strings.ToLower(s), searchTarget)
		}
	}
	return s
}