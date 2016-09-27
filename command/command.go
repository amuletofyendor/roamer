package command

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type Command struct {
	history           []string
	instruction       string
	ready             bool
	choice            rune
	validChoices      []rune
	postChoiceCommand string
	choiceDesc        string
	mode              int
}

func (com *Command) Processed() {
	if len(com.instruction) > 0 {
		com.history = append(com.history, com.instruction)
	}

	com.instruction = ""
	com.ready = false
}

func (com *Command) Ready() {
	com.ready = true

	if com.mode == SingleCommandMode ||
		com.mode == ChoiceMode {
		com.mode = MovementMode
	}
}

func (com *Command) IsReady() bool {
	return com.ready
}

func (com *Command) ToggleMode() {
	com.mode++
	if com.mode > CommandMode {
		com.mode = MovementMode
	}
}

func (com *Command) ChoiceMode(validChoices []rune, postChoiceCommand, choiceDesc string) {
	com.choice = '?'
	com.validChoices = nil
	com.postChoiceCommand = ""
	com.choiceDesc = ""

	if len(validChoices) > 0 {
		com.validChoices = validChoices
		com.mode = ChoiceMode
		com.postChoiceCommand = postChoiceCommand
		com.choiceDesc = choiceDesc
	}
}

func (com *Command) Append(ch rune) {
	com.instruction = com.instruction + string(ch)
}

func (com *Command) Backspace() {
	if len(com.instruction) >= 1 {
		com.instruction = com.instruction[:len(com.instruction)-1]
	}
}

func (com *Command) SetValidChoices(choices []rune) {
	com.validChoices = choices
}

func (com *Command) GetValidChoices() []rune {
	return com.validChoices
}

func (com *Command) Format() string {
	switch com.mode {
	case MovementMode:
		return "[movement mode]"
	case SingleCommandMode:
		return "[single command mode]"
	case ChoiceMode:
		return fmt.Sprintf("[choice mode] %v %c", com.choiceDesc, com.validChoices)
	case CommandMode:
		return "[multiple command mode]"
	default:
		return "[unknown mode]"
	}
}

func (com *Command) Accept(ev *termbox.Event) {
	if com.AcceptControlSequence(ev) {
		return
	}

	switch com.mode {
	case MovementMode:
		com.AcceptMovement(ev)
	case SingleCommandMode, CommandMode:
		com.AcceptCommand(ev)
	case ChoiceMode:
		com.AcceptChoice(ev)
	}
}

func (com *Command) AcceptControlSequence(ev *termbox.Event) bool {
	switch ev.Key {
	case termbox.KeyCtrlC:
		com.instruction = "quit"
		com.Ready()
		com.mode = CommandMode
	case termbox.KeySpace:
		com.ToggleMode()
	default:
		return false
	}

	return true
}

func (com *Command) AcceptMovement(ev *termbox.Event) bool {
	switch ev.Key {
	case termbox.KeyArrowUp:
		com.instruction = "up"
		com.Ready()
	case termbox.KeyArrowDown:
		com.instruction = "down"
		com.Ready()
	case termbox.KeyArrowLeft:
		com.instruction = "left"
		com.Ready()
	case termbox.KeyArrowRight:
		com.instruction = "right"
		com.Ready()
	default:
		switch ev.Ch {
		case 'q', 'Q', 'y', 'Y', '7':
			com.instruction = "ul"
			com.Ready()
		case 'w', 'W', 'k', 'K', '8':
			com.instruction = "up"
			com.Ready()
		case 'e', 'E', 'u', 'U', '9':
			com.instruction = "ur"
			com.Ready()
		case 'a', 'A', 'h', 'H', '4':
			com.instruction = "left"
			com.Ready()
		case 'd', 'D', 'l', 'L', '6':
			com.instruction = "right"
			com.Ready()
		case 'z', 'Z', 'b', 'B', '1':
			com.instruction = "dl"
			com.Ready()
		case 'x', 'X', 'j', 'J', '2':
			com.instruction = "down"
			com.Ready()
		case 'c', 'C', 'n', 'N', '3':
			com.instruction = "dr"
			com.Ready()
		case '>':
			com.instruction = "descend"
			com.Ready()
		case '<':
			com.instruction = "ascend"
			com.Ready()
		case 's', '.', '5':
			com.instruction = "selfwards"
			com.Ready()
		case ',':
			com.instruction = "take"
			com.Ready()
		case ':':
			com.instruction = "lookhere"
			com.Ready()
		case 'i', 'I':
			com.instruction = "inventory"
			com.Ready()
		default:
			return false
		}
	}

	return true
}

