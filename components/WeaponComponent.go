package components

//WeaponComponent keep the values associated with attacking
type WeaponComponent struct {
	MinCharge     float32
	MaxCharge     float32
	CurrentCharge float32
	RechargeSpeed float32
	Damage        int32
	Range         int32
}
