package chapter07

import "github.com/gin-gonic/gin"

func Router(chap07 *gin.RouterGroup) {
	/*------------------------------第七章-------------------------------*/

	/*------------------------------初始化、DDL-------------------------------*/

	// 创建表
	chap07.GET("/create_table", CreateTable)

	// 删除表
	chap07.GET("/del_table", DeleteTable)

	// 表是否存在
	chap07.GET("/has_table", HasTable)

	// 添加数据
	chap07.GET("/add_data", AddData)

	// 查询数据
	chap07.GET("/search_data", SearchData)

	// 更新数据 [先查后改]
	chap07.GET("/update_data", UpdateData)

	// 删：先查再改
	chap07.GET("/delete_data", DeleteData)

	// 外键：一对一
	chap07.GET("/foreign_key", ForeignKey)

}
