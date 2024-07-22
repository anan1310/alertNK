package response

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"

	"net/http"
)

// TODO: 返回0为正常信息 返回1为错误信息
func returnJson(Context *gin.Context, httpCode int, dataCode int, msg string, data interface{}) {

	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	Context.JSON(httpCode, gin.H{
		"code": dataCode,
		"msg":  msg,
		"data": data,
	})
}

// ReturnJsonFromString 将json字符窜以标准json格式返回（例如，从redis读取json、格式的字符串，返回给浏览器json格式）
func ReturnJsonFromString(Context *gin.Context, httpCode int, jsonStr string) {
	Context.Header("Content-Type", "application/json; charset=utf-8")
	Context.String(httpCode, jsonStr)
}

func returnJsonTotal(Context *gin.Context, httpCode int, dataCode int, total int64, msg string, data interface{}) {

	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	Context.JSON(httpCode, gin.H{
		"code":  dataCode,
		"msg":   msg,
		"data":  data,
		"total": total,
	})
}

// SuccessTotal 返回带有统计的数据
func SuccessTotal(c *gin.Context, msg string, total int64, data interface{}) {
	returnJsonTotal(c, http.StatusOK, 0, total, msg, data)
}

// Success 返回成功信息
func Success(c *gin.Context, msg string, data interface{}) {
	returnJson(c, http.StatusOK, 0, msg, data)
}

// Fail 返回失败的信息
func Fail(c *gin.Context, msg string, data interface{}) {
	returnJson(c, http.StatusOK, 1, msg, data)
	c.Abort()
}

func ErrPrintln(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func Service(ctx *gin.Context, fu func() (interface{}, interface{})) {
	data, err := fu()
	if err != nil {
		Fail(ctx, err.(error).Error(), "failed")
		ctx.Abort()
		return
	} else {
		Success(ctx, "success", data)
	}

}

func ServiceTotal(ctx *gin.Context, fu func() (interface{}, interface{}, interface{})) {
	data, total, err := fu()
	if err != nil {
		Fail(ctx, err.(error).Error(), "failed")
		ctx.Abort()
		return
	} else {
		SuccessTotal(ctx, "success", total.(int64), data)
	}
}

func BindJson(ctx *gin.Context, req interface{}) error {
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		Fail(ctx, err.Error(), "failed")
		ctx.Abort()
		return err
	}
	return nil
}

func BindQuery(ctx *gin.Context, req interface{}) error {
	err := ctx.ShouldBindQuery(req)
	if err != nil {
		Fail(ctx, err.Error(), "failed")
		ctx.Abort()
		return err
	}
	return nil
}

func TokenFail(ctx *gin.Context) {
	code := 401
	returnJson(ctx, code, code, "鉴权失败", "")
}
