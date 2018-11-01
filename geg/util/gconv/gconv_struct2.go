package main

import (
    "fmt"
    "gitee.com/johng/gf/g"
    "gitee.com/johng/gf/g/util/gconv"
)


// 使用默认映射规则绑定属性值到对象
func main() {
    type User struct {
        Uid   int
        Name  string
        Pass1 string
        Pass2 string
    }
    user    := new(User)
    params  := g.Map {
        "uid"   : 1,
        "Name"  : "john",
        "PASS1" : "123",
        "PASS2" : "456",
    }
    if err := gconv.Struct(params, user); err == nil {
        fmt.Println(user)
    }
}