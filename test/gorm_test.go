package test

import (
	"bosom_friend/models"
	"fmt"
	"testing"
)

func TestGorm(t *testing.T) {

	count, activitys := models.GetActivityTest("1", "1")
	//count := tx.(map[string] interface{})["count"]
	fmt.Println("xxxxxxxxxxx")
	fmt.Println(count)
	fmt.Println(activitys)
	fmt.Println("xxxxxxxxxxx")

	//db.Model(&models.Activity{}).Where("number = 2").Count(&count).Scan(&activity)
	//fmt.Println(activity)
	//fmt.Println(count)
	//dict := map[string] interface {}{"test":1}
	//dict["test2"] = "test"
	//fmt.Println(dict)

	//type DataMap[string,any] map[K] V
	//
	//m := DataMap[string,int]{"fang":1999}
	//fmt.Println(m)
	//var hobbyid string = "a"
	//var locationid string = "b"
	//db = models.GetActivityList(hobbyid,locationid)
	//data :=  make([] *models.Activity,0)
	//for _,v := range data{
	//	fmt.Println(v,'\n')
	//}
	//var count int64
	//err := db.Count(&count).Find(&data).Error
	//
	//if err != err{
	//	fmt.Println("xxxxxxxxx")
	//	fmt.Println("查询失败")
	//	fmt.Println("xxxxxxxxx")
	//
	//}
	//fmt.Println("======")
	//fmt.Println(count)
	//fmt.Println("======")

	//username := "root"  //账号
	//password := "beiwei1010" //密码
	//host := "49.232.124.98" //数据库地址，可以是Ip或者域名
	//port := 3306 //数据库端口
	//Dbname := "bosomfriend" //数据库名
	//timeout := "10s" //连接超时，10秒
	////拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	////连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic("连接数据库失败, error=" + err.Error())
	//}
	////延时关闭数据库连接
	//var activity *models.Activity
	//
	//err = db.First(&activity,1).Error
	//if err != err{
	//	t.Fatal(err)
	//}
	//fmt.Println(activity)
	//data := make([] *models.Activity,0)
	//err = db.Find(&data).Error
	//if err != nil{
	//	t.Fatal(err)
	//	fmt.Println("查询语句出错")
	//}
	//
	//for _,v := range(data){
	//	fmt.Println(v)
	//}

}
