package sdk

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gucooing/lolo/config"
	"gucooing/lolo/gdconf"
)

func (s *Server) Router() {
	s.router.Any("/", HandleDefault)

	dispatch := s.router.Group("/dispatch")
	{
		dispatch.POST("/region_info", regionInfo)
	}
}

func HandleDefault(c *gin.Context) {
	c.String(200, "BanGK!")
}

type RegionInfoRequest struct {
	Version         string `form:"version" binding:"required"`
	Version2        string `form:"version2" binding:"required"`
	AccountType     string `form:"accountType" binding:"required"`
	OS              string `form:"os" binding:"required"`
	LastLoginSDKUID string `form:"lastloginsdkuid" binding:"required"`
}

type RegionInfo struct {
	Status           bool   `json:"status"`
	Message          string `json:"message"`
	GateTcpIp        string `json:"gate_tcp_ip"`
	GateTcpPort      int    `json:"gate_tcp_port"`
	IsServerOpen     bool   `json:"is_server_open"`
	Text             string `json:"text"`
	ClientLogTcpIp   string `json:"client_log_tcp_ip"`
	ClientLogTcpPort int    `json:"client_log_tcp_port"`
	CurrentVersion   string `json:"currentVersion"`
	PhotoShareCdnUrl string `json:"photo_share_cdn_url"`
}

func regionInfo(c *gin.Context) {
	var req RegionInfoRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request parameters",
		})
		return
	}

	conf := config.GetGateWay()
	info := &RegionInfo{
		Status:           true,
		Message:          "success",
		GateTcpIp:        conf.GetOuterIp(),
		GateTcpPort:      conf.GetOuterPort(),
		IsServerOpen:     true,
		Text:             "",
		ClientLogTcpIp:   config.GetLogServer().GetOuterIp(),
		ClientLogTcpPort: config.GetLogServer().GetOuterPort(),
		CurrentVersion:   gdconf.GetClientVersion(req.Version),
		PhotoShareCdnUrl: "https://cdn-photo-of.inutan.com/cn_prod_main",
	}

	c.JSONP(http.StatusOK, info)
}
