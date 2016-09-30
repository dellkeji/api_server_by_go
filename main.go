package main

import (
	_ "bee_api/routers"

	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		fmt.Println("[DEV] start")
	} else {
		fmt.Println("[PROD] start")
		fmt.Println(beego.BConfig.Listen.HTTPAddr, beego.BConfig.Listen.HTTPPort)
	}
	var validate_app_code_secret = func(ctx *context.Context) {
		// 获取请求参数
		req_method := ctx.Request.Method
		var app_code string
		var app_secret string
		if req_method == "GET" {
			params_func := ctx.Input.Query
			app_code = params_func("app_code")
			app_secret = params_func("app_secret")
		} else if req_method == "POST" {
			var data_map interface{}
			json.Unmarshal(ctx.Input.RequestBody, &data_map)
			data := data_map.(map[string]interface{})
			for key, val := range data {
				switch val_val := val.(type) {
				case string:
					if key == "app_code" {
						app_code = val_val
					}
					if key == "app_secret" {
						app_secret = val_val
					}
				}
			}
		}
		error_response := map[string]interface{}{"result": false, "message": "没有权限访问当前接口，请检查app_code和app_secret"}
		if (app_code == "") && (app_secret == "") {
			ctx.Output.JSON(error_response, true, true)
		}
		// 校验app_code和app_secret在可访问的配置中
		global_app_code_secret := beego.AppConfig.String("global_app_code_secret::" + app_code)
		// fmt.Println(app_code, app_secret)
		if global_app_code_secret != app_secret {
			// error_response := map[string]interface{}{"result": false, "message": "没有权限访问当前接口"}
			//ctx.Output.Body([]byte("The app_code and app_secret are not allowed!"))
			ctx.Output.JSON(error_response, true, true)
		}
	}

	// 校验app_code和app_secret
	beego.InsertFilter("*", beego.BeforeRouter, validate_app_code_secret)

	beego.Run()
}
