// Created by alex at 2023/06/24 21:10
// leetgo: dev
// https://leetcode.cn/problems/ti-huan-kong-ge-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	// s := " We are happy."
	ans := replaceSpace(s)

	fmt.Println("\noutput:", Serialize(ans))
}
