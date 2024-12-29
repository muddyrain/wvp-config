package internal

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/google/uuid"
	"gopkg.in/ini.v1"
	_ "gopkg.in/ini.v1"
	"gopkg.in/yaml.v2"
	"os"
	"strconv"
	"strings"
	"wvp-config/router/internal/model"
)

func StandaloneConversion(req Request) []ConfigDetails {
	wvpConfig := parseWvpConfig(req, "./template/standalone/application.yml")
	zlmConfig := parseZlmConfig(wvpConfig, "./template/standalone/config.ini")
	wvpData, _ := yaml.Marshal(wvpConfig)
	cfg := ini.Empty()
	_ = cfg.ReflectFrom(&zlmConfig)
	name := fmt.Sprintf("%s_zlm.ini", zlmConfig.API.Secret)
	_ = cfg.SaveTo(name)
	zlmData, _ := os.ReadFile(name)
	defer func() {
		//_ = os.Remove(name)
		os.WriteFile("wvp.yaml", wvpData, os.ModePerm)
	}()
	hint := createHit(req, wvpConfig)
	b, _ := json.MarshalIndent(hint, "", "  ")
	return []ConfigDetails{
		{
			FileType: WVP,
			Content:  string(wvpData),
			Data:     wvpConfig,
		},
		{
			FileType: ZLM,
			Content:  string(zlmData),
			Data:     zlmConfig,
		},
		{
			FileType: TXT,
			Content:  string(b),
			Data:     hint,
		},
	}
}

func parseZlmConfig(wvpConfig model.WvpConfig, path string) model.ZLMConfig {
	var zlmConfig model.ZLMConfig
	cfg, _ := ini.Load(path)
	_ = cfg.MapTo(&zlmConfig)
	zlmConfig.API.Secret = wvpConfig.Media.Secret
	const port = "12336"
	targetPort := strconv.Itoa(wvpConfig.Server.Port)
	zlmConfig.Hook.OnPlay = strings.ReplaceAll(zlmConfig.Hook.OnPlay, port, targetPort)
	zlmConfig.Hook.OnPublish = strings.ReplaceAll(zlmConfig.Hook.OnPublish, port, targetPort)
	zlmConfig.Hook.OnRecordMP4 = strings.ReplaceAll(zlmConfig.Hook.OnRecordMP4, port, targetPort)
	zlmConfig.Hook.OnRtpServerTimeout = strings.ReplaceAll(zlmConfig.Hook.OnRtpServerTimeout, port, targetPort)
	zlmConfig.Hook.OnSendRtpStopped = strings.ReplaceAll(zlmConfig.Hook.OnSendRtpStopped, port, targetPort)
	zlmConfig.Hook.OnServerKeepalive = strings.ReplaceAll(zlmConfig.Hook.OnServerKeepalive, port, targetPort)
	zlmConfig.Hook.OnServerStarted = strings.ReplaceAll(zlmConfig.Hook.OnServerStarted, port, targetPort)
	zlmConfig.Hook.OnStreamChanged = strings.ReplaceAll(zlmConfig.Hook.OnStreamChanged, port, targetPort)
	zlmConfig.Hook.OnStreamNoneReader = strings.ReplaceAll(zlmConfig.Hook.OnStreamNoneReader, port, targetPort)
	zlmConfig.Hook.OnStreamNotFound = strings.ReplaceAll(zlmConfig.Hook.OnStreamNotFound, port, targetPort)
	zlmConfig.Hook.StreamChangedSchemas = strings.ReplaceAll(zlmConfig.Hook.StreamChangedSchemas, port, targetPort)
	zlmConfig.HTTP.Port = wvpConfig.Media.HttpPort
	zlmConfig.RTC.ExternIP = wvpConfig.Media.StreamIP
	zlmConfig.RTPProxy.PortRange = strings.ReplaceAll(wvpConfig.Media.Rtp.PortRange, ",", "-")
	return zlmConfig
}

func parseWvpConfig(req Request, path string) model.WvpConfig {
	data, _ := os.ReadFile(path)
	var wvpConfig model.WvpConfig
	_ = yaml.Unmarshal(data, &wvpConfig)
	wvpConfig.Server.Port = req.ExternalPort
	wvpConfig.Media.SdpIP = req.ExternalIP
	wvpConfig.Media.StreamIP = req.ExternalIP
	wvpConfig.Media.Secret = uuid.New().String()
	wvpConfig.UserSettings.AllowedOrigins = append(wvpConfig.UserSettings.AllowedOrigins,
		fmt.Sprintf("http://%s:%d", req.ExternalIP, req.ExternalPort))
	if req.WvpConfigOption.RedisConfig.Enable {
		wvpConfig.Spring.Redis = req.WvpConfigOption.RedisConfig
	}
	if req.WvpConfigOption.DatasourceConfig.Enable {
		wvpConfig.Spring.Datasource = req.WvpConfigOption.DatasourceConfig
	}
	if req.WvpConfigOption.SipConfig.Enable {
		wvpConfig.Sip = req.WvpConfigOption.SipConfig
	}
	if req.WvpConfigOption.MediaConfig.Enable {
		wvpConfig.Media = req.WvpConfigOption.MediaConfig
	}

	return wvpConfig
}

func createHit(req Request, wvpConfig model.WvpConfig) model.Hint {
	return model.Hint{
		Contents: []string{
			fmt.Sprintf("web页面地址：http://%s:%d", req.ExternalIP, req.ExternalPort),
			fmt.Sprintf("信令端口: %d", wvpConfig.Sip.Port),
			fmt.Sprintf("流媒体请求端口: %d", wvpConfig.Media.HttpPort),
			fmt.Sprintf("流媒体收流端口: %s", wvpConfig.Media.Rtp.PortRange),
		},
	}
}
