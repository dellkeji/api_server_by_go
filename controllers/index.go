package controllers

import (
	"bee_api/redis_lib"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

type IndexController struct {
	beego.Controller
}

func (o *IndexController) Get() {
	rc := redis_lib.RedisClient.Get()
	err := rc.Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := redis.String(rc.Do("GET", "username"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
		fmt.Println(val)
	}
	// close the link
	defer rc.Close()
	index := map[string]interface{}{"result": true, "message": "欢迎使用CC Cache API！"}
	o.Data["json"] = index
	o.ServeJSON()
}
