// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"hf_api/controllers"

	"github.com/astaxie/beego"
)

func init() {

	ns := beego.NewNamespace("/v1",

		// /v1/user
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		// /v1/recipe
		beego.NSNamespace("/recipe",
			beego.NSInclude(
				&controllers.RecipeController{},
			),
		),

		// /v1/ingredient
		beego.NSNamespace("/ingredient",
			beego.NSInclude(
				&controllers.IngredientController{},
			),
		),

		// /v1/menu
		beego.NSNamespace("/menu",
			beego.NSInclude(
				&controllers.MenuController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
