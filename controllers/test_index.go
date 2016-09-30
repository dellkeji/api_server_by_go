package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type TestIndex struct {
	beego.Controller
}

func (t *TestIndex) test_index() {
	fmt.Println("test")
}
