package Settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*MongoConfig `mapstructure:"mongo"`
	*KafkaConfig `mapstructure:"kafka"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

type MongoConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type KafkaConfig struct {
	HostPort string `mapstructure:"hostport"`
}

type EsConfig struct {
	Addresses []string `mapstructure:"addresses"`
}

func Init() (err error) {
	viper.SetConfigFile("config.yaml") // 指定配置文件路径
	//viper.SetConfigName("config") // 配置文件名称(无扩展名)
	//viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项
	//viper.AddConfigPath(".")      // 查找配置文件所在的路径
	//viper.AddConfigPath("$HOME/.appname") // 多次调用以添加多个搜索路径
	//viper.AddConfigPath(".")   // 还可以在工作目录中查找配置
	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		fmt.Printf("viper.ReadInConfig failed:%v\n", err)
		return
	}
	//把读取的信息反序列化到Conf中
	if err := viper.Unmarshal(&Conf); err != nil {
		fmt.Printf("vip.Unmarshal failed,err:%v\n", err)
	}
	//实时监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:")
		if err := viper.Unmarshal(&Conf); err != nil {
			fmt.Printf("vip.Unmarshal failed,err:%v\n", err)
		}
	})
	return
}
