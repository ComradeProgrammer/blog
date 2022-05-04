package prettyprint

import (
	"encoding/json"
	"fmt"
)

func PrintJsonln(obj interface{}) {
	data, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(data))
	}
}
