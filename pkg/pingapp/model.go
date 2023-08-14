package pingapp

type PongInfo struct {
	HttpMethod string `json:"method"`
	UserAgent  string `json:"userAgent"`
	ReqId      string `json:"reqId"`
	Version    string `json:"version"`
	CommitId   string `json:"commitId"`
}
