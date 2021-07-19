package serializer

import (
	"github.com/fatih/color"
)

type kv struct {
	key   string
	value string
}

// pprint takes in a title and a list of key value pairs and pretty prints it as a table with 2 columns.
func pprint(title string, kvs ...kv) {
	color.Red(title)
	for _, data := range kvs {
		color.Cyan("%-20s | %s", data.key, data.value)
	}
}
