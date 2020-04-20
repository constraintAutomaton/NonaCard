package remotecommuncation

// RemotreCommuncation interface of module that communicate to remote source
type RemoteCommuncation interface {
	Fetch(query ParameterQuery) error
}

// ParameterQuery parameter of remote communication query data is send to the out parameter
type ParameterQuery struct {
	Url           string
	Query         string
	Variables     *map[string]string
	Out           interface{}
	Authorization []string
}
