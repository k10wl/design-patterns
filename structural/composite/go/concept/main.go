package main

import (
	"fmt"
	"strings"
)

type Composite interface {
	Report() string
}

type Member struct {
	Name string
	Role string
}

type Team struct {
	Captain Composite
	Members []Composite
}

func (member Member) Report() string {
	return fmt.Sprintf("my name is %s, my role - %s", member.Name, member.Role)
}

func (team Team) Report() string {
	var stringBuilder strings.Builder
	stringBuilder.WriteString(fmt.Sprintf("team captain: \n-%s\n\nteam members:\n", team.Captain.Report()))
	for _, member := range team.Members {
		stringBuilder.WriteString(fmt.Sprintf("- %s;\n", member.Report()))
	}
	return stringBuilder.String()
}

func main() {
	fmt.Printf(">>> Composite demo\n\n")
	team := &Team{
		Captain: Member{
			Name: "Joe",
			Role: "Leader",
		},
		Members: []Composite{
			Member{
				Name: "Rendy",
				Role: "Gudard",
			},
			Member{
				Name: "Linda",
				Role: "Imposter",
			},
			Member{
				Name: "Khorn",
				Role: "Warrior",
			},
		},
	}
	fmt.Print(team.Report())
}
