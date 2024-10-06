package sort

import (
	"os"
)

var Debug bool

func init() {
	if os.Getenv("DEBUG") == "1" {
		Debug = true
	}
}
