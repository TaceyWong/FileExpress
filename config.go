package main

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

var defaultConfig = `{
	"FileStorage": "local",
	"Name": "文件快递 - FileExpress",
	"Description": "开箱即用的文件快传系统",
	"NotifyTitle": "系统通知",
	"NotifyContent": "欢迎使用 FileExpress，本程序开源于 <a href=\"https://github.com/TaceyWong/FileExpress\" target=\"_blank\">Github</a> ，欢迎Star和Fork。",
	"PageExplain": "请勿上传或分享违法内容。根据《中华人民共和国网络安全法》、《中华人民共和国刑法》、《中华人民共和国治安管理处罚法》等相关规定。 传播或存储违法、违规内容，会受到相关处罚，严重者将承担刑事责任。本站坚决配合相关部门，确保网络内容的安全，和谐，打造绿色网络环境。",
	"Keywords": "FileExpress, 文件快递, 文件快传, 文件分享, 文件",
	"S3AccessKeyId": "",
	"S3SecretAccessKey": "",
	"S3BucketName": "",
	"S3EndpointUrl": "",
	"S3RegionName": "auto",
	"S3SignatureVersion": "s3v2",
	"S3Hostname": "",
	"S3Proxy": 0,
	"MaxSaveSeconds": 0,
	"AWSSessionToken": "",
	"OnedriveDomain": "",
	"OnedriveClientId": "",
	"OnedriveUsername": "",
	"OnedrivePassword": "",
	"OnedriveRootPath": "filebox_storage",
	"OnedriveProxy": 0,
	"AdminToken": "FileCodeBox2023",
	"OpenUpload": 1,
	"UploadSize": 1024 * 1024 * 10,
	"UploadMinute": 1,
	"Opacity": 0.9,
	"Background": "",
	"UploadCount": 10,
	"ErrorMinute": 1,
	"ErrorCount": 1,
	"Port": 12345,
	"ShowAdminAddr": 0,
	"RobotsText": "User-agent: *\nDisallow: /",
	"ExpireStyle": ["day", "hour", "minute", "forever", "count"],
}`

func LoadConfig() {
	// 设置默认配置
	var config map[string]interface{}
	json.Unmarshal([]byte(defaultConfig), &config)
	for key, value := range config {
		viper.SetDefault(key, value)
	}
	// 设置配置文件路径
	viper.SetConfigName("feconfig")       // name of config file (without extension)
	viper.SetConfigType("json")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")

	// 读取配置文件
	viper.ReadInConfig()
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found; ignore error if desired")
		} else {
			fmt.Println("Config file was found but another error was produced")
		}
	}
}
