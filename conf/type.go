package conf

type ArchitectureOptions uint

const (
	Standalone ArchitectureOptions = iota + 1
	StandaloneIntercom
)

func (a ArchitectureOptions) String() string {
	switch a {
	case Standalone:
		return "单机版"
	case StandaloneIntercom:
		return "单机-对讲版"
	default:
		return ""
	}
}

type FileType uint

const (
	WVP FileType = iota + 1
	ZLM
	Nginx
	TXT
)

func (f FileType) String() string {
	switch f {
	case WVP:
		return "wvp"
	case ZLM:
		return "zlm"
	case Nginx:
		return "nginx"
	case TXT:
		return "txt"
	default:
		return ""
	}
}
