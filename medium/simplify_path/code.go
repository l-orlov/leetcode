package simplify_path

import "strings"

func simplifyPath(path string) string {
	res := make([]byte, 0, len(path))

	i := 0
	for i < len(path) {
		if path[i] == '/' {
			j := i + 1
			for j < len(path) && path[j] != '/' {
				j++
			}
			num := j - i
			if num == 1 {
				if len(res) == 0 && j == len(path) {
					res = append(res, '/')
				}
				i++
			} else if num == 2 {
				if path[i+1] != '.' {
					res = append(res, path[i], path[i+1])
				}
				i += 2
			} else if num == 3 {
				if path[i+1] == '.' && path[i+2] == '.' {
					// delete element
					if len(res) > 1 {
						slashIdx := len(res) - 1
						for slashIdx >= 0 && res[slashIdx] != '/' {
							slashIdx--
						}
						res = res[:slashIdx]
					}
				} else {
					res = append(res, path[i], path[i+1], path[i+2])
				}
				i += 3
			} else {
				for i < j {
					res = append(res, path[i])
					i++
				}
			}
		} else {
			i++
		}
	}

	if len(res) == 0 {
		res = append(res, '/')
	}

	return string(res)
}

func simplifyPathV2(path string) string {
	segments := strings.Split(path, "/")
	var stack []string

	for _, segment := range segments {
		switch segment {
		case "", ".":
			continue
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, segment)
		}
	}

	return "/" + strings.Join(stack, "/")
}
