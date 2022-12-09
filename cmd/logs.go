/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"spug-cli/api"
	"strconv"
)

func init() {
	// logsCmd represents the logs command
	var logsCmd = &cobra.Command{
		Use:   "logs [applyId]",
		Short: "Show apply logs",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				panic("申请单ID不能为空！")
			}
			applyId, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}
			api.Login(&s)
			log, err := s.RequestLog(applyId)
			if err != nil {
				panic(err)
			}
			fmt.Println(log.Data)
		},
	}
	rootCmd.AddCommand(logsCmd)
}
