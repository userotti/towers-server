package components

import (
	"time"

	"engo.io/engo"
)

//MoveTweenComponent manage the tweening of a entity
type MoveTweenComponent struct {
	StartPosition       engo.Point
	DestinationPosition engo.Point
	Range               float32
	Speed               float32
	StartTime           time.Time
	EndTime             time.Time
	Tweening            bool
}
