package util

import (
	"SE_Project/internal/model"
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type node struct {
	p      int
	ls, rs int
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func getCnt(path string, size uint64) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	r := bufio.NewReader(file)
	p := make([]byte, 4096)
	if err != nil {
		return nil, err
	}
	cnt := make([]int, 256)
	for i := 0; i < 256; i++ {
		cnt[i] = 0
	}
	for {
		len, err := r.Read(p)
		for i := 0; i < len; i++ {
			cnt[uint8(p[i])]++
		}
		if err != nil {
			break
		}
	}

	return cnt, nil
}

func dfs(u int, s string, mp []string, nd []node) {
	if u < 256 {
		mp[u] = s
		return
	}
	if nd[u].ls != -1 {
		dfs(nd[u].ls, s+"0", mp, nd)
	}
	if nd[u].rs != -1 {
		dfs(nd[u].rs, s+"1", mp, nd)
	}
}

func genMap(cnt []int) ([]string, error) {
	c := 255
	nd := make([]node, 256)
	Q := make(PriorityQueue, 256)
	for i := 0; i < 256; i++ {
		Q[i] = &Item{
			index:    i,
			priority: cnt[i],
			value:    i,
		}
		nd[i].ls = -1
		nd[i].rs = -1
	}
	heap.Init(&Q)
	for Q.Len() > 1 {
		x := heap.Pop(&Q).(*Item)
		y := heap.Pop(&Q).(*Item)
		c++
		nd[x.value].p = c
		nd[y.value].p = c
		nd = append(nd, node{p: -1, ls: x.value, rs: y.value})
		heap.Push(&Q, &Item{value: c, priority: x.priority + y.priority})
	}
	root := heap.Pop(&Q).(*Item).value
	mp := make([]string, 256)
	dfs(root, "", mp, nd)
	return mp, nil
}

func writeFile(mp []string, path string, size uint64) error {
	desPath := path[:len(path)-len(model.CloudTempType)]
	wfile, err := os.OpenFile(desPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer wfile.Close()
	rfile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer rfile.Close()

	r := bufio.NewReader(rfile)
	w := bufio.NewWriter(wfile)
	w.WriteString(fmt.Sprint(size) + "\n")
	for i := 0; i < 256; i++ {
		w.WriteString(mp[i] + "\n")
	}

	p := make([]byte, 4096)
	if err != nil {
		return err
	}
	var byteToWrite uint8
	byteToWrite = 0
	var cnt int
	cnt = 0
	str := ""
	for {
		len, err := r.Read(p)
		for i := 0; i < len; i++ {
			str = mp[uint8(p[i])]
			for _, bit := range str {
				cnt++
				byteToWrite <<= 1
				if bit == '1' {
					byteToWrite |= 1
				}
				if cnt&8 != 0 {
					cnt = 0
					w.WriteByte(byteToWrite)
					byteToWrite = 0
				}
			}
		}
		w.Flush()
		if err != nil {
			break
		}
	}
	if cnt > 0 {
		for cnt < 8 {
			byteToWrite <<= 1
			cnt++
		}
		w.WriteByte(byteToWrite)
	}
	w.Flush()
	return nil
}

func Compress(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return nil
	}
	//生成字符频次统计表
	cnt, err := getCnt(path, uint64(info.Size()))
	if err != nil {
		return err
	}
	//采用优先队列、dfs生成字符转换表
	mp, err := genMap(cnt)
	if err != nil {
		return err
	}
	//写压缩文件
	if err := writeFile(mp, path, uint64(info.Size())); err != nil {
		return err
	}
	return nil
}
