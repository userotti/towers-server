package components

//SingleValueTweenComponent testing tweening
type SingleValueTweenComponent struct {
	Min                float32 //0 is good choice for this
	Max                float32 //1 is good choice for this,
	Current            float32 //is gonna change every game tick
	Speed              float32 //is the amount of seconds to complete the tween
	Done               bool    //bool to indicate if we should still tween
	CalculatedDuration float32
}
