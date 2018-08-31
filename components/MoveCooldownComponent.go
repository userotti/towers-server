package components

//MoveCooldownComponent keep the values assciated with moving
type MoveCooldownComponent struct {
	MinCharge     float32
	MaxCharge     float32
	CurrentCharge float32
	RechargeSpeed float32
	Ready         bool
}
