package helpers

import (
	"fmt"
	"strings"
)

func FormatRupiah(amount float32) string {
	formatted := fmt.Sprintf("%.2f", amount)
	parts := strings.Split(formatted, ".")
	integerPart := parts[0]

	var result strings.Builder
	for i, digit := range integerPart {
		if i > 0 && (len(integerPart)-i)%3 == 0 {
			result.WriteString(",")
		}

		result.WriteRune(digit)
	}

	return "Rp. " + result.String()
}
