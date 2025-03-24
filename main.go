// This file is all you need to start a project.
// Save it somewhere, install the `gd` command and use `gd run` to get started.
package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/LineEdit"
	"graphics.gd/classdb/Node"
	"graphics.gd/startup"
)

type Main struct {
	classdb.Extension[Main, Node.Instance] `gd:"MainThing"`
	MyLabel                                Label.Instance
	MyButton                               Button.Instance
	MyInput                                LineEdit.Instance
}

var TextValue string = ""

func (m *Main) Ready() {
	m.MyInput.AsLineEdit().SetPlaceholderText("Enter your name")
	m.MyLabel.AsLabel().SetText("Fill out your name")
	m.MyButton.AsButton().SetText("Click me!")
	m.MyButton.AsBaseButton().OnPressed(m.OnPressed)
	m.MyInput.AsLineEdit().OnTextChanged(m.OnTextChanged)
}

func (m *Main) OnTextChanged(text string) {
	TextValue = text
}

func (m *Main) OnPressed() {
	m.MyLabel.AsLabel().SetText("Hello " + TextValue)
}

func main() {
	// startup.LoadingScene() // setup the SceneTree and wait until we have access to engine functionality
	classdb.Register[Main]()
	startup.Scene() // starts up the scene and blocks until the engine shuts down.
}
