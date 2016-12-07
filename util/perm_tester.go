// +build darwin,linux
package util

import (
	"os"
	"syscall"
)

func CanRead(file string) bool {
	stat, err := os.Stat(file)
	if err != nil {
		return false
	}
	fm := stat.Mode()
	if fm & (1 << 2) != 0 {
		return true
	} else if (fm & (1 << 5)) != 0 && (os.Getegid() == int(stat.Sys().(*syscall.Stat_t).Gid)) {
		return true
	} else if (fm & (1 << 8)) != 0 && (os.Geteuid() == int(stat.Sys().(*syscall.Stat_t).Uid)) {
		return true
	}
	return false
}
