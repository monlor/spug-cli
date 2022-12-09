/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
	"spug-cli/api"

	"github.com/spf13/cobra"
)

var (
	version  = "v0.0.1"
	s        api.Spug
	login    Login
	userPath = fmt.Sprintf("%s/.spug", os.Getenv("HOME"))
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "spug-cli",
	Short: "Spug命令行工具",
	Long:  `Spug命令行工具，支持提申请单发布，查询日志，审核等操作`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

type Login struct {
	Url          string
	Username     string
	Password     string
	LarkBotToken string
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Spug cli %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	readUser()
}

func readUser() {
	if !PathExists(userPath) {
		return
	}
	file, err := os.OpenFile(userPath, os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}
	dec := gob.NewDecoder(file)
	err = dec.Decode(&login)
	if err != nil {
		panic(err)
	}
	s = api.Spug{
		Url:      login.Url,
		Username: login.Username,
		Password: login.Password,
	}
}

func GetEnvId(environment string) (envId int) {
	envs, err := s.Env()
	if err != nil {
		panic(err)
	}
	for _, env := range envs {
		if env.Key == environment {
			envId = env.Id
			break
		}
	}
	if envId == -1 {
		panic("找不到该环境！")
	}
	return
}

func GetAppId(appKey string) (appId int) {
	appId = -1
	apps, err := s.App()
	if err != nil {
		panic(err)
	}
	for _, v := range apps {
		if v.Key == appKey {
			appId = v.Id
			break
		}
	}
	if appId == -1 {
		panic("找不到该应用！")
	}
	return
}

func SendMessage(text string) error {
	if login.LarkBotToken == "" {
		return nil
	}
	client := resty.New()
	var result FeishuBotResp

	_, err := client.R().
		SetBody(fmt.Sprintf(`{"msg_type":"text","content":{"text":"%s"}}`, text)).
		SetHeader("Content-Type", "application/json").
		SetResult(&result).
		Post("https://open.feishu.cn/open-apis/bot/v2/hook/" + login.LarkBotToken)
	if result.StatusCode != 0 {
		return errors.New(result.StatusMessage)
	}
	return err
}
