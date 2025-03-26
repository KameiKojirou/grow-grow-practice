// This file is all you need to start a project.
// Save it somewhere, install the `gd` command and use `gd run` to get started.
package main

import (
	"graphics.gd/classdb"
	"graphics.gd/startup"
	"main/examples"
)


func main() {
	// startup.LoadingScene() // setup the SceneTree and wait until we have access to engine functionality
	classdb.Register[examples.InputExample]()
	classdb.Register[examples.MainMenu]()
	classdb.Register[examples.CreateChildrenExample]()
	startup.Scene() // starts up the scene and blocks until the engine shuts down.
}
