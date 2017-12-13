// 测试内容
// 每个文件一个测试内容

package main

import (
	"encoding/json"
	"net/http"
	"testing"
	"bytes"

	. "github.com/smartystreets/goconvey/convey"

	"time"
)

// 获取测试结果。网上公开的。
func getTestResult() {

}


func TestSelfInAppEventReportHandler(t *testing.T) {

	Convey("测试post接口", t, func() {

		// 手动启动服务 ./run.sh

		// 构造传输数据
		var postEvents []InAppEvent
		var deviceInfo DeviceInfo

		for _, tempEventName := range []string{"k_live", "k_close", "k_click"} {
			eventValue := map[string]string{"user_id":"wocaocoa", "我草草": "我草拟大爷"}
			postEvents = append(postEvents, InAppEvent{
				SessionId:"wocaocoa",
				DeviceId: "bodao",
				UserId: "1000",
				EventKey: tempEventName,
				EventValue: eventValue,
				Ctime: string(int64(time.Now().Unix())*1000),
			})
		}
		deviceInfo = DeviceInfo{
			SessionId: "wocaocoa",
			// 设备信息
			DeviceID: "bodao",
			Platform: "ios",
			SystemVersion: "10.2.12",
			SystemModel: "iphone 7",
			Lang: "thai",
			// 用户相关信息
			UserId: "1000",
			Country: "TH",
			// 网络信息
			Operator: "中国移动",
			Network: "4g",
			Ipv4: "124.45.251.214",
			// 软件信息等
			AppVersion: "57",
			InstallChannel: "organic",
		}

		postData := InAppEventData{DeviceInfo:deviceInfo, Events:postEvents}

		// 发送post请求
		localTestUrl := "http://localhost:9000/v1/data/report"
		jsonString, jsonErr := json.Marshal(postData)
		So(jsonErr, ShouldEqual, nil)

		postResult, postErr := http.Post(localTestUrl, "application-json; charset=utf-8", bytes.NewReader(jsonString))

		So(postErr, ShouldEqual, nil)
		So(postResult, ShouldNotEqual, nil)

		// 验证相关数据在mongo里边的存储情况

		// // 记录条数 event 应该是3条, deviceinfo应该是一条

		eventDocs := make([]InAppEvent, 0)
		eventDocsReadErr := db.MongoClient.Find(&db.ReportDataBase, &EventCollectionName, &bson.M{}, &eventDocs, true)
		So(eventDocsReadErr, ShouldEqual, nil)
		//fmt.Println(eventDocs)
		So(len(eventDocs), ShouldEqual, len(postEvents))

		// // 看看数据是否能对应上.
		deviceDocs := make([]DeviceInfo, 0)
		deviceDocsReadErr := db.MongoClient.Find(&db.ReportDataBase, &DeviceInfoCollectionName, &bson.M{}, &deviceDocs, true)
		So(deviceDocsReadErr, ShouldEqual, nil)
		//fmt.Println(deviceDocs)
		So(len(deviceDocs), ShouldEqual, 1)

		// 查看deviceInfo的数据
		//So(deviceDocs[0], ShouldEqual, deviceInfo)

		// 删除mongo里边的数据
		// 有一个bug. 如果convey执行不通过, 这个函数不会执行.
		Reset(func() {

			for _, collectionName := range []string{DeviceInfoCollectionName, EventCollectionName} {
				db.MongoClient.DeleteSync(&db.ReportDataBase, &collectionName, &bson.M{})
			}
		})

	})


}