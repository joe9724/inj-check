package main

import (
	"net/http"
	"log"
	_"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"encoding/json"
	"inj-init/utils"
	"inj-init/model"
)


func main() {
	http.HandleFunc("/startUp", StartUp)
	http.HandleFunc("/home", Home)
	err := http.ListenAndServe(":106", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("start server at 0.0.0.0:106")

}

func StartUp(w http.ResponseWriter, req *http.Request) {
	db,err := utils.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()
	var start model.StartUpModel
	db.Raw("select cover,coverx,`force`,number,msg,downloadUrl as download_url from btk_Version where client=?",req.URL.Query().Get("client")).First(&start)
	start.CoverX = "http://tingting-resource.bitekun.xin/resource/image/icon/phonex.jpg"
	start.ShareUrl = "http://www.qq.com"
	fmt.Println("start is",start)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "A Go Web Server")
	data, err := json.Marshal(start)
	if err != nil {
		log.Fatal("err get data: ", err)
	}

	w.Write(data)


}

func Home(w http.ResponseWriter, req *http.Request) {
	db,err := utils.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()
	var home model.Home
	//get banners
	var banners []model.Banner
	db.Raw("select id,title,type,cover,target_id,web_url,event_id from btk_Banner").Find(&banners)
	fmt.Println("banners is",banners)
	//get icons
	var icons []model.Icon
	db.Table("btk_Icon").Select("id, title,icon,target_id,web_url,type").Find(&icons)
	fmt.Println("icons is",icons)
	var cityBanner model.CityToutiao
	db.Raw("select id,icon,type,status,title,event_id,target_id,web_url,cover,style,sub_title from btk_Toutiao where style='banner' order by id desc limit 0,1").Find(&cityBanner)
	var cityTopList []model.CityToutiao
	db.Raw("select id,icon,type,status,title,target_id,web_url,cover,style,sub_title from btk_Toutiao  order by id desc limit 0,12 ").Find(&cityTopList)
	//get zones
	var zones []model.ZoneItem
	db.Raw("SELECT Cover as cover,ZoneID as zone_id, Name as name, Logo as logo, Brief as brief, MemberCount as member_count, Level as level, Tag as tag, CreateAt as create_at FROM btk_Zone WHERE Status = 0 ORDER BY CreateAt DESC").Limit(6).Offset(0).Find(&zones)
	fmt.Println("zones is",zones)
	//get goods 待开发
	//get activity
	var activities []model.Activity
    db.Raw("select tag,event_id,event_title,event_summary,event_thumb from t_event limit 0,4").Find(&activities)
    fmt.Println("activies is",activities)
    //get news
   /* var news []model.News
    db.Raw("select id,thumb,title,source,create_time,web_url from btk_News limit 0,6").Find(&news)
    fmt.Println("news is",news)*/

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "A Go Web Server")
	home.Banners = banners
	home.Zones = zones
	home.Icons = icons
	home.CityBanner = cityBanner
	home.CityTopList = cityTopList
	home.Activities = activities
	//home.News = news

	data, err := json.Marshal(home)
	if err != nil {
		log.Fatal("err get data: ", err)
	}

	w.Write(data)
}