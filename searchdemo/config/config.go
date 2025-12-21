package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var Cfg = &Config{}

type Config struct {
	App           App           `mapstructure:"app"`
	Mysql         Mysql         `mapstructure:"mysql"`
	MongoDB       MongoDB       `mapstructure:"mongodb"`
	Elasticsearch Elasticsearch `mapstructure:"elasticsearch"`
	Redis         Redis         `mapstructure:"redis"`
}

type App struct {
	AppSignExpire   int64         `mapstructure:"app_sign_expire" yaml:"app_sign_expire"`
	RunMode         string        `mapstructure:"run_mode" yaml:"run_mode"`
	HttpPort        int           `mapstructure:"http_port" yaml:"http_port"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout" yaml:"read_timeout"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout" yaml:"write_timeout"`
	RuntimeRootPath string        `mapstructure:"runtime_root_path" yaml:"runtime_root_path"`
	AppLogPath      string        `mapstructure:"app_log_path" yaml:"app_log_path"`
}

type Mysql struct {
	Dbname            string        `mapstructure:"dbname"`
	User              string        `mapstructure:"user"`
	Password          string        `mapstructure:"password"`
	Host              string        `mapstructure:"host"`
	MaxOpenConn       int           `mapstructure:"max_open_conn"`
	MaxIdleConn       int           `mapstructure:"max_idle_conn"`
	ConnMaxLifeSecond time.Duration `mapstructure:"conn_max_life_second"`
	TablePrefix       string        `mapstructure:"table_prefix"`
}

type MongoDB struct {
	DBname   string `mapstructure:"dbname"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
}

type Elasticsearch struct {
	Url            string `mapstructure:"url"`
	User           string `mapstructure:"user"`
	Password       string `mapstructure:"password"`
	BulkActionNum  int    `mapstructure:"bulk_action_num"`
	BulkActionSize int    `mapstructure:"bulk_action_size"`
	BulkWorkersNum int    `mapstructure:"bulk_workers_num"`
}

type Redis struct {
	Host        string `mapstructure:"host"`
	Db          string `mapstructure:"db"`
	Password    string `mapstructure:"password"`
	MinIdleConn int    `mapstructure:"min_idle_conn"`
	PoolSize    int    `mapstructure:"pool_size"`
	MaxRetries  int    `mapstructure:"max_retries"`
}

func LoadConfig() {
	viper := viper.New()
	// 1 设置配置文件路径
	viper.SetConfigFile("config.yml")

	// 2、读取配置
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	//3.将配置映射成结构体
	if err := viper.Unmarshal(&Cfg); err != nil {
		logrus.Error(err)
		panic(err)
	}

	//4.监听配置文件变动,重新解析配置（静态文件配置一般不推荐）
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file change", e.Name)
		if err := viper.Unmarshal(&Cfg); err != nil {
			logrus.Error(err)
			panic(err)
		}
	})
}
