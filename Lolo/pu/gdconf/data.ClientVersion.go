package gdconf

type ClientVersion struct {
	VersionMap map[string]string `json:"VersionMap"`
}

func (g *GameConfig) loadClientVersion() {
	g.ClientVersion = new(ClientVersion)
	ReadJson(g.dataPath, "ClientVersion.json", &g.ClientVersion)
}

func GetClientVersion(v string) string {
	return cc.ClientVersion.VersionMap[v]
}
