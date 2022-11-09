package routing

import (
	_os "go/UI/insider/usecase/osset/windows"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MiddleDocker() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Redirect(http.StatusOK, "/setting/docker/")
	}
}

func docker(c *gin.Context) {
	c.HTML(http.StatusOK, "docker.html", _os.Windowsseting)
}

func setip(c *gin.Context) {
	ipv4 := c.DefaultQuery("ip", "192.168.0.251")
	prefixLength := c.DefaultQuery("subnet", "24")
	gateway := c.DefaultQuery("gateway", "192.168.0.1")
	DHCPS := c.DefaultQuery("DHCP", "false")
	DHCP, err := strconv.ParseBool(DHCPS)
	if err != nil {
		print("DHCP injection")
	}
	_os.Windowsseting.SetDHCP(DHCP)
	_os.Windowsseting.Setipv4(ipv4, prefixLength, gateway)
}

func setsql(c *gin.Context) {
	//sqltype := c.DefaultQuery("sqltype", "mysql")
	//sqlport := c.DefaultQuery("sqlport", "3306")
	//_os.Windowsseting.Sqltype = sqltype
	//_os.Windowsseting.Sqlport = sqlport
}

func rebuildDB(c *gin.Context) {

}
