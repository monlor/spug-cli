/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"spug-cli/api"
)

func init() {
	var (
		environment string
	)
	// approveCmd represents the approve command
	var approveCmd = &cobra.Command{
		Use:   "approve",
		Short: "Approve apply",
		Run: func(cmd *cobra.Command, args []string) {
			api.Login(&s)
		},
	}

	rootCmd.AddCommand(approveCmd)
	approveCmd.PersistentFlags().StringVarP(&environment, "environment", "e", "", "Publish Environment Key, eg: dev,test,uat,saas...")
	approveCmd.MarkPersistentFlagRequired("environment")
}
