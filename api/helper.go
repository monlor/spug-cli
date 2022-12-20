package api

import "errors"

func (s Spug) GetEnvId(environment string) (envId int, err error) {
	envId = -1
	if environment == "" {
		return envId, nil
	}
	envs, err := s.Env()
	if err != nil {
		return envId, err
	}
	for _, env := range envs {
		if env.Key == environment {
			envId = env.Id
			break
		}
	}
	if envId == -1 {
		return envId, errors.New("找不到该环境：" + environment)
	}
	return envId, err
}

func (s Spug) GetAppId(appKey string) (appId int, err error) {
	appId = -1
	apps, err := s.App()
	if err != nil {
		return appId, err
	}
	for _, v := range apps {
		if v.Key == appKey {
			appId = v.Id
			break
		}
	}
	if appId == -1 {
		return appId, errors.New("找不到该应用：" + appKey)
	}
	return appId, err
}
