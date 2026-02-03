package helper

import (
	"fmt"
	"strings"
)

func FirstN(s string, limit int) string {
	var sb strings.Builder

	for i := 0; i < len(s); i++ {
		if i < limit {
			sb.WriteString(string(s[i]))
		}
	}
	return sb.String()
}

func HexFromUUID(s string) string {
	return fmt.Sprintf("#%s", s)
}
