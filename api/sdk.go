package api

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

// Login 登录
func Login(s *Spug) error {
	client := resty.New()
	var result LoginResp

	_, err := client.R().
		SetBody(fmt.Sprintf(`{"username":"%s","password":"%s","type":"ldap"}`, s.Username, s.Password)).
		SetResult(&result).
		Post(s.Url + SpugLoginApi)

	if result.Error != "" {
		return errors.New(result.Error)
	}

	s.Token = result.Data.AccessToken

	return err
}

// App 查询app列表
func (s Spug) App() ([]AppData, error) {
	client := resty.New()
	var result AppResp

	_, err := client.R().
		SetHeader("X-Token", s.Token).
		SetResult(&result).
		Get(s.Url + SpugAppApi)

	return result.Data, err
}

// Deploy 查询部署配置列表
func (s Spug) Deploy(appId int) ([]DeployData, error) {
	client := resty.New()
	var result DeployResp
	var data []DeployData

	_, err := client.R().
		SetHeader("X-Token", s.Token).
		SetQueryParam("app_id", fmt.Sprint(appId)).
		SetResult(&result).
		Get(s.Url + SpugDeployApi)
	if err != nil {
		return nil, err
	}
	data = result.Data

	// 获取环境信息
	envs, err := s.Env()
	if err != nil {
		return nil, err
	}

	// 匹配环境名称
	for k, v := range data {
		for _, e := range envs {
			if v.EnvId == e.Id {
				data[k].EnvName = e.Name
			}
		}
	}

	return data, err
}

// Request 提交发布申请
func (s Spug) Request(title string, release string, hostIds []int, deployId int) (r ApplyData, err error) {
	client := resty.New()
	var result CommonVo

	_, err = client.R().
		SetHeader("X-Token", s.Token).
		SetBody(RequestReq{
			Name:     title,
			Version:  release,
			HostIds:  hostIds,
			DeployId: deployId,
		}).
		SetResult(&result).
		Post(s.Url + SpugRequestApi)

	if err != nil {
		return r, err
	}

	if result.Error != "" {
		return r, errors.New(result.Error)
	}

	// 查询申请单
	applies, err := s.Apply()
	if err != nil {
		return r, err
	}

	for _, apply := range applies[:10] {
		if apply.DeployId == deployId && apply.Version == release && apply.Name == title && apply.StatusAlias == "待发布" {
			return apply, err
		}
	}

	return r, err

}

// Env 查询环境列表
func (s Spug) Env() ([]EnvData, error) {
	client := resty.New()
	var result EnvResp

	_, err := client.R().
		SetHeader("X-Token", s.Token).
		SetResult(&result).
		Get(s.Url + SpugEnvApi)

	return result.Data, err
}

// Publish 发布申请单
func (s Spug) Publish(applyId int) error {
	client := resty.New()
	var result CommonVo

	_, err := client.R().
		SetHeader("X-Token", s.Token).
		SetBody(`{"mode":"all"}`).
		SetPathParam("applyId", fmt.Sprint(applyId)).
		SetResult(&result).
		Post(s.Url + SpugPublishApi)

	if err != nil {
		return err
	}

	if result.Error != "" {
		return errors.New(result.Error)
	}

	return err
}

// Apply 查询所有申请单信息
func (s Spug) Apply() ([]ApplyData, error) {
	client := resty.New()
	var result ApplyResp

	_, err := client.R().
		SetHeader("X-Token", s.Token).
		SetResult(&result).
		Get(s.Url + SpugApplyApi)

	return result.Data, err
}

// RequestInfo 查询申请单信息，状态
func (s Spug) RequestInfo(applyId int) (RequestInfoData, error) {
	client := resty.New()
	var result RequestInfoResp

	_, err := client.R().
		SetHeader("X-Token", s.Token).
		SetQueryParam("id", fmt.Sprint(applyId)).
		SetResult(&result).
		Get(s.Url + SpugRequestInfoApi)

	return result.Data, err
}

// RequestLog 获取申请单日志
func (s Spug) RequestLog(applyId int) (RequestLogInfo, error) {
	client := resty.New()
	var result RequestLogResp

	_, err := client.R().
		SetHeader("X-Token", s.Token).
		SetPathParam("applyId", fmt.Sprint(applyId)).
		SetResult(&result).
		Get(s.Url + SpugRequestLogApi)

	return result.Data.Output.Log, err
}

// Approve 审批申请单
func (s Spug) Approve(applyId int) error {
	client := resty.New()
	var result CommonVo

	_, err := client.R().
		SetHeader("X-Token", s.Token).
		SetBody(`{"is_pass":true}`).
		SetPathParam("applyId", fmt.Sprint(applyId)).
		SetResult(&result).
		Patch(s.Url + SpugApproveApi)

	if err != nil {
		return err
	}

	if result.Error != "" {
		return errors.New(result.Error)
	}

	return err
}

// ApproveAll 审批所有申请单
func (s Spug) ApproveAll() error {
	applies, err := s.Apply()
	if err != nil {
		return err
	}
	// 现在查询申请单列表是全量数据，所以只取前200条过滤需要审批的单子
	for _, apply := range applies[:200] {
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
}
