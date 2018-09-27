package systems

import (
	"fmt"
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

	var removeList []MoveTweenEntity

	for _, e := range mts.entities {
		if !e.MoveTweenComponent.Done {
			fmt.Println(dt)
			fmt.Println(e.MoveTweenComponent.Current)

			e.MoveTweenComponent.Current = e.MoveTweenComponent.Current + (e.MoveTweenComponent.Speed * dt) // this is not neccesary at this stage,but might be valueble at a later stage.
			linearResult := utility.Linear(e.MoveTweenComponent.MoveStartTime, e.MoveTweenComponent.MoveStartTime.Add(time.Duration(e.MoveTweenComponent.CalculatedDuration)*time.Second), time.Now())

			if linearResult == 1 {
				// removeList = append(removeList, e) //this is where we remove the entity from this system. but I don't think that is neccesary for a perf yet
				e.MoveTweenComponent.Current = e.MoveTweenComponent.Max
				e.MoveTweenComponent.Done = true
			}

			e.SpaceComponent.Position.X = e.MoveTweenComponent.StartPosition.X + (e.MoveTweenComponent.DestinationPosition.X-e.MoveTweenComponent.StartPosition.X)*linearResult
			e.SpaceComponent.Position.Y = e.MoveTweenComponent.StartPosition.Y + (e.MoveTweenComponent.DestinationPosition.Y-e.MoveTweenComponent.StartPosition.Y)*linearResult

		}
	}

	if len(removeList) != 0 {
		for _, e := range removeList {
			mts.Remove(*e.BasicEntity)
		}
	}

}
