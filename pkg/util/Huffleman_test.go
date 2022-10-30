package util

import "testing"

func TestCompress(t *testing.T) {
	if res, err := Compress("/home/lush/tmp/static.cloud"); err != nil {
		println(err, res)
	}

	if res, err := Decompress("/home/lush/tmp/static.cloud_tmp_"); err != nil {
		println(err, res)
	}
}
