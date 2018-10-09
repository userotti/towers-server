package systems

import (
	"math"
	"time"
	"towers-server/components"
	"towers-server/utility"

	"engo.io/ecs"
)

//WeaponRechargeEntity system entity
type WeaponRechargeEntity struct {
	*ecs.BasicEntity
	*components.WeaponComponent
}

//WeaponRechargeSystem mange move cooldown
type WeaponRechargeSystem struct {
	entities []WeaponRechargeEntity
}

//New whne the system is created
func (wrs *WeaponRechargeSystem) New(w *ecs.World) {

}

//Add new entity to the system
func (wrs *WeaponRechargeSystem) Add(basic *ecs.BasicEntity, move *components.WeaponComponent) {
	wrs.entities = append(wrs.entities, WeaponRechargeEntity{basic, move})
}

//Remove removes them from the system
func (wrs *WeaponRechargeSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range wrs.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		wrs.entities = append(wrs.entities[:delete], wrs.entities[delete+1:]...)
	}
}

//Update updates all entities
func (wrs *WeaponRechargeSystem) Update(dt float32) {

	for _, e := range wrs.entities {
		if !e.WeaponComponent.Loaded {

			//Calculate the tween values
			e.WeaponComponent.NanosecondsTotalDuration = int64((math.Pow(10, 9) * e.WeaponComponent.Cooldown) / e.WeaponComponent.Recharge)
			e.WeaponComponent.NanosecondsFromStart = e.WeaponComponent.NanosecondsFromStart + int64(math.Pow(10, 9)*float64(dt))

			var startTime = e.WeaponComponent.StartTime
			var endTime = e.WeaponComponent.StartTime.Add(time.Duration(e.WeaponComponent.NanosecondsTotalDuration))

			//Apply the desired tween function
			tweenFunctionResult := utility.Linear(startTime, endTime, startTime.Add(time.Duration(e.WeaponComponent.NanosecondsFromStart))) //Change NanosecondsFromStart of time.Now for time travel stuff

			//Stop tweening
			if tweenFunctionResult == 1 {
				e.WeaponComponent.Loaded = true
			}

		}
	}

}
