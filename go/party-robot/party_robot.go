package partyrobot

import (
	"fmt"
	"strings"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	welcome := Welcome(name)

	tableAssignMsg := fmt.Sprintf("You have been assigned to table %03d.", table)
	tableDirections := fmt.Sprintf("Your table is %s, exactly %.1f meters from here.", direction, distance)
	tableInfo := strings.Join([]string{tableAssignMsg, tableDirections}, " ")

	neighborIntro := fmt.Sprintf("You will be sitting next to %s.", neighbor)

	return strings.Join([]string{welcome, tableInfo, neighborIntro}, "\n")
}
