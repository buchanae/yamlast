package yamlast

import "fmt"

type yamlError struct {
	err error
}

func failf(format string, args ...interface{}) {
	panic(yamlError{fmt.Errorf("yaml: "+format, args...)})
}
