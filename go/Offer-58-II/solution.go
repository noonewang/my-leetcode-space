// Created by alex at 2023/06/24 22:14
// leetgo: dev
// https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseLeftWords(s string, n int) string {
	bytes := []byte(s)
	sb := strings.Builder{}
	str1 := bytes[n:]
	str2 := bytes[:n]
	sb.WriteString(string(str1))
	sb.WriteString(string(str2))
	return sb.String()
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	n := Deserialize[int](ReadLine(stdin))
	ans := reverseLeftWords(s, n)

	fmt.Println("\noutput:", Serialize(ans))
}
