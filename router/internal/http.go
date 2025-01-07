package internal

import (
	"wvp-config/conf"
	"wvp-config/router/internal/model"
)

type (
	Request struct {
		Architecture conf.ArchitectureOptions `json:"mode"`
		// ExternalIP 外网IP
		ExternalIP string `json:"externalIP" binding:"required"`
		// ExternalPort 外网端口 指wvp的web页面的
		ExternalPort int `json:"externalPort" binding:"required"`
		WvpConfigOption
		secret string
	}

	WvpConfigOption struct {
		// 默认这些选项前端传过来就有默认值
		model.RedisConfig      `json:"redis"`
		model.DatasourceConfig `json:"datasource"`
		model.SipConfig        `json:"sip"`
		model.MediaConfig      `json:"media"`
	}
)

type (
	Response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    any    `json:"data"`
	}

	TemplateDetails struct {
		Architecture conf.ArchitectureOptions `json:"mode"`
		ImageApi     string                   `json:"imageApi"`
	}

	ConfigDetails struct {
		FileType conf.FileType `json:"fileType"`
		Content  string        `json:"content"`
		Url      string        `json:"url"`
	}
)
