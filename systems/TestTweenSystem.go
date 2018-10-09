package systems

import (
	"fmt"
	"towers-server/components"

	"engo.io/ecs"
)

//TestTweenEnity system entity
type TestTweenEnity struct {
	*ecs.BasicEntity
	*components.TestTweenComponent
}

//TestTweenSystem mange move cooldown
type TestTweenSystem struct {
	entities []TestTweenEnity
}

//New whne the system is created
func (tts *TestTweenSystem) New(w *ecs.World) {

}

//Add new entity to the system
func (tts *TestTweenSystem) Add(basic *ecs.BasicEntity, move *components.TestTweenComponent) {
	tts.entities = append(tts.entities, TestTweenEnity{basic, move})
}

//Remove removes them from the system
func (tts *TestTweenSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range tts.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		tts.entities = append(tts.entities[:delete], tts.entities[delete+1:]...)
	}
}

//Update updates all entities
func (tts *TestTweenSystem) Update(dt float32) {

	for _, e := range tts.entities {
		if !e.TestTweenComponent.Charged {
			fmt.Println(dt)
			fmt.Println(e.TestTweenComponent.CurrentCharge)

			e.TestTweenComponent.CurrentCharge = e.TestTweenComponent.CurrentCharge + (e.TestTweenComponent.RechargeSpeed * dt)
			if e.TestTweenComponent.CurrentCharge >= e.TestTweenComponent.MaxCharge {
				e.TestTweenComponent.CurrentCharge = e.TestTweenComponent.MaxCharge
				e.TestTweenComponent.Charged = true
			}
		}
	}

}
