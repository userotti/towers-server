package systems

import (
	"math"
	"math/rand"
	"time"
	"towers/components"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

//MoveAIEntity system entity
type MoveAIEntity struct {
	*ecs.BasicEntity
	*components.AIComponent
	*components.MoveCooldownComponent
	*components.MoveTweenComponent
	*common.SpaceComponent
}

//MoveAISystem mange move cooldown
type MoveAISystem struct {
	entities []MoveAIEntity
}

//New whne the system is created
func (mas *MoveAISystem) New(w *ecs.World) {

}

//Add new entity to the system
func (mas *MoveAISystem) Add(basic *ecs.BasicEntity, moveAi *components.AIComponent, moveCooldown *components.MoveCooldownComponent, moveTween *components.MoveTweenComponent, space *common.SpaceComponent) {
	mas.entities = append(mas.entities, MoveAIEntity{basic, moveAi, moveCooldown, moveTween, space})
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
		if e.MoveCooldownComponent.Done && e.MoveTweenComponent.Done {

			//Now you're tweening
			e.MoveCooldownComponent.Done = false
			e.MoveCooldownComponent.StartTime = time.Now()
			e.MoveCooldownComponent.NanosecondsFromStart = 0

			e.MoveTweenComponent.Done = false
			e.MoveTweenComponent.StartTime = time.Now()
			e.MoveTweenComponent.NanosecondsFromStart = 0

			//But where are you tweening to? ...
			if e.AIComponent.Type == components.Crazy {
				sin, cos := math.Sincos((rand.Float64() * (2 * math.Pi)) - math.Pi)

				var differenceVector = engo.Point{X: float32(sin * e.MoveTweenComponent.Range), Y: float32(cos * e.MoveTweenComponent.Range)}
				e.MoveTweenComponent.DestinationPosition.X = e.SpaceComponent.Position.X + differenceVector.X
				e.MoveTweenComponent.DestinationPosition.Y = e.SpaceComponent.Position.Y + differenceVector.Y

				// fmt.Print("X:")
				// fmt.Print(float32(float32(sin) * e.MoveTweenComponent.Range))
				// fmt.Println()
				// fmt.Print("Y:")
				// fmt.Print(float32(float32(cos) * e.MoveTweenComponent.Range))
				// fmt.Println()
			}

		}
	}

}
