package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

// GetPrelude decodes and remove the prelude from the input string.
func GetPrelude(s string) (string, map[string]interface{}, error) {

	var v map[string]interface{}
	k := strings.Split(string(s), "\n")
	i := 0
	prelude := ""
	if len(k) > 1 && k[0] == "---" {
		i++
		for _, l := range k[1:] {
			i++
			if l == "---" {
				break
			}
			z := strings.Index(l, ":")
			name := l[0:z]
			val := l[z+1:]
			val = strings.TrimLeft(val, " ")
			// quote raw values, it allows easier writing of strings containing strings.
			if val[0] != '[' && val[0] != '"' && val[0] != '\'' {
				val = fmt.Sprintf("%q", val)
			}
			prelude += fmt.Sprintf("%q:%v,\n", name, val)
		}

		prelude = "{" + prelude[:len(prelude)-2] + "}"

		if err := json.Unmarshal([]byte(prelude), &v); err != nil {
			return "", v, err
		}
	}

	// rebuild the tempalte without prelude
	c := ""
	for _, l := range k[i:] {
		c += fmt.Sprintf("%v\n", l)
	}

	return c, v, nil
}
