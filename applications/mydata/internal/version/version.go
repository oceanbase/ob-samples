package version

import (
	"fmt"
	"log"
)

var (
	ReleaseVersion = "None"
	BuildTS        = "None"
	GitHash        = "None"
	GitBranch      = "None"
	GoVersion      = "None"
)

func LogVersionInfo() {
	log.Printf("welcome to use the command, release-version:%s, git-hash:%s, git-branch:%s, utc-build-time:%s, go-version:%s",
		ReleaseVersion, GitHash, GitBranch, BuildTS, GoVersion)
}

func GetRawInfo() string {
	var info string
	info += fmt.Sprintf("Release version   : %s\n", ReleaseVersion)
	info += fmt.Sprintf("Git Commit Hash   : %s\n", GitHash)
	info += fmt.Sprintf("Git Branch        : %s\n", GitBranch)
	info += fmt.Sprintf("Build Time        : %s\n", BuildTS)
	info += fmt.Sprintf("Go Version        : %s\n", GoVersion)
	return info
}
