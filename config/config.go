package config

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
	"github.com/spf13/viper"
	"gohash/utils/encode"
	"log"
	"os"
)

var (
	AppDebug            bool
	MysqlDns            string
	BotToken            string
	BotUsername         string
	AppName             string
	HashKey             string
	BotGameGroupID      int64
	BotServiceGroupID   int64
	RedisHost           string
	RedisPort           string
	RedisDB             int
	ChannelPrefix       string
	TronPrivateKey      string
	TronOwnerAddress    string
	TronOwnerAddressHex string
	TronToAddress       string
	TronToAddressHex    string
	TronApiKey          string
)

func Load() {
	if file, err := os.OpenFile("./logs/"+carbon.Now(carbon.Shanghai).ToDateString()+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666); err != nil {
		panic("open log file error: " + err.Error())
	} else {
		log.SetOutput(file)
	}
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	AppDebug = viper.GetBool("APP_DEBUG")
	MysqlDns = parseMysql()
	BotToken = viper.GetString("BOT_TOKEN")
	AppName = viper.GetString("APP_NAME")
	BotUsername = viper.GetString("BOT_USERNAME")
	fmt.Println("当前机器人：", BotUsername)
	RedisHost = viper.GetString("REDIS_HOST")
	RedisPort = viper.GetString("REDIS_PORT")
	RedisDB = viper.GetInt("REDIS_DB")
	if AppName == "" {
		panic("app name not set!")
	}
	HashKey = AppName + "_config"
	BotGameGroupID = viper.GetInt64("BOT_GAME_GROUP")
	BotServiceGroupID = viper.GetInt64("BOT_SERVICE_GROUP")
	ChannelPrefix = AppName + "_channel:"
	parseTronConfig()
}

func parseMysql() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_PASSWORD"),
		fmt.Sprintf(
			"%s:%s",
			viper.GetString("DB_HOST"),
			viper.GetString("DB_PORT")),
		viper.GetString("DB_DATABASE"))
}

func parseTronConfig() {
	TronPrivateKey = viper.GetString("TRON_PRIVATE_KEY")
	TronApiKey = viper.GetString("TRON_API_KEY")
	var err error
	TronOwnerAddress = viper.GetString("TRON_OWNER_ADDRESS")
	TronOwnerAddressHex, err = encode.Base58ToHexV41(TronOwnerAddress)
	fmt.Println("owner address: ", TronOwnerAddress)
	if err != nil {
		panic("生成TRON发送地址失败" + err.Error())
	}
	TronToAddress = viper.GetString("TRON_TO_ADDRESS")
	TronToAddressHex, err = encode.Base58ToHexV41(TronToAddress)
	if err != nil {
		panic("生成TRON接收地址失败" + err.Error())
	}
}
