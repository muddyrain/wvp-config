package model

// ZLMConfig 整体配置结构体
type ZLMConfig struct {
	API       APIConfig       `mapstructure:"api" ini:"api"`
	Cluster   ClusterConfig   `mapstructure:"cluster" ini:"cluster"`
	FFMPEG    FFMPEGConfig    `mapstructure:"ffmpeg" ini:"ffmpeg"`
	General   GeneralConfig   `mapstructure:"general" ini:"general"`
	HLS       HLSConfig       `mapstructure:"hls" ini:"hls"`
	Hook      HookConfig      `mapstructure:"hook" ini:"hook"`
	HTTP      HTTPConfig      `mapstructure:"http" ini:"http"`
	Multicast MulticastConfig `mapstructure:"multicast" ini:"multicast"`
	Protocol  ProtocolConfig  `mapstructure:"protocol" ini:"protocol"`
	Record    RecordConfig    `mapstructure:"record" ini:"record"`
	RTC       RTCConfig       `mapstructure:"rtc" ini:"rtc"`
	RTMP      RTMPConfig      `mapstructure:"rtmp" ini:"rtmp"`
	RTP       RTPConfig       `mapstructure:"rtp" ini:"rtp"`
	RTPProxy  RTPProxyConfig  `mapstructure:"rtp_proxy" ini:"rtp_proxy"`
	RTSP      RTSPConfig      `mapstructure:"rtsp" ini:"rtsp"`
	Shell     ShellConfig     `mapstructure:"shell" ini:"shell"`
	SRT       SRTConfig       `mapstructure:"srt" ini:"srt"`
}

// APIConfig api相关配置
type APIConfig struct {
	ApiDebug     int    `mapstructure:"apiDebug" ini:"apiDebug"`
	DefaultSnap  string `mapstructure:"defaultSnap" ini:"defaultSnap"`
	DownloadRoot string `mapstructure:"downloadRoot" ini:"downloadRoot"`
	Secret       string `mapstructure:"secret" ini:"secret"`
	SnapRoot     string `mapstructure:"snapRoot" ini:"snapRoot"`
}

// ClusterConfig cluster相关配置
type ClusterConfig struct {
	OriginURL  string `mapstructure:"origin_url" ini:"origin_url"`
	RetryCount int    `mapstructure:"retry_count" ini:"retry_count"`
	TimeoutSec int    `mapstructure:"timeout_sec" ini:"timeout_sec"`
}

// FFMPEGConfig ffmpeg相关配置
type FFMPEGConfig struct {
	Bin        string `mapstructure:"bin" ini:"bin"`
	Cmd        string `mapstructure:"cmd" ini:"cmd"`
	Log        string `mapstructure:"log" ini:"log"`
	RestartSec int    `mapstructure:"restart_sec" ini:"restartSec"`
	Snap       string `mapstructure:"snap" ini:"snap"`
}

// GeneralConfig general相关配置
type GeneralConfig struct {
	BroadcastPlayerCountChanged int    `mapstructure:"broadcast_player_count_changed" ini:"broadcast_player_count_changed"`
	CheckNvidiaDev              int    `mapstructure:"check_nvidia_dev" ini:"check_nvidia_dev"`
	EnableVhost                 int    `mapstructure:"enableVhost" ini:"enableVhost"`
	EnableFFmpegLog             int    `mapstructure:"enable_ffmpeg_log" ini:"enable_ffmpeg_log"`
	FlowThreshold               int    `mapstructure:"flowThreshold" ini:"flowThreshold"`
	MaxStreamWaitMS             int    `mapstructure:"maxStreamWaitMS" ini:"maxStreamWaitMS"`
	MediaServerId               string `mapstructure:"mediaServerId" ini:"mediaServerId"`
	MergeWriteMS                int    `mapstructure:"mergeWriteMS" ini:"mergeWriteMS"`
	ResetWhenRePlay             int    `mapstructure:"resetWhenRePlay" ini:"resetWhenRePlay"`
	StreamNoneReaderDelayMS     int    `mapstructure:"streamNoneReaderDelayMS" ini:"streamNoneReaderDelayMS"`
	UnreadyFrameCache           int    `mapstructure:"unready_frame_cache" ini:"unready_frame_cache"`
	WaitAddTrackMS              int    `mapstructure:"wait_add_track_ms" ini:"wait_add_track_ms"`
	WaitTrackReadyMS            int    `mapstructure:"wait_track_ready_ms" ini:"wait_track_ready_ms"`
}

