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

type Login struct {
	Url          string
	Username     string
	Password     string
	LarkBotToken string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "spug-cli",
	Short: "Spug command line tool, support application release, log query, audit and so on",
	Example: `  spug-cli login -u [your user name] -p [your password]
  spug-cli publish -e dev -a job
  spug-cli publish -e dev -a base -v dev-latest -w
  spug-cli logs 6634
  spug-cli status 6634`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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

func GetEnvId(environment string) (envId int, err error) {
	envId = -1
	if environment == "" {
		return envId, nil
	}
	envs, err := s.Env()
	if err != nil {
		return envId, err
	}
	for _, env := range envs {
		if env.Key == environment {
			envId = env.Id
			break
		}
	}
	if envId == -1 {
		return envId, errors.New("找不到该环境：" + environment)
	}
	return envId, err
}

func GetAppId(appKey string) (appId int, err error) {
	appId = -1
	apps, err := s.App()
	if err != nil {
		return appId, err
	}
	for _, v := range apps {
		if v.Key == appKey {
			appId = v.Id
			break
		}
	}
	if appId == -1 {
		return appId, errors.New("找不到该应用：" + appKey)
	}
	return appId, err
}

func SelectEnv(result *string) error {
	if *result != "" {
		return nil
	}
	envs, err := s.Env()
	if err != nil {
		return err
	}
	var options []string
	for _, env := range envs {
		options = append(options, env.Name)
	}
	var index int
	err = SelectOptions("请选择环境:", &index, options)
	*result = envs[index].Key
	return err
}

func SelectApp(result *string) error {
	if *result != "" {
		return nil
	}
	apps, err := s.App()
	if err != nil {
		return err
	}
	var options []string
	for _, app := range apps {
		options = append(options, app.Name)
	}
	var index int
	err = SelectOptions("请选择应用:", &index, options)
	*result = apps[index].Key
	return err
}

func SelectApply(result *int) error {
	applies, err := s.Apply()
	if err != nil {
		return err
	}
	var options []string
	for _, apply := range applies[:20] {
		if apply.Version == "" {
			apply.Version = "空"
		}
		options = append(options, fmt.Sprintf("%d-%s-%s-%s-%s-%s",
			apply.Id, apply.AppName, apply.EnvName, apply.Version, apply.CreatedByUser, apply.StatusAlias))
	}
	var index int
	err = SelectOptions("请选择申请单:", &index, options)
	*result = applies[index].Id
	return err
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
