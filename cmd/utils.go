package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"os"
)

/*
判断文件或文件夹是否存在
如果返回的错误为nil,说明文件或文件夹存在
如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
如果返回的错误为其它类型,则不确定是否在存在
*/
func PathExists(path string) bool {

	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func InputString(message string, result *string, defaultValue string) error {
	prompt := &survey.Input{
		Message: message,
		Default: defaultValue,
	}
	err := survey.AskOne(prompt, result)
	return err
}

func InputStringRequired(message string, result *string, defaultValue string) error {
	prompt := &survey.Input{
		Message: message,
		Default: defaultValue,
	}
	err := survey.AskOne(prompt, result, survey.WithValidator(survey.Required))
	return err
}

func InputPassword(message string, result *string) error {
	prompt := &survey.Password{
		Message: message,
	}
	err := survey.AskOne(prompt, result, survey.WithValidator(survey.Required))
	return err
}

func SelectOptions(message string, result *int, options []string) error {
	prompt := &survey.Select{
		Message: message,
		Options: options,
	}
	err := survey.AskOne(prompt, result, survey.WithValidator(survey.Required))
	return err
}

func Confirm(message string, result *bool) error {
	prompt := &survey.Confirm{
		Message: message,
	}
	err := survey.AskOne(prompt, result)
	return err
}
