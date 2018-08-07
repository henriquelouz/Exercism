package bob

import "strings"

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Hey returns bob's response depending on what's being said
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)
	remarkUpper := strings.ToUpper(remark)

	if remarkUpper == remark && strings.HasSuffix(remark, "?") && strings.ContainsAny(remarkUpper, alphabet) {
		return "Calm down, I know what I'm doing!"
	} else if strings.ContainsAny(remarkUpper, alphabet) && remarkUpper == remark {
		return "Whoa, chill out!"
	} else if strings.HasSuffix(remark, "?") {
		return "Sure."
	} else if remark == "" {
		return "Fine. Be that way!"
	}
	return "Whatever."
}
