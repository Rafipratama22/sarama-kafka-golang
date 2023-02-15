package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodul/envy"
)

type databaseItem struct {
	Drivername     string
	Datasourcename string
}

type database struct {
	MySQL databaseItem
}

var DB = database{
	MySQL: databaseItem{
		Drivername: envy.Get("MYSQL_DRIVER", "mysql"),
		Datasourcename: fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?parseTime=true&charset=utf8",
			envy.Get("MYSQL_USERNAME", "root"),
			envy.Get("MYSQL_PASSWORD", "abcd"),
			envy.Get("MYSQL_HOST", "10.5.99.203"),
			envy.Get("MYSQL_PORT", "3306"),
			envy.Get("MYSQL_NAME", "mncmbank"),
		),
	},
}
