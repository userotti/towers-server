package components

//ShieldComponent keep the values associated with the shield
type ShieldComponent struct {
	MinCharge     float32
	MaxCharge     float32
	CurrentCharge float32
	RechargeSpeed float32
	ShieldUp      bool
}
