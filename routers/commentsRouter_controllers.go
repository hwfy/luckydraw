package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsConditionsTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsConditionsTableController"],
		beego.ControllerComments{
			Method: "URLMapping",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsConditionsTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsConditionsTableController"],
		beego.ControllerComments{
			Method: "URLMapping",
			Router: `/:id`,
			AllowHTTPMethods: []string{"options"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsConditionsTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsConditionsTableController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsConditionsTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsConditionsTableController"],
		beego.ControllerComments{
			Method: "GetDictionary",
			Router: `/data_dictionary`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsConditionsTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsConditionsTableController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsConditionsTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsConditionsTableController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsConditionsTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsConditionsTableController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsConditionsTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsConditionsTableController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsStageTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsStageTableController"],
		beego.ControllerComments{
			Method: "URLMapping",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsStageTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsStageTableController"],
		beego.ControllerComments{
			Method: "URLMapping",
			Router: `/:id`,
			AllowHTTPMethods: []string{"options"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsStageTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsStageTableController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsStageTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsStageTableController"],
		beego.ControllerComments{
			Method: "GetDictionary",
			Router: `/data_dictionary`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsStageTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsStageTableController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsStageTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsStageTableController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsStageTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsStageTableController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsStageTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsStageTableController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsTableController"],
		beego.ControllerComments{
			Method: "URLMapping",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsTableController"],
		beego.ControllerComments{
			Method: "URLMapping",
			Router: `/:id`,
			AllowHTTPMethods: []string{"options"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsTableController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsTableController"],
		beego.ControllerComments{
			Method: "GetDictionary",
			Router: `/data_dictionary`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsTableController"],
		beego.ControllerComments{
			Method: "GetLottery",
			Router: `/lottery`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsTableController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:name`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsTableController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsTableController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:AwardsTableController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:AwardsTableController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:PersonnelBasicInformationController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:PersonnelBasicInformationController"],
		beego.ControllerComments{
			Method: "URLMapping",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:PersonnelBasicInformationController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:PersonnelBasicInformationController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:PersonnelBasicInformationController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:PersonnelBasicInformationController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:PersonnelBasicInformationController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:PersonnelBasicInformationController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:PersonnelBasicInformationController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:PersonnelBasicInformationController"],
		beego.ControllerComments{
			Method: "GetDictionary",
			Router: `/data_dictionary`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:PersonnelBasicInformationController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:PersonnelBasicInformationController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:PersonnelBasicInformationController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:PersonnelBasicInformationController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:WinningFormController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:WinningFormController"],
		beego.ControllerComments{
			Method: "URLMapping",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:WinningFormController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:WinningFormController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:WinningFormController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:WinningFormController"],
		beego.ControllerComments{
			Method: "GetDictionary",
			Router: `/data_dictionary`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:WinningFormController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:WinningFormController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:WinningFormController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:WinningFormController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:WinningFormController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:WinningFormController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["luckyDraw/controllers:WinningFormController"] = append(beego.GlobalControllerRouter["luckyDraw/controllers:WinningFormController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
