package json

import (
	"encoding/json"
	"fmt"
	"os"
)

func Json2Map(s []byte) map[string]interface{} {

	map2 := make(map[string]interface{})
	err := json.Unmarshal(s, &map2)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	return map2
}
