package args

import "strings"

func Parse(args []string) map[string]string {
	argMap := make(map[string]string)
	for _, arg := range args {
		pair := strings.Split(arg, "=")
		if len(pair) == 2 {
			argMap[pair[0]] = pair[1]
		}
	}

	return argMap
}
