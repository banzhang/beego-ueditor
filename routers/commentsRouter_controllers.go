package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["blog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["blog/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["blog/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["blog/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["blog/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["blog/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:CategoryController"] = append(beego.GlobalControllerRouter["blog/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:CategoryController"] = append(beego.GlobalControllerRouter["blog/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:CategoryController"] = append(beego.GlobalControllerRouter["blog/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:CategoryController"] = append(beego.GlobalControllerRouter["blog/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:CategoryController"] = append(beego.GlobalControllerRouter["blog/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:UeditorController"] = append(beego.GlobalControllerRouter["blog/controllers:UeditorController"],
		beego.ControllerComments{
			Method: "Config",
			Router: `/config`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:UeditorController"] = append(beego.GlobalControllerRouter["blog/controllers:UeditorController"],
		beego.ControllerComments{
			Method: "Uploadimage",
			Router: `/uploadimage`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:UeditorController"] = append(beego.GlobalControllerRouter["blog/controllers:UeditorController"],
		beego.ControllerComments{
			Method: "Uploadscrawl",
			Router: `/uploadscrawl`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:UeditorController"] = append(beego.GlobalControllerRouter["blog/controllers:UeditorController"],
		beego.ControllerComments{
			Method: "Uploadvideo",
			Router: `/uploadvideo`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:UeditorController"] = append(beego.GlobalControllerRouter["blog/controllers:UeditorController"],
		beego.ControllerComments{
			Method: "Uploadfile",
			Router: `/uploadfile`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:UeditorController"] = append(beego.GlobalControllerRouter["blog/controllers:UeditorController"],
		beego.ControllerComments{
			Method: "Listimage",
			Router: `/listimage`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:UeditorController"] = append(beego.GlobalControllerRouter["blog/controllers:UeditorController"],
		beego.ControllerComments{
			Method: "Listfile",
			Router: `/listfile`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:UeditorController"] = append(beego.GlobalControllerRouter["blog/controllers:UeditorController"],
		beego.ControllerComments{
			Method: "Catchimage",
			Router: `/catchimage`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:UeditorController"] = append(beego.GlobalControllerRouter["blog/controllers:UeditorController"],
		beego.ControllerComments{
			Method: "Catchimage",
			Router: `/catchimage`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:UeditorController"] = append(beego.GlobalControllerRouter["blog/controllers:UeditorController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get","post"},
			Params: nil})

}
