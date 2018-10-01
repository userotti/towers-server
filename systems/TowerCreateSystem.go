package systems

import (
	"fmt"
	"image/color"
	"towers/components"
	"towers/entities"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

//MouseTracker tracker entity
type MouseTracker struct {
	ecs.BasicEntity
	common.MouseComponent
}

//TowerCreateSystem Create Towers here
type TowerCreateSystem struct {
	world        *ecs.World
	mouseTracker MouseTracker
}

// New is the initialisation of the System
func (tb *TowerCreateSystem) New(w *ecs.World) {
	fmt.Println("TowerCreateSystem was added to the Scene")

	tb.world = w

	tb.mouseTracker.BasicEntity = ecs.NewBasic()
	tb.mouseTracker.MouseComponent = common.MouseComponent{Track: true}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.MouseSystem:
			sys.Add(&tb.mouseTracker.BasicEntity, &tb.mouseTracker.MouseComponent, nil, nil)
		}
	}
}

//Remove is called whenever an Entity is removed from the World, in order to remove it from this sytem as well
func (*TowerCreateSystem) Remove(ecs.BasicEntity) {}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (tb *TowerCreateSystem) Update(dt float32) {
	if engo.Input.Button("AddTower").JustPressed() {
		fmt.Println("The gamer pressed d")
		// fmt.Println(dt)

		tower := entities.Tower{BasicEntity: ecs.NewBasic()}

		tower.SpaceComponent = common.SpaceComponent{
			Position: engo.Point{X: tb.mouseTracker.MouseX, Y: tb.mouseTracker.MouseY},
			Width:    10,
			Height:   10,
		}

		tower.RenderComponent = common.RenderComponent{
			Drawable: common.Circle{
				BorderColor: color.White,
				BorderWidth: 0,
			},
		}

		tower.MoveTweenComponent = components.MoveTweenComponent{
			StartPosition: tower.SpaceComponent.Position,
			Speed:         400, //Travels 400 units per second.
			Range:         150, //units of distance
			Done:          true,
		}

		tower.MoveCooldownComponent = components.MoveCooldownComponent{
			Recharge: 200, // Recharges 200 unit per second
			Cooldown: 100,
			Done:     false,
		}

		tower.MoveAIComponent = components.MoveAIComponent{
			Type: "RandomMover",
		}

		tower.WeaponComponent = components.WeaponComponent{
			Recharge: 100, // Recharges 100 unit per second
			Cooldown: 100,
			Damage:   1,
			Range:    100,
			Loaded:   false,
		}

		for _, system := range tb.world.Systems() {
			switch sys := system.(type) {
			case *common.RenderSystem:
				sys.Add(&tower.BasicEntity, &tower.RenderComponent, &tower.SpaceComponent)
			case *MoveTweenSystem:
				sys.Add(&tower.BasicEntity, &tower.SpaceComponent, &tower.MoveTweenComponent)
			case *MoveCooldownSystem:
				sys.Add(&tower.BasicEntity, &tower.MoveCooldownComponent)
			case *MoveAISystem:
				sys.Add(&tower.BasicEntity, &tower.MoveAIComponent, &tower.MoveCooldownComponent, &tower.MoveTweenComponent, &tower.SpaceComponent)
			case *WeaponRechargeSystem:
				sys.Add(&tower.BasicEntity, &tower.WeaponComponent)
			}

		}

	}
}
