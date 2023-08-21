// Created by alex at 2023/08/21 21:06
// leetgo: dev
// https://leetcode.cn/problems/move-pieces-to-obtain-a-string/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func canChange(start string, target string) bool {
	if strings.ReplaceAll(start, "_", "") != strings.ReplaceAll(target, "_", "") {
		return false
	}
	j := 0
	for i, c := range start {
		if c != '_' {
			for target[j] == '_' {
				j++
			}
			if i != j && c == 'L' == (i < j) {
				return false
			}
			j++
		}
	}
	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	start := Deserialize[string](ReadLine(stdin))
	target := Deserialize[string](ReadLine(stdin))
	ans := canChange(start, target)

	fmt.Println("\noutput:", Serialize(ans))
}
