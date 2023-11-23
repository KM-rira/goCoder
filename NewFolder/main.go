package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// https://atcoder.jp/contests/abc261/tasks/abc261_c
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	// N を取得
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	var slice []string
	pattern := `\([^)]*\)`
	re := regexp.MustCompile(pattern)
	for i := 0; i < n; i++ {
		sc.Scan()
		input := sc.Text()
		// 配列を逆順にループ
		length := len(slice)
		for i := length - 1; i >= 0; i-- {
			s := slice[i]
			reS := re.ReplaceAllString(s, "")
			fmt.Println("reS: ", reS)
			if reS == input {
				num, _ := extractNumber(s)
				addNum := strconv.Itoa(num + 1)
				input = input + "(" + addNum + ")"
				break
			}
		}
		slice = append(slice, input)
	}
	for _, a := range slice {
		println(a)
	}
}

func extractNumber(input string) (int, error) {
	// 括弧内の数字にマッチする正規表現パターン
	pattern := `\((\d+)\)`
	re := regexp.MustCompile(pattern)

	// 正規表現にマッチする部分を探す
	match := re.FindStringSubmatch(input)

	if len(match) > 1 {
		// グループ1（括弧内の数字）を取得し、intに変換
		number, err := strconv.Atoi(match[1])
		if err != nil {
			return 0, err
		}
		return number, nil
	}

	return 0, fmt.Errorf("数字が見つかりませんでした")
}
