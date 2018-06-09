package simplify_path

import (
	"strings"
)

func simplifyPath(path string) string {
	ps := strings.Split(path, "/")
	p := make([]string, 0, len(ps))
	for _, one := range ps {
		switch one {
		case "", ".":

		case "..":
			if len(p) > 0 {
				p = p[:len(p)-1]
			}

		default:
			p = append(p, one)
		}
	}

	return "/" + strings.Join(p, "/")
}
