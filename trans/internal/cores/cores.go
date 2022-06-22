package cores

import (
	"github.com/dtm-labs/dtmcli"
	"xa/trans/internal/config"
)

func ToConfig(mysql config.Mysql) dtmcli.DBConf {
	return dtmcli.DBConf{
		Driver:   mysql.Driver,
		Host:     mysql.Host,
		Port:     mysql.Port,
		User:     mysql.User,
		Password: mysql.Password,
		Db:       mysql.Db,
	}
}
