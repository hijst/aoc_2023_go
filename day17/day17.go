package main

import (
	"aoc_2023_go/util"
	"container/heap"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type Dir struct {
	v int
	h int
}

func main() {
	lines, _ := util.ReadFileLines("")
	fmt.Println("______part_1______")
	start := time.Now()
	one(lines)
	fmt.Printf("took %s\n", time.Since(start))

	fmt.Println("______part_2______")
	start = time.Now()
	two(lines)
	fmt.Printf("took %s\n", time.Since(start))
}

func one(lines []string) {
	grid := readGrid(lines)
	res := dijkstra(grid, 0, 0, 0, 0, 0, 0, 3)
	fmt.Printf("answer: %d\n", res)
}

func two(lines []string) {
	grid := readGrid(lines)
	res := dijkstra(grid, 0, 0, 0, 0, 0, 4, 10)
	fmt.Printf("answer: %d\n", res)
}

func readGrid(lines []string) [][]int {
	var res [][]int
	for _, line := range lines {
		s := strings.Split(line, "")
		r := util.Map(s, atoi)
		res = append(res, r)
	}
	return res
}

func dijkstra(grid [][]int, r, c, vdir, hdir, same, minsame, maxsame int) int {
	dist := make([][][][][]int, len(grid))
	for i := range dist {
		dist[i] = make([][][][]int, len(grid[0]))
		for j := range dist[i] {
			dist[i][j] = make([][][]int, 4)
			for k := range dist[i][j] {
				dist[i][j][k] = make([][]int, 4)
				for l := range dist[i][j][k] {
					dist[i][j][k][l] = make([]int, maxsame+1)
					for m := range dist[i][j][k][l] {
						dist[i][j][k][l][m] = math.MaxInt
					}
				}
			}
		}
	}
	dist[0][0][0][0][0] = 0
	dist[0][0][1][1][0] = 0
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{value: hash(r, c, vdir, hdir, same), priority: 0})
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		cr, cc, cvdir, chdir, csame := unhash(item.value)
		nbs := getOptions(grid, cr, cc, cvdir, chdir, csame, minsame, maxsame)

		for _, nb := range nbs {
			alt := dist[cr][cc][cvdir+1][chdir+1][csame] + grid[cr+nb.v][cc+nb.h]
			s := 1
			if cvdir == nb.v && chdir == nb.h {
				s += csame
			}
			if alt < dist[cr+nb.v][cc+nb.h][nb.v+1][nb.h+1][s] {
				if !(cr+nb.v == len(grid)-1 && cc+nb.h == len(grid[0])-1 && s < minsame) {
					dist[cr+nb.v][cc+nb.h][nb.v+1][nb.h+1][s] = alt
					heap.Push(&pq, &Item{value: hash(cr+nb.v, cc+nb.h, nb.v, nb.h, s), priority: alt})
				}
			}
		}
	}
	return getMin(dist[len(grid)-1][len(grid[0])-1])
}

func getMin(vdirs [][][]int) int {
	res := math.MaxInt
	for _, hdirs := range vdirs {
		for _, sames := range hdirs {
			for _, value := range sames {
				res = min(res, value)
			}
		}
	}
	return res
}

func getOptions(grid [][]int, r, c, vdir, hdir, same, minsame, maxsame int) []Dir {
	var res []Dir
	if same < minsame && !(r == 0 && c == 0) {
		if (c < len(grid[0])-1 && hdir == 1) || (c > 0 && hdir == -1) || (r < len(grid)-1 && vdir == 1) || (r > 0 && vdir == -1) {
			res = append(res, Dir{v: vdir, h: hdir})
		}
		return res
	}
	if c < len(grid[0])-1 && hdir != -1 && !(hdir == 1 && same == maxsame) {
		res = append(res, Dir{v: 0, h: 1})
	}
	if r < len(grid)-1 && vdir != -1 && !(vdir == 1 && same == maxsame) {
		res = append(res, Dir{v: 1, h: 0})
	}
	if c > 0 && hdir != 1 && !(hdir == -1 && same == maxsame) {
		res = append(res, Dir{v: 0, h: -1})
	}
	if r > 0 && vdir != 1 && !(vdir == -1 && same == maxsame) {
		res = append(res, Dir{v: -1, h: 0})
	}
	return res
}

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func hash(nums ...int) string {
	var res string
	for _, num := range nums {
		res = res + strconv.Itoa(num) + ","
	}
	return res
}

func unhash(s string) (int, int, int, int, int) {
	res := strings.Split(s, ",")
	return atoi(res[0]), atoi(res[1]), atoi(res[2]), atoi(res[3]), atoi(res[4])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Priority queue implementation from go docs https://pkg.go.dev/container/heap

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// only modification from docs -> make it MIN priority (was max)
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
