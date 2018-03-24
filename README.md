# gosoro
Igisoro and Omweso are games of the mancala family.

In this project I aim to implement Igisoro and Omweso AIs in the Go programming language.

# How to play

To start a game:

currently, the only supported mancala-family game is "igisoro".

```
go run igisoro/main.go
```
The computer will play the first move.

Instruction format: `<row><column><direction><row><column><direction>...`

For example, `02C16A` corresponds to the move row 0, column 2, move clockwise, capture and move counterclockwise from row 1, column 6 row A. When the computer asks a user for their move it is expected to be in this format.


# GOPATH

GOPATH should point to a directoy of which src/gosoro is a subdirectory. This is go convention and the best way to make sure that tooling works.


## TODO

1. Overall structure - does it fit in well with a Go project?
2. Decide how the game controller will fix the next player. Right now the AI never gets to play, but it is necessary to make it so that a branch move works too. Probably the ruleset should, given a move, and a board before and after the move has been executed, be able to decide whether it is a leaf move or not and if not what the available moves are.
