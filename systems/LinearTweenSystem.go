package systems

import (
	"time"
	"towers/components"

	"engo.io/ecs"
	"engo.io/engo/common"
)

//LinearTweenEntity is an example of a entity that needs to tween
type LinearTweenEntity struct {
	*ecs.BasicEntity
	*components.TweenComponent
	*common.SpaceComponent
}

//LinearTweenSystem is responsible for the tweening of towers
type LinearTweenSystem struct {
	entities []LinearTweenEntity
}

//New whne the system is created
func (lts *LinearTweenSystem) New(w *ecs.World) {
}

//Add uAdd new entity
func (lts *LinearTweenSystem) Add(basic *ecs.BasicEntity, space *common.SpaceComponent, tween *components.TweenComponent) {
	lts.entities = append(lts.entities, LinearTweenEntity{basic, tween, space})
}

//Remove removes them from the system
func (lts *LinearTweenSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range lts.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		lts.entities = append(lts.entities[:delete], lts.entities[delete+1:]...)
	}
}

//Update updates all entities
func (lts *LinearTweenSystem) Update(dt float32) {

	var removeList []LinearTweenEntity

	for _, e := range lts.entities {

		linearResult := linear(e.TweenComponent.StartTime, e.TweenComponent.EndTime, time.Now())
		if e.TweenComponent.Tweening {
			e.SpaceComponent.Position.X = e.TweenComponent.StartPosition.X + (e.TweenComponent.DestinationPosition.X-e.TweenComponent.StartPosition.X)*linearResult
			e.SpaceComponent.Position.Y = e.TweenComponent.StartPosition.Y + (e.TweenComponent.DestinationPosition.Y-e.TweenComponent.StartPosition.Y)*linearResult

			if linearResult == 1 {
				removeList = append(removeList, e)
				e.TweenComponent.Tweening = false
			}
		}

	}

	if len(removeList) != 0 {
		for _, e := range removeList {
			lts.Remove(*e.BasicEntity)
		}
	}

	// fmt.Println(len(lts.entities))
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
