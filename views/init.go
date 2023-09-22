package views

import (
	"fmt"
	"github.com/fatih/color"
	"passguard/components"
	"passguard/utils"
)

func RenderInitView() {
	utils.CreateHorizontalLine()
	fmt.Println("")

	utils.PrintWithColor(utils.CenterText("WELCOME TO PASSGUARD"), []color.Attribute{color.FgHiRed}...)
	utils.PrintWithColor(utils.CenterText("[press s to start]"), []color.Attribute{color.FgHiBlue}...)

	fmt.Println("")
	utils.CreateHorizontalLine()

	fmt.Printf("\033[?25l")
	i := utils.GetInput()

	if i == utils.START {
		utils.ClearCmd()

		RenderInitActions()
	}
}

func RenderInitActions() {
	o := []*components.Option{
		{
			Name:   "List all passwords",
			Id:     utils.GenerateId(10),
			Action: listPasswords,
		},
		{
			Name: "Add passwords",
			Id:   utils.GenerateId(10),
			Action: func() {
				fmt.Println("To be continued")
			},
		},
	}

	s := components.Select{
		Title:   "What do you wanna do?",
		Options: o,
	}

	s.DisplaySelect()
}
