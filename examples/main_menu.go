// This file is all you need to start a project.
// Save it somewhere, install the `gd` command and use `gd run` to get started.
package examples

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Node"
	// "graphics.gd/classdb/SceneTree"
)

type MainMenu struct {
	classdb.Extension[MainMenu, Node.Instance] `gd:"MainMenu"`
	MyInputExampleButton                           Button.Instance
}



func (m *MainMenu) Ready() {
	m.MyInputExampleButton.AsBaseButton().OnPressed(m.OnPressed)
}

func (m *MainMenu) OnPressed() {
	// SceneTree.AsSceneTree().ChangeScene("res://input_example.tscn")
}
