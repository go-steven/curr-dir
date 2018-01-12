package curr_dir

import (
	"path"
	"runtime"
	"strings"
)

func GetCurrDir() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}

	if runtime.GOOS == "windows" {
		file = strings.Replace(file, "\\", "/", -1)
	}

	return path.Dir(file)
}
