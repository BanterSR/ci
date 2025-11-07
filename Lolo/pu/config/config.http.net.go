package config

type HttpNet struct {
	InnerIp   string `json:"InnerIp"`
	InnerPort string `json:"InnerPort"`
}

var defaultHttpNet = &HttpNet{
	InnerIp:   "0.0.0.0",
	InnerPort: "18881",
}

func GetHttpNet() *HttpNet {
	conf := GetConfig()
	if conf.HttpNet == nil {
		return defaultHttpNet
	}
	return conf.HttpNet
}

func (x *HttpNet) GetInnerIp() string {
	return x.InnerIp
}

func (x *HttpNet) GetInnerPort() string {
	return x.InnerPort
}
