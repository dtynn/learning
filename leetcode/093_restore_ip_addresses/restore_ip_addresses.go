package restore_ip_addresses

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	appendIP([]string{}, s, &res)
	return res
}

func appendIP(prefix []string, left string, res *[]string) {
	if len(prefix) == 4 && len(left) == 0 {
		*res = append(*res, strings.Join(prefix, "."))
		return
	}

	if !(len(prefix) < 4 && len(left) > 0) {
		return
	}

	if isValid(left[:1]) {
		nextPrefix := make([]string, len(prefix)+1)
		copy(nextPrefix, prefix)
		nextPrefix[len(prefix)] = left[:1]
		appendIP(nextPrefix, left[1:], res)
	}

	if len(left) >= 2 && isValid(left[:2]) {
		nextPrefix := make([]string, len(prefix)+1)
		copy(nextPrefix, prefix)
		nextPrefix[len(prefix)] = left[:2]
		appendIP(nextPrefix, left[2:], res)
	}

	if len(left) >= 3 && isValid(left[:3]) {
		nextPrefix := make([]string, len(prefix)+1)
		copy(nextPrefix, prefix)
		nextPrefix[len(prefix)] = left[:3]
		appendIP(nextPrefix, left[3:], res)
	}
}

func isValid(part string) bool {
	if len(part) == 0 || len(part) > 3 {
		return false
	}

	if part == "0" {
		return true
	}

	if len(part) == 2 && part[:1] == "0" {
		return false
	}

	if len(part) == 3 && (part[:1] == "0" || part[:2] == "00") {
		return false
	}

	i, err := strconv.Atoi(part)
	if err != nil {
		return false
	}

	return i > 0 && i < 256
}
