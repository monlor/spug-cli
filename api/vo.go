package api

type Spug struct {
	Token    string
	Url      string
	Username string
	Password string
}

type CommonVo struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

// LoginResp 登录响应体
type LoginResp struct {
	Data  LoginData `json:"data"`
	Error string    `json:"error"`
}

type LoginData struct {
	Id          int    `json:"id"`
	AccessToken string `json:"access_token"`
	NickName    string `json:"nickname"`
	IsSupper    bool   `json:"is_supper"`
}

// AppResp 应用响应体
type AppResp struct {
	Data  []AppData `json:"data"`
	Error string    `json:"error"`
}

type AppData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
	Desc string `json:"desc"`
}

// DeployResp 发布配置响应体
type DeployResp struct {
	Data  []DeployData `json:"data"`
	Error string       `json:"error"`
}

type DeployData struct {
	Id      int    `json:"id"`
	AppId   int    `json:"app_id"`
	EnvId   int    `json:"env_id"`
	EnvName string `json:omitempty`
	HostIds []int  `json:"host_ids"`
	IsAudit bool   `json:"is_audit"`
	AppKey  string `json:"app_key"`
	AppName string `json:"app_name"`
}

// RequestReq 发布申请请求体
type RequestReq struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	HostIds  []int  `json:"host_ids"`
	DeployId int    `json:"deploy_id"`
}

// EnvResp 环境信息响应体
type EnvResp struct {
	Data  []EnvData `json:"data"`
	Error string    `json:"error"`
}

type EnvData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

// ApplyResp 申请单响应体
type ApplyResp struct {
	Data  []ApplyData `json:"data"`
	Error string      `json:"error"`
}

type ApplyData struct {
	Id            int    `json:"id"`
	DeployId      int    `json:"deploy_id"`
	Name          string `json:"name"'`
	Version       string `json:"version"`
	Type          string `json:"type"`
	envId         int    `json:"env_id"`
	EnvName       string `json:"env_name"`
	Plan          string `json:"plan"`
	AppId         int    `json:"app_id"`
	AppName       string `json:"app_name"`
	Status        string `json:"status"`
	StatusAlias   string `json:"status_alias"`
	CreatedByUser string `json:"created_by_user"`
}

// RequestInfoResp 发布请求信息响应体
type RequestInfoResp struct {
	Data  RequestInfoData `json:"data"`
	Error string          `json:"error"`
}

type RequestInfoData struct {
	Status      string `json:"status"`
	Reason      string `json:"reason"`
	StatusAlias string `json:"status_alias"`
}

// RequestLogResp 查询申请单运行日志响应体
type RequestLogResp struct {
	Data  RequestLogData `json:"data"`
	Error string         `json:"error"`
}

type RequestLogInfo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Data  string `json:"data"`
	Step  int    `json:"step"`
}

type RequestLogData struct {
	Output struct {
		Log RequestLogInfo `json:"1"`
	} `json:"outputs"`
	Status string `json:"status"`
	Index  int    `json:"index"`
}
