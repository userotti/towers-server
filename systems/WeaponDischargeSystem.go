package systems

import (
	"../components"

	"engo.io/ecs"
	"engo.io/engo/common"
)

//WeaponAIEntity system entity
type WeaponAIEntity struct {
	*ecs.BasicEntity
	*components.MoveTweenComponent
	*common.SpaceComponent
	*components.AIComponent
	*components.TeamComponent
	*components.WeaponComponent
}

//WeaponAISystem mange move cooldown
type WeaponAISystem struct {
	entities []WeaponAIEntity
}

//New whne the system is created
func (wds *WeaponAISystem) New(w *ecs.World) {

}

//Add new entity to the system
func (wds *WeaponAISystem) Add(basic *ecs.BasicEntity, moveTween *components.MoveTweenComponent, space *common.SpaceComponent, ai *components.AIComponent, team *components.TeamComponent, weapon *components.WeaponComponent) {
	wds.entities = append(wds.entities, WeaponAIEntity{basic, moveTween, space, ai, team, weapon})
}

//Remove removes them from the system
func (wds *WeaponAISystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range wds.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		wds.entities = append(wds.entities[:delete], wds.entities[delete+1:]...)
	}
}

//Update updates all entities
func (wds *WeaponAISystem) Update(dt float32) {

	for _, shooter := range wds.entities {

		if shooter.AIComponent.Type == components.Crazy {
			// var targets []WeaponAIEntity
			for _, other := range wds.entities {
				if other.TeamComponent.Name != shooter.TeamComponent.Name && shooter.WeaponComponent.Loaded {
					// append(targets, other)
				}
			}
		}

	}

}
