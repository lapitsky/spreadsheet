# Spreadsheet Programming Task Description
Write a program to parse a given CSV file and evaluate each cell by these rules
1. Each cell is an expression in ​postfix notation​. Please refer to the wikipedia page for a full description.
2. Each number or operation will always be separated by one ​or more​ spaces.
3. A cell can refer to another cell, via the LETTER NUMBER notation (A2, B4, etc - letters
refer to columns, numbers to rows). You can assume it’s a single letter followed by a
number to make parsing a bit easier.
4. Support the basic arithmetic operators +, -, *, /
The output will be a CSV file of the same dimensions, where each cell is evaluated to its final value. If any cell is an invalid expression, then ​for that cell only​ print #ERR.
For example, the following CSV input:
```csv
b1 b2 +,2 b2 3 * -,3 ,+
a1,5 ,,72/
c23*,12 , ,512+4*+3-
```
Might output something like this:
```csv
-8,-13,3,#ERR
-8,5,0,3.5
0,#ERR,0,14
```
The program should take its input from a command line file argument, and should print its output to ​STDOUT
# Requirements
* Do not over-engineer your solution, you should aim to finish the task in ~4 hours so don’t make it too complex. We are looking for simple solutions and clearly written code.
* The timeframe is just a guide to help you plan. You are not being timed, there is no bonus point for finishing quickly. Do not rush, do not write spaghetti code.
* Use only what is available in the standard library. This problem is simple enough that you shouldn’t need 3rd party libraries.
* Specific details about the behaviour of the application are left for you to decide how best to handle. Limitations in your implementation are fine but please document them.
* ​NOTE: ​The above test data is just an example. As part of this task, you are required to thoroughly test your application yourself.
Deliverables
* Your source code and test input data
* A ​short​ report (1⁄2 page max) containing:
    * Brief description of the code structure
    * Any limitations of your implementation (or other things you’d like to point out)
    * Any trade-offs or design decisions you made that are worth noting (optional).