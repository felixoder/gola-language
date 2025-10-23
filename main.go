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
	variables map[string]interface{}
}

// NewInterpreter initializes an interpreter
func NewInterpreter() *Interpreter {
	return &Interpreter{variables: make(map[string]interface{})}
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

		// Check if the argument is a quoted string
		if strings.HasPrefix(arg, "\"") && strings.HasSuffix(arg, "\"") {
			fmt.Println(arg[1 : len(arg)-1]) // Remove the quotes for display
			return
		}

		// If not a string, evaluate the expression
		value, err := i.evaluateExpression(strings.Fields(arg))
		if err == nil {
			fmt.Println(value)
		} else {
			// Check if it's a variable
			if value, ok := i.variables[arg]; ok {
				fmt.Println(value)
			} else {
				fmt.Println("bhul hoye gelo vai check kor ekbar")
			}
		}

	case "bol":
		if len(tokens) != 3 || tokens[1] != "bhai" {
			fmt.Println("bhul hoye gelo vai check kor ekbar")
			return
		}
		varName := tokens[2]
		fmt.Printf(`Value bolo %s er =>: `, varName)
		var input string
		fmt.Scanln(&input)

		// Handle string input properly (strip the quotes if present)
		if strings.HasPrefix(input, "\"") && strings.HasSuffix(input, "\"") {
			i.variables[varName] = input[1 : len(input)-1] // Store the string without quotes
		} else if value, err := strconv.Atoi(input); err == nil {
			i.variables[varName] = value // Store the integer value
		} else {
			fmt.Println("bhul hoye gelo vai check kor ekbar: invalid value")
			return
		}

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

		condValue, err := i.evaluateCondition(condition)
		if err != nil {
			fmt.Println("bhul hoye gelo vai check kor ekbar: Error evaluating condition")
			return
		}

		if condValue {
			i.Execute(ifCommand)
		} else {
			if strings.HasPrefix(elsePart, "\"") && strings.HasSuffix(elsePart, "\"") {
				fmt.Println(elsePart[1 : len(elsePart)-1]) // Remove quotes for printing
			} else {
				i.Execute(elsePart)
			}
		}

	default:
		fmt.Println("bhul hoye gelo vai check kor ekbar: Unknown command")
	}
}

func (i *Interpreter) evaluateExpression(tokens []string) (int, error) {
	if len(tokens) == 1 {
		if value, ok := i.variables[tokens[0]]; ok {
			switch v := value.(type) {
			case int:
				return v, nil
			case string:
				return 0, fmt.Errorf("expected an integer, but got a string")
			default:
				return 0, fmt.Errorf("invalid type for evaluation")
			}
		}
		return strconv.Atoi(tokens[0])
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

func (i *Interpreter) evaluateCondition(condition string) (bool, error) {
	tokens := strings.Fields(condition)
	if len(tokens) != 3 {
		return false, fmt.Errorf("invalid condition syntax")
	}

	left, err := i.getValue(tokens[0])
	if err != nil {
		return false, err
	}
	right, err := i.getValue(tokens[2])
	if err != nil {
		return false, err
	}

	switch tokens[1] {
	case ">":
		return left > right, nil
	case "<":
		return left < right, nil
	case ">=":
		return left >= right, nil
	case "<=":
		return left <= right, nil
	case "==":
		return left == right, nil
	case "!=":
		return left != right, nil
	default:
		return false, fmt.Errorf("unknown operator '%s'", tokens[1])
	}
}

func (i *Interpreter) getValue(token string) (int, error) {
	if value, err := strconv.Atoi(token); err == nil {
		return value, nil
	}
	if value, exists := i.variables[token]; exists {
		switch v := value.(type) {
		case int:
			return v, nil
		case string:
			return 0, fmt.Errorf("expected an integer, but got a string")
		default:
			return 0, fmt.Errorf("unknown variable '%s'", token)
		}
	}
	return 0, fmt.Errorf("unknown variable '%s'", token)
}

func main() {
	version := flag.Bool("v", false, "Print the version")
	help := flag.Bool("h", false, "Display help")
	flag.Parse()

	if *version {
		fmt.Println("Gola Compiler v1.0.2")
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
	fmt.Println(`Welcome to Gola - Bengali Edition!

     ::::::::        ::::::::       :::            :::
    :+:    :+:     :+:    :+:      :+:          :+: :+:
   +:+            +:+    +:+      +:+         +:+   +:+
  :#:            +#+    +:+      +#+        +#++:++#++:
 +#+   +#+#     +#+    +#+      +#+        +#+     +#+
#+#    #+#     #+#    #+#      #+#        #+#     #+#
########       ########       ########## ###     ###

				`)
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
