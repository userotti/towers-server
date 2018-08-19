package systems

import (
	"fmt"
	"image/color"
	"time"
	"towers/components"
	"towers/entities"
	"towers/systems"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type TowerCreator struct {
}

// CreateTower make new towers
func (tc *TowerCreator) CreateTower(tb *systems.TowerCreateSystem) {

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

// Remove is called whenever an Entity is removed from the World, in order to remove it from this sytem as well
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

//LinearTweenEntity is an example of a entity that needs to tween
type LinearTweenEntity struct {
	*ecs.BasicEntity
	*components.TweenComponent
	*common.SpaceComponent
}

//LinearTweenSystem is responsible for the tweening of towers
type LinearTweenSystem struct {
	entities []LinearTweenEntity
}

//New whne the system is created
func (lts *LinearTweenSystem) New(w *ecs.World) {

}

//Add uAdd new entity
func (lts *LinearTweenSystem) Add(basic *ecs.BasicEntity, space *common.SpaceComponent, tween *components.TweenComponent) {
	lts.entities = append(lts.entities, LinearTweenEntity{basic, tween, space})
}

//Remove removes them from the system
func (lts *LinearTweenSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range lts.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		lts.entities = append(lts.entities[:delete], lts.entities[delete+1:]...)
	}
}

func linear(t0 time.Time, t1 time.Time, t time.Time) float32 {
	if t.After(t1) || t.Equal(t1) {
		return 1
	}

	if t.Before(t0) || t.Equal(t0) {
		return 0
	}

	return float32(float32(t.Sub(t0).Nanoseconds()) / float32(t1.Sub(t0).Nanoseconds()))
}

//Update updates all entities
func (lts *LinearTweenSystem) Update(dt float32) {

	var removeList []LinearTweenEntity

	for _, e := range lts.entities {

		linearResult := linear(e.TweenComponent.StartTime, e.TweenComponent.EndTime, time.Now())
		if e.TweenComponent.Tweening {
			e.SpaceComponent.Position.X = e.TweenComponent.StartPosition.X + (e.TweenComponent.DestinationPosition.X-e.TweenComponent.StartPosition.X)*linearResult
			e.SpaceComponent.Position.Y = e.TweenComponent.StartPosition.Y + (e.TweenComponent.DestinationPosition.Y-e.TweenComponent.StartPosition.Y)*linearResult

			if linearResult == 1 {
				removeList = append(removeList, e)
				e.TweenComponent.Tweening = false
			}
		}

	}

	if len(removeList) != 0 {
		for _, e := range removeList {
			lts.Remove(*e.BasicEntity)
		}
	}

	// fmt.Println(len(lts.entities))

}
