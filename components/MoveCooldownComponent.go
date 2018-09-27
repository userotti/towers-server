package components

import "time"

//MoveCooldownComponent keep the values assciated with moving
type MoveCooldownComponent struct {
	StartTime          time.Time
	MinCharge          float32
	MaxCharge          float32
	CurrentCharge      float32
	RechargeSpeed      float32
	Ready              bool
	CalculatedDuration float32
}
