package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*
{
    "service":{
        "SdMon" : {
            "Type" : "Default",
            "ServiceGroup":"SdMonGroup",
            "ServiceBin":"SdMon.dll",
            "Description":"",
            "Scope":"",
            "Support64":false,
            "InstallNow":true,
			"AutoStart":true,
			"Tag":"4ebc56dabc8d1efebcfd10c3ed89205c"
        },
        "SdDoc" : {
            "Type" : "Default",
            "ServiceGroup":"SdDocGroup",
            "ServiceBin":"DocumentAudit.dll",
            "ServiceBin64":"DocumentAudit64.dll",
            "Description":"",
            "Scope":"",
            "Support64":true,
            "InstallNow":true,
			"AutoStart":true,
			"Tag":"78a399b14017a3c6f921fa66a5c874c1",
			"Tag64":"83ca84846f646699252c519ea7aaf140",
			"Depends": {
				"DocumentAuditDriver.sys":{
					"Platform":"x86",
					"Tag":"ad5097011a101ab6fe6744f7cf67d2ae"
				},
				"DocumentAuditDriver64.sys":{
					"Platform":"x64",
					"Tag":"86d851614932da6db570b1d03837572e"
				}
			}
        },
        "SdUp" : {
            "Type" : "Default",
            "ServiceGroup":"SdUpGroup",
            "ServiceBin":"update.dll",
            "Description":"",
            "Scope":"",
            "Support64":false,
            "InstallNow":true,
			"AutoStart":true,
			"Tag":"34832845d20b013316b38a82bebd51fe"
        },
        "SdIc" : {
            "Type" : "Default",
            "ServiceGroup":"SdIcGroup",
            "ServiceBin":"IC.dll",
            "Description":"",
            "Scope":"",
            "Support64":false,
            "InstallNow":true,
			"AutoStart":true,
			"Tag":"32bc25682833b600f30b2e1582436bb3",
			// 服务依赖的文件
			"Depends": {
				"www\\conf\\server.pem":{
					"Tag":"714dac0dec2356afc6990890921f55c2"
				},
				"www\\conf\\server.key":{
					"Tag":"1c50f67c3c7451245faa450da3f28bd8"
				},
				"www\\html\\index.html":{
					"Tag":"cb0252dffc5690ab5ab1c1fc06a1553c"
				},
				"Proxy.dll":{
					"Tag":"4f9e1123ad6dddc6bbe7aa07672ade92"
				},
				"WinDivert32.sys":{
					"Platform":"x86",
					"Tag":"c68af5c064be4a7dafa041f50f3862ee"
				},
				"WinDivert64.sys":{
					"Platform":"x64",
					"Tag":"6a33620de63bccaf5e5314ee49cd58fb"
				}
			}
        }
    },
	"extres": {
		"WxMan.dll":{
			"Tag":"24e5c93415c81c43441fcf3502c03dc8"
		}
	},
	"version": "1.0.0.4",
	"package":{
		"1.0.0.4":"sd_top_version.zip"
	}
}


type ModleDllST struct {
	DllName string
}

// 对内容进行过滤处理，去除\r, 去除对应的注释
func StrContentSimple(data []byte) ([]byte, error) {
	data = bytes.Replace(data, []byte("\r"), []byte(""), 0) // 把windwos中的\r去除
	lines := bytes.Split(data, []byte("\n"))
	simple_data := make([][]byte, 0)

	for _, line := range lines {
		match, err := regexp.Match(`^\s*#`, line)
		if err != nil {
			return nil, err
		}
		if !match {
			simple_data = append(simple_data, line)
		}
	}

	return bytes.Join(simple_data, []byte("\n")), nil
}

func LoadConfig(path string) {

	config_file, err := os.Open(path)
	if err != nil {
		fmt.Printf("failed to open file %s : %s\n", path, err)
		return
	}

	fileinfos, _ := config_file.Stat()
	if fileinfos.Size() == 0 {
		fmt.Printf("config file (%s) is empty, skipping \n", path)
		return
	}

	buffer := make([]byte, fileinfos.Size())
	_, err = config_file.Read(buffer)

	buffer, err = StrContentSimple(buffer)
	if err != nil {
		fmt.Printf("failed to simple content from json, err:%s\n", err)
		return
	}

}
*/

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

/*
type MailInfos struct {
	SendAccount string       // 发送账号
	ToAccount   []string     // 接收账号
	CcAccount   []string     // 抄送账号
	BccAccount  []string     // 加密抄送
	MailContent string       // 邮件内容
	MailTitle   string       // 邮件标题
	Fileinfo    FileinfosMap // 附件的文件名
	MailBoxFlag int          // 邮箱标识，网页版本，客户端版本，和不同的厂商的邮箱
	TimeStamp   int          // 发送时间戳
}

*/

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
	generateJson()
	testMd5Str()
}
