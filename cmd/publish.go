package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"spug-cli/api"
	"time"
)

func init() {

	var (
		appKey      string
		version     string
		environment string
		wait        bool
	)

	var publishCmd = &cobra.Command{
		Use:   "publish",
		Short: "Publish your application",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := api.Login(&s)
			if err != nil {
				panic(err)
			}
			// 获取环境id
			envId := GetEnvId(environment)
			// 先获取appId
			appId := GetAppId(appKey)
			// 查询发布配置
			deploys, err := s.Deploy(appId)
			deploy := api.DeployData{}
			for _, v := range deploys {
				if v.EnvId == envId {
					deploy = v
					break
				}
			}
			if deploy.EnvName == "" {
				panic("找不到发布配置！")
			}
			// 提交发布申请
			apply, err := s.Request("Spug Cli 工具自动提交", version, deploy.HostIds, deploy.Id)
			if err != nil {
				panic(err)
			}
			if !deploy.IsAudit {
				fmt.Println("申请单不用审批，直接执行发布...")
				err := s.Publish(apply.Id)
				if err != nil {
					panic(err)
				}
			}
			fmt.Printf("发布申请提交成功！\n查看日志：spug-cli logs %d\n查看状态：spug-cli status %d\n", apply.Id, apply.Id)
			// 开启了等待，并且不用审批的申请，执行等待
			if wait && !deploy.IsAudit {
				waitFinish(apply)
			}
			return nil
		},
	}
	rootCmd.AddCommand(publishCmd)
	publishCmd.PersistentFlags().StringVarP(&appKey, "appKey", "a", "", "Application name, eg: base,data-web... (required)")
	publishCmd.PersistentFlags().StringVarP(&version, "version", "v", "", "Application branch/tag to publish, eg: dev-latest,v1.0.0...")
	publishCmd.PersistentFlags().StringVarP(&environment, "environment", "e", "", "Publish Environment Key, eg: dev,test,uat,saas... (required)")
	publishCmd.PersistentFlags().BoolVarP(&wait, "wait", "w", false, "Wait for the release to complete")
	publishCmd.MarkPersistentFlagRequired("appKey")
	publishCmd.MarkPersistentFlagRequired("environment")
}

func waitFinish(data api.ApplyData) {
	count := 0
	maxCount := 100
	fmt.Printf("等待发布完成...")
	for true {
		info, err := s.RequestInfo(data.Id)
		if err != nil {
			panic(err)
		}
		if info.Status == "3" {
			fmt.Println("发布成功！")
			SendMessage(fmt.Sprintf("模块 %s 发布 %s 成功！", data.AppName, data.Version))
			break
		}
		if info.Status != "2" {
			fmt.Printf("发布失败！%s:%s\n", info.StatusAlias, info.Reason)
			SendMessage(fmt.Sprintf("模块 %s 发布 %s 失败！", data.AppName, data.Version))
			break
		}
		if count >= maxCount {
			fmt.Println("检查超时！请手动检查状态！")
			SendMessage(fmt.Sprintf("模块 %s 发布 %s 超时，请手动检查状态！", data.AppName, data.Version))
			break
		}
		count++
		fmt.Printf(".")
		time.Sleep(time.Second * 5)
	}
}
