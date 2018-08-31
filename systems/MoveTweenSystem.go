package systems

import (
	"time"
	"towers/components"

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
	mts.entities = append(mts.entities, MoveTweenEntity{basic, tween, space})
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

		linearResult := linear(e.MoveTweenComponent.StartTime, e.MoveTweenComponent.EndTime, time.Now())
		if e.MoveTweenComponent.Tweening {
			e.SpaceComponent.Position.X = e.MoveTweenComponent.StartPosition.X + (e.MoveTweenComponent.DestinationPosition.X-e.MoveTweenComponent.StartPosition.X)*linearResult
			e.SpaceComponent.Position.Y = e.MoveTweenComponent.StartPosition.Y + (e.MoveTweenComponent.DestinationPosition.Y-e.MoveTweenComponent.StartPosition.Y)*linearResult

			if linearResult == 1 {
				// removeList = append(removeList, e)
				e.MoveTweenComponent.Tweening = false
			}
		}

	}

	if len(removeList) != 0 {
		for _, e := range removeList {
			mts.Remove(*e.BasicEntity)
		}
	}

	// fmt.Println(len(mts.entities))
}

// linear does a interpolation returns [0..1] floats
func linear(t0 time.Time, t1 time.Time, t time.Time) float32 {
	if t.After(t1) || t.Equal(t1) {
		return 1
	}

	if t.Before(t0) || t.Equal(t0) {
		return 0
	}

	return float32(float32(t.Sub(t0).Nanoseconds()) / float32(t1.Sub(t0).Nanoseconds()))
}
