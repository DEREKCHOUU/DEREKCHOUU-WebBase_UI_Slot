//  "dbUser:dbPass@tcp(dbHost:dbPort)/dbName"
1.	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
//  "dbUser:dbPass@tcp(dbHost:dbPort)/dbName?parseTime=1&loc=Asia/Jakarta"
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())