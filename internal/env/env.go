package env

import "os"

var DEV bool = os.Getenv("DEV") == "1"
