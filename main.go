package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var m int
	_, err := fmt.Scanln(&m)
	if err != nil {
		stdin.ReadString('\n')
	}
	return m, err
}

func main() {
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(100)
	var cnt int = 0
	for {
		fmt.Printf("숫자값을 입력하세요 : ")
		cnt++
		m, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else if n > m {
			fmt.Println("입력하신 숫자보다 큽니다.")
		} else if n < m {
			fmt.Println("입력하신 숫자보다 작습니다.")
		} else {
			fmt.Println(n, "정답입니다.")
			fmt.Println("시도한 횟수는 ", cnt)
			break
		}
	}
}
