// @APIVersion 1.0.0
// @Title 抽奖项目 API
// @Description 人事基础资料、奖项表、中奖表、概率表管理
// @Contact luckyfanyang@gmail.com
package routers

import (
	"luckyDraw/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	ns := beego.NewNamespace("/v1/luckydraw",
		beego.NSNamespace("/awards_table",
			beego.NSInclude(
				&controllers.AwardsTableController{},
			),
		),

		beego.NSNamespace("/personnel_basic_information",
			beego.NSInclude(
				&controllers.PersonnelBasicInformationController{},
			),
		),

		beego.NSNamespace("/winning_form",
			beego.NSInclude(
				&controllers.WinningFormController{},
			),
		),
		beego.NSNamespace("/awards_conditions_table",
			beego.NSInclude(
				&controllers.AwardsConditionsTableController{},
			),
		),

		beego.NSNamespace("/awards_stage_table",
			beego.NSInclude(
				&controllers.AwardsStageTableController{},
			),
		),
	)
	beego.AddNamespace(ns.Filter("before", func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
		ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type")
		ctx.Output.Header("Access-Control-Allow-Methods", "DELETE,PUT")
	}))
}
