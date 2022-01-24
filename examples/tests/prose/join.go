package prose

import "strings"

func JoinWithCommas(phrase []string) string {
	if len(phrase) == 0 {
		return ""
	} else if len(phrase) == 1 {
		return phrase[0]
	} else if len(phrase) == 2 {
		return phrase[0] + " and " + phrase[1]
	} else {
		result := strings.Join(phrase[:len(phrase)-1], ", ")
		result += ", and "
		result += phrase[len(phrase)-1]
		return result
	}
}
