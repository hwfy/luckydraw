package controllers

import (
	"luckyDraw/models"

	"encoding/json"
	"html/template"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// 人事基本资料控制器
type PersonnelBasicInformationController struct {
	beego.Controller
}

// URLMapping ...
// @router / [options]
func (c *PersonnelBasicInformationController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)

	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)

}

// Post ...
// @Title Post
// @Description 新增数据到人事基本资料
// @Param	import			query	bool								false	"whether to import data ?"
// @Param	body			body	models.PersonnelBasicInformation	true		"body for PersonnelBasicInformation content"
// @Success 200 {object} 	models.PersonnelBasicInformation
// @Failure 400 参数错误
// @Failure 409 保存失败
// @router / [post]
func (c *PersonnelBasicInformationController) Post() {
	isImport, _ := c.GetBool("import")
	if isImport {
		count, err := models.ImportPersonnelBasicInformation()
		if err == nil {
			c.Data["json"] = count
		} else {
			c.Ctx.Output.SetStatus(409)
			c.Data["json"] = err.Error()
		}
	} else {
		var v models.PersonnelBasicInformation

		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
			v.IsItPresent = true //默认在场
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
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description 根据id获取人事基本资料数据
// @Param	id				path 	string	true		"The key for staticblock"
// @Success 200 {object} 	models.PersonnelBasicInformation
// @Failure 400 参数错误
// @Failure 409 获取失败
// @router /:id [get]
func (c *PersonnelBasicInformationController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err == nil {
		v, err := models.GetPersonnelBasicInformationByPK(id)
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
// @Description 获取人事基本资料所有数据
// @Param	query			query	string	false	"filter e.g. col1:v1,col2:v2 ..."
// @Param	joins			query	string	false	"joins e.g. inner join t1 on t1.col1=t.col,left join t2 on t2.col1=t.col"
// @Param	fields			query	string	false	"fields e.g. col1,col2 ..."
// @Param	sortby			query	string	false	"Sorted-by fields. e.g. col1 desc,col2 asc ..."
// @Param	limit			query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset			query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} 	[]models.PersonnelBasicInformation
// @Failure 409 获取失败
// @router / [get]
func (c *PersonnelBasicInformationController) GetAll() {
	var query = make(map[string]interface{})
	var fields = []string{"*"}
	var joins = []string{}
	var sortby string
	var limit int64 = 1000
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
	l, err := models.GetPersonnelBasicInformations(query, joins, fields, sortby, offset, limit)
	if err == nil {
		c.Data["json"] = l
	} else {
		c.Ctx.Output.SetStatus(409)
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetDictionary ...
// @Title Get DATA Dictionary
// @Description 获取人事基本资料数据字典
// @Success 200 {object} 	models.DataDictionary
// @Failure 409 获取失败
// @router /data_dictionary [get]
func (c *PersonnelBasicInformationController) GetDictionary() {
	personnel := models.NewPersonnelBasicInformation()

	data_dict, err := models.GetDataDictionary(personnel.TableName())
	if err == nil {
		c.Data["json"] = data_dict
	} else {
		c.Ctx.Output.SetStatus(409)
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description 根据id更新人事基本资料数据
// @Param	id				path 	string						true		"The id you want to update"
// @Param	body				body 	models.PersonnelBasicInformation		true		"body for PersonnelBasicInformation content"
// @Success 200 {string} OK
// @Failure 400 参数错误
// @Failure 409 更新失败
// @router /:id [put]
func (c *PersonnelBasicInformationController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err == nil {
		v := models.PersonnelBasicInformation{ID: id}
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
// @Description 根据id删除人事基本资料数据
// @Param	id				path 	string	true		"The id you want to delete"
// @Success 200 {string} OK
// @Failure 400 参数错误
// @Failure 409 删除失败
// @router /:id [delete]
func (c *PersonnelBasicInformationController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err == nil {
		v := models.PersonnelBasicInformation{ID: id}
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
