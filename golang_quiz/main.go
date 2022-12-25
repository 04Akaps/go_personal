package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func problemPuller(fileName string) ([]problem, error) {
	// read the problem

	// 1 파일을 연다
	if file, err := os.Open(fileName); err == nil {
		// 2 리더를 추가한다.
		csv := csv.NewReader(file)

		// 3. 파일을 읽는다.
		if cl, err := csv.ReadAll(); err == nil {
			// .4. 해결 함수를 호출한다.
			return parseProblem(cl), nil
		} else {
			return nil, fmt.Errorf("Error while ReadAll method")
		}
	} else {
		panic("Can't Open File")
		return nil, nil
	}

}

func parseProblem(line [][]string) []problem {
	r := make([]problem, len(line))

	for i := 0; i < len(line); i++ {
		r[i] = problem{q: line[i][0], a: line[i][1]}

	}

	return r
}

func main() {
	fName := flag.String("f", "quiz.csv", "path of csv file")

	timer := flag.Int("t", 30, "timer of the quiz")
	flag.Parse()

	problem, err := problemPuller(*fName)

	if err != nil {
		panic("Can't Read File")
	}

	correctAns := 0

	tObj := time.NewTimer(time.Duration(*timer) * time.Second)
	channel := make(chan string)

problemLoop:

	for i, p := range problem {
		var answer string
		fmt.Printf("Problem %d: %s=", i+1, p.q)

		go func() {
			fmt.Scanf("%s", &answer)
		}()

		select {
		case <-tObj.C:
			fmt.Println()
			break problemLoop

		case iAns := <-channel:
			if iAns == p.a {
				correctAns++
			}
			if i == len(problem)-1 {
				close(channel)
			}
		}
	}

	fmt.Printf("Youer correct Number is %d", correctAns)
}

type problem struct {
	q string
	a string
}
