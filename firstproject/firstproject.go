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
	var number int
	_, err := fmt.Scanln(&number)
	if err != nil {
		stdin.ReadString('\n')
	}
	return number, err
}

func main() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100)
	count := 0
	for {
		fmt.Printf("숫자를 입력하세요:")
		number, err := InputIntValue()
		count++
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
			continue
		}
		if number == target {
			fmt.Printf("숫자를 맞췄습니다. 축하합니다. 시도한 횟수: %d\n", count)
			break
		} else if number > target {
			fmt.Println("입력하신 숫자가 더 큽니다.")
			continue
		} else {
			fmt.Println("입력하신 숫자가 더 작습니다.")
			continue
		}
	}
	


}
