package util

import "testing"

func TestAES128_CBC_Encrypt(t *testing.T) {
	input := []byte("1234567890abcdABCD!")
	leng := len(input)
	output := make([]byte, (leng+16)/16*16)
	pkey := "123abcABC"
	key := PasswordPadding(pkey)
	AES128_CBC_Encrypt(input, output, key, uint64(len(input)))
	res := make([]byte, len(output))
	AES128_CBC_Decrypt(output, res, key, uint64(len(output)))
	println(string(res))
}

func TestAES128_CBC(t *testing.T) {
	var i, j int
	tmp := [][]byte{
		{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4},
	}
	state := [][]byte{
		{1, 1, 1, 1}, {2, 2, 2, 2}, {3, 3, 3, 3}, {4, 4, 4, 4},
	}
	for i = 0; i < 4; i++ {
		for j = 0; j < 4; j++ {
			state[i][j] = tmp[i][(j+i)%4]
		}
	}
	for i = 0; i < 4; i++ {
		for j = 0; j < 4; j++ {
			tmp[i][j] = state[i][(j-i+4)%4]
		}
	}
	print(tmp)
}
