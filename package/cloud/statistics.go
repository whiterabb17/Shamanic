package cloud

import gophersocket "github.com/whiterabb17/gopher-socket"

var (
	ClientIdList   []string
	CommsHistory   []string
	ExecHistory    []string
	RebroadHistory []string
	Status         bool
	Keeper         *gophersocket.Server
	Channels       []*gophersocket.Channel
)

type Spell struct {
	Key string `json:"key"`
	Val string `json:"val"`
}
type Dispatch struct {
	Cmd    string `json:"cmd"`
	Args   string `json:"arg"`
	ArgCnt string `json:"cnt"`
	Tag    string `json:"tag"`
}

type Resp struct {
	Resp string `json:"res"`
	Tag  string `json:"tag"`
}
