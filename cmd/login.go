/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/gob"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	var (
		url          string
		username     string
		password     string
		larkBotToken string
	)

	// loginCmd represents the login command
	var loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Login Spug",
		RunE: func(cmd *cobra.Command, args []string) error {
			// 输入参数
			if url == "" || username == "" || password == "" {
				if err := InputStringRequired("请输入 Spug 地址:", &url, login.Url); err != nil {
					return err
				}
				if err := InputStringRequired("请输入 Spug 用户名:", &username, login.Username); err != nil {
					return err
				}
				if err := InputPassword("请输入 Spug 密码:", &password); err != nil {
					return err
				}
				if err := InputString("请输入飞书机器人密钥[非必输]:", &larkBotToken, login.LarkBotToken); err != nil {
					return err
				}
			}
			// 开始登录
			login := Login{
				Url:          url,
				Username:     username,
				Password:     password,
				LarkBotToken: larkBotToken,
			}
			if err := writeUser(login); err != nil {
				return err
			}
			fmt.Println("登录成功！")
			return nil
		},
	}
	rootCmd.AddCommand(loginCmd)

	loginCmd.PersistentFlags().StringVarP(&url, "url", "", "", "Spug url (required)")
	loginCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Spug username (required)")
	loginCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Spug password (required)")
	loginCmd.PersistentFlags().StringVarP(&larkBotToken, "larkBotToken", "l", "", "Lark bot id, push message")
	//loginCmd.MarkPersistentFlagRequired("url")
	//loginCmd.MarkPersistentFlagRequired("username")
	//loginCmd.MarkPersistentFlagRequired("password")
}

func writeUser(login Login) error {
	file, err := os.OpenFile(userPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	enc := gob.NewEncoder(file)
	err = enc.Encode(login)
	if err != nil {
		return err
	}
	return nil
}
