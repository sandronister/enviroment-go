package load

import (
	"encoding/json"
	"fmt"
)

func (r *environment) LoadList(field string, result interface{}) error {
	if field == "" {
		return fmt.Errorf("field cannot be empty")
	}

	if _, ok := r.variables[field]; !ok {
		return fmt.Errorf("field %s not found", field)
	}

	err := json.Unmarshal([]byte(r.variables[field]), result)

	if err != nil {
		return fmt.Errorf("cannot unmarshal %s: %v", r.variables[field], err)
	}

	return nil
}
