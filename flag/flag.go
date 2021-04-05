package flag

import "flag"

// IsTesting is testing
func IsTesting() bool {
	return flag.Lookup("test.v") != nil
}
