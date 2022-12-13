package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"spug-cli/api"
)

func init() {
	var (
		environment string
		applyId     int
		appKey      string
	)

	var approveCmd = &cobra.Command{
		Use:   "approve",
		Short: "Approve spug apply",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := api.Login(&s); err != nil {
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
			// 现在查询申请单列表是全量数据，为了提升效率，只取前200条过滤需要审批的单子
			for _, apply := range applies[:200] {
				if envId != -1 && envId != apply.EnvId {
					continue
				}
				if applyId != -1 && applyId != apply.Id {
					continue
				}
				if appKey != "" && apply.AppName != appKey {
					continue
				}
				if apply.Status == "0" {
					fmt.Printf("审核申请单：%s，模块：%s, 发布环境：%s, 版本：%s，申请人：%s，发布时间：%s\n",
						apply.Name, apply.AppName, apply.EnvName, apply.Version, apply.CreatedByUser, apply.Plan)
					err := s.Approve(apply.Id)
					if err != nil {
						fmt.Printf("> 审核失败！%s\n", err.Error())
					}
				}
			}
			return nil
		},
	}

	rootCmd.AddCommand(approveCmd)

	approveCmd.PersistentFlags().StringVarP(&environment, "environment", "e", "", "Environment Key, eg: dev,test,uat,saas...")
	approveCmd.PersistentFlags().StringVarP(&appKey, "appKey", "a", "", "Application name, eg: base,data-web... (required)")
	approveCmd.PersistentFlags().IntVarP(&applyId, "applyId", "", -1, "Apply id")
}
