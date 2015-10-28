# gosoro
Igisoro AI, implemented in GO

# Further thoughts on the game:

The AI is currently very board-centric, that is to say it evaluates moves based on the resultant board.
In actual fact the instruction or move is the central concept, a game should be seen as a sequence of moves
not board states. A move can either be a parent move (other moves must follow from it) or a leaf move.
The AI can then find the path through the move tree which terminates in the best leaf move.
A leaf move can either be one which leads to a capture, forcing another move, or one which
causes automatic victory. A leaf move can also be evaluated in a minimax-style manner by
evaluating what leaf moves the oponent may play based on the eventual board
associated with the computer's leaf move.

# Notes

 Initial rules: No reversing, only victory is if other player can't move

 Need the following methods:
 1. Board initialiser - this is a method which just returns a board.
 2. Scoring function
 3. Board mmutator
 4. A string-based notation for the move.
 5. A player-perspective on the board, so that confusions regarding clockwise/counterclockwise etc are easily resolved.

 Go Questions:
 Q 1. What is a slice?

 Q 2. How do you add a method to a type
 A func (this Type) func_name(func_param param_type) (return_type) {...}

 Q 3. What is an interface
 A If a type has all the correct function names, including signatures, then it satisfies an interface
   And can be passed to another method which takes that interface.
   Still not sure about the details of this but it sounds quite interesting.
 Q how do maps work in go?


 UPDATE: I've decided that the 4x8 board representation is very unhelpful.
 The 2x2x8 is better, that is to say, the board is represented as two player views.
 Each player view is a 2x8 board.
 For calcuations of scores this can still be mapped to the 4x8 representation if need be,
 But for carrying out mutations this is a much more simple strategy.

 Thoughts on mutators:
 In some positions a decision can be made whether to go clockwise or counterclockwise.
 Some sort of tree is going to be needed to keep track of decisions, OR we could not
 implement this aspect of it for now.

 Instruction Format: RCD (row, column, direction)
 Example: 02C or 16A


