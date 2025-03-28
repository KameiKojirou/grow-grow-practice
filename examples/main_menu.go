package examples

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/SceneTree"
	// "graphics.gd/classdb/SceneTree"
	// "graphics.gd/classdb/Scene"
	// "graphics.gd/classdb/ResourceLoader"
)

type MainMenu struct {
	classdb.Extension[MainMenu, Node.Instance] `gd:"MainMenu"`
	MyMenuLabel                                Label.Instance
	MyInputExampleButton                       Button.Instance
	MyCreateChildrenExampleButton              Button.Instance
}

func (m *MainMenu) Ready() {
	m.MyMenuLabel.AsLabel().SetText("Grow Examples")
	m.MyInputExampleButton.AsBaseButton().OnPressed(m.OnPressed)
	m.MyCreateChildrenExampleButton.AsBaseButton().OnPressed(m.OnPressedcc)
}

func (m *MainMenu) OnPressed() {
	var tree SceneTree.Instance = m.Super().GetTree()
	tree.ChangeSceneToFile("res://input_example.tscn")
}

func (m *MainMenu) OnPressedcc() {
	var tree SceneTree.Instance = m.Super().GetTree()
	tree.ChangeSceneToFile("res://create_children_example.tscn")
}
