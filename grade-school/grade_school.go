package school

import (
	"sort"
)

const testVersion = 1

// Grade is a roster of student names for each grade level
// should this be a dictionary?
type Grade struct {
	Level int
	Names []string
}

// School contains a list of grades
type School struct {
	Grades []Grade
}

// New creates a new school object
func New() *School {
	return &School{Grades: []Grade{}}
}

// Add adds a student to the given grade level
func (s *School) Add(name string, level int) {
	levelPosn := findGradePosition(level, s.Grades)
	if levelPosn == -1 {
		// no level currently exists, create one and find position in array
		s.Grades = append(s.Grades, Grade{Level: level})
		levelPosn = findGradePosition(level, s.Grades)
	}

	if s.Grades[levelPosn].Names == nil {
		s.Grades[levelPosn].Names = make([]string, 0)
	}
	s.Grades[levelPosn].Names = append(s.Grades[levelPosn].Names, name)
}

// Grade returns a list of student names in the specified grade level
func (s *School) Grade(level int) []string {
	if len(s.Grades) == 0 {
		return []string{}
	}
	levelPosn := findGradePosition(level, s.Grades)
	return s.Grades[levelPosn].Names
}

// Enrollment returns a sorted list of students in all grades
func (s *School) Enrollment() []Grade {
	sortedGrades := sortGrades(s.Grades)
	return sortedGrades
}

// helper function to find the location of a given level in the grades array,
// since we are not using a map
func findGradePosition(level int, grades []Grade) int {
	for i := 0; i < len(grades); i++ {
		if grades[i].Level == level {
			return i
		}
	}
	return -1
}

// helper function to sort grades and students within grades
func sortGrades(grades []Grade) []Grade {
	// sort grades by level
	// sorting by a struct property requires implementation of sort's less function
	sort.Slice(grades, func(i, j int) bool { return grades[i].Level < grades[j].Level })

	// sort names alphabetically within grades
	for i := 0; i < len(grades); i++ {
		sort.Strings(grades[i].Names)
	}
	return grades
}
