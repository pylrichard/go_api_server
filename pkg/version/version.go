package version.

import (
	"fmt"
	"runtime"
)

// Info contains version info
type Info struct {
	GitTag       string `json:"gitTag"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

func (info Info) String() string {
	return info.GitTag + ":" + info.GitCommit + ":" 
			+ info.GitTreeState + ":" + info.BuildDate + ":" 
			+ info.GoVersion + ":" + info.Compiler + ":" + info.Platform
}

func Get() Info {
	return Info {
		GitTag:       gitTag,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
	}
}