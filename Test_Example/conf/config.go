package appconfig

type AppConfig struct {
	EtcdServerEndpoints []string

	SourceMappingNewSID      string
	SourceMappingNewHost     string
	SourceMappingNewPort     string
	SourceMappingNewProtocol string
}

var Config *AppConfig