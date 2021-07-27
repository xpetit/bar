package bar

import (
	"fmt"
	"io"
	"math"
	"sort"
	"strings"
)

func Draw(width int, value, max float64) (s string) {
	if width < 0 || value < 0 || max < 0 {
		panic("the arguments can't be negative")
	}
	if value > max {
		value = max
	}
	barWidth := int(math.Round(float64(width) * 8 * value / max))
	s = strings.Repeat("█", barWidth/8)
	if rem := barWidth % 8; rem > 0 {
		s += map[int]string{
			1: "▏",
			2: "▎",
			3: "▍",
			4: "▌",
			5: "▋",
			6: "▊",
			7: "▉",
		}[rem]
		width--
	}
	return s + strings.Repeat(" ", width-(barWidth/8))
}

func Graph(w io.Writer, maxWidth int, data map[string]float64) {
	if maxWidth == 0 {
		maxWidth = 80
	}
	var keys []string
	var keyLen int
	for key := range data {
		keys = append(keys, key)
		if len(key) > keyLen {
			keyLen = len(key)
		}
	}
	availableWidth := maxWidth - keyLen - 1
	if availableWidth < 0 {
		panic("maximal width too small")
	}
	sort.Slice(keys, func(i, j int) bool {
		if data[keys[i]] == data[keys[j]] {
			return keys[i] < keys[j]
		}
		return data[keys[i]] > data[keys[j]]
	})
	for _, key := range keys {
		fmt.Fprintf(w, "%*s ", keyLen, key)
		io.WriteString(w, Draw(availableWidth, data[key], data[keys[0]]))
		fmt.Fprintln(w)
	}
}
