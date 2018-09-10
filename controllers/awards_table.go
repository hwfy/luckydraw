package controllers

import (
	"luckyDraw/models"

	"encoding/json"
	"html/template"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// 奖项表控制器
type AwardsTableController struct {
	beego.Controller
}

// URLMapping ...
// @router / [options]
// @router /:id [options]
func (c *AwardsTableController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)

	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)

}

// Post ...
// @Title Post
// @Description 新增数据到奖项表
// @Param	body				body		models.AwardsTable	true		"body for AwardsTable content"
// @Success 200 {object} 	models.AwardsTable
// @Failure 400 参数错误
// @Failure 409 保存失败
// @router / [post]
func (c *AwardsTableController) Post() {
	var awards []*models.AwardsTable

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &awards); err == nil {
		if err = models.InsertAwards(awards); err == nil {
			c.Data["json"] = awards
		} else {
			c.Ctx.Output.SetStatus(409)
			c.Data["json"] = err.Error()
		}
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetDictionary ...
// @Title Get DATA Dictionary
// @Description 获取奖项表数据字典
// @Success 200 {object} 	models.DataDictionary
// @Failure 409 获取失败
// @router /data_dictionary [get]
func (c *AwardsTableController) GetDictionary() {
	awards := models.NewAwardsTable()

	data_dict, err := models.GetDataDictionary(awards.TableName())
	if err == nil {
		c.Data["json"] = data_dict
	} else {
		c.Ctx.Output.SetStatus(409)
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetLottery ...
// @Title Get Lottery Name
// @Description 获取当前抽奖名称
// @Success 200 {string} 抽奖名称
// @Failure 409 获取失败
// @router /lottery [get]
func (c *AwardsTableController) GetLottery() {
	name, err := models.GetAwardsName()
	if err == nil {
		c.Data["json"] = name
	} else {
		c.Ctx.Output.SetStatus(409)
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description 根据奖项名获取所有抽奖人员
// @Param	name				path 	string	true		"The key for staticblock"
// @Success 200 {object} 	[]models.LotteryStaffTemporaryTable
// @Failure 400 参数错误
// @Failure 409 获取失败
// @router /:name [get]
func (c *AwardsTableController) GetOne() {
	name := c.Ctx.Input.Param(":name")
	if name != "{name}" {
		v, err := models.GetAwardsTableByName(name)
		if err == nil {
			c.Data["json"] = v
		} else {
			c.Ctx.Output.SetStatus(409)
			c.Data["json"] = err.Error()
		}
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = "名称为空"
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description 获取奖项表所有数据
// @Param	query			query	string	false	"filter e.g. col1:v1,col2:v2 ..."
// @Param	joins			query	string	false	"joins e.g. inner join t1 on t1.col1=t.col,left join t2 on t2.col1=t.col"
// @Param	fields			query	string	false	"fields e.g. col1,col2 ..."
// @Param	sortby			query	string	false	"Sorted-by fields. e.g. col1 desc,col2 asc ..."
// @Param	limit			query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset			query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} 	[]models.AwardsTable
// @Failure 409 获取失败
// @router / [get]
func (c *AwardsTableController) GetAll() {
	var query = make(map[string]interface{})
	var fields = []string{"*"}
	var joins = []string{}
	var sortby string
	var limit int64 = 100
	var offset int64

	if v := c.GetString("joins"); v != "" {
		joins = strings.Split(v, ",")
	}
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// sortby: col1 desc,col2 asc
	if v := c.GetString("sortby"); v != "" {
		sortby = template.HTMLEscapeString(v)
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = "无效的查询符:" + kv[0] + "，正确格式：key:value"
				c.ServeJSON()
				return
			}
			query[kv[0]] = kv[1]
		}
	}
	l, err := models.GetAwardsTables(query, joins, fields, sortby, offset, limit)
	if err == nil {
		c.Data["json"] = l
	} else {
		c.Ctx.Output.SetStatus(409)
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description 根据id更新奖项表抽奖状态
// @Param	id				path 	string					true		"The id you want to update"
// @Param	body			body 	models.AwardsTable		true		"body for AwardsTable content"
// @Success 200 {string} OK
// @Failure 400 参数错误
// @Failure 409 更新失败
// @router /:id [put]
func (c *AwardsTableController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err == nil {
		v := models.AwardsTable{ID: id}
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
			if err := v.Update(); err == nil {
				c.Data["json"] = "OK"
			} else {
				c.Ctx.Output.SetStatus(409)
				c.Data["json"] = err.Error()
			}
		} else {
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = err.Error()
		}
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description 根据id删除奖项表数据
// @Param	id				path 	string	true		"The id you want to delete"
// @Success 200 {string} OK
// @Failure 400 参数错误
// @Failure 409 删除失败
// @router /:id [delete]
func (c *AwardsTableController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err == nil {
		v := models.AwardsTable{ID: id}
		if err := v.Delete(); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Ctx.Output.SetStatus(409)
			c.Data["json"] = err.Error()
		}
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
