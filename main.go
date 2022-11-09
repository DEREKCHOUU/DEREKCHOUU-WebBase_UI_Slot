package main

import (
	_Config "go/UI/config"
	"go/UI/insider/app"
	"log"
)

func main() {
	var cfg, err = _Config.NewConfig()
	if err != nil {
		log.Fatal("Config global missing: ", err)
	}
	app.Run(cfg)

}

//func stratuprun() {}

//func selfrestart() {}

// func initDB() {
// 	network, err := os.ReadFile("./UI/json/docker.json")
// 	if err != nil {
// 		log.Fatal("Error when opening file: ", err)
// 	}
// 	err = json.Unmarshal(network, &windows.Windowsseting)
// 	if err != nil {
// 		log.Fatal("Error during Unmarshal(): ", err)
// 	}
// 	dburl := fmt.Sprintf("%s:%s@tcp(%s:%s)/UI", windows.Windowsseting.Sqlacc, windows.Windowsseting.Sqlpw, windows.Windowsseting.Ip, windows.Windowsseting.Sqlport)
// 	val := url.Values{}
// 	val.Add("parseTime", "1")
// 	val.Add("loc", windows.Windowsseting.UTC)
// 	dsn := fmt.Sprintf("%s?%s", dburl, val.Encode())
// 	err = _mysql.Opensql(dsn)
// 	if err != nil {
// 		log.Fatal("sql connect fail: ", err)
// 	}
// }
