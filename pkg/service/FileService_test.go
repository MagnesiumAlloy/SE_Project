package service

import (
	"testing"
)

func TestBackupPackedData(t *testing.T) {
	BackupPackedData("/home/lush/SE_Project/web/html", "/")
}
func TestRecoverPackedData(t *testing.T) {
	RecoverPackedData("/home/lush/tmp/static.cloud", "/home/lush/tmp/newdir")
}
