package settings

var Local *local

type local struct {
	AppId  string
	Facets []string
}

func init() {
	Local = &local{}
}
