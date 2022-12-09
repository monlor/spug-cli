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
			login := Login{
				Url:          url,
				Username:     username,
				Password:     password,
				LarkBotToken: larkBotToken,
			}
			writeUser(login)
			fmt.Println("登录成功！")
			return nil
		},
	}
	rootCmd.AddCommand(loginCmd)

	loginCmd.PersistentFlags().StringVarP(&url, "url", "", "https://spug.byteplan.com", "Spug url (required)")
	loginCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Spug username (required)")
	loginCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Spug password (required)")
	loginCmd.PersistentFlags().StringVarP(&larkBotToken, "larkBotToken", "l", "", "Lark bot id, push message")
	loginCmd.MarkPersistentFlagRequired("url")
	loginCmd.MarkPersistentFlagRequired("username")
	loginCmd.MarkPersistentFlagRequired("password")
}

func writeUser(login Login) {
	file, err := os.OpenFile(userPath, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	enc := gob.NewEncoder(file)
	err = enc.Encode(login)
	if err != nil {
		panic(err)
	}
}
