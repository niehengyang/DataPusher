package utils

import (
	"dm/ebyte/logger"
	"fmt"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"mime/multipart"
	"net/http"
)

// UploadExcel 文件上传
func UploadExcel(req *http.Request) (fileInfo map[string]string, err error) {

	// 设置表单最大10MB
	errS := req.ParseMultipartForm(10 << 20)
	if errS != nil {
		logger.Error("文件大小超出范围:", zap.Error(errS))
		return fileInfo, errS
	}

	// 读取文件
	file, _, errF := req.FormFile("file")
	if errF != nil {
		logger.Error("文件读取失败:", zap.Error(errS))
		return fileInfo, errF
	}
	defer func(file multipart.File) {
		if errC := file.Close(); errC != nil {
			logger.Error("文件关闭失败:", zap.Error(errC))
		}
	}(file)

	f, err := excelize.OpenReader(file)
	if err != nil {
		return
	}

	// 选择要读取的工作表
	sheetName := "Sheet1"
	rows, err := f.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 创建一个 map 用于存储数据
	dataMap := make(map[string]string)

	// 遍历每一行，将数据存储到 map 中
	for index, row := range rows {
		// 假设第一列是键，第二列是值
		if index > 0 {
			if len(row) >= 2 {
				key := row[1]
				value := row[2]
				dataMap[key] = value
			}
		}
	}

	return dataMap, nil
}