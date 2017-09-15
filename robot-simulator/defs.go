package robot

import "fmt"

const testVersion = 3

// compass directions, 0-3
const (
	N = iota
	E
	S
	W
)

// definitions used in step 1

// Step1Robot contains the current coordinates and direction for the robot
var Step1Robot struct {
	X, Y int
	Dir
}

// Dir is the direction the robot is facing
type Dir int

func (d Dir) String() string {
	return fmt.Sprintf("%s")
}

var _ fmt.Stringer = Dir(1729)

// Advance will move the robot forward in whatever direction they are facing
func Advance() {
	if Step1Robot.Dir == 0 {
		Step1Robot.Y++
	} else if Step1Robot.Dir == 1 {
		Step1Robot.X++
	} else if Step1Robot.Dir == 2 {
		Step1Robot.Y--
	} else if Step1Robot.Dir == 3 {
		Step1Robot.X--
	}
}

// Right will change direction to the right
func Right() {
	if Step1Robot.Dir < 3 {
		Step1Robot.Dir++
	} else {
		Step1Robot.Dir = 0
	}
}

// Left will change direction to the left
func Left() {
	if Step1Robot.Dir > 0 {
		Step1Robot.Dir--
	} else {
		Step1Robot.Dir = 3
	}
}

// additional definitions used in step 2

// Command has valid values of 'R', 'L', 'A'
type Command byte
type RU int

// Pos is x,y coords
type Pos struct{ Easting, Northing RU }
type Rect struct{ Min, Max Pos }
type Step2Robot struct {
	Dir
	Pos
}
type Action interface {
	Advance()
	Left()
	Right()
}

// Room has walls and a robot
func Room(rect Rect, robot Step2Robot, action chan Action, rep chan Step2Robot) {
	// stop robot from running into walls
}

// StartRobot starts the robot
// given a command, do the corresponding action
func StartRobot(cmd chan Command, action chan Action) {

}

// additional definition used in step 3

type Step3Robot struct {
	Name string
	Step2Robot
}
