package components

//TeamName enum for teams
type TeamName int

//TeamNames enum
const (
	Sprite  TeamName = 0
	Monster TeamName = 1
	Rock    TeamName = 2
	// strong  TeamName = 3
	// proud   TeamName = 4
	// green   TeamName = 5
	// black   TeamName = 6
)

//Return the
func (teamName TeamName) String() string {
	names := [...]string{
		"Sprite",
		"Monster",
		"Rock"}

	if teamName < Sprite || teamName > Rock {
		return "Unknown"
	}

	return names[teamName]
}

//TeamComponent decideds which team the tower is in, also details about that team
type TeamComponent struct {
	Name TeamName
}
