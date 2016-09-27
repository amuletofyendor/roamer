package interop

type Command interface {
	Ready()
	IsReady() bool
	ToggleMode()
	Processed()
	ChoiceMode(validChoices []rune, postChoiceCommand, choiceDesc string)
	SetValidChoices(choices []rune)
	GetValidChoices() []rune
	Format() string
	Choice() rune

	IsQuit() bool
	IsUp() bool
	IsDown() bool
	IsLeft() bool
	IsRight() bool
	IsMove() bool
	IsTake() bool
	IsDescend() bool
	IsAscend() bool
	IsSelfwards() bool
	IsLookHere() bool
	IsEat() bool
	IsCraft() bool
	IsWield() bool
	IsWieldChoice() bool
	IsInventory() bool
	IsInventoryChoice() bool
	IsEatingChoice() bool
	IsQuitChoice() bool
}
