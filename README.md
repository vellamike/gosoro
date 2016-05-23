# gosoro
Igisoro and Omweso are games of the mancala family.

In this project I aim to implement Igisoro and Omweso AIs in the Go programming language.

# How to play

To start a game:
```
go run main.go
```
The computer will play the first move.

Instruction format: `<row><column><direction><row><column><direction>...`

For example, `02C16A` corresponds to the move row 0, column 2, move clockwise, capture and move counterclockwise from row 1, column 6 row A. When the computer asks a user for their move it is expected to be in this format.


# GOPATH

GOPATH should point to a directoy of which src/gosoro is a subdirectory. This is go convention and the best way to make sure that tooling works.
