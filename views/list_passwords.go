package views

import (
	"fmt"
	"github.com/fatih/color"
	"passguard/components"
	"passguard/utils"
	"strconv"
)

type Password struct {
	name     string
	password string
}

func listPasswords() {
	dummyPs := []Password{{
		name:     "Gmail",
		password: "asdaszxc",
	}, {
		name:     "Facebook",
		password: "Onasdsddsa",
	}, {
		name:     "Twitter",
		password: "adasasdasjkasgdgk",
	}, {
		name:     "Asana",
		password: "sdfkjdsjdfhksdjhsdf",
	}}

	utils.ClearCmd()

	var o []*components.Option

	for _, p := range dummyPs {
		o = append(o, &components.Option{
			Name: p.name,
			Id:   utils.GenerateId(10),
			Action: func() {
				fmt.Println("asdasdads")
			},
		})
	}

	s := components.Select{
		Title: fmt.Sprintf("All your passwords %v",
			color.CyanString(fmt.Sprintf("(%v)", strconv.Itoa(len(o)))),
		),
		Options: o,
		Description: fmt.Sprintf("\n%v\n%v\n%v\n",
			color.BlueString("[C] - copy password"),
			color.RedString("[D] - delete password"),
			color.YellowString("[E] - edit password"),
		),
		BackView: RenderInitActions,
	}

	s.DisplaySelect()

}
