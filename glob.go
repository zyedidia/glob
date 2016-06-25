package glob

import "regexp"

type Glob struct {
	*regexp.Regexp
}

func Compile(pattern string) (*Glob, error) {
	r, err := globToRegex(pattern)
	return &Glob{r}, err
}

func globToRegex(glob string) (*regexp.Regexp, error) {
	regex := ""
	inGroup := 0
	inClass := 0
	firstIndexInClass := -1
	arr := []byte(glob)

	for i := 0; i < len(arr); i++ {
		ch := arr[i]

		switch ch {
		case '\\':
			i++
			if i >= len(arr) {
				regex += "\\"
			} else {
				next := arr[i]
				switch next {
				case ',':
					// Nothing
				case 'Q', 'E':
					regex += "\\\\"
				default:
					regex += "\\"
				}
				regex += string(next)
			}
		case '*':
			if inClass == 0 {
				regex += ".*"
			} else {
				regex += "*"
			}
		case '?':
			if inClass == 0 {
				regex += "."
			} else {
				regex += "?"
			}
		case '[':
			inClass++
			firstIndexInClass = i + 1
			regex += "["
		case ']':
			inClass--
			regex += "]"
		case '.', '(', ')', '+', '|', '^', '$', '@', '%':
			if inClass == 0 || (firstIndexInClass == i && ch == '^') {
				regex += "\\"
			}
			regex += string(ch)
		case '!':
			if firstIndexInClass == i {
				regex += "^"
			} else {
				regex += "!"
			}
		case '{':
			inGroup++
			regex += "("
		case '}':
			inGroup--
			regex += ")"
		case ',':
			if inGroup > 0 {
				regex += "|"
			} else {
				regex += ","
			}
		default:
			regex += string(ch)
		}
	}
	return regexp.Compile(regex)
}
