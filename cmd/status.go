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
	var applyId int
	// statusCmd represents the status command
	var statusCmd = &cobra.Command{
		Use:   "status [applyId]",
		Short: "Show apply status",
		RunE: func(cmd *cobra.Command, args []string) error {
			api.Login(&s)
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
			info, err := s.RequestInfo(applyId)
			if err != nil {
				return err
			}
			fmt.Println(info.StatusAlias)
			if info.Reason != "" {
				fmt.Println(info.Reason)
			}
			return nil
		},
	}
	rootCmd.AddCommand(statusCmd)
}
