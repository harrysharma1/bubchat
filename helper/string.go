package helper

import (
	"fmt"
	"strings"
)

// Limits string based on given limit.
func FirstN(s string, limit int) string {
	var sb strings.Builder

	for i := 0; i < len(s); i++ {
		if i < limit {
			sb.WriteString(string(s[i]))
		}
	}
	return sb.String()
}

// Converts string (which should be a uuid) to hex value for colour styling.
func HexFromUUID(s string) string {
	return fmt.Sprintf("#%s", s)
}
