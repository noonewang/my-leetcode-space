// Created by alex at 2023/08/09 17:20
// leetgo: dev
// https://leetcode.cn/problems/average-of-levels-in-binary-tree/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type data struct {
	sum   int
	count int
}

func averageOfLevels(root *TreeNode) (ans []float64) {
	levelData := []data{}
	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		// 当前层已经遍历过
		if level < len(levelData) {
			levelData[level].sum += root.Val
			levelData[level].count++
		} else {
			levelData = append(levelData, data{root.Val, 1})
		}
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)
	for _, datum := range levelData {
		avg := float64(datum.sum) / float64(datum.count)
		ans = append(ans, avg)
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := averageOfLevels(root)

	fmt.Println("\noutput:", Serialize(ans))
}
