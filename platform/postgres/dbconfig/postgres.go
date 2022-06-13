package dbconfig

import "fmt"

const PostgresDriver = "postgres"

const User = "postgres"

const Host = "localhost"

const Port = "5432"

const Password = "vapordev123"

const DbName = "MobyDick"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
