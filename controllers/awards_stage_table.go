package controllers

import (
	"luckyDraw/models"

	"encoding/json"
	"html/template"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// 奖项阶段表控制器
type AwardsStageTableController struct {
	beego.Controller
}

// URLMapping ...
// @router / [options]
// @router /:id [options]
func (c *AwardsStageTableController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)

	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)

}

// Post ...
// @Title Post
// @Description 新增数据到奖项阶段表
// @Param	body				body		models.AwardsStageTable	true		"body for AwardsStageTable content"
// @Success 200 {object} 	models.AwardsStageTable
// @Failure 400 参数错误
// @Failure 409 保存失败
// @router / [post]
func (c *AwardsStageTableController) Post() {
	var v models.AwardsStageTable

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err = v.Insert(); err == nil {
			c.Data["json"] = v
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
// @Description 获取奖项阶段数据字典
// @Success 200 {object} 	models.DataDictionary
// @Failure 409 获取失败
// @router /data_dictionary [get]
func (c *AwardsStageTableController) GetDictionary() {
	awards := models.NewAwardsStageTable()

	data_dict, err := models.GetDataDictionary(awards.TableName())
	if err == nil {
		c.Data["json"] = data_dict
	} else {
		c.Ctx.Output.SetStatus(409)
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description 根据id获取奖项阶段表数据
// @Param	id				path 	string	true		"The key for staticblock"
// @Success 200 {object} 	models.AwardsStageTable
// @Failure 400 参数错误
// @Failure 409 获取失败
// @router /:id [get]
func (c *AwardsStageTableController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err == nil {
		v, err := models.GetAwardsStageTableByPK(id)
		if err == nil {
			c.Data["json"] = v
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

// GetAll ...
// @Title Get All
// @Description 获取奖项阶段表所有数据
// @Param	query			query	string	false	"filter e.g. col1:v1,col2:v2 ..."
// @Param	joins			query	string	false	"joins e.g. inner join t1 on t1.col1=t.col,left join t2 on t2.col1=t.col"
// @Param	fields			query	string	false	"fields e.g. col1,col2 ..."
// @Param	sortby			query	string	false	"Sorted-by fields. e.g. col1 desc,col2 asc ..."
// @Param	limit			query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset			query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} 	[]models.AwardsStageTable
// @Failure 409 获取失败
// @router / [get]
func (c *AwardsStageTableController) GetAll() {
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
	l, err := models.GetAwardsStageTables(query, joins, fields, sortby, offset, limit)
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
// @Description 根据id更新奖项阶段表数据
// @Param	id				path 	string						true		"The id you want to update"
// @Param	body				body 	models.AwardsStageTable		true		"body for AwardsStageTable content"
// @Success 200 {string} OK
// @Failure 400 参数错误
// @Failure 409 更新失败
// @router /:id [put]
func (c *AwardsStageTableController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err == nil {
		v := models.AwardsStageTable{ID: id}
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
// @Description 根据id删除奖项阶段表数据
// @Param	id				path 	string	true		"The id you want to delete"
// @Success 200 {string} OK
// @Failure 400 参数错误
// @Failure 409 删除失败
// @router /:id [delete]
func (c *AwardsStageTableController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err == nil {
		v := models.AwardsStageTable{ID: id}
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
