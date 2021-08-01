package controller

import (
	"github.com/gin-gonic/gin"
	"lesson04/models"
	"net/http"
)


/**   url   -> controller ->  logic  ->   models
 *  请求来了 ->   控制器   -> 业务逻辑 -> 模型层的增删改查
 * @Description 访问首页
 * @Param
 * @return
 **/

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK,"index.html",nil)
}

func CreateMenu(c *gin.Context) {
	// 从请求中把数据取出来
	var menu models.Menu
	c.BindJSON(&menu)
	// 存入数据库
	err := models.InsertData(menu)
	// 响应
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"error" : err.Error(),
		})
	}else {
		c.JSON(http.StatusOK,menu)
	}

}

func GetAllMenus(c *gin.Context) {
	var menuList [] models.Menu
	// 查询所有数据
	err := models.FindAllData(&menuList)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"error" : err.Error(),
		})
	}else {
		c.JSON(http.StatusOK,menuList)
	}
}

func UpdateMenu(c *gin.Context) {
	id , ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK,gin.H{
			"error" : "id不存在",
		})
	}else {
		//查找相应的数据
		var menu models.Menu
		err := models.FindOneData(id, &menu)
		if err != nil{
			c.JSON(http.StatusOK,gin.H{
				"error" : err.Error(),
			})
		}else {
			// 将内容进行修改
			c.BindJSON(&menu)
			err := models.UpdateData(&menu)
			if err != nil {
				c.JSON(http.StatusOK,gin.H{
					"error" : err.Error(),
				})
			}else {
				c.JSON(http.StatusOK,menu)
			}
		}
	}
}

func DeleteMenu(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK,gin.H{
			"error" : "id不存在",
		})
	}else {
		err := models.DeleteDataById(id)
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"error" : err,
			})
		}else {
			c.JSON(http.StatusOK,gin.H{id : "deleted"})
		}
	}

}