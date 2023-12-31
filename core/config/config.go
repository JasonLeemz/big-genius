package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

var GlobalConfig *Config

type Config struct {
	App      AppCfg      `yaml:"app"`
	Database DatabaseCfg `yaml:"database"`
	Redis    RedisCfg    `yaml:"redis"`
	MQ       MQCfg       `yaml:"mq"`
	Log      LogCfg      `yaml:"log"`
	OpenAI   OpenAICfg   `yaml:"openai"`
	Proxy    ProxyCfg    `yaml:"proxy"`
	WeChat   WeChatCfg   `yaml:"wechat"`
}

type AppCfg struct {
	Port string `yaml:"port"`
}

type DatabaseCfg struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	DB       string `yaml:"db"`
}
type RedisCfg struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type MQCfg struct {
	Schema   string `yaml:"schema"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type LogCfg struct {
	LogLevel int    `yaml:"logLevel"`
	Path     string `yaml:"path"`
}

type OpenAICfg struct {
	ChatGPT struct {
		Enable  bool   `yaml:"enable"`
		Token   string `yaml:"token"`
		BaseURL string `yaml:"baseURL"`
	} `yaml:"chatgpt"`
	Azure struct {
		Enable      bool   `yaml:"enable"`
		Token       string `yaml:"token"`
		BaseURL     string `yaml:"baseURL"`
		Deployments string `yaml:"deployments"`
		ApiVersion  string `yaml:"apiVersion"`
	} `yaml:"azure"`
	TimeOut time.Duration `yaml:"timeout"`
}

type ProxyCfg struct {
	Schema string `yaml:"schema"`
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
}

type WeChatCfg struct {
	Token             string `yaml:"token"`
	CorpID            string `yaml:"corpID"`
	EncodingAesKey    string `yaml:"encodingAesKey"`
	CorpSecret        string `yaml:"corpSecret"`
	GetAccessTokenUrl string `yaml:"getAccessTokenUrl"`
	SendMsgUrl        string `yaml:"sendMsgUrl"`
	AgentID           int    `yaml:"agentid"`

	ProxyHost   string `yaml:"proxyHost"`
	SendMsgPath string `yaml:"sendMsgPath"`
}

func Init() {
	path := "./config/app.dev.yaml"
	viper.SetConfigFile(path) // 指定配置文件路径
	//viper.SetConfigName("config")         // 配置文件名称(无扩展名)
	//viper.SetConfigType("yaml")           // 如果配置文件的名称中没有扩展名，则需要配置此项
	//viper.AddConfigPath("/etc/appname/")  // 查找配置文件所在的路径
	//viper.AddConfigPath("$HOME/.appname") // 多次调用以添加多个搜索路径
	//viper.AddConfigPath(".")              // 还可以在工作目录中查找配置
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = viper.Unmarshal(&GlobalConfig)
	if err != nil {
		panic(fmt.Errorf("Failed to unmarshal config: %s \n", err))
		return
	}
}
