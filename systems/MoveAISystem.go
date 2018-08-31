package systems

import (
	"math"
	"math/rand"
	"towers/components"

	"engo.io/ecs"
	"engo.io/engo"
)

//MoveAIEntity system entity
type MoveAIEntity struct {
	*ecs.BasicEntity
	*components.MoveAIComponent
	*components.MoveCooldownComponent
	*components.MoveTweenComponent
}

//MoveAISystem mange move cooldown
type MoveAISystem struct {
	entities []MoveAIEntity
}

//New whne the system is created
func (mas *MoveAISystem) New(w *ecs.World) {

}

//Add new entity to the system
func (mas *MoveAISystem) Add(basic *ecs.BasicEntity, moveAi *components.MoveAIComponent, moveCooldown *components.MoveCooldownComponent, moveTween *components.MoveTweenComponent) {
	mas.entities = append(mas.entities, MoveAIEntity{basic, moveAi, moveCooldown, moveTween})
}

//Remove removes them from the system
func (mas *MoveAISystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range mas.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		mas.entities = append(mas.entities[:delete], mas.entities[delete+1:]...)
	}
}

//Update updates all entities
func (mas *MoveAISystem) Update(dt float32) {

	for _, e := range mas.entities {
		if !e.MoveCooldownComponent.Ready {

			e.MoveCooldownComponent.Ready = false

			if e.MoveAIComponent.Type == "RandomMover" {
				e.MoveTweenComponent.Tweening = true
				sin, cos := math.Sincos(rand.Float64())
				e.MoveTweenComponent.DestinationPosition = engo.Point{X: float32(float32(sin) * e.MoveTweenComponent.Range), Y: float32(float32(cos) * e.MoveTweenComponent.Range)}
			}

		}
	}

}
