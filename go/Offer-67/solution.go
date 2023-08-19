// Created by alex at 2023/08/19 20:48
// leetgo: dev
// https://leetcode.cn/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func strToInt(str string) (ans int) {
	c := strings.Trim(str, " ")
	if len(c) == 0 {
		return
	}

	res, bndry := 0, math.MaxInt32/10
	i, sign := 1, 1
	if c[0] == '-' {
		sign = -1
	} else if c[0] != '+' {
		i = 0
	}
	for j := i; j < len(c); j++ {
		if c[j] < '0' || c[j] > '9' {
			break
		}
		if res > bndry || res == bndry && c[j] > '7' {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		res = res*10 + int(c[j]-'0')
	}
	return res * sign
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	str := Deserialize[string](ReadLine(stdin))
	ans := strToInt(str)

	fmt.Println("\noutput:", Serialize(ans))
}
