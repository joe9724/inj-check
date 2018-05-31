package main

import (
	"net/http"
	"log"
	_"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"encoding/json"
	"inj-check/utils"
	"inj-check/model"
)


func main() {
	http.HandleFunc("/check/status", Status)
	http.HandleFunc("/check/do", Do)
	http.HandleFunc("/check/getCheckHistory",GetCheckHistory)
	err := http.ListenAndServe(":1066", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("start server at 0.0.0.0:1066")

}

func GetCheckHistory (w http.ResponseWriter, req *http.Request) {
	db,err := utils.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()

	var msg string
	var code int64
	var response utils.Response
	//签到,先检查是否已签到
	var list []model.CheckModel
	db.Raw("select id,user_id,`status`,DATE_FORMAT(create_time,'%Y-%m-%d %h:%m:%s') as create_time from btk_Check where user_id=? and DATE_SUB(CURDATE(), INTERVAL 31 DAY) <= date(create_time) order by id desc",utils.DeUserID(req.URL.Query().Get("euid"))).Find(&list)
	if len(list)==0{ //当天还没有签到，可以正常签到
		//db.Exec("insert into btk_Check(user_id) values(?)",utils.DeUserID(req.URL.Query().Get("euid")))
		msg = "31天内没有签到记录"
		code = 200
	}else{
		msg = "31天内签到记录"
		code = 300
		response.Data = list

	}
	for k:=0;k<len(list) ;k++  {
		list[k].RewardName = "5"
		list[k].RewardUnit = "积分"
	}
	response.Msg = msg
	response.Code = code

	//fmt.Println("start is",start)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "A Go Web Server")
	data, err := json.Marshal(response)
	if err != nil {
		log.Fatal("err get data: ", err)
	}

	w.Write(data)

}

func Status(w http.ResponseWriter, req *http.Request) {
	db,err := utils.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()
	var msg string
	var code int64
	var response utils.Response
	//签到,先检查是否已签到
	var list []model.CheckModel
	db.Raw("select id,user_id,create_time,`status` from btk_Check where user_id=? and  to_days(create_time) = to_days(now())",utils.DeUserID(req.URL.Query().Get("euid"))).Find(&list)
	if len(list)==0{ //当天还没有签到，可以正常签到
		//db.Exec("insert into btk_Check(user_id) values(?)",utils.DeUserID(req.URL.Query().Get("euid")))
		msg = "今日未签到"
		code = 200
	}else{
		msg = "今日已签到"
		code = 300
	}
	response.Msg = msg
	response.Code = code

	//fmt.Println("start is",start)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "A Go Web Server")
	data, err := json.Marshal(response)
	if err != nil {
		log.Fatal("err get data: ", err)
	}

	w.Write(data)


}

func Do(w http.ResponseWriter, req *http.Request) {
	db,err := utils.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()
	var msg string
	var code int64
	var response utils.Response
	//签到,先检查是否已签到
	var list []model.CheckModel
	db.Raw("select id,user_id,create_time,`status` from btk_Check where user_id=? and to_days(create_time) = to_days(now())",utils.DeUserID(req.URL.Query().Get("euid"))).Find(&list)
	fmt.Println("select id,user_id,create_time,`status` from btk_Check where user_id=? and to_days(create_time) = to_days(now())",utils.DeUserID(req.URL.Query().Get("euid")))
	if len(list)==0{ //当天还没有签到，可以正常签到
		db.Exec("insert into btk_Check(user_id) values(?)",utils.DeUserID(req.URL.Query().Get("euid")))
		msg = "签到成功"
		code = 200
	}else{
		msg = "今日已签到过了,亲"
		code = 300
	}
	response.Msg = msg
	response.Code = code

	//fmt.Println("start is",start)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "A Go Web Server")
	data, err := json.Marshal(response)
	if err != nil {
		log.Fatal("err get data: ", err)
	}

	w.Write(data)
}