package systems

import (
	"math"
	"time"
	"../components"
	"../utility"

	"engo.io/ecs"
	"engo.io/engo/common"
)

//MoveTweenEntity is an example of a entity that needs to tween
type MoveTweenEntity struct {
	*ecs.BasicEntity
	*components.MoveTweenComponent
	*common.SpaceComponent
}

//MoveTweenSystem is responsible for the tweening of towers
type MoveTweenSystem struct {
	entities []MoveTweenEntity
}

//New whne the system is created
func (mts *MoveTweenSystem) New(w *ecs.World) {
}

//Add uAdd new entity
func (mts *MoveTweenSystem) Add(basic *ecs.BasicEntity, space *common.SpaceComponent, tween *components.MoveTweenComponent) {
	var entity = MoveTweenEntity{basic, tween, space}
	mts.entities = append(mts.entities, entity)
}

//Remove removes them from the system
func (mts *MoveTweenSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range mts.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		mts.entities = append(mts.entities[:delete], mts.entities[delete+1:]...)
	}
}

//Update updates all entities
func (mts *MoveTweenSystem) Update(dt float32) {

	//*** Perf optimatization 1 *** Remove items from the system if its not tweening.
	// var removeList []MoveTweenEntity

	for _, e := range mts.entities {

		if !e.MoveTweenComponent.Done {

			//Calculate the tween values
			e.MoveTweenComponent.NanosecondsTotalDuration = int64(math.Pow(10, 9) * float64(e.MoveTweenComponent.StartPosition.PointDistance(e.MoveTweenComponent.DestinationPosition)) / e.MoveTweenComponent.Speed)
			e.MoveTweenComponent.NanosecondsFromStart = e.MoveTweenComponent.NanosecondsFromStart + int64(math.Pow(10, 9)*float64(dt))

			var startTime = e.MoveTweenComponent.StartTime
			var endTime = e.MoveTweenComponent.StartTime.Add(time.Duration(e.MoveTweenComponent.NanosecondsTotalDuration))

			//Apply the desired tween function
			tweenFunctionResult := utility.Linear(startTime, endTime, startTime.Add(time.Duration(e.MoveTweenComponent.NanosecondsFromStart))) //Change NanosecondsFromStart of time.Now for time travel stuff

			//Move the the space component
			e.SpaceComponent.Position.X = e.MoveTweenComponent.StartPosition.X + (e.MoveTweenComponent.DestinationPosition.X-e.MoveTweenComponent.StartPosition.X)*tweenFunctionResult
			e.SpaceComponent.Position.Y = e.MoveTweenComponent.StartPosition.Y + (e.MoveTweenComponent.DestinationPosition.Y-e.MoveTweenComponent.StartPosition.Y)*tweenFunctionResult

			//Stop tweening
			if tweenFunctionResult == 1 {
				e.MoveTweenComponent.StartPosition.X = e.MoveTweenComponent.DestinationPosition.X
				e.MoveTweenComponent.StartPosition.Y = e.MoveTweenComponent.DestinationPosition.Y
				e.MoveTweenComponent.Done = true
			}

		}
	}

	////*** Perf optimatization 1 ***
	// if len(removeList) != 0 {
	// 	for _, e := range removeList {
	// 		mts.Remove(*e.BasicEntity)
	// 	}
	// }

}
