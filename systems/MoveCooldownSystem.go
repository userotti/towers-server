package systems

import (
	"fmt"
	"towers/components"

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
	var entity = MoveCooldownEntity{basic, move}
	entity.MoveCooldownComponent.CalculatedDuration = (entity.MoveCooldownComponent.MaxCharge - entity.MoveCooldownComponent.MinCharge) / entity.MoveCooldownComponent.RechargeSpeed
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
		if !e.MoveCooldownComponent.Ready {
			e.MoveCooldownComponent.CurrentCharge = e.MoveCooldownComponent.CurrentCharge + (e.MoveCooldownComponent.RechargeSpeed * dt)
			fmt.Println(e.MoveCooldownComponent.CurrentCharge)
			if e.MoveCooldownComponent.CurrentCharge >= e.MoveCooldownComponent.MaxCharge {
				e.MoveCooldownComponent.CurrentCharge = e.MoveCooldownComponent.MaxCharge
				e.MoveCooldownComponent.Ready = true
			}
		} else {
			//Its ready
			e.MoveCooldownComponent.CalculatedDuration = (e.MoveCooldownComponent.MaxCharge - e.MoveCooldownComponent.MinCharge) / e.MoveCooldownComponent.RechargeSpeed

		}
	}

}
