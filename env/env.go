package env

import "os"

var DEV bool = os.Getenv("DEV") == "1"
var DB string = os.Getenv("DB")