// HLSConfig hls相关配置
type HLSConfig struct {
	BroadcastRecordTs int `mapstructure:"broadcastRecordTs" ini:"broadcastRecordTs"`
	DeleteDelaySec    int `mapstructure:"deleteDelaySec" ini:"deleteDelaySec"`
	FastRegister      int `mapstructure:"fastRegister" ini:"fastRegister"`
	FileBufSize       int `mapstructure:"fileBufSize" ini:"fileBufSize"`
	SegDelay          int `mapstructure:"segDelay" ini:"segDelay"`
	SegDur            int `mapstructure:"segDur" ini:"segDur"`
	SegKeep           int `mapstructure:"segKeep" ini:"segKeep"`
	SegNum            int `mapstructure:"segNum" ini:"segNum"`
	SegRetain         int `mapstructure:"segRetain" ini:"segRetain"`
}

// HookConfig hook相关配置
type HookConfig struct {
	AliveInterval        float64 `mapstructure:"alive_interval" ini:"alive_interval"`
	Enable               int     `mapstructure:"enable" ini:"enable"`
	OnFlowReport         string  `mapstructure:"on_flow_report" ini:"on_flow_report"`
	OnHTTPAccess         string  `mapstructure:"on_http_access" ini:"on_http_access"`
	OnPlay               string  `mapstructure:"on_play" ini:"on_play"`
	OnPublish            string  `mapstructure:"on_publish" ini:"on_publish"`
	OnRecordMP4          string  `mapstructure:"on_record_mp4" ini:"on_record_mp4"`
	OnRecordTs           string  `mapstructure:"on_record_ts" ini:"on_record_ts"`
	OnRtpServerTimeout   string  `mapstructure:"on_rtp_server_timeout" ini:"on_rtp_server_timeout"`
	OnRtspAuth           string  `mapstructure:"on_rtsp_auth" ini:"on_rtsp_auth"`
	OnRtspRealm          string  `mapstructure:"on_rtsp_realm" ini:"on_rtsp_realm"`
	OnSendRtpStopped     string  `mapstructure:"on_send_rtp_stopped" ini:"on_send_rtp_stopped"`
	OnServerExited       string  `mapstructure:"on_server_exited" ini:"on_server_exited"`
	OnServerKeepalive    string  `mapstructure:"on_server_keepalive" ini:"on_server_keepalive"`
	OnServerStarted      string  `mapstructure:"on_server_started" ini:"on_server_started"`
	OnShellLogin         string  `mapstructure:"on_shell_login" ini:"on_shell_login"`
	OnStreamChanged      string  `mapstructure:"on_stream_changed" ini:"on_stream_changed"`
	OnStreamNoneReader   string  `mapstructure:"on_stream_none_reader" ini:"on_stream_none_reader"`
	OnStreamNotFound     string  `mapstructure:"on_stream_not_found" ini:"on_stream_not_found"`
	Retry                int     `mapstructure:"retry" ini:"retry"`
	RetryDelay           float64 `mapstructure:"retry_delay" ini:"retry_delay"`
	StreamChangedSchemas string  `mapstructure:"stream_changed_schemas" ini:"stream_changed_schemas"`
	TimeoutSec           int     `mapstructure:"timeoutSec" ini:"timeoutSec"`
}

// HTTPConfig http相关配置
type HTTPConfig struct {
	AllowCrossDomains int    `mapstructure:"allow_cross_domains" ini:"allow_cross_domains"`
	AllowIPRange      string `mapstructure:"allow_ip_range" ini:"allow_ip_range"`
	CharSet           string `mapstructure:"charSet" ini:"charSet"`
	DirMenu           int    `mapstructure:"dirMenu" ini:"dirMenu"`
	ForbidCacheSuffix string `mapstructure:"forbidCacheSuffix" ini:"forbidCacheSuffix"`
	ForwardedIPHeader string `mapstructure:"forwarded_ip_header" ini:"forwarded_ip_header"`
	KeepAliveSecond   int    `mapstructure:"keepAliveSecond" ini:"keepAliveSecond"`
	MaxReqSize        int    `mapstructure:"maxReqSize" ini:"maxReqSize"`
	NotFound          string `mapstructure:"notFound" ini:"notFound"`
	Port              int    `mapstructure:"port" ini:"port"`
	RootPath          string `mapstructure:"rootPath" ini:"rootPath"`
	SendBufSize       int    `mapstructure:"sendBufSize" ini:"sendBufSize"`
	SSLPort           int    `mapstructure:"sslport" ini:"sslport"`
	VirtualPath       string `mapstructure:"virtualPath" ini:"virtualPath"`
}

