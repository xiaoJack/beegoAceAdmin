package controllers

type IndexController struct {
	BaseController
}


// URLMapping ...
func (c *IndexController) URLMapping() {
	c.Mapping("get", c.Index)
}


// @router / [get]
func (this *IndexController) Index() {
	this.Data["pageTitle"] = "系统概况"
	this.display()
}


