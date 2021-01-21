package app

import (
	"fmt"
	"strings"
)

func subVariables(s string, variables map[string]string) string {
	for varName, varVal := range variables {
		searchTarget := fmt.Sprintf("{%s}", strings.ToLower(varName))
		idx := strings.Index(strings.ToLower(s), searchTarget)
		s = s[:idx] + varVal + s[idx+len(searchTarget):]
	}
	return s
}