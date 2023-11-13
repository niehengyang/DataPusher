package apilogic

import (
	"bytes"
	"context"
	"dm/service/resp"
	"dm/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UploadLogic struct {
	ctx    context.Context
	req    *http.Request
}

func NewUploadLogic(ctx context.Context, req *http.Request) *UploadLogic {
	return &UploadLogic{
		ctx:    ctx,
		req:    req,
	}
}

func (l *UploadLogic) UploadImage() (res *resp.Response, err error) {
	fileInfo, errU := utils.UploadExcel(l.req)
	if errU != nil {
		return nil, resp.NewError(resp.Err_InternalServerError, errU.Error())
	}

	for key,item := range fileInfo{
		subData(key, item)
		fmt.Println("key:", key)
		fmt.Println("item:", item)
	}

	return resp.Success("生成成功！"), nil
}

type DataF struct {
	Devicelv1 string `json:"devicelv1"`
	Devicelv2 string `json:"devicelv2"`
	DictServerValue int `json:"dictServerValue"`
	GbPlayerUrl string `json:"gbPlayerUrl"`
	Name string `json:"name"`
	SchoolId string `json:"schoolId"`
	Status int `json:"status"`
}

func subData(deviceId string, deviceName string) error {
	url := "https://h5.139zhxy.cn/base/apiCamera/api/camera/b/addCamera" // 替换为目标API的URL

	fileCount := DataF{"7", "0", 1,
		deviceId, deviceName, "1723723760696602625", 1}

	data, _ := json.Marshal(fileCount)
	// 创建自定义请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return err
	}

	// 添加Authorization标头
	token := "eyJhbGciOiJIUzUxMiJ9.eyJwbGF0Zm9ybVR5cGUiOjQsInN1YiI6IjE3MjMyODA0NzUyMzQ1NDU2NjYiLCJleHAiOjE3MDA0NDIwNDAsImlhdCI6MTY5OTgzNzI0MCwianRpIjoiNGM0MTRjZjctMDI4MS00MGQ4LWJkNTEtNTNiMzg5YTNkYTU5IiwiYWNjb3VudFR5cGUiOjZ9.lRu91pwYpBV6VuVh2yrhB-0ukOZMeOHDmyfSezVWWHmwYjQW-2XDN0xWGU31qRASFh-wDZyGMdLdfd4Uai4m3g" // 替换为实际的访问令牌
	req.Header.Set("Authorization", "Bearer "+token)

	// 设置请求的Content-Type
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败:", err)
		return err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应体失败:", err)
		return err
	}

	// 输出响应结果
	fmt.Println("响应状态码:", resp.Status)
	fmt.Println("响应体:", string(body))

	return nil
}