package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
	"wvp-config/conf"
	"wvp-config/router/internal/model"
)

func StandaloneIntercomConversion(req Request) []ConfigDetails {
	config := conf.GetData().GetConfig(conf.StandaloneIntercom)
	id, nginxName := parseNginxConfig(req, config.NginxPath)
	req.ExternalPort++
	req.secret = id
	wvpConfig := parseWvpConfig(req, config.WvpPath)
	zlmConfig := parseZlmConfig(wvpConfig, config.ZlmPath)
	wvpName := fmt.Sprintf("%s_wvp.yaml", zlmConfig.API.Secret)
	wvpData, _ := yaml.Marshal(wvpConfig)
	_ = os.WriteFile(config.Dir+wvpName, wvpData, os.ModePerm)
	cfg := ini.Empty()
	_ = cfg.ReflectFrom(&zlmConfig)
	zlmName := fmt.Sprintf("%s_zlm.ini", zlmConfig.API.Secret)
	_ = cfg.SaveTo(config.Dir + zlmName)
	zlmData, _ := os.ReadFile(config.Dir + zlmName)
	nginxData, _ := os.ReadFile(config.Dir + nginxName)
	hint := createHit2(req, wvpConfig)
	b, _ := json.MarshalIndent(hint, "", "  ")
	return []ConfigDetails{
		{
			FileType: conf.WVP,
			Content:  string(wvpData),
			Url:      config.HTTPPrefix + wvpName,
		},
		{
			FileType: conf.ZLM,
			Content:  string(zlmData),
			Url:      config.HTTPPrefix + zlmName,
		},
		{
			FileType: conf.Nginx,
			Content:  string(nginxData),
			Url:      config.HTTPPrefix + nginxName,
		},
		{
			FileType: conf.TXT,
			Content:  string(b),
		},
	}
}

func parseNginxConfig(req Request, path string) (string, string) {
	// 打开配置文件，这里假设配置文件名为config.txt，你可以根据实际情况修改文件名
	file, _ := os.Open(path)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	// 创建一个切片用于存储替换后的行内容
	var (
		newLines   []string
		port       = req.ExternalPort
		recordPort = 13039
	)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		newLine := strings.Replace(line, fmt.Sprintf("%d", recordPort),
			fmt.Sprintf("%d", port), -1)
		newLines = append(newLines, newLine)
		port++
		recordPort++
	}

	id := uuid.New().String()
	name := fmt.Sprintf("%s_nginx.conf", id)
	config := conf.GetData().GetConfig(conf.StandaloneIntercom)
	newFile, _ := os.Create(config.Dir + name)
	for _, line := range newLines {
		_, _ = newFile.WriteString(line + "\n")
	}
	_ = newFile.Close()
	return id, name
}

func createHit2(req Request, wvpConfig model.WvpConfig) model.Hint {
	return model.Hint{
		Contents: []string{
			fmt.Sprintf("web页面地址：http://%s:%d", req.ExternalIP, req.ExternalPort-1),
			fmt.Sprintf("信令端口: %d", wvpConfig.Sip.Port),
			fmt.Sprintf("流媒体请求端口: %d", req.ExternalPort+1),
			fmt.Sprintf("流媒体收流端口: %s", wvpConfig.Media.Rtp.PortRange),
			fmt.Sprintf("nginx中的配置需要修改成自己的证书"),
		},
	}
}
