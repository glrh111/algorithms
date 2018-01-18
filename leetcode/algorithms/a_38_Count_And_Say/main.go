package main

import "fmt"

/*
   Count And Say
   从 1 开始
 */
func countAndSay(n int) string {
	var (
		news, s = "", "1"
		count, lastc rune
	)
	for i:=2; i<=n; i++ {
		for _, c := range s {
			if c != lastc && count > 0 {
				news += fmt.Sprintf("%d%s", count, string(lastc))
				count = 0
			}
			count++
			lastc = c
		}
		if count > 0 {
			news += fmt.Sprintf("%d%s", count, string(lastc))
		}
		count = 0
		lastc = 0
		s = news
		fmt.Println(news)
		news = ""
	}
	return s
}

func main() {
	fmt.Println(countAndSay(50))
}
