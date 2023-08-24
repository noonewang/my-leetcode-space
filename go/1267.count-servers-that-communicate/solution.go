// Created by alex at 2023/08/24 17:52
// leetgo: dev
// https://leetcode.cn/problems/count-servers-that-communicate/

package main

import (
	"bufio"
	"fmt"
	. "github.com/j178/leetgo/testutils/go"
	"os"
)

// @lc code=begin

func countServers(grid [][]int) (ans int) {
	m := len(grid)
	n := len(grid[0])
	row := make(map[int]int)
	col := make(map[int]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				row[i] = row[i] + 1
				col[j] = col[j] + 1
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && (row[i] > 1 || col[j] > 1) {
				ans++
			}
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	grid := Deserialize[[][]int](ReadLine(stdin))
	//grid := [][]int{{1, 0}, {1, 1}}
	ans := countServers(grid)

	fmt.Println("\noutput:", Serialize(ans))
}