// MulticastConfig multicast相关配置
type MulticastConfig struct {
	AddrMax string `mapstructure:"addrMax" ini:"addrMax"`
	AddrMin string `mapstructure:"addrMin" ini:"addrMin"`
	UdpTTL  int    `mapstructure:"udpTTL" ini:"udpTTL"`
}

// ProtocolConfig protocol相关配置
type ProtocolConfig struct {
	AddMuteAudio   int    `mapstructure:"add_mute_audio" ini:"add_mute_audio"`
	AutoClose      int    `mapstructure:"auto_close" ini:"auto_close"`
	ContinuePushMS int    `mapstructure:"continue_push_ms" ini:"continue_push_ms"`
	EnableAudio    int    `mapstructure:"enable_audio" ini:"enable_audio"`
	EnableFMP4     int    `mapstructure:"enable_fmp4" ini:"enable_fmp4"`
	EnableHLS      int    `mapstructure:"enable_hls" ini:"enable_hls"`
	EnableHLSFMP4  int    `mapstructure:"enable_hls_fmp4" ini:"enable_hls_fmp4"`
	EnableMP4      int    `mapstructure:"enable_mp4" ini:"enable_mp4"`
	EnableRTMP     int    `mapstructure:"enable_rtmp" ini:"enable_rtmp"`
	EnableRTSP     int    `mapstructure:"enable_rtsp" ini:"enable_rtsp"`
	EnableTS       int    `mapstructure:"enable_ts" ini:"enable_ts"`
	FMP4Demand     int    `mapstructure:"fmp4_demand" ini:"fmp4_demand"`
	HLSDemand      int    `mapstructure:"hls_demand" ini:"hls_demand"`
	HLSSavePath    string `mapstructure:"hls_save_path" ini:"hls_save_path"`
	ModifyStamp    int    `mapstructure:"modify_stamp" ini:"modify_stamp"`
	MP4AsPlayer    int    `mapstructure:"mp4_as_player" ini:"mp4_as_player"`
	MP4MaxSecond   int    `mapstructure:"mp4_max_second" ini:"mp4_max_second"`
	MP4SavePath    string `mapstructure:"mp4_save_path" ini:"mp4_save_path"`
	PacedSenderMS  int    `mapstructure:"paced_sender_ms" ini:"paced_sender_ms"`
	RTMPSDemand    int    `mapstructure:"rtmp_demand" ini:"rtmp_demand"`
	RTSPDemand     int    `mapstructure:"rtsp_demand" ini:"rtsp_demand"`
	TSIDemand      int    `mapstructure:"ts_demand" ini:"ts_demand"`
}

// RecordConfig record相关配置
type RecordConfig struct {
	AppName     string `mapstructure:"appName" ini:"appName"`
	EnableFmp4  int    `mapstructure:"enableFmp4" ini:"enableFmp4"`
	FastStart   int    `mapstructure:"fastStart"  ini:"fastStart"`
	FileBufSize int    `mapstructure:"fileBufSize" ini:"fileBufSize"`
	FileRepeat  int    `mapstructure:"fileRepeat" ini:"fileRepeat"`
	SampleMS    int    `mapstructure:"sampleMS" ini:"sampleMS"`
}

// RTCConfig rtc相关配置
type RTCConfig struct {
	ExternIP          string  `mapstructure:"externIP" ini:"externIP"`
	MaxRtpCacheMS     int     `mapstructure:"maxRtpCacheMS" ini:"maxRtpCacheMS"`
	MaxRtpCacheSize   int     `mapstructure:"maxRtpCacheSize" ini:"maxRtpCacheSize"`
	MaxBitrate        int     `mapstructure:"max_bitrate" ini:"max_bitrate"`
	MinBitrate        int     `mapstructure:"min_bitrate" ini:"min_bitrate"`
	NackIntervalRatio float64 `mapstructure:"nackIntervalRatio" ini:"nackIntervalRatio"`
	NackMaxCount      int     `mapstructure:"nackMaxCount" ini:"nackMaxCount"`
	NackMaxMS         int     `mapstructure:"nackMaxMS" ini:"nackMaxMS"`
	NackMaxSize       int     `mapstructure:"nackMaxSize" ini:"nackMaxSize"`
	NackRtpSize       int     `mapstructure:"nackRtpSize" ini:"nackRtpSize"`
	Port              int     `mapstructure:"port" ini:"port"`
	PreferredCodecA   string  `mapstructure:"preferredCodecA" ini:"preferredCodecA"`
	PreferredCodecV   string  `mapstructure:"preferredCodecV" ini:"preferredCodecV"`
	RembBitRate       int     `mapstructure:"rembBitRate" ini:"rembBitRate"`
	StartBitrate      int     `mapstructure:"start_bitrate" ini:"start_bitrate"`
	TcpPort           int     `mapstructure:"tcpPort" ini:"tcpPort"`
	TimeoutSec        int     `mapstructure:"timeoutSec" ini:"timeoutSec"`
}

