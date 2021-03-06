package entities

import (
	"towers-server/components"

	"engo.io/ecs"
	"engo.io/engo/common"
)

//Tower base entity for the simulation
type Tower struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	components.AIComponent
	components.MoveCooldownComponent
	components.MoveTweenComponent
	components.TeamComponent
	components.WeaponComponent
	components.ShieldComponent
	components.HitpointsComponent
}

//Rock destructable object
// type Rock struct {
// 	ecs.BasicEntity
// 	common.RenderComponent
// 	common.SpaceComponent
//
// }
