package utils

import (
	"crypto/md5"
	"hash"
	"strings"
)

func ContainString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func RemoveString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}

func ConcatString(slice []string, sep string) string {

	if slice == nil || len(slice) == 0 {
		return ""
	}

	ns := ""
	for _, s := range slice {
		ns = ns + s + sep
	}

	return strings.TrimSuffix(ns, sep)
}

func HashCode(msg string) string {
	var h hash.Hash = md5.New()
	h.Write([]byte(msg))
	return string(h.Sum(nil))
}
