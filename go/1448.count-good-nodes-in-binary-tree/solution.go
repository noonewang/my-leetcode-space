// Created by alex at 2023/08/25 15:05
// leetgo: dev
// https://leetcode.cn/problems/count-good-nodes-in-binary-tree/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func goodNodes(root *TreeNode) (ans int) {
	// 根节点当前的最大值是根节点本身
	var dfs func(root *TreeNode, max int)
	dfs = func(root *TreeNode, max int) {
		if root == nil {
			return
		}
		// 这是一个好节点
		if root.Val >= max {
			max = root.Val
			ans++
		}
		dfs(root.Left, max)
		dfs(root.Right, max)
	}
	dfs(root, math.MinInt)
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := goodNodes(root)

	fmt.Println("\noutput:", Serialize(ans))
}
