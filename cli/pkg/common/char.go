/*
 * @Date: 2022-05-23 11:01:31
 * @LastEditors: Jason Chen
 * @LastEditTime: 2022-09-09 16:26:55
 * @FilePath: /cli/pkg/common/char.go
 */

package common

import (
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
	"unsafe"
)

const (
	chars    = "0123456789abcdefghijklmnopqrstuvwxyz"
	charsLen = len(chars)
	mask     = 1<<6 - 1
)

var rng = rand.NewSource(time.Now().UnixNano())

// RandLowStr 返回指定长度的随机字符串
func RandLowStr(ln int) string {
	/* chars 38个字符
	 * rng.Int63() 每次产出64bit的随机数,每次我们使用6bit(2^6=64) 可以使用10次
	 */
	buf := make([]byte, ln)
	for idx, cache, remain := ln-1, rng.Int63(), 10; idx >= 0; {
		if remain == 0 {
			cache, remain = rng.Int63(), 10
		}
		buf[idx] = chars[int(cache&mask)%charsLen]
		cache >>= 6
		remain--
		idx--
	}
	return *(*string)(unsafe.Pointer(&buf))
}

// 删除连续的空行
func RemoveWhiteLines(filePath string) {
	bytes, _ := os.ReadFile(filePath)
	newLine := GetNewline()
	content := regexp.MustCompile(`\s+(\n|\r\n){2,}`).ReplaceAllString(strings.TrimSpace(string(bytes)), newLine)
	content += newLine
	os.WriteFile(filePath, []byte(content), 0700)
}
