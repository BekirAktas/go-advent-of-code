package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type partTwoAnswer struct {
	greatestOne    int
	secondGreatest int
	thirdGreatest  int
}

func main() {
	fmt.Printf("part one answer is: %d\n", partOne())
	topThreeElves := partTwo()
	fmt.Println(topThreeElves)
	fmt.Println(
		topThreeElves.greatestOne + topThreeElves.secondGreatest + topThreeElves.thirdGreatest)
}

func partOne() int {
	file, err := os.Open("adventofcode.com_2022_day_1_input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	oneElvesTotalValue := 0
	greatestElvesTotalValue := 0
	for scanner.Scan() {
		row := scanner.Text()
		value, err := strconv.Atoi(row)

		if err != nil {
			if oneElvesTotalValue > greatestElvesTotalValue {
				greatestElvesTotalValue = oneElvesTotalValue
			}
			oneElvesTotalValue = 0
			continue
		}

		oneElvesTotalValue = value + oneElvesTotalValue
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return greatestElvesTotalValue
}

func partTwo() partTwoAnswer {
	file, err := os.Open("adventofcode.com_2022_day_1_input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	oneElvesTotalValue := 0
	greatestElvesTotalValue := 0
	secondGreatestElvesTotalValue := 0
	thirdGreatestElvesTotalValue := 0

	for scanner.Scan() {
		row := scanner.Text()
		value, err := strconv.Atoi(row)

		if err != nil {

			if oneElvesTotalValue > thirdGreatestElvesTotalValue {
				if oneElvesTotalValue < secondGreatestElvesTotalValue {
					thirdGreatestElvesTotalValue = oneElvesTotalValue
				} else if oneElvesTotalValue > secondGreatestElvesTotalValue {
					thirdGreatestElvesTotalValue = secondGreatestElvesTotalValue
					secondGreatestElvesTotalValue = oneElvesTotalValue
					if oneElvesTotalValue > greatestElvesTotalValue {
						secondGreatestElvesTotalValue = greatestElvesTotalValue
						greatestElvesTotalValue = oneElvesTotalValue
					}
				}
			}
			oneElvesTotalValue = 0
			continue
		}

		oneElvesTotalValue = value + oneElvesTotalValue
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	answer := partTwoAnswer{
		greatestOne:    greatestElvesTotalValue,
		secondGreatest: secondGreatestElvesTotalValue,
		thirdGreatest:  thirdGreatestElvesTotalValue,
	}
	return answer
}
