package util

import "testing"

func TestCompress(t *testing.T) {
	if err := Compress("/home/lush/tmp/ccloud.zip_tmp_"); err != nil {
		println(err)
	}

}
