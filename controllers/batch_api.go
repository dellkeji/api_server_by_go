package controllers

import (
	"bee_api/redis_lib"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"reflect"
	"regexp"
)

type SetProperty struct {
	beego.Controller
}

// 获取请求参数
func get_params(params interface{}) []string {
	var params_list []string
	switch val_val := params.(type) {
	case string:
		reg := regexp.MustCompile(`[^,;]+`)
		params_list = reg.FindAllString(val_val, -1)
		fmt.Println(reflect.TypeOf(params_list))
		fmt.Println(params_list[1])
	}
	return params_list
}

func insert_data(slice, insertion []interface{}, index int) []interface{} {
	result := make([]interface{}, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}

func get_redis_data(key string, field []string) []string {
	rc := redis_lib.RedisClient.Get()
	defer rc.Close()
	err := rc.Err()
	if err != nil {
		fmt.Println("link redis error", err)
	}
	fmt.Println(field)
	fmt.Println("test")
	args := make([]interface{}, len(field)+1)
	args[0] = "pet"
	for index, val := range field {
		args[index+1] = val
	}
	val, err := redis.Strings(rc.Do("HMGET", args...))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}
	return val
}

func (o *SetProperty) Get() {
	params := o.GetString("set_ids")
	params_list := get_params(params)
	fmt.Println(params_list)
	ret_data := get_redis_data("SetNameAppIDToSetID", params_list)
	fmt.Println(ret_data)
	index := map[string]interface{}{"result": true, "data": ret_data}
	o.Data["json"] = index
	o.ServeJSON()
}

func (o *SetProperty) Post() {
	index := map[string]interface{}{"result": true, "message": "欢迎使用CC Cache API！"}
	o.Data["json"] = index
	o.ServeJSON()
}
