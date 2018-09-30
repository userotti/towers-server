package components

import (
	"time"

	"engo.io/engo"
)

//MoveTweenComponent manage the tweening of a entity
type MoveTweenComponent struct {
	StartPosition            engo.Point
	DestinationPosition      engo.Point
	StartTime                time.Time
	NanosecondsFromStart     int64
	NanosecondsTotalDuration int64
	Range                    float64
	Speed                    float64 //is the amount of seconds to complete the tween
	Done                     bool    //bool to indicate if we should still tween
}
