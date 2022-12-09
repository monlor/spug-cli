package api

import (
	"testing"
	"time"
)

var s Spug
var appId int
var deploy DeployData
var applyId int

func TestLogin(t *testing.T) {
	s := Spug{
		Url:      "",
		Username: "",
		Password: "",
	}
	err := Login(&s)
	if err == nil {
		t.Log(s.Token)
	} else {
		t.Error(err)
	}

}

func TestApp(t *testing.T) {
	apps, err := s.App()
	if err == nil {
		appId = apps[5].Id
		t.Log(apps)
	} else {
		t.Error(err)
	}
}

func TestDeploy(t *testing.T) {
	t.Logf("查询id为%d的应用发布配置...", appId)
	deploys, err := s.Deploy(appId)
	if err == nil {
		deploy = deploys[0]
		t.Log(deploys)
	} else {
		t.Error(err)
	}

}

func TestRequest(t *testing.T) {
	t.Logf("发布配置id：%v", deploy)
	apply, r := s.Request("测试", "dev-latest", deploy.HostIds, deploy.Id)
	if r == nil {
		applyId = apply.Id
		t.Logf("提交发布申请成功！%v", apply)
		if !deploy.IsAudit {
			t.Log("申请单不用审批，直接执行发布...")
			err := s.Publish(apply.Id)
			if err == nil {
				t.Log("执行发布成功！")
			}
		}
	} else {
		t.Error(r)
	}
}

func TestRequestInfo(t *testing.T) {
	r, err := s.RequestInfo(applyId)
	if err == nil {
		t.Log(r.StatusAlias)
	} else {
		t.Error(err)
	}
}

func TestRequestLog(t *testing.T) {
	time.Sleep(time.Second * 10)
	r, err := s.RequestLog(applyId)
	if err == nil {
		t.Log(r.Data)
	} else {
		t.Error(err)
	}
}

func TestApproveAll(t *testing.T) {
	err := s.ApproveAll()
	if err != nil {
		t.Error(err)
	}
}
