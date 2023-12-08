package utils

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var (
	Prefix   = "netdoop"
	MaxProcs = 1

	VerboseMode = false
	DebugMode   = false
	ConfigFile  = ""

	WorkDir           = "/opt/netdoop"
	DataPath          = "data"
	ConfPath          = "conf"
	DefaultConfigFile = WorkDir + "/" + ConfPath + "/netdoop.yaml"
)
var envOnce sync.Once

func initEnv() {
	MaxProcs = runtime.GOMAXPROCS(0)
	if MaxProcs < 1 {
		MaxProcs = 1
	}

	v := viper.GetViper()
	if Prefix != "" {
		v.SetEnvPrefix(Prefix)
	}
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if ConfigFile != "" && strings.HasSuffix(ConfigFile, ".yaml") {
		v.SetConfigType("yaml")
		v.SetConfigFile(ConfigFile)
	}
	if _, err := os.Stat(ConfigFile); err == nil {
		if err := v.ReadInConfig(); err != nil {
			fmt.Println("read config: " + err.Error())
		}
	}

	v.SetDefault("work_dir", "/opt/netdoop")
	v.SetDefault("data_path", "data")
	v.SetDefault("conf_path", "conf")
	WorkDir = v.GetString("work_dir")
	DataPath = v.GetString("data_path")
	ConfPath = v.GetString("conf_path")

	v.SetDefault("app_name", "NetDoop")
	v.SetDefault("secret", "ann2017")

	v.SetDefault("server_external_url", "https://dev.netdoop.com")
	v.SetDefault("server_address", ":9176")
	v.SetDefault("server_public_path", "/usr/share/netdoop/html")
	v.SetDefault("server_cors_enable", 0)
	v.SetDefault("server_basicauth_enable", 0)
	v.SetDefault("server_basicauth_username", "test")
	v.SetDefault("server_basicauth_password", "netdoop2023")

	v.SetDefault("admin_password", "ann2017")

	v.SetDefault("db_driver", "postgres")
	v.SetDefault("db_host", "pg")
	v.SetDefault("db_port", "5432")
	v.SetDefault("db_db", "netdoop")
	v.SetDefault("db_username", "netdoop")
	v.SetDefault("db_password", "ann2022")
	v.SetDefault("db_sslmode", "disable")
	v.SetDefault("db_max_idle_conns", 0)
	v.SetDefault("db_max_open_conns", 0)
	v.SetDefault("db_max_conn_lifetime", 0)

	v.SetDefault("mongodb_uri", "mongodb://netdoop:ann2022@mongodb:27017")
	v.SetDefault("mongodb_database", "netdoop")

	v.SetDefault("redis_host", "redis")
	v.SetDefault("redis_port", "6379")

	v.SetDefault("stun_addr", "0.0.0.0:3478")
	v.SetDefault("stun_network", "udp")

	v.SetDefault("s3_storage_type", "mongo")

	v.SetDefault("data_retention_period", 14*24*3600)
	v.SetDefault("iam_audit_log_enable", 1)

	v.SetDefault("omc_base_url", "http://43.143.43.123:9176")
	v.SetDefault("omc_upload_url", "http://43.143.43.123:9176/upload")
	v.SetDefault("omc_upload_username", "")
	v.SetDefault("omc_upload_password", "")

}

func GetEnv() *viper.Viper {
	envOnce.Do(initEnv)
	return viper.GetViper()
}
