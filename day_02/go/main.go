package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 02")
	solve()
}

// Position of the submarine
type Position struct {
	Horizontal int
	Depth      int
	Aim        int
}

// Command to move the submarine
type Command struct {
	Cmd   string
	Value int
}

func (p *Position) MovePartOne(c *Command) {
	switch c.Cmd {
	case "forward":
		p.Horizontal += c.Value
	case "down":
		p.Depth += c.Value
	case "up":
		p.Depth -= c.Value
	}
}

func (p *Position) MovePartTwo(c *Command) {
	switch c.Cmd {
	case "forward":
		p.Horizontal += c.Value
		p.Depth += p.Aim * c.Value
	case "down":
		p.Aim += c.Value
	case "up":
		p.Aim -= c.Value
	}
}

func solve() {
	commands := loadInput("../input_1.txt")
	partOnePos := &Position{}
	partTwoPos := &Position{}
	for _, command := range commands {
		partOnePos.MovePartOne(command)
		partTwoPos.MovePartTwo(command)
	}

	fmt.Println("Part One Ans: ", partOnePos.Horizontal*partOnePos.Depth)
	fmt.Println("Part Two Ans: ", partTwoPos.Horizontal*partTwoPos.Depth)

}

func loadInput(path string) []*Command {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cmds []*Command
	for scanner.Scan() {
		value := scanner.Text()
		split := strings.Split(value, " ")
		cmd, value := split[0], split[1]
		val, _ := strconv.Atoi(value)
		cmds = append(cmds, &Command{Cmd: cmd, Value: val})
	}
	return cmds
}
