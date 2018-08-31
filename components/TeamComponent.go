package components

//TeamName enum for teams
type TeamName int

const (
	sprite  TeamName = 0
	monster TeamName = 1
	rock    TeamName = 2
	strong  TeamName = 3
	proud   TeamName = 4
	green   TeamName = 5
	black   TeamName = 6
)

//TeamComponent decideds which team the tower is in, also details about that team
type TeamComponent struct {
	TeamName    string
	FriendTeams string
	TargetTeams string
}
