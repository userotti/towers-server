package systems

import (
	"fmt"
	"image/color"
	"time"
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
		fmt.Println(dt)

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

		tower.TweenComponent = components.TweenComponent{
			StartPosition:       tower.SpaceComponent.Position,
			DestinationPosition: engo.Point{X: 500, Y: 500},
			StartTime:           time.Now(),
			EndTime:             time.Now().Add(time.Duration(2 * time.Second)),
			Tweening:            true,
		}
		//
		// fmt.Println(tower.TweenComponent.StartTime.String())
		// fmt.Println(tower.TweenComponent.EndTime.String())
		// fmt.Println(time.Duration(2 * time.Second).Nanoseconds())
		// fmt.Println(tower.TweenComponent.EndTime.Sub(tower.TweenComponent.StartTime).Nanoseconds())

		for _, system := range tb.world.Systems() {
			switch sys := system.(type) {
			case *common.RenderSystem:
				sys.Add(&tower.BasicEntity, &tower.RenderComponent, &tower.SpaceComponent)
			case *LinearTweenSystem:
				sys.Add(&tower.BasicEntity, &tower.SpaceComponent, &tower.TweenComponent)
			}

		}

	}
}
