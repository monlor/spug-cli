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
	var applyId int
	// logsCmd represents the logs command
	var logsCmd = &cobra.Command{
		Use:   "logs [applyId]",
		Short: "Show apply logs",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := api.Login(&s); err != nil {
				return err
			}
			if len(args) == 0 {
				if err := SelectApply(&applyId); err != nil {
					return err
				}
			} else {
				if r, err := strconv.Atoi(args[0]); err != nil {
					return err
				} else {
					applyId = r
				}
			}
			log, err := s.RequestLog(applyId)
			if err != nil {
				return err
			}
			fmt.Println(log.Data)
			return nil
		},
	}
	rootCmd.AddCommand(logsCmd)
}
