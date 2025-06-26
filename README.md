# PPY (PHP-Python) Compiler

A custom templating language compiler that combines HTML with embedded Python code, similar to PHP but using Python syntax.

## Overview

PPY is a templating language that allows you to embed Python code directly into HTML files. It uses a syntax similar to PHP but executes Python code instead. The compiler converts `.ppy` files into executable Python scripts that generate HTML output.

## Features

- **Embedded Python Code**: Use `<py? ... ?>` blocks to embed Python code
- **Variable Interpolation**: Use `<?= variable ?>` syntax for easy variable output
- **Loop Support**: Automatic indentation handling for Python loops and control structures
- **HTML Generation**: Converts templates to Python scripts that output HTML
- **Command Line Interface**: Simple CLI for compiling `.ppy` files

## Syntax

### Python Code Blocks
```html
<py?
# Python code goes here
for i in range(5):
    # Your Python logic
?>
```

### Variable Output
```html
<?= variable_name ?>
```

### Mixed Example
```html
<py?
for i in range(1, 6):
?>
    <p>Item number: <?= i ?></p>
<py?
# end loop
?>
```

## Installation

1. Ensure you have Go installed on your system
2. Clone or download this project
3. Build the compiler:
   ```bash
   go build -o ppy-compiler main.go
   ```

## Usage

### Basic Usage
```bash
# Compile and output to console
./ppy-compiler -i input.ppy

# Compile to a specific output file
./ppy-compiler -i input.ppy -o output.py

# Alternative syntax (positional argument)
./ppy-compiler input.ppy
```

### Running the Generated Python Code
After compilation, run the generated Python file:
```bash
python output.py
```

## Example

### Input File (`index.ppy`)
```html
<!DOCTYPE html>
<html>
<head>
    <title>Multiplication Table</title>
</head>
<body>
    <h1>Multiplication Table</h1>
    <table>
        <tr>
            <th>×</th>
            <py?
            for i in range(1, 11):
            ?>
                <th><?= i ?></th>
            <py?
            # end headers loop
            ?>
        </tr>
        <py?
        for row in range(1, 11):
        ?>
            <tr>
                <th><?= row ?></th>
                <py?
                for col in range(1, 11):
                    product = row * col
                ?>
                    <td><?= product ?></td>
                <py?
                # end column loop
                ?>
            </tr>
        <py?
        # end row loop
        ?>
    </table>
</body>
</html>
```

### Generated Python Code
The compiler converts the above into Python print statements with proper indentation handling for loops and variable interpolation.

## How It Works

1. **Parsing**: The compiler uses regular expressions to identify Python code blocks (`<py? ... ?>`) and variable expressions (`<?= ... ?>`)
2. **Code Generation**: HTML content is converted to Python print statements
3. **Indentation Handling**: The compiler automatically manages Python indentation for loops and control structures
4. **Variable Interpolation**: Variable expressions are converted to string concatenation in print statements

## Command Line Options

- `-i <file>`: Input PPY file path (required)
- `-o <file>`: Output Python file path (optional, defaults to console output)

## Error Handling

The compiler provides clear error messages for:
- Missing input files
- File read/write errors
- Invalid command line arguments

## Contributing

This is a custom templating language project. Feel free to:
- Report bugs
- Suggest improvements
- Add new features
- Improve documentation

## License

MIT License

Copyright (c) 2025 sanjaysagar12

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

## Author

Created by sanjaysagar12

---

*PPY makes it easy to generate dynamic HTML using Python's powerful syntax and libraries while maintaining the familiar feel of server-side templating languages.*
