package main

import (
	"./systems"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type myScene struct{}

// Type uniquely defines your game type
func (*myScene) Type() string { return "myGame" }

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (*myScene) Preload() {}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (*myScene) Setup(u engo.Updater) {
	engo.Input.RegisterButton("AddTower", engo.KeyD)
	world, _ := u.(*ecs.World)
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})
	world.AddSystem(&systems.TowerCreateSystem{})
	world.AddSystem(&systems.MoveTweenSystem{})
	world.AddSystem(&systems.MoveCooldownSystem{})
	world.AddSystem(&systems.MoveAISystem{})
	world.AddSystem(&systems.WeaponRechargeSystem{})

	// world.AddSystem(&systems.TestTweenSystem{})

}

func main() {
	opts := engo.RunOptions{
		Title:  "Towers",
		Width:  800,
		Height: 800,
	}
	engo.Run(opts, &myScene{})
}
