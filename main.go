package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Interpreter struct {
	variables map[string]int
}

// NewInterpreter initializes an interpreter
func NewInterpreter() *Interpreter {
	return &Interpreter{variables: make(map[string]int)}
}

// Execute parses and executes a single line of code
func (i *Interpreter) Execute(line string) {
	tokens := strings.Fields(line)

	if len(tokens) == 0 {
		return
	}

	switch tokens[0] {
	case "kemon":
		if len(tokens) < 3 || tokens[1] != "achis" {
			fmt.Println("bhul hoye gelo vai check kor ekbar")
			return
		}
		arg := strings.Join(tokens[2:], " ")
		if strings.HasPrefix(arg, "\"") && strings.HasSuffix(arg, "\"") {
			message := strings.Trim(arg, "\"")
			fmt.Println(message)
		} else {
			value, exists := i.variables[arg]
			if !exists {
				fmt.Printf("bhul hoye gelo vai check kor ekbar: Variable '%s' not defined\n", arg)
				return
			}
			fmt.Println(value)
		}
		fmt.Println("weee ko peyechis ebar porte bos")
	case "bol":
		if len(tokens) != 3 || tokens[1] != "bhai" {
			fmt.Println("bhul hoye gelo vai check kor ekbar")
			return
		}
		varName := tokens[2]
		fmt.Printf("Enter value for %s: ", varName)
		var input string
		fmt.Scanln(&input)
		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("bhul hoye gelo vai check kor ekbar: invalid number")
			return
		}
		i.variables[varName] = value
		// fmt.Println("weee ko peyechis ebar porte bos")
	case "dyakh":
		if len(tokens) < 4 || tokens[1] != "jodi" || !strings.Contains(line, "ar nahole:") {
			fmt.Println("bhul hoye gelo vai check kor ekbar: Invalid conditional syntax")
			return
		}
		parts := strings.Split(line, "ar nahole:")
		if len(parts) != 2 {
			fmt.Println("bhul hoye gelo vai check kor ekbar: Missing 'ar nahole'")
			return
		}
		ifPart := strings.TrimSpace(parts[0])
		elsePart := strings.TrimSpace(parts[1])
		ifTokens := strings.SplitN(ifPart, ":", 2)
		if len(ifTokens) != 2 {
			fmt.Println("bhul hoye gelo vai check kor ekbar: Missing ':' in 'dyakh jodi'")
			return
		}
		condition := strings.TrimSpace(strings.Join(strings.Fields(ifTokens[0])[2:], " "))
		ifCommand := strings.TrimSpace(ifTokens[1])
		condValue, err := i.evaluateExpression(strings.Fields(condition))
		if err != nil {
			fmt.Println("bhul hoye gelo vai check kor ekbar: Error evaluating condition:", err)
			return
		}
		if condValue != 0 {
			i.Execute(ifCommand)
		} else {
			i.Execute(elsePart)
		}
	default:
		if len(tokens) < 3 || tokens[1] != "=" {
			fmt.Println("bhul hoye gelo vai check kor ekbar")
			return
		}
		varName := tokens[0]
		expression := tokens[2:]
		value, err := i.evaluateExpression(expression)
		if err != nil {
			fmt.Println("bhul hoye gelo vai check kor ekbar:", err)
			return
		}
		i.variables[varName] = value
		fmt.Println("weee ko peyechis ebar porte bos")
	}
}

func (i *Interpreter) evaluateExpression(tokens []string) (int, error) {
	if len(tokens) == 1 {
		return i.getValue(tokens[0])
	}
	if len(tokens) == 3 {
		left, err := i.getValue(tokens[0])
		if err != nil {
			return 0, err
		}
		right, err := i.getValue(tokens[2])
		if err != nil {
			return 0, err
		}
		switch tokens[1] {
		case "+":
			return left + right, nil
		case "-":
			return left - right, nil
		case "*":
			return left * right, nil
		case "/":
			if right == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			return left / right, nil
		default:
			return 0, fmt.Errorf("unknown operator '%s'", tokens[1])
		}
	}
	return 0, fmt.Errorf("invalid expression")
}

func (i *Interpreter) getValue(token string) (int, error) {
	if value, err := strconv.Atoi(token); err == nil {
		return value, nil
	}
	if value, exists := i.variables[token]; exists {
		return value, nil
	}
	return 0, fmt.Errorf("unknown variable '%s'", token)
}

func main() {
	version := flag.Bool("v", false, "Print the version")
	help := flag.Bool("h", false, "Display help")
	flag.Parse()

	if *version {
		fmt.Println("Gola Compiler v1.0")
		return
	}
	if *help {
		fmt.Println("Usage: gola [options] <file>")
		fmt.Println("Options:")
		fmt.Println("  -v, --version     Print the version of the compiler")
		fmt.Println("  -h, --help        Display help information")
		return
	}

	args := flag.Args()
	if len(args) > 0 {
		fileName := args[0]
		if !strings.HasSuffix(fileName, ".gola") {
			fmt.Println("Error: Only .gola files are supported")
			return
		}
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("Error: Unable to open file '%s'\n", fileName)
			return
		}
		defer file.Close()
		interpreter := NewInterpreter()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			interpreter.Execute(strings.TrimSpace(line))
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading file '%s': %v\n", fileName, err)
		}
		fmt.Println("Jay Shree Ram Bhai!")
		return
	}

	fmt.Println("Welcome to MiniLang - Bengali Edition!")
	fmt.Println("Type 'exit' to quit.")
	scanner := bufio.NewScanner(os.Stdin)
	interpreter := NewInterpreter()
	for {
		fmt.Print(">> ")
		scanner.Scan()
		line := scanner.Text()
		if strings.TrimSpace(line) == "exit" {
			break
		}
		interpreter.Execute(line)
	}
	fmt.Println("Jay Shree Ram Bhai!")
}
