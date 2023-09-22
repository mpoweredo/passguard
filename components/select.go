package components

import (
	"fmt"
	"github.com/fatih/color"
	"passguard/utils"
)

type Option struct {
	Name   string
	Id     string
	Action func()
}

type Select struct {
	Title       string
	Options     []*Option
	Description string
	BackView    func()
}

func (s *Select) renderSelect(arrowAtIndex int) {
	underline := color.New(color.Underline).Add(color.FgGreen).SprintFunc()

	fmt.Println(color.MagentaString(s.Title))

	if s.BackView != nil {
		fmt.Println("")
		fmt.Printf("Press %v or %v to go back", color.BlueString("backspace"), color.HiBlueString("left arrow"))
	}

	if s.Description != "" {
		fmt.Println(s.Description)
	}

	for i, o := range s.Options {
		if i == arrowAtIndex {
			fmt.Printf("%v %v", color.GreenString(">"), underline(o.Name))
			fmt.Println("")
			continue
		}
		fmt.Printf("  %v", o.Name)
		fmt.Println("")
	}
}

func (s *Select) DisplaySelect() {
	i := 0
	oLength := len(s.Options)

	s.renderSelect(0)

	// Turn the terminal cursor off
	fmt.Printf("\033[?25l")

	for {
		keyCode := utils.GetInput()

		if keyCode == utils.ESCAPE {
			break
		}

		if keyCode == utils.UP {
			if i == 0 {
				continue
			}

			i--
		}

		if keyCode == utils.DOWN {
			if i == oLength-1 {
				continue
			}

			i++
		}

		if keyCode == utils.BACKSPACE || keyCode == utils.LEFT_ARROW {
			if s.BackView == nil {
				continue
			}

			utils.ClearCmd()
			s.BackView()
		}

		if keyCode == utils.ENTER {
			s.Options[i].Action()

			break
		}

		utils.ClearCmd()
		s.renderSelect(i)
	}
}
