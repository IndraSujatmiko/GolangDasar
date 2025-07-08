package database

var connection string

func init() {  //otomatis terpanggil ketika packagenya terpanggil
	connection = "MySql"
}

func GetDatabase() string {
	return connection
}
