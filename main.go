package main

import (
	"fmt"

	"github.com/bitly/go-simplejson"
)

//https://github.com/bitly/go-simplejson.git
func main() {
	jsonValue := `{"err_code":"ORDERREFUND",
        "err_code_des":"已经退款的订单不允许撤销","mch_id":"1293941701",
        "nonce_str":"E84YHlmuEKvQJeI9","recall":"N","result_code":"FAIL",
        "return_code":"SUCCESS","return_msg":"OK"}`
	js, _ := simplejson.NewJson([]byte(jsonValue))
	recall, _ := js.Get("recall").String()
	fmt.Println(recall)

	err_code_des, _ := js.Get("err_code_des").String()
	fmt.Println(err_code_des)

}
