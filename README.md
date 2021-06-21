Tic-tac-toe CLI written in Go
=======================
## Description 
Simple tic tac toe game, utilizing the negamax algorithm for the opponent/AI player.
##### Key features:
Doesn't rely on greedy algorithms, and can support bigger boards than 3x3.
Uses a negamax algorithm with alpha-beta pruning.
Has a Plugable Player interface.

## Installation
To use clone the repo, and build the binary for your platform:
```
git clone https://github.com/buni/ttt-task.git
make build 
```
Optionally you can cross compile the project by running:
```
make build_all
```
To run all tests:
```
make tests
```
## Usage
Run the binary, chose your symbol, and starting position
```
 ./ttt
Starting a new game
Chose your symbol, X or O:
X
Chose your starting position, 1 or 2:
1
Chose your next move, in format 1x2 (rowXcolumn)
1x2

 O  X  .
 .  .  .
 .  .  .
```
Your moves should be in format 1x2, where 1 is the row, 2 is the column. 
Moves range is [1:3], for both row and column, where 1 is the minimum and 3 is the maximum possible value.

