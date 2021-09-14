package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	balance := 1000
	rand.Seed(time.Now().UnixNano())
	var n int
	for {
		fmt.Println("1~5사이의 값을 입력하세요")
		random := rand.Intn(5) + 1
		fmt.Scanln(&n)

		if n == random {
			balance = balance + 500
			fmt.Println("축하합니다 잔액은 ", balance)
			if 5000 <= balance || balance <= 0 {
				fmt.Println("게임을 종료합니다.")
				break
			}
		} else if n != random {
			balance = balance - 100
			fmt.Println("아쉽습니다 잔액은", balance)
			if 5000 <= balance || balance <= 0 {
				fmt.Println("게임을 종료합니다.")
				break
			}
		} else {
			break
		}

	}
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"time"
// )

// const (
// 	//초기 잔액
// 	Balance = 1000
// 	//맞췄을 때 버는 양
// 	EarnPoint = 500
// 	//틀렸을 때 잃는 양
// 	LosePoint = 100
// 	//게임 승리 포인트
// 	VictoryPoint = 5000
// 	//게임 오버 포인트
// 	GameoverPoint = 0
// )

// var stdin = bufio.NewReader(os.Stdin)

// func InputIntValue() (int, error) {
// 	var n int
// 	_, err := fmt.Scanln(&n)
// 	if err != nil {
// 		stdin.ReadString('\n')
// 	}
// 	return n, err
// }
// func main() {
// 	rand.Seed(time.Now().UnixNano())
// 	balance := Balance
// 	for {
// 		fmt.Println("1~5 사이의 값을 입력하세요")
// 		n, err := InputIntValue()
// 		if err != nil {
// 			fmt.Println("숫자만 입력하세요")
// 		} else {
// 			r := rand.Intn(5) + 1
// 			if n == r {
// 				balance += EarnPoint
// 				fmt.Println("축하합니다. 맞추셨습니다. 남은 돈 :", balance)
// 				if balance >= VictoryPoint {
// 					fmt.Println("게임 승리")
// 					break
// 				}
// 			} else {
// 				balance -= LosePoint
// 				fmt.Println("꽝 아쉽지만 다음 기회를.. 남은 돈 :", balance)
// 				if balance <= GameoverPoint {
// 					fmt.Println("게임 오버")
// 					break
// 				}
// 			}
// 		}
// 	}
// }
