package model

type WvpConfig struct {
	Spring       SpringConfig       `yaml:"spring"`
	PageHelper   PageHelperConfig   `yaml:"pagehelper"`
	Server       ServerConfig       `yaml:"server"`
	Sip          SipConfig          `yaml:"sip"`
	Media        MediaConfig        `yaml:"media"`
	UserSettings UserSettingsConfig `yaml:"user-settings"`
	Logging      LoggingConfig      `yaml:"logging"`
	SpringDoc    SpringDocConfig    `yaml:"springdoc"`
}

type SpringConfig struct {
	MVC        MVCConfig        `yaml:"mvc"`
	Thymeleaf  ThymeleafConfig  `yaml:"thymeleaf"`
	Servlet    ServletConfig    `yaml:"servlet"`
	Cache      CacheConfig      `yaml:"cache"`
	Redis      RedisConfig      `yaml:"redis"`
	Datasource DatasourceConfig `yaml:"datasource"`
}

type MVCConfig struct {
	Async AsyncConfig `yaml:"async"`
}

type AsyncConfig struct {
	RequestTimeout int `yaml:"request-timeout"`
}

type ThymeleafConfig struct {
	Cache bool `yaml:"cache"`
}

type ServletConfig struct {
	Multipart MultipartConfig `yaml:"multipart"`
}

type MultipartConfig struct {
	MaxFileSize    string `yaml:"max-file-size"`
	MaxRequestSize string `yaml:"max-request-size"`
}

type CacheConfig struct {
	Type string `yaml:"type"`
}

type RedisConfig struct {
	Host     string `yaml:"host" json:"host,omitempty"`
	Port     int    `yaml:"port" json:"port,omitempty"`
	Database int    `yaml:"database" json:"database,omitempty"`
	Password string `yaml:"password,omitempty" json:"password,omitempty"`
	Timeout  int    `yaml:"timeout" json:"timeout,omitempty"`
	Enable   bool   `yaml:"-" json:"enable,omitempty"`
}

type DatasourceConfig struct {
	Dynamic DynamicConfig `yaml:"dynamic" json:"dynamic"`
	Enable  bool          `yaml:"-" json:"enable,omitempty"`
}

type DynamicConfig struct {
	Primary    string                `yaml:"primary" json:"primary,omitempty"`
	Datasource map[string]DataSource `yaml:"datasource" json:"datasource,omitempty"`
}

type DataSource struct {
	Type            string       `yaml:"type" json:"type,omitempty"`
	DriverClassName string       `yaml:"driver-class-name" json:"driverClassName,omitempty"`
	Url             string       `yaml:"url" json:"url,omitempty"`
	Username        string       `yaml:"username" json:"username,omitempty"`
	Password        string       `yaml:"password" json:"password,omitempty"`
	Hikari          HikariConfig `yaml:"hikari" json:"hikari"`
}

type HikariConfig struct {
	ConnectionTimeout int `yaml:"connection-timeout"`
	InitialSize       int `yaml:"initialSize"`
	MaximumPoolSize   int `yaml:"maximum-pool-size"`
	MinimumIdle       int `yaml:"minimum-idle"`
	IdleTimeout       int `yaml:"idle-timeout"`
	MaxLifetime       int `yaml:"max-lifetime"`
}

type PageHelperConfig struct {
	HelperDialect string `yaml:"helper-dialect"`
}

type ServerConfig struct {
	Port int       `yaml:"port"`
	SSL  SSLConfig `yaml:"ssl"`
}

type SSLConfig struct {
	Enabled               bool   `yaml:"enabled"`
	KeyStore              string `yaml:"key-store"`
	KeyStorePassword      string `yaml:"key-store-password"`
	Certificate           string `yaml:"certificate"`
	CertificatePrivateKey string `yaml:"certificate-private-key"`
}

type SipConfig struct {
	IP       string `yaml:"ip" json:"IP,omitempty"`
	Port     int    `yaml:"port" json:"port,omitempty"`
	Domain   string `yaml:"domain" json:"domain,omitempty"`
	ID       string `yaml:"id" json:"ID,omitempty"`
	Password string `yaml:"password" json:"password,omitempty"`
	Alarm    bool   `yaml:"alarm" json:"alarm,omitempty"`
	Enable   bool   `yaml:"-" json:"enable,omitempty"`
}

