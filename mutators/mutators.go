package mutators

import "gosoro/ds"
//import "fmt"


type Mutator struct {
}


func pop_instruction_stack(stack *[][]ds.Instruction) []ds.Instruction {
	// An instruction stack is slice of slices of instructions.
	// This method returns a value off the stack, and removes that value
	// From the stack
	len_stack := len(*stack)
	val := (*stack)[len_stack-1]
	*stack = (*stack)[:len_stack-1]
	return val
}



// This method needs rethinking because it relies on the ExecuteInstruction method
//func all_moves(board ds.Board, player_number int) (boards []ds.Board, instructions [][]ds.Instruction) {
//	// first find all rows and columns with more than two
//	p, _ := ds.PlayersFromName(player_number, &board)
//	var coords []ds.Coord
//	for row_index, row := range p.Positions {
//		for column_index, num_seeds := range row {
//			if num_seeds > 1 {
//				coord := ds.Coord{row_index, column_index, &board, p}
//				coords = append(coords, coord)
//			}
//		}
//	}
//
//	if len(coords) <= 0 {
//		fmt.Println("COMPUTER HAS BEEN DEFEATED")
//	}
//
//	var initial_instructions []ds.Instruction // list of all initial instructions, not yet popualting the stack
//	for _, c := range coords {
//		instruction := ds.Instruction{c.Row, c.Column, "A", board}       // Perform the counterclockwise move
//		initial_instructions = append(initial_instructions, instruction) // Append instruction to initial instruction list
//		if board.Is_bidirectional(c.Row, c.Column) {                     // If we can move counterclockwise from this position
//			instruction = ds.Instruction{c.Row, c.Column, "C", board}
//			initial_instructions = append(initial_instructions, instruction)
//		}
//	}
//
//	var instructions_stack [][]ds.Instruction // list of lists of instructions
//
//	for _, i_instruction := range initial_instructions { // populate the instruction stack
//		instructions_stack = append(instructions_stack, []ds.Instruction{i_instruction})
//	}
//
//	for len(instructions_stack) > 0 { // While the stack is not emtpy
//		instruction_set := (pop_instruction_stack(&instructions_stack)) // Get a list of instructions
//
//		instruction := instruction_set[len(instruction_set)-1]  // take the last instruction
//		b, i := ds.ExecuteInstruction(instruction, player_number) // Get the resulting instruction
//		if len(i) > 0 {
//			for _, ins := range i { // If more instructions, populate the instructions stack
//				new_instruction_set := append(instruction_set, ins)
//				instructions_stack = append(instructions_stack, new_instruction_set)
//			}
//		} else {
//			boards = append(boards, b)
//			instructions = append(instructions, instruction_set)
//		}
//	}
//
//	return boards, instructions // return the boards and full corresponding instruction sets
//}
