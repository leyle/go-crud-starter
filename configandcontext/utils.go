package configandcontext

import (
	"errors"
	"io/fs"
	"os"
	"syscall"
)

func CheckPathExist(path string, permission int) error {
	// minPermission:
	// 4 -> only check if it can read
	// 4 + 2 = 6 -> check if it can read and write

	// first check if exist
	if _, err := os.Stat(path); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			// if os.IsNotExist(err) {
			// todo
		} else {
		}
		return err
	}

	// then check if it can read or read/write
	var bit uint32 = syscall.O_RDWR
	if permission < 6 {
		bit = syscall.O_RDONLY
	}

	err := syscall.Access(path, bit)
	if err != nil {
		return err
	}
	return nil
}
