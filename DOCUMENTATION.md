# Gola Language Interpreter Documentation

Gola is a lightweight and fun programming language inspired by simplicity and Bengali expressions. This documentation provides an overview of its syntax, usage, and functionality.

---

## Features
- Lightweight and interactive interpreter.
- Uses Bengali-inspired keywords and syntax.
- Supports basic arithmetic, variable assignment, and conditional execution.
- Includes interactive and file-based execution modes.

---

## Installation

### Prerequisites
- Go programming language installed.

### Build
Run the following commands to build the interpreter for your platform:

```bash
# Build for Linux
GOOS=linux GOARCH=amd64 go build -o gola-linux-x86_64

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o gola-darwin-x86_64
```

---

## Usage

### Command-line Options
- `-v` or `--version`: Display the version of the interpreter.
- `-h` or `--help`: Display help information.

Example:
```bash
gola -v
```

### Running Gola Code
You can execute `.gola` files or use the interactive mode.

#### File-based Execution
```bash
gola <file_name>.gola
```
Ensure the file has a `.gola` extension. For example:
```bash
gola sample.gola
```

#### Interactive Mode
Run the interpreter without arguments:
```bash
gola
```
Type commands interactively and press Enter to execute.

---

## Syntax

### Print Statements
- **Command**: `kemon achis "<message>"`
- **Description**: Prints the specified message to the console.

Example:
```gola
kemon achis "Hello, world!"
```

### Variable Input
- **Command**: `bol bhai <variable>`
- **Description**: Prompts the user to input a value for the specified variable.

Example:
```gola
bol bhai x
```

### Variable Assignment
- **Command**: `<variable> = <expression>`
- **Description**: Assigns the result of an expression to a variable.

Example:
```gola
x = 10 + 5
```

### Conditional Execution
- **Command**: `dyakh jodi <condition>: <command> ar nahole: <command>`
- **Description**: Executes the first command if the condition is true; otherwise, executes the second command.

Example:
```gola
dyakh jodi x > 5: kemon achis "Greater" ar nahole: kemon achis "Smaller or Equal"
```

---

## Error Handling
- **Error Messages**: The interpreter provides Bengali-inspired error messages for invalid commands or conditions, such as:
  - `bhul hoye gelo vai check kor ekbar` (Check your syntax.)

---

## Implementation Overview

### Key Components

#### `Interpreter`
The core interpreter structure contains:
- `variables`: A map to store variable values.

#### Methods
- `NewInterpreter()`: Initializes the interpreter.
- `Execute(line string)`: Parses and executes a single line of code.
- `evaluateExpression(tokens []string)`: Evaluates arithmetic expressions.
- `getValue(token string)`: Retrieves a variable's value or parses integers.

### Commands
Commands are executed based on the first token in a line:
1. **`kemon achis`**: Print messages or variables.
2. **`bol bhai`**: Accept user input for a variable.
3. **`dyakh jodi`**: Execute conditional logic.
4. **Assignment (`=`)**: Assigns values to variables.

---

## Sample Code
```gola
bol bhai x
y = x + 5
dyakh jodi y > 10: kemon achis "Value is greater" ar nahole: kemon achis "Value is smaller"
```

Expected Output:
```
Enter value for x: 6
Value is greater
```

---

## Interactive Mode Example
```
Welcome to Gola - Bengali Edition!
Type 'exit' to quit.
>> bol bhai x
Enter value for x: 10
>> kemon achis "x is "
>> kemon achis x
10
>> exit
Jay Shree Ram Bhai!
```

---

## New in Gola

### Variable Assignment Fix
Assignments now support expressions on the right-hand side, correctly updating the variable:
- `x = x + 5`

### Loop Statement (`bolte thak`)
A basic file-mode loop is introduced. The loop header is followed by a single-line body on the next line:
- **Syntax:** `bolte thak <iterator> [<start>] -> [<end>] [<increment>] :`
- **Example:**
  ```gola
  bolte thak i [0] -> [5] [i++] :
  kemon achis i

## Contributing
Feel free to contribute to this project by submitting issues or pull requests to the GitHub repository.

---

## License
This project is licensed under the MIT License.