type MediaConfig struct {
	ID               string    `yaml:"id" json:"ID,omitempty"`
	IP               string    `yaml:"ip" json:"IP,omitempty"`
	HttpPort         int       `yaml:"http-port" json:"httpPort,omitempty"`
	StreamIP         string    `yaml:"stream-ip" json:"streamIP,omitempty"`
	SdpIP            string    `yaml:"sdp-ip" json:"sdpIP,omitempty"`
	HookIP           string    `yaml:"hook-ip" json:"hookIP,omitempty"`
	HttpSslPort      string    `yaml:"http-ssl-port" json:"httpSslPort,omitempty"`
	AutoConfig       bool      `yaml:"auto-config" json:"autoConfig,omitempty"`
	Secret           string    `yaml:"secret" json:"secret,omitempty"`
	Rtp              RtpConfig `yaml:"rtp" json:"rtp"`
	RecordAssistPort int       `yaml:"record-assist-port" json:"recordAssistPort,omitempty"`
	Enable           bool      `yaml:"-" json:"enable,omitempty"`
}

type RtpConfig struct {
	Enable        bool   `yaml:"enable" json:"enable,omitempty"`
	PortRange     string `yaml:"port-range" json:"portRange,omitempty"`
	SendPortRange string `yaml:"send-port-range" json:"sendPortRange,omitempty"`
}

type UserSettingsConfig struct {
	ServerID                             string   `yaml:"server-id"`
	AutoApplyPlay                        bool     `yaml:"auto-apply-play"`
	SeniorSdp                            bool     `yaml:"senior-sdp"`
	SavePositionHistory                  bool     `yaml:"save-position-history"`
	PlayTimeout                          int      `yaml:"play-timeout"`
	PlatformPlayTimeout                  int      `yaml:"platform-play-timeout"`
	InterfaceAuthentication              bool     `yaml:"interface-authentication"`
	InterfaceAuthenticationExcludes      []string `yaml:"interface-authentication-excludes"`
	RecordPushLive                       bool     `yaml:"record-push-live"`
	RecordSip                            bool     `yaml:"record-sip"`
	LogInDatabase                        bool     `yaml:"logInDatabase"`
	UsePushingAsStatus                   bool     `yaml:"use-pushing-as-status"`
	UseSourceIpAsStreamIp                bool     `yaml:"use-source-ip-as-stream-ip"`
	StreamOnDemand                       bool     `yaml:"stream-on-demand"`
	PushAuthority                        bool     `yaml:"push-authority"`
	SyncChannelOnDeviceOnline            bool     `yaml:"sync-channel-on-device-online"`
	SipUseSourceIpAsRemoteAddress        bool     `yaml:"sip-use-source-ip-as-remote-address"`
	SipLog                               bool     `yaml:"sip-log"`
	SqlLog                               bool     `yaml:"sql-log"`
	SendToPlatformsWhenIdLost            bool     `yaml:"send-to-platforms-when-id-lost"`
	RefuseChannelStatusChannelFormNotify bool     `yaml:"refuse-channel-status-channel-form-notify"`
	MaxNotifyCountQueue                  int      `yaml:"max-notify-count-queue"`
	DeviceStatusNotify                   bool     `yaml:"device-status-notify"`
	UseCustomSsrcForParentInvite         bool     `yaml:"use-custom-ssrc-for-parent-invite"`
	RegisterAgainAfterTime               int      `yaml:"register-again-after-time"`
	RegisterKeepIntDialog                bool     `yaml:"register-keep-int-dialog"`
	AllowedOrigins                       []string `yaml:"allowed-origins"`
}

type LoggingConfig struct {
	Config string `yaml:"config"`
}

type SpringDocConfig struct {
	ApiDocs   ApiDocsConfig   `yaml:"api-docs"`
	SwaggerUI SwaggerUIConfig `yaml:"swagger-ui"`
}

type ApiDocsConfig struct {
	Enabled bool `yaml:"enabled"`
}

type SwaggerUIConfig struct {
	Enabled bool   `yaml:"enabled"`
	Path    string `yaml:"path"`
}
