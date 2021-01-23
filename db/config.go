package db

type dbConfigModel struct {
	mongoUri  string
	mongoDb   string
	ServerUri string
}

var DbConfig = dbConfigModel{
	mongoUri:  "mongodb://mongo:27017/CloudWriteAPI",
	mongoDb:   "CloudWriteAPI",
	ServerUri: ":8080",
}
