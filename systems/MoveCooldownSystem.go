package systems

import (
	"math"
	"time"
	"towers-server/components"
	"towers-server/utility"

	"engo.io/ecs"
)

//MoveCooldownEntity system entity
type MoveCooldownEntity struct {
	*ecs.BasicEntity
	*components.MoveCooldownComponent
}

//MoveCooldownSystem mange move cooldown
type MoveCooldownSystem struct {
	entities []MoveCooldownEntity
}

//New whne the system is created
func (mcs *MoveCooldownSystem) New(w *ecs.World) {

}

//Add new entity to the system
func (mcs *MoveCooldownSystem) Add(basic *ecs.BasicEntity, move *components.MoveCooldownComponent) {
	mcs.entities = append(mcs.entities, MoveCooldownEntity{basic, move})
}

//Remove removes them from the system
func (mcs *MoveCooldownSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range mcs.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		mcs.entities = append(mcs.entities[:delete], mcs.entities[delete+1:]...)
	}
}

//Update updates all entities
func (mcs *MoveCooldownSystem) Update(dt float32) {

	for _, e := range mcs.entities {
		if !e.MoveCooldownComponent.Done {

			//Calculate the tween values
			e.MoveCooldownComponent.NanosecondsTotalDuration = int64((math.Pow(10, 9) * e.MoveCooldownComponent.Cooldown) / e.MoveCooldownComponent.Recharge)
			e.MoveCooldownComponent.NanosecondsFromStart = e.MoveCooldownComponent.NanosecondsFromStart + int64(math.Pow(10, 9)*float64(dt))

			var startTime = e.MoveCooldownComponent.StartTime
			var endTime = e.MoveCooldownComponent.StartTime.Add(time.Duration(e.MoveCooldownComponent.NanosecondsTotalDuration))

			//Apply the desired tween function
			tweenFunctionResult := utility.Linear(startTime, endTime, startTime.Add(time.Duration(e.MoveCooldownComponent.NanosecondsFromStart))) //Change NanosecondsFromStart of time.Now for time travel stuff

			//Stop tweening
			if tweenFunctionResult == 1 {
				e.MoveCooldownComponent.Done = true
			}

		}
	}

}
