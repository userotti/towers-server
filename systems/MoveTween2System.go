package systems

import (
	"fmt"
	"time"
	"towers/components"

	"engo.io/ecs"
	"engo.io/engo/common"
)

//MoveTween2Entity is an example of a entity that needs to tween
type MoveTween2Entity struct {
	*ecs.BasicEntity
	*components.MoveTween2Component
	*common.SpaceComponent
}

//MoveTween2System is responsible for the tweening of towers
type MoveTween2System struct {
	entities []MoveTween2Entity
}

//New whne the system is created
func (mts *MoveTween2System) New(w *ecs.World) {
}

//Add uAdd new entity
func (mts *MoveTween2System) Add(basic *ecs.BasicEntity, space *common.SpaceComponent, tween *components.MoveTween2Component) {
	var entity = MoveTween2Entity{basic, tween, space}
	fmt.Println("hi2")
	entity.MoveTween2Component.CalculatedDuration = (entity.MoveTween2Component.Max - entity.MoveTween2Component.Min) / entity.MoveTween2Component.Speed
	mts.entities = append(mts.entities, entity)
}

//Remove removes them from the system
func (mts *MoveTween2System) Remove(basic ecs.BasicEntity) {
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
func (mts *MoveTween2System) Update(dt float32) {

	var removeList []MoveTween2Entity

	for _, e := range mts.entities {
		if !e.MoveTween2Component.Done {
			fmt.Println(dt)
			fmt.Println(e.MoveTween2Component.Current)

			linearResult := linear(e.MoveTween2Component.MoveStartTime, e.MoveTween2Component.MoveStartTime.Add(time.Duration(e.MoveTween2Component.CalculatedDuration)*time.Second), time.Now())

			e.MoveTween2Component.Current = e.MoveTween2Component.Current + (e.MoveTween2Component.Speed * dt)
			if e.MoveTween2Component.Current >= e.MoveTween2Component.Max {
				e.MoveTween2Component.Current = e.MoveTween2Component.Max
				e.MoveTween2Component.Done = true
			}

			e.SpaceComponent.Position.X = e.MoveTween2Component.StartPosition.X + (e.MoveTween2Component.DestinationPosition.X-e.MoveTween2Component.StartPosition.X)*linearResult
			e.SpaceComponent.Position.Y = e.MoveTween2Component.StartPosition.Y + (e.MoveTween2Component.DestinationPosition.Y-e.MoveTween2Component.StartPosition.Y)*linearResult

		}
	}

	if len(removeList) != 0 {
		for _, e := range removeList {
			mts.Remove(*e.BasicEntity)
		}
	}

}
