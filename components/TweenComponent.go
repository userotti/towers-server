package components

import (
	"time"

	"engo.io/engo"
)

//TweenComponent manage the tweening of a entity
type TweenComponent struct {
	StartPosition       engo.Point
	DestinationPosition engo.Point
	StartTime           time.Time
	EndTime             time.Time
	Tweening            bool
}
