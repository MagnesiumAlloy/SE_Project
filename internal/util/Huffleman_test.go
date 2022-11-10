package util

import "testing"

func TestCompress(t *testing.T) {
	//if err := Compress("/home/lush/Cloud_Backup/1/114.mp3.cloud.ctmp"); err != nil {
	//	println(err)
	//}
	if err := Decompress("/home/lush/Cloud_Backup/1/114.mp3.cloud"); err != nil {
		println(err)
	}

}
