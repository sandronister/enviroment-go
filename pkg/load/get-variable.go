package load

import (
	"fmt"
)

func (r *environment) GetVariable(key string) (string, error) {
	value, ok := r.variables[key]
	if !ok {
		return "", fmt.Errorf("variable %s not found", key)
	}

	return value, nil
}
