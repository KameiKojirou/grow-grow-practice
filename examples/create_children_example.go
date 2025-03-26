package examples

import (
	"fmt"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/VBoxContainer"
)

var count int = 0

type CreateChildrenExample struct {
	classdb.Extension[InputExample, Control.Instance] `gd:"CreateChildrenExample"`
	MyContainer                                       VBoxContainer.Instance
	MyButton                                          Button.Instance
}

func (c *CreateChildrenExample) Ready() {
	c.MyButton.AsBaseButton().OnPressed(c.OnPressed)
}

func (c *CreateChildrenExample) OnPressed() {
	count++
	newlabel := Label.New()
	newlabel.SetText(fmt.Sprintf("Child %d", count))
	c.MyContainer.AsNode().AddChild(newlabel.AsNode())
}
