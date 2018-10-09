package systems

import (
	"towers-server/components"

	"engo.io/ecs"
)

//ShieldEntity system entity
type ShieldEntity struct {
	*ecs.BasicEntity
	*components.ShieldComponent
}

//ShieldRechargeSystem mange move cooldown
type ShieldRechargeSystem struct {
	entities []ShieldEntity
}

//New whne the system is created
func (srs *ShieldRechargeSystem) New(w *ecs.World) {

}

//Add new entity to the system
func (srs *ShieldRechargeSystem) Add(basic *ecs.BasicEntity, move *components.ShieldComponent) {
	srs.entities = append(srs.entities, ShieldEntity{basic, move})
}

//Remove removes them from the system
func (srs *ShieldRechargeSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range srs.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		srs.entities = append(srs.entities[:delete], srs.entities[delete+1:]...)
	}
}

//Update updates all entities
func (srs *ShieldRechargeSystem) Update(dt float32) {

	for _, e := range srs.entities {
		if !e.ShieldComponent.ShieldUp {
			e.ShieldComponent.CurrentCharge = e.ShieldComponent.CurrentCharge + e.ShieldComponent.RechargeSpeed
			if e.ShieldComponent.CurrentCharge >= e.ShieldComponent.MaxCharge {
				e.ShieldComponent.CurrentCharge = e.ShieldComponent.MaxCharge
				e.ShieldComponent.ShieldUp = true
			}
		}
	}

}
