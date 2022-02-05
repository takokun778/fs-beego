package routers

import (
	"fs-beego/controllers"
	v1 "fs-beego/controllers/v1"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	nsv1 := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			beego.NSRouter("/users/", &v1.UserController{}, "get:GetAll;post:Post"),
			beego.NSRouter("/users/:id", &v1.UserController{}, "get:GetOne;put:Put;delete:Delete"),
		))

	beego.AddNamespace(nsv1)
}
