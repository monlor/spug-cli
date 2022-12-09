/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"spug-cli/api"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	// statusCmd represents the status command
	var statusCmd = &cobra.Command{
		Use:   "status [applyId]",
		Short: "Show apply status",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				panic("申请单ID不能为空！")
			}
			applyId, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}
			api.Login(&s)
			info, err := s.RequestInfo(applyId)
			if err != nil {
				panic(err)
			}
			fmt.Println(info.StatusAlias)
			if info.Reason != "" {
				fmt.Println(info.Reason)
			}
		},
	}
	rootCmd.AddCommand(statusCmd)
}
