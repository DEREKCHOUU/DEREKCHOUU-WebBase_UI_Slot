package app

import (
	"encoding/hex"
	"os"

	_global "go/UI/config"
	"go/UI/insider/routing"
	_mysql "go/UI/outsider/SQLDB/MYSQL"
	"go/UI/outsider/logger"

	"github.com/denisbrodbeck/machineid"
	"github.com/gin-gonic/gin"
)

func Run(cfg *_global.Global) {
	i := logger.New(cfg.Log.Level)

	//DB open
	DB, err := _mysql.Opensql(cfg.DB, cfg.App.UTC)
	if err != nil {
		i.Fatal("sql connect fail: ", err)
	}
	defer func() {
		err := DB.Close()
		if err != nil {
			i.Fatal(err)
		}
	}()

	//usecasehere
	//uiusecase := _os.NewWindowSetup(cfg.Os, cfg.App)

	//RibbitMQ for Device
	//mq := amqprpc.NewRouter(uiusecase)

	//router engine start
	router := gin.Default()

	//license key check
	lock := "_endme"
	id, _ := machineid.ID()
	key, _ := os.ReadFile("./os/key.bin")
	keystr := hex.EncodeToString(key)
	if lock+id != keystr {
		router.Use( /*forcedockerfunc*/ )
	}

	routing.Routersetup(router)
}
