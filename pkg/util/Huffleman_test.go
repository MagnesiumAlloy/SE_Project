package util

import "testing"

func TestCompress(t *testing.T) {
	if res, err := Compress("/home/lush/SE_Project/Input.pack"); err != nil {
		println(err, res)
	}

	if res, err := Decompress("/home/lush/SE_Project/Input.pack_tmp_"); err != nil {
		println(err, res)
	}
}
