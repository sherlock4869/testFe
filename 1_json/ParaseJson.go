package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// 邮件附件信息
type FileinfosMap map[string]int

// 邮箱结构体相关字段信息
type MailInfos struct {
	SendAccount   string       `json:"send"`          // 发送账号
	ToAccount     []string     `json:"to"`            // 接收账号
	CcAccount     []string     `json:"cc"`            // 抄送账号
	BccAccount    []string     `json:"bcc"`           // 加密抄送
	MailContent   string       `json:"content"`       // 邮件内容
	MailTitle     string       `json:"title"`         // 邮件标题
	MailTimeStamp int          `json:"mailtimestamp"` // 发送时间戳
	Fileinfo      FileinfosMap `json:"attach"`        // 附件的文件名
	MailBoxFlag   int          `json:"flag"`          // 邮箱标识，网页版本，客户端版本，和不同的厂商的邮箱
}

type ServerPostJsonSt struct {
	ClientId  string `json:"clientid"`
	Timestamp string `json:"timestamp"`
	Tag       string `json:"tag"`
	MailInfos
}

type fileInfoSt struct {
	FileName string `json:"name"`
	FileSize string `json:"size"`
}

func readfile(filename string) []byte {
	if filename == "" {
		fmt.Println(" filename is null")
		return nil
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ioutil.ReadFile is failed, err:", err)
		return nil
	}
	return data
}

func parseJsondData() {

	fileArr := []fileInfoSt{}
	data := readfile("test.json")
	err := json.Unmarshal(data, &fileArr)
	if err != nil {
		fmt.Println(" json.Unmarshal is failed, err:", err)
		return
	}
	fmt.Println("fileArr:", fileArr)
}

func generateJson() {
	postServer := ServerPostJsonSt{
		ClientId:  "clientid123456789",
		Timestamp: "1641471693",
		Tag:       "qwertyuiopasdfghjklzxcvbnm1234567890",
	}

	postServer.MailInfos = MailInfos{
		SendAccount:   "zhushaopeng555@163.com",
		ToAccount:     []string{"To1111xxx@126.com", "To223456789@qq.com"},
		CcAccount:     []string{"Cc1111122222@126.com", "CC111223333332@qq.com"},
		BccAccount:    []string{"Bcc111122444444@163.com", "Bcc5556667778@qq.com"},
		MailContent:   "this is a test mail",
		MailTitle:     "test_zsp",
		MailBoxFlag:   1,
		MailTimeStamp: 1641000000,
		Fileinfo: map[string]int{
			"123456.txt": 123,
			"67890.txt":  111,
		},
	}

	jsonStr, err := json.Marshal(postServer)
	if err != nil {
		fmt.Println("json.Marshal is failed, err:", err)
		return
	}

	fmt.Println(string(jsonStr))
}

func testMd5Str() {
	str := "this is a test"
	has := md5.Sum([]byte(str))
	md5Str := fmt.Sprintf("%x", has)
	fmt.Println("str:", str, " | md5Sum:", md5Str)
}

func main() {
	//parseJsondData()
	fmt.Println("test")
	generateJson()
	testMd5Str()
}