// RTMPConfig rtmp相关配置
type RTMPConfig struct {
	DirectProxy     int `mapstructure:"directProxy" ini:"directProxy"`
	Enhanced        int `mapstructure:"enhanced" ini:"enhanced"`
	HandshakeSecond int `mapstructure:"handshakeSecond" ini:"handshakeSecond"`
	KeepAliveSecond int `mapstructure:"keepAliveSecond" ini:"keepAliveSecond"`
	Port            int `mapstructure:"port" ini:"port"`
	SSLPort         int `mapstructure:"sslport" ini:"SSLPort"`
}

// RTPConfig rtp相关配置
type RTPConfig struct {
	AudioMtuSize int `mapstructure:"audioMtuSize" ini:"audioMtuSize"`
	H264StapA    int `mapstructure:"h264_stap_a" ini:"h264_stap_a"`
	LowLatency   int `mapstructure:"lowLatency" ini:"lowLatency"`
	RtpMaxSize   int `mapstructure:"rtpMaxSize" ini:"rtpMaxSize"`
	VideoMtuSize int `mapstructure:"videoMtuSize" ini:"videoMtuSize"`
}

// RTPProxyConfig rtp_proxy相关配置
type RTPProxyConfig struct {
	DumpDir             string `mapstructure:"dumpDir" ini:"dumpDir"`
	GopCache            int    `mapstructure:"gop_cache" ini:"gop_cache"`
	H264Pt              int    `mapstructure:"h264_pt" ini:"h264_pt"`
	H265Pt              int    `mapstructure:"h265_pt" ini:"h265_pt"`
	OpusPt              int    `mapstructure:"opus_pt" ini:"opus_pt"`
	Port                int    `mapstructure:"port" ini:"port"`
	PortRange           string `mapstructure:"port_range" ini:"port_range"`
	PsPt                int    `mapstructure:"ps_pt" ini:"ps_pt"`
	RtpG711DurMS        int    `mapstructure:"rtp_g711_dur_ms" ini:"rtp_g711_dur_ms"`
	TimeoutSec          int    `mapstructure:"timeoutSec" ini:"timeoutSec"`
	UdpRecvSocketBuffer int    `mapstructure:"udp_recv_socket_buffer" ini:"udp_recv_socket_buffer"`
}

// RTSPConfig rtsp相关配置
type RTSPConfig struct {
	AuthBasic        int `mapstructure:"authBasic" ini:"authBasic"`
	DirectProxy      int `mapstructure:"directProxy" ini:"directProxy"`
	HandshakeSecond  int `mapstructure:"handshakeSecond" ini:"handshakeSecond"`
	KeepAliveSecond  int `mapstructure:"keepAliveSecond" ini:"keepAliveSecond"`
	LowLatency       int `mapstructure:"lowLatency" ini:"lowLatency"`
	Port             int `mapstructure:"port" ini:"port"`
	RtpTransportType int `mapstructure:"rtpTransportType" ini:"rtpTransportType"`
	SSLPort          int `mapstructure:"sslport" ini:"sslport"`
}

// ShellConfig shell相关配置
type ShellConfig struct {
	MaxReqSize int `mapstructure:"maxReqSize" ini:"maxReqSize"`
	Port       int `mapstructure:"port" ini:"port"`
}

// SRTConfig srt相关配置
type SRTConfig struct {
	LatencyMul int `mapstructure:"latencyMul" ini:"latencyMul"`
	PktBufSize int `mapstructure:"pktBufSize" ini:"pktBufSize"`
	Port       int `mapstructure:"port" ini:"port"`
	TimeoutSec int `mapstructure:"timeoutSec" ini:"timeoutSec"`
}
