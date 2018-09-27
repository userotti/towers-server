package systems

import (
	"time"
	"towers/components"
	"towers/utility"

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
	entity.MoveTweenComponent.CalculatedDuration = (entity.MoveTweenComponent.Max - entity.MoveTweenComponent.Min) / entity.MoveTweenComponent.Speed
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

		//Don't update the duration during the tween, but make sure the next tween is with the latest states. Speed, Min and Max...
		if e.MoveTweenComponent.Done {
			e.MoveTweenComponent.CalculatedDuration = (e.MoveTweenComponent.Max - e.MoveTweenComponent.Min) / e.MoveTweenComponent.Speed
		}

		if !e.MoveTweenComponent.Done {

			//Calculate the tween values
			e.MoveTweenComponent.Current = e.MoveTweenComponent.Current + (e.MoveTweenComponent.Speed * dt) // this is not neccesary at this stage,but might be valueble at a later stage.
			tweenFunctionResult := utility.Linear(e.MoveTweenComponent.StartTime, e.MoveTweenComponent.StartTime.Add(time.Duration(e.MoveTweenComponent.CalculatedDuration)*time.Second), time.Now())

			//Stop tweening
			if tweenFunctionResult == 1 {
				//removeList = append(removeList, e) //*** Perf optimatization 1 ***
				e.MoveTweenComponent.Current = e.MoveTweenComponent.Max
				e.MoveTweenComponent.StartPosition.X = e.MoveTweenComponent.DestinationPosition.X
				e.MoveTweenComponent.StartPosition.Y = e.MoveTweenComponent.DestinationPosition.Y
				e.MoveTweenComponent.Done = true
			}

			//Move the the space component
			e.SpaceComponent.Position.X = e.MoveTweenComponent.StartPosition.X + (e.MoveTweenComponent.DestinationPosition.X-e.MoveTweenComponent.StartPosition.X)*tweenFunctionResult
			e.SpaceComponent.Position.Y = e.MoveTweenComponent.StartPosition.Y + (e.MoveTweenComponent.DestinationPosition.Y-e.MoveTweenComponent.StartPosition.Y)*tweenFunctionResult

		}
	}

	////*** Perf optimatization 1 ***
	// if len(removeList) != 0 {
	// 	for _, e := range removeList {
	// 		mts.Remove(*e.BasicEntity)
	// 	}
	// }

}
