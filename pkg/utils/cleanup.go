package utils

import "os"

func DeleteLocalVidoeFile(key string) error {
	err := os.Remove(key)
	return err
}
