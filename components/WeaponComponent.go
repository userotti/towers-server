package components

import "time"

//WeaponComponent keep the values associated with attacking
type WeaponComponent struct {
	StartTime                time.Time //For the reload tween
	NanosecondsFromStart     int64     //For the reload tween
	NanosecondsTotalDuration int64     //For the reload tween
	Loaded                   bool      //For the reload tween

	Cooldown float64 //Cooldown max
	Recharge float64 //Cooldown recharge per second

	Damage int32 //Damage
	Range  int32 //Range
}
