package components

import (
	"time"

	"engo.io/engo"
)

//MoveTweenComponent manage the tweening of a entity
type MoveTweenComponent struct {
	StartPosition       engo.Point
	DestinationPosition engo.Point
	MoveStartTime       time.Time
	Range               float32
	Min                 float32 //0 is good choice for this
	Max                 float32 //1 is good choice for this,
	Current             float32 //is gonna change every game tick
	Speed               float32 //is the amount of seconds to complete the tween
	Done                bool    //bool to indicate if we should still tween
	CalculatedDuration  float32
}
