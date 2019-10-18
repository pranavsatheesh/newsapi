package controllers

type NewsController struct {
	BaseController
}
type Data struct {
	Age     int
	Code    int
	Name    string
	Address string
}

func (c *NewsController) NewsRead() {
	// w := c.Ctx.ResponseWriter

	data := Data{Age: 42, Code: 200, Name: "Pranav", Address: "sasdad"}
	// slice := []interface{}{data}
	// sliceToClient, _ := json.Marshal(slice[0])
	// w.Write(sliceToClient)
	c.Data["json"] = data
	c.ServeJSON()
}
