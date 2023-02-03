package chapter11

import "github.com/gin-gonic/gin"

func Router(chap11 *gin.RouterGroup) {
	chap11.GET("/test_axios", TestAxios)
	chap11.GET("/test_string", User)
	chap11.GET("/test_struct", UserInfoStruct)
	chap11.GET("/test_arr", ArrController)
	chap11.GET("/test_arr_struct", ArrStructController)
	chap11.GET("/test_map", MapController)
	chap11.GET("/test_map_map", MapMapController)
	chap11.GET("/test_map_struct", MapStructController)
	chap11.GET("/test_slice", SliceController)
	chap11.GET("/test_slice_struct", SliceStructController)
}
