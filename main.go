package main

import "fmt"

/*
Reverse Polish Notation (RPN) is a mathematical notation in which every operator follows all of its
operands. It does not need any parentheses as long as both the operator and operand are fixed.
An RPN calculator processes expressions from left to right, pushing numbers onto a stack, and
applying operations to operands by popping them from the stack.

Your task is to implement an RPN calculator that supports basic arithmetic operations:
- addition
- subtraction,
- multiplication
- division

The calculator should be handling errors gracefully.

Input Format

A single line containing the RPN expression. The expression will consist of integers and arithmetic
operators (+, -, *, /), and spaces as separators. The input will be a string where tokens (numbers
and operators) are separated by spaces, for example, "3 4 +", "10 2 /".

If the expression is valid and can be successfully evaluated, print the result of the expression.
If the expression is invalid or an error occurs during evaluation (e.g., division by zero,
insufficient operands), print an appropriate error message. Only the final result or error message
should be output, not the intermediate steps of the calculation.

Sample Inputs and Expected Outputs
Valid Inputs

Input: "3 4 +"
Expected Output: Result: 7

Input: "5 1 2 + 4 * + 3 -"
Expected Output: Result: 14

Invalid Inputs

Input: "10 0 /"
Expected Output: Error: division by zero

Input: "+"
Expected Output: Error: not enough elements in the stack for operation

Notes for the Candidate:

- You may assume that all numerical inputs fit within the range of a 64-bit signed integer.
- Focus on writing clean, efficient code, and handle edge cases and potential errors thoughtfully.

Other examples for testing:

 validExpressions := []string{
"3 4 +",             // Addition, expected result: 7
"10 4 -",            // Subtraction, expected result: 6
"6 7 *",             // Multiplication, expected result: 42
"8 2 /",             // Division, expected result: 4

"5 1 2 + 4 * + 3 -", // Mixed operations, expected result: 14


// When find symbol, calculate and add value to the stack
5 (1 2 +) 4 * + 3 -
5 3 4 * + 3 -
5 12 + 3 -
(5 + 12) 3 -

3 (2 1 +) -
3 3 -
0

// TODO
"9 sqrt",            // Square root, expected result: 3
"9 sqrt 3 /",        // Unary followed by binary operation, expected result: 1
}

invalidExpressions := []string{
"+",       // Error: not enough elements for operation
"3 -",     // Error: not enough elements for operation
"sqrt",    // Error: not enough elements for operation
"10 0 /",  // Error: division by zero
"-4 sqrt", // Error: square root of negative number
"3 3",     // Error: the user input does not form a valid RPN expression
"abc",     // Error: unknown token
"4 5 &",   // Error: unknown token/operator
}
*/

func main() {
	fmt.Println("Hello, world!")
}
