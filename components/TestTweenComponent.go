package components

//TestTweenComponent testing tweening
type TestTweenComponent struct {
	MinCharge     float32 //0 is good choice for this
	MaxCharge     float32 //1 is good choice for this,
	CurrentCharge float32 //is gonna change every game tick
	RechargeSpeed float32 //is the amount of seconds to complete the tween
	Charged       bool    //bool to indicate if we should still tween
}
