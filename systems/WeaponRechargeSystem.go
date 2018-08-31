package systems

import (
	"towers/components"

	"engo.io/ecs"
)

//WeaponEnity system entity
type WeaponEnity struct {
	*ecs.BasicEntity
	*components.WeaponComponent
}

//WeaponRechargeSystem mange move cooldown
type WeaponRechargeSystem struct {
	entities []WeaponEnity
}

//New whne the system is created
func (wrs *WeaponRechargeSystem) New(w *ecs.World) {

}

//Add new entity to the system
func (wrs *WeaponRechargeSystem) Add(basic *ecs.BasicEntity, move *components.WeaponComponent) {
	wrs.entities = append(wrs.entities, WeaponEnity{basic, move})
}

//Remove removes them from the system
func (wrs *WeaponRechargeSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range wrs.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		wrs.entities = append(wrs.entities[:delete], wrs.entities[delete+1:]...)
	}
}

//Update updates all entities
func (wrs *WeaponRechargeSystem) Update(dt float32) {

	for _, e := range wrs.entities {
		if !e.WeaponComponent.Loaded {
			e.WeaponComponent.CurrentCharge = e.WeaponComponent.CurrentCharge + e.WeaponComponent.RechargeSpeed
			if e.WeaponComponent.CurrentCharge >= e.WeaponComponent.MaxCharge {
				e.WeaponComponent.CurrentCharge = e.WeaponComponent.MaxCharge
				e.WeaponComponent.Loaded = true
			}
		}
	}

}
