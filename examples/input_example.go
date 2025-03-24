// This file is all you need to start a project.
// Save it somewhere, install the `gd` command and use `gd run` to get started.
package examples

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/LineEdit"
	"graphics.gd/classdb/Node"
)

type InputExample struct {
	classdb.Extension[InputExample, Node.Instance] `gd:"InputExample"`
	MyLabel                                Label.Instance
	MyButton                               Button.Instance
	MyInput                                LineEdit.Instance
}

var TextValue string = ""

func (i *InputExample) Ready() {
	i.MyInput.AsLineEdit().SetPlaceholderText("Enter your name")
	i.MyLabel.AsLabel().SetText("Fill out your name")
	i.MyButton.AsButton().SetText("Click me!")
	i.MyButton.AsBaseButton().OnPressed(i.OnPressed)
	i.MyInput.AsLineEdit().OnTextChanged(i.OnTextChanged)
}

func (i *InputExample) OnTextChanged(text string) {
	TextValue = text
}

func (m *InputExample) OnPressed() {
	m.MyLabel.AsLabel().SetText("Hello " + TextValue)
}

