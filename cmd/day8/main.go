package main

import (
	"aod2020/pkg/config"
	"bufio"
	"io"
	"strconv"
	"strings"
)

func main() {
	part, in := config.ParseDefaultFlags()
	program := parseProgram(in)
	if part == 1 {
		result, _ := runProgram(program, 0)
		println("The result of the program is:", result)
	} else {
		result := testProgramVariants(program)
		println("The result of the fixed program is:", result)
	}
}

type Program = []Operation
type Operation interface {
	Execute(state State) (movement int, newState State)
}
type State int

type Nop int
func (n Nop) Execute(state State) (int, State) {
	return 1, state
}

type Acc int
func (a Acc) Execute(state State) (movement int, newState State) {
	return 1, state + State(a)
}

type Jmp int

func (j Jmp) Execute(state State) (movement int, newState State) {
	return int(j), state
}

func parseProgram(in io.Reader) Program {
	s := bufio.NewScanner(in)
	var program Program
	for s.Scan() {
		fields := strings.Fields(s.Text())
		i, err := strconv.ParseInt(fields[1], 10, 64)
		if err != nil { panic(err.Error()) }
		switch fields[0] {
		case "nop":
			program = append(program, Nop(i))
		case "acc":
			program = append(program, Acc(i))
		case "jmp":
			program = append(program, Jmp(i))
		}
	}
	return program
}

func runProgram(program Program, state State) (State, bool) {
	marker := 0
	visitedMarks := make(map[int]bool)
	for !visitedMarks[marker] && marker < len(program) {
		visitedMarks[marker] = true
		var mov int
		mov, state = program[marker].Execute(state)
		marker += mov
	}
	return state, marker >= len(program)
}

func testProgramVariants(program Program) State {
	for i, op := range program {
		switch op.(type) {
		case Nop:
			program[i] = Jmp(op.(Nop))
		case Jmp:
			program[i] = Nop(op.(Jmp))
		}
		result, ok := runProgram(program, 0)
		if ok {
			return result
		}
		program[i] = op // cleaning!
	}
	return 0
}
