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
		// Ensure that the syntax is correct
		if len(tokens) < 3 || tokens[1] != "achis" {
			fmt.Println("bhul hoye gelo vai check kor ekbar")
			return
		}

		// Get the argument after 'kemon achis'
		arg := strings.Join(tokens[2:], " ")

		// Case 1: If the argument is a valid expression (e.g., a + b)
		// We will split the argument based on spaces and evaluate it as an expression
		if strings.Contains(arg, "+") || strings.Contains(arg, "-") || strings.Contains(arg, "*") || strings.Contains(arg, "/") {
			tokens := splitExpression(arg)
			result, err := i.evaluateExpression(tokens)
			if err != nil {
				fmt.Println("bhul hoye gelo vai check kor ekbar")
			} else {
				fmt.Println(result)
			}
		} else if strings.HasPrefix(arg, `"`) && strings.HasSuffix(arg, `"`) {
			// Case 2: Argument is a string literal (enclosed in double quotes)
			// Remove the surrounding quotes and print the value
			fmt.Println(arg[1 : len(arg)-1])
		} else if _, err := strconv.Atoi(arg); err == nil {
			// Case 3: Argument is a number (valid as an integer)
			fmt.Println(arg)
		} else {
			// Case 4: Invalid expression or unknown token
			fmt.Println("bhul hoye gelo vai check kor ekbar")
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

	case "dyakh":
		// Syntax: dyakh jodi <condition>: <command> ar nahole: <command>
		if len(tokens) < 4 || tokens[1] != "jodi" || !strings.Contains(line, "ar nahole:") {
			fmt.Println("bhul hoye gelo vai check kor ekbar: Invalid conditional syntax")
			return
		}

		// Split into condition and commands
		parts := strings.Split(line, "ar nahole:")
		if len(parts) != 2 {
			fmt.Println("bhul hoye gelo vai check kor ekbar: Missing 'ar nahole'")
			return
		}

		ifPart := strings.TrimSpace(parts[0])   // "dyakh jodi <condition>: <command>"
		elsePart := strings.TrimSpace(parts[1]) // "<command>"

		// Extract the condition and command for the "if" part
		ifTokens := strings.SplitN(ifPart, ":", 2)
		if len(ifTokens) != 2 {
			fmt.Println("bhul hoye gelo vai check kor ekbar: Missing ':' in 'dyakh jodi'")
			return
		}

		condition := strings.TrimSpace(strings.Join(strings.Fields(ifTokens[0])[2:], " ")) // Extract the condition
		ifCommand := strings.TrimSpace(ifTokens[1])                                        // Command after ':'

		// Evaluate the condition
		condValue, err := i.evaluateExpression(strings.Fields(condition))
		if err != nil {
			fmt.Println("bhul hoye gelo vai check kor ekbar: Error evaluating condition:", err)
			return
		}

		// Execute the appropriate command
		if condValue != 0 {
			i.Execute(ifCommand) // Execute the command after 'dyakh jodi'
		} else {
			// If the elsePart is a string surrounded by quotes, remove them
			if strings.HasPrefix(elsePart, `"`) && strings.HasSuffix(elsePart, `"`) {
				// Remove the surrounding quotes and print the result
				fmt.Println(elsePart[1 : len(elsePart)-1])
			} else {
				// Otherwise, just execute the command in elsePart
				i.Execute(elsePart)
			}
		}

	case "ar":
		// Syntax: ar nahole
		if len(tokens) != 2 || tokens[1] != "nahole" {
			fmt.Println("bhul hoye gelo vai check kor ekbar: Invalid else syntax")
			return
		}

		// Execute the next block
		// Now it's just printing the raw string contents without quotes
		fmt.Println("ar nahole: Execute the else block")

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

func splitExpression(expression string) []string {
	// Split based on operators (+, -, *, /)
	var tokens []string
	var currentToken string
	for _, ch := range expression {
		if ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '>' || ch == '<' || ch == '=' || ch == '|' || ch == '&' {
			if len(currentToken) > 0 {
				tokens = append(tokens, currentToken)
			}
			tokens = append(tokens, string(ch))
			currentToken = ""
		} else if ch != ' ' {
			currentToken += string(ch)
		}
	}
	if len(currentToken) > 0 {
		tokens = append(tokens, currentToken)
	}
	return tokens
}

func (i *Interpreter) evaluateExpression(tokens []string) (int, error) {
	// If there's only one token, return its value
	if len(tokens) == 1 {
		// Handle the case where the token is a string
		if tokens[0][0] == '"' && tokens[0][len(tokens[0])-1] == '"' {
			return 1, nil // Consider string as "true" for conditional purposes
		}
		return i.getValue(tokens[0])
	}

	// Handle binary operations like a + b
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
		case ">":
			if left > right {
				return 1, nil
			}
			return 0, nil
		case "<":
			if left < right {
				return 1, nil
			}
			return 0, nil
		case "==":
			if left == right {
				return 1, nil
			}
			return 0, nil
		case "||":
			if left != 0 || right != 0 {
				return 1, nil
			}
			return 0, nil
		case "&&":
			if left != 0 && right != 0 {
				return 1, nil
			}
			return 0, nil
		default:
			return 0, fmt.Errorf("unknown operator '%s'", tokens[1])
		}
	}

	return 0, fmt.Errorf("invalid expression")
}

func (i *Interpreter) getValue(token string) (int, error) {
	// Check if token is a number
	if value, err := strconv.Atoi(token); err == nil {
		return value, nil
	}
	// Check if token is a variable
	if value, exists := i.variables[token]; exists {
		return value, nil
	}
	// Return error if token is unknown
	return 0, fmt.Errorf("unknown variable '%s'", token)
}

func main() {
	version := flag.Bool("v", false, "Print the version")
	help := flag.Bool("h", false, "Display help")
	flag.Parse()

	if *version {
		fmt.Println("Gola Compiler v1.0.1")
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
