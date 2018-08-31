package components

//DeathCounterComponent is counter manager
type DeathCounterComponent struct {
	MinCharge     float32
	MaxCharge     float32
	CurrentCharge float32
	RechargeSpeed float32
	TimeToDie     bool
}
