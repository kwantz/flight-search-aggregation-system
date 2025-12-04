package utils

import (
	"fmt"
	"strings"
	"time"
)

func FormatDuration(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60

	if hours > 0 && minutes > 0 {
		return fmt.Sprintf("%dh %dm", hours, minutes)
	}

	if hours > 0 {
		return fmt.Sprintf("%dh", hours)
	}

	if minutes > 0 {
		return fmt.Sprintf("%dm", minutes)
	}

	return "0m"
}

func FormatCurrency(currency string, amount int64) string {
	s := fmt.Sprintf("%d", amount)
	n := len(s)

	if n <= 3 {
		return fmt.Sprintf("%s %s", currency, s)
	}

	var result strings.Builder
	pre := n % 3

	if pre > 0 { // Note: get 1-2 beginning nums
		result.WriteString(s[:pre])
		if n > pre {
			result.WriteString(".")
		}
	}

	for i := pre; i < n; i += 3 {
		result.WriteString(s[i : i+3])
		if i+3 < n {
			result.WriteString(".")
		}
	}

	return fmt.Sprintf("%s %s", currency, result.String())
}