func (com *Command) AcceptCommand(ev *termbox.Event) bool {
	switch ev.Ch {
	case ',':
		com.instruction = "take"
		com.Ready()
	case 'q', 'Q':
		com.instruction = "quit"
		com.Ready()
	case ':':
		com.instruction = "lookhere"
		com.Ready()
	case 'e', 'E':
		com.instruction = "eat"
		com.Ready()
	case 'w', 'W':
		com.instruction = "wield"
		com.Ready()
	case 'c', 'C':
		com.instruction = "craft"
		com.Ready()
	case 'i', 'I':
		com.instruction = "inventory"
		com.Ready()
	default:
		return false
	}

	return true
}

func (com *Command) AcceptChoice(ev *termbox.Event) bool {
	for _, r := range com.validChoices {
		if ev.Ch == r {
			com.choice = r
			com.instruction = com.postChoiceCommand
			com.Ready()
			return true
		}
	}

	return false
}

func (com *Command) Choice() rune {
	return com.choice
}

func (com *Command) IsQuit() bool {
	return com.instruction == "quit"
}

func (com *Command) IsUp() bool {
	return com.instruction == "u" ||
		com.instruction == "up" ||
		com.instruction == "ur" ||
		com.instruction == "up right" ||
		com.instruction == "ul" ||
		com.instruction == "up left"
}

func (com *Command) IsDown() bool {
	return com.instruction == "d" ||
		com.instruction == "down" ||
		com.instruction == "dr" ||
		com.instruction == "down right" ||
		com.instruction == "dl" ||
		com.instruction == "down left"
}

func (com *Command) IsLeft() bool {
	return com.instruction == "l" ||
		com.instruction == "left" ||
		com.instruction == "ul" ||
		com.instruction == "up left" ||
		com.instruction == "dl" ||
		com.instruction == "down left"
}

func (com *Command) IsRight() bool {
	return com.instruction == "r" ||
		com.instruction == "right" ||
		com.instruction == "dr" ||
		com.instruction == "down right" ||
		com.instruction == "ur" ||
		com.instruction == "up right"
}

func (com *Command) IsMove() bool {
	return com.IsUp() || com.IsDown() || com.IsLeft() || com.IsRight()
}

func (com *Command) IsTake() bool {
	return com.instruction == "take"
}

func (com *Command) IsDescend() bool {
	return com.instruction == "descend"
}

func (com *Command) IsAscend() bool {
	return com.instruction == "ascend"
}

func (com *Command) IsSelfwards() bool {
	return com.instruction == "selfwards"
}

func (com *Command) IsLookHere() bool {
	return com.instruction == "lookhere"
}

func (com *Command) IsEat() bool {
	return com.instruction == "eat"
}

func (com *Command) IsCraft() bool {
	return com.instruction == "craft"
}

func (com *Command) IsWield() bool {
	return com.instruction == "wield"
}

func (com *Command) IsWieldChoice() bool {
	return com.instruction == "wield choice"
}

func (com *Command) IsInventory() bool {
	return com.instruction == "inventory"
}

func (com *Command) IsInventoryChoice() bool {
	return com.instruction == "inventory choice"
}

func (com *Command) IsEatingChoice() bool {
	return com.instruction == "eating choice"
}

func (com *Command) IsQuitChoice() bool {
	return com.instruction == "quit choice"
}
