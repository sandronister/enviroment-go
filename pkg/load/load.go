package load

func (r *environment) Load(vars ...interface{}) []error {
	var errs []error
	for _, v := range vars {
		err := r.LoadVariable(v)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}
