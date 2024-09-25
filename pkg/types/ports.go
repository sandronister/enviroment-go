package types

type EnvironmentPorts interface {
	LoadVariable(result interface{}) error
	LoadList(field string, result interface{}) error
	GetVariable(key string) (string, error)
}
