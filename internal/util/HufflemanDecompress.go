package util

import (
	"SE_Project/internal/model"
	"bufio"
	"os"
	"strconv"
)

var data []byte

type Trie struct {
	p, ls, rs, v int
}

var T []Trie
var tot int
var rt int
var Len int

func readRawData(path string) (map[int]string, uint64, error) {
	rfile, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}
	info, err := os.Stat(path)
	if err != nil {
		return nil, 0, err
	}
	defer rfile.Close()
	r := bufio.NewReader(rfile)
	mp := make(map[int]string)
	var size uint64
	str, err := r.ReadString('\n')
	if err != nil {
		return nil, 0, err
	}
	x, err := strconv.Atoi(str[:len(str)-1])
	if err != nil {
		return nil, 0, err
	}
	size = uint64(x)
	for i := 0; i < 256; i++ {
		str, err := r.ReadString('\n')
		if err != nil {
			return nil, 0, err
		}
		mp[i] = str[:len(str)-1]
	}
	data = make([]byte, info.Size()*8)
	p := make([]byte, 4096)
	c := 0
	for {
		len, err := r.Read(p)
		for i := 0; i < len; i++ {
			for j := 0; j < 8; j++ {
				data[c+7-j] = p[i] & 1
				p[i] >>= 1
			}
			c += 8
		}
		if err != nil {
			break
		}
	}
	Len = c
	return mp, size, nil
}

func insert(s string, id int) {
	cur := rt
	for _, c := range s {
		if c == '0' {
			if T[cur].ls == -1 {
				T[cur].ls = tot
				tot++
				T = append(T, Trie{p: cur, ls: -1, rs: -1, v: -1})
			}
			cur = T[cur].ls
		} else {
			if T[cur].rs == -1 {
				T[cur].rs = tot
				tot++
				T = append(T, Trie{p: cur, ls: -1, rs: -1, v: -1})
			}
			cur = T[cur].rs
		}
	}
	T[cur].v = id
}

func buildTrie(mp map[int]string) {
	T = []Trie{}
	T = append(T, Trie{p: -1, ls: -1, rs: -1, v: -1})
	rt = 0
	tot = 1
	for i := 0; i < 256; i++ {
		insert(mp[i], i)
	}
}

func decode(path string) error {
	desPath := path + model.CloudTempType
	wfile, err := os.OpenFile(desPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer wfile.Close()
	w := bufio.NewWriter(wfile)
	ptr := 0
	cur := rt
	var cnt uint64 = 0
	for ptr < Len {
		if data[ptr] == 0 {
			if T[cur].ls == -1 {
				if T[cur].v == -1 {
					break
				} else {
					w.WriteByte(byte(T[cur].v))
					cur = rt
					cnt++
					ptr--
				}
			} else {
				cur = T[cur].ls
			}
		} else {
			if T[cur].rs == -1 {
				if T[cur].v == -1 {
					break
				} else {
					w.WriteByte(byte(T[cur].v))
					cur = rt
					cnt++
					ptr--
				}
			} else {
				cur = T[cur].rs
			}
		}
		ptr++
	}
	w.Flush()
	return nil
}

func Decompress(path string) error {

	mp, _, err := readRawData(path)
	if err != nil {
		return err
	}
	buildTrie(mp)
	if err := decode(path); err != nil {
		return err
	}
	return nil
}
