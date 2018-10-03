package components

//AI enum for teams
type AI int

//AIs enum
const (
	Crazy  AI = 0
	Walker AI = 1
	Seeker AI = 2
	// strong  AI = 3
	// proud   AI = 4
	// green   AI = 5
	// black   AI = 6
)

//Return the
func (ai AI) String() string {
	names := [...]string{
		"Crazy",
		"Walker",
		"Seeker"}

	if ai < Crazy || ai > Seeker {
		return "Unknown"
	}

	return names[ai]
}

//AIComponent where moving to next
type AIComponent struct {
	Type AI
}
