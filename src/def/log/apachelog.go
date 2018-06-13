package log

// ApacheLog struct
type ApacheLog struct {
	RemoteHost  string
	UserIdentd  string
	UserID      string
	ReqMethod   string
	ReqResource string
	ReqProtocol string
	ReqTime     string
	StatusCode  string
	SizeByte    uint64
}

// Init - Initialize
func Init() *ApacheLog {
	log := new(ApacheLog)
	return log
}
