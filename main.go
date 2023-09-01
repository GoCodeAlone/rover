package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Setting plateau...")
	// create a plateau
	plateau := Plateau{5, 5}
	fmt.Println("Plateau size: ", plateau.x, plateau.y)

	fmt.Println("Instantiating rovers...")
	// create a rover
	rover1 := Rover{plateau, Position{1, 2}, Direction{"N"}}
	rover2 := Rover{plateau, Position{3, 3}, Direction{"E"}}
	fmt.Println("Rover 1 starting position: ", rover1.getPosition())
	fmt.Println("Rover 2 starting position: ", rover2.getPosition())

	fmt.Println("Building instruction sets...")
	// create instructions
	instruction1 := Instruction{0, "LMLMLMLMM"}
	instruction2 := Instruction{0, "MMRMMRMRRM"}

	fmt.Println("Pairing rovers with instructions...")
	// create rover instructions
	roverInstruction1 := RoverInstruction{rover1, instruction1}
	roverInstruction2 := RoverInstruction{rover2, instruction2}
	fmt.Println("Rover 1 instruction set: ", roverInstruction1.instruction.instruction)
	fmt.Println("Rover 2 instruction set: ", roverInstruction2.instruction.instruction)

	fmt.Println("Executing instruction sets...")
	// execute instructions
	rover1 = roverInstruction1.execute()
	rover2 = roverInstruction2.execute()

	// print final position of the rover
	fmt.Println("Rover 1 final position: ", rover1.getPosition())
	fmt.Println("Rover 2 final position: ", rover2.getPosition())
}

type Rover struct {
	plateau   Plateau
	position  Position
	direction Direction
}

type RoverInstruction struct {
	rover       Rover
	instruction Instruction
}

// execute instructions against the rover
func (ri RoverInstruction) execute() Rover {
	for instruction := ri.instruction.getInstruction(); instruction != ""; instruction = ri.instruction.getInstruction() {
		//fmt.Println("Instruction: ", instruction)
		switch instruction {
		case "L":
			ri.rover = ri.rover.turnLeft()
		case "R":
			ri.rover = ri.rover.turnRight()
		case "M":
			// when moving, ensure that the rover does not move beyond the plateau
			if ri.rover.direction.direction == "N" && ri.rover.position.y < ri.rover.plateau.y {
				ri.rover = ri.rover.forward()
			} else if ri.rover.direction.direction == "E" && ri.rover.position.x < ri.rover.plateau.x {
				ri.rover = ri.rover.forward()
			} else if ri.rover.direction.direction == "S" && ri.rover.position.y > 0 {
				ri.rover = ri.rover.forward()
			} else if ri.rover.direction.direction == "W" && ri.rover.position.x > 0 {
				ri.rover = ri.rover.forward()
			} else {
				fmt.Println("Rover cannot move beyond the plateau")
			}
		}
	}
	return ri.rover
}

// move rover's x,y coordinates based on the direction the rover is facing when the rover is moved
func (r Rover) forward() Rover {
	switch r.direction.direction {
	case "N":
		r.position.y++
	case "E":
		r.position.x++
	case "S":
		r.position.y--
	case "W":
		r.position.x--
	}
	return r
}

func (r Rover) turnLeft() Rover {
	r.direction = r.direction.setDirection("L")
	return r
}

func (r Rover) turnRight() Rover {
	r.direction = r.direction.setDirection("R")
	return r
}

func (r Rover) getPosition() string {
	return fmt.Sprintf("%d %d %s", r.position.x, r.position.y, r.direction.direction)
}

type Plateau struct {
	x int
	y int
}

type Direction struct {
	direction string
}

// set the direction of the rover based on which direction the rover is to turn
// rover can turn L or R (left or right), with the new direction being derived from the current direction
func (d Direction) setDirection(turn string) Direction {
	currentDir := "NSEW"
	leftDir := []string{"W", "E", "N", "S"}
	rightDir := []string{"E", "W", "S", "N"}

	currentDirIndex := strings.Index(currentDir, d.direction)

	switch turn {
	case "L":
		d.direction = leftDir[currentDirIndex]
	case "R":
		d.direction = rightDir[currentDirIndex]
	}

	return d
}

type Instruction struct {
	step        int
	instruction string
}

// iterate through the instruction string, returning the current instruction to execute
// move pointer to the next character in the instruction string as we iterate
func (i *Instruction) getInstruction() string {
	if i.step < len(i.instruction) {
		instruction := string(i.instruction[i.step])
		i.step++
		return instruction
	}
	return ""
}

type Position struct {
	x int
	y int
}
