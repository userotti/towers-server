package components

import "time"

//MoveCooldownComponent keep the values assciated with moving
type MoveCooldownComponent struct {
	StartTime                time.Time
	NanosecondsFromStart     int64
	NanosecondsTotalDuration int64
	Cooldown                 float64 //Cooldown max
	Recharge                 float64 //Cooldown recharge per second
	Done                     bool    //bool to indicate if we should still tween
}
