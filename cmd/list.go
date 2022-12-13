package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {

	var (
		count       int
		environment string
		status      string
		appKey      string
	)

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List spug apply",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := LoginSpug(); err != nil {
				return err
			}
			applies, err := s.Apply()
			if err != nil {
				return err
			}
			envId, err := GetEnvId(environment)
			if err != nil {
				return err
			}
			fmt.Println("行号: 申请单ID 模块 环境 版本 状态")
			for i, apply := range applies[:count] {
				if envId != -1 && envId != apply.EnvId {
					continue
				}
				if status != "" && apply.StatusAlias != status {
					continue
				}
				if appKey != "" && apply.AppName != appKey {
					continue
				}
				if apply.Version == "" {
					apply.Version = "空"
				}
				fmt.Printf("%d: %d %s %s %s %s\n",
					i, apply.Id, apply.AppName, apply.EnvName, apply.Version, apply.StatusAlias)
			}
			return nil
		},
	}

	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().StringVarP(&environment, "environment", "e", "", "Environment Key, eg: dev,test,uat,saas...")
	listCmd.PersistentFlags().StringVarP(&appKey, "appKey", "a", "", "Application name, eg: base,data-web... (required)")
	listCmd.PersistentFlags().StringVarP(&status, "status", "s", "", "Status alias")
	listCmd.PersistentFlags().IntVarP(&count, "count", "c", 50, "Maximum number of rows to filter")
}
