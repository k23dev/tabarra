package tabarra

import (
	"github.com/k23dev/go4it"
	"github.com/k23dev/go4it/interact"
)

type Tabarra struct {
	App               go4it.App
	IsCronJob         bool   `toml:"is_cron_job"`
	ExecutionInterval int    `toml:"execution_interval"`
	IsIntervalSecods  bool   `toml:"is_interval_seconds"`
	DataFile          string `toml:"data_file"`
	// Add instace of the datas struct
}

func NewTabarraApp() *Tabarra {

	// building configuration

	// Cargar la configuración de la app
	appconfigFile := "./config/appconfig"
	app_config := go4it.NewApp(appconfigFile)

	app_config.Connect2Db("local")
	app_config.DB.SetPrimaryDB(0)

	appTabarra := &Tabarra{
		App: app_config,
	}

	interact.ReadAndParseToml(appconfigFile+".toml", &appTabarra)

	// read datas config
	// configPath := "./config/"
	// fmt.Printf("Loading config file > %s \n", configPath+appTabarra.DataFile)
	// interact.ReadAndParseJson(configPath+appTabarra.DataFile, &appTabarra.DataInstance)

	return appTabarra

}
