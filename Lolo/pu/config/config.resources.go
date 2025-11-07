package config

type Resources struct {
	ExcelPath  string `json:"ExcelPath"`
	DataPath   string `json:"DataPath"`
	ConfigPath string `json:"ConfigPath"`
}

var defaultResources = &Resources{
	ExcelPath:  "./Resource/Excel",
	DataPath:   "./data",
	ConfigPath: "./Resource/Config",
}

func GetResources() *Resources {
	return GetConfig().Resources
}

func (x *Resources) GetExcelPath() string {
	return x.ExcelPath
}

func (x *Resources) GetDataPath() string {
	return x.DataPath
}

func (x *Resources) GetConfigPath() string {
	return x.ConfigPath
}
