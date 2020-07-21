package models
import (
	appconfig "Test_Example/conf"
	"log"
	"sync"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

var (
	bigsetIf StringBigsetService.StringBigsetServiceIf
	Once     sync.Once
)

func InitBigSetIf() {
	log.Println(appconfig.Config, "appconfig.Config")
	Once.Do(func() {
		log.Println(appconfig.Config, "appconfig.Config")
		bigsetIf = StringBigsetService.NewStringBigsetServiceModel(appconfig.Config.SourceMappingNewSID,
			appconfig.Config.EtcdServerEndpoints,
			GoEndpointBackendManager.EndPoint{
				Host:      appconfig.Config.SourceMappingNewHost,
				Port:      appconfig.Config.SourceMappingNewPort,
				ServiceID: appconfig.Config.SourceMappingNewSID,
			})
		log.Println(bigsetIf, "bigsetIf")
	})
}

func GetBigsetIf() StringBigsetService.StringBigsetServiceIf {
	return bigsetIf
}