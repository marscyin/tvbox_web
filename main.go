package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func RequestClient(URL string, METHOD string, HEADER string, DATA string) *http.Response {
	HeaderMap := FormatStr(HEADER)
	DataMap := FormatStr(DATA)
	client := &http.Client{}
	if METHOD == "get" {
		METHOD = http.MethodGet
	} else if METHOD == "post" {
		METHOD = http.MethodPost
	}
	FormatData := ""
	for i, j := range DataMap {
		FormatData = FormatData + i + "=" + j + "&"
	}
	if FormatData != "" {
		FormatData = FormatData[:len(FormatData)-1]
	}
	requset, _ := http.NewRequest(
		METHOD,
		URL,
		strings.NewReader(FormatData),
	)
	if METHOD == http.MethodPost && requset.Header.Get("Content-Type") == "" {
		requset.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	requset.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.71 Safari/537.36")
	for i, j := range HeaderMap {
		requset.Header.Set(i, j)
	}
	resp, err := client.Do(requset)
	if err != nil {
		return nil
	}
	return resp
}

func FormatStr(jsonstr string) map[string]string {
	DataMap := make(map[string]string)
	Nslice := strings.Split(jsonstr, "\n")
	for i := 0; i < len(Nslice); i++ {
		if strings.Index(Nslice[i], ":") != -1 {
			if strings.TrimSpace(Nslice[i])[:6] == "Origin" {

				a := strings.TrimSpace(Nslice[i][:strings.Index(Nslice[i], ":")])
				b := strings.TrimSpace(Nslice[i][strings.Index(Nslice[i], ":")+1:])
				c := strings.Trim(a, "\"")
				d := strings.Trim(b, "\"")
				DataMap[c] = d
			} else {
				a := strings.TrimSpace(Nslice[i][:strings.LastIndex(Nslice[i], ":")])
				b := strings.TrimSpace(Nslice[i][strings.LastIndex(Nslice[i], ":")+1:])
				c := strings.Trim(a, "\"")
				d := strings.Trim(b, "\"")
				DataMap[c] = d
			}
		}
	}
	return DataMap
}

func Go_HomeContent(etd string, filter bool, file_name string) string {
	M := make(map[string]interface{})
	R := make(map[string]interface{})
	flr := "False"
	if filter {
		flr = "True"
	}
	head, file_name := filepath.Split(file_name)
	File_Name_Remove_py := strings.TrimSuffix(file_name, ".py")
	cmd := exec.Command("python3", "-c", "import sys;sys.path.append(\""+head+"\");from "+File_Name_Remove_py+" import homeContent,init;init(\""+etd+"\");homeContent("+flr+")")
	// fmt.Println(cmd)
	content, err := cmd.Output()
	e := json.Unmarshal([]byte(content), &R)

	if err != nil || e != nil {
		M["code"] = 0
		M["message"] = "homeContent运行出错，请检查!!!"
		M["data"] = err.Error()
		fmt.Printf("%v\n", cmd)

	} else {
		M["code"] = 1
		M["message"] = "success"
		M["data"] = R
	}
	jstr, err := json.Marshal(M)
	if err != nil {
		return "{\"code\":0,\"message\":\"格式化json出错，请检查!!!函数名为Go_HomeContent!!\",\"data\":\"\"}"
	} else {
		return string(jstr)
	}
}

func Go_CategoryContent(etd string, tid string, pg string, filter bool, extend string, file_name string) string {
	M := make(map[string]interface{})
	R := make(map[string]interface{})
	flr := "False"
	if filter {
		flr = "True"
	}
	head, file_name := filepath.Split(file_name)
	File_Name_Remove_py := strings.TrimSuffix(file_name, ".py")
	cmd := exec.Command("python3", "-c", "import sys;sys.path.append(\""+head+"\");from "+File_Name_Remove_py+" import categoryContent,init;init(\""+etd+"\");categoryContent(\""+tid+"\",\""+pg+"\","+flr+",'"+extend+"')")
	content, err := cmd.Output()
	e := json.Unmarshal([]byte(content), &R)
	if err != nil || e != nil {
		M["code"] = 0
		M["message"] = "categoryContent运行出错，请检查!!!"
		M["data"] = err.Error()
		fmt.Printf("%v\n", cmd)

	} else {
		M["code"] = 1
		M["message"] = "success"
		M["data"] = R
	}
	jstr, err := json.Marshal(M)
	if err != nil {
		return "{\"code\":0,\"message\":\"格式化json出错，请检查!!!函数名为Go_CategoryContent!!\",\"data\":\"\"}"
	} else {
		return string(jstr)
	}
}

func Go_DetailContent(etd string, ids string, file_name string) string {
	M := make(map[string]interface{})
	R := make(map[string]interface{})
	head, file_name := filepath.Split(file_name)
	File_Name_Remove_py := strings.TrimSuffix(file_name, ".py")
	cmd := exec.Command("python3", "-c", "import sys;sys.path.append(\""+head+"\");from "+File_Name_Remove_py+" import detailContent,init;init(\""+etd+"\");detailContent(\""+ids+"\")")
	content, err := cmd.Output()
	e := json.Unmarshal([]byte(content), &R)
	if err != nil || e != nil {
		M["code"] = 0
		M["message"] = "detailContent运行出错，请检查!!!"
		M["data"] = err.Error()
		fmt.Printf("%v\n", cmd)

	} else {
		M["code"] = 1
		M["message"] = "success"
		M["data"] = R
	}
	jstr, err := json.Marshal(M)
	if err != nil {
		return "{\"code\":0,\"message\":\"格式化json出错，请检查!!!函数名为Go_DetailContent!!\",\"data\":\"\"}"
	} else {
		return string(jstr)
	}
}

func Go_PlayerContent(etd string, flag string, id string, file_name string) string {
	M := make(map[string]interface{})
	R := make(map[string]interface{})
	head, file_name := filepath.Split(file_name)
	File_Name_Remove_py := strings.TrimSuffix(file_name, ".py")
	cmd := exec.Command("python3", "-c", "import sys;sys.path.append(\""+head+"\");from "+File_Name_Remove_py+" import playerContent,init;init(\""+etd+"\");playerContent(\""+flag+"\",\""+id+"\")")
	content, err := cmd.Output()
	e := json.Unmarshal([]byte(content), &R)
	if err != nil || e != nil {
		M["code"] = 0
		M["message"] = "playerContent运行出错，请检查!!!"
		M["data"] = err.Error()
		fmt.Printf("%v\n", cmd)

	} else {
		M["code"] = 1
		M["message"] = "success"
		M["data"] = R
	}
	jstr, err := json.Marshal(M)
	if err != nil {
		return "{\"code\":0,\"message\":\"格式化json出错，请检查!!!函数名为Go_PlayerContent!!\",\"data\":\"\"}"
	} else {
		return string(jstr)
	}
}

func Base64ToStr(b64 string) string {
	//
	b, e := base64.StdEncoding.DecodeString(b64)
	if e != nil {
		return ""
	}
	return string(b)
}

func Go_SearchContent(etd string, key string, file_name string) string {
	M := make(map[string]interface{})
	R := make(map[string]interface{})
	head, file_name := filepath.Split(file_name)
	File_Name_Remove_py := strings.TrimSuffix(file_name, ".py")
	cmd := exec.Command("python3", "-c", "import sys;sys.path.append(\""+head+"\");from "+File_Name_Remove_py+" import searchContent,init;init(\""+etd+"\");searchContent(\""+key+"\")")
	content, err := cmd.Output()
	e := json.Unmarshal([]byte(content), &R)
	if err != nil || e != nil {
		M["code"] = 0
		M["message"] = "searchContent运行出错，请检查!!!"
		M["data"] = err.Error()
		fmt.Printf("%v\n", cmd)

	} else {
		M["code"] = 1
		M["message"] = "success"
		M["data"] = R
	}
	jstr, err := json.Marshal(M)
	if err != nil {
		return "{\"code\":0,\"message\":\"格式化json出错，请检查!!!函数名为Go_SearchContent!!\",\"data\":\"\"}"
	} else {
		return string(jstr)
	}
}

func main() {
	r := gin.Default()
	r.GET("/homeContent", func(c *gin.Context) {
		etd := c.Query("etd")
		spider_file_path := c.Query("spider_file_path")
		if spider_file_path == "" {
			c.JSON(200, gin.H{
				"code":    0,
				"message": "没有spider_file_path的参数!",
			})
			return
		} else {
			res := Go_HomeContent(etd, true, spider_file_path)
			c.String(200, res)
			return
		}
	})
	r.GET("/categoryContent", func(c *gin.Context) {
		extend := c.Query("extend")
		etd := c.Query("etd")
		tid := c.Query("tid")
		pg := c.Query("pg")
		spider_file_path := c.Query("spider_file_path")

		if spider_file_path == "" {
			c.JSON(200, gin.H{
				"code":    0,
				"message": "没有spider_file_path的参数!",
			})
			return
		}
		if pg == "" {
			c.JSON(200, gin.H{
				"code":    0,
				"message": "没有pg的参数!",
			})
			return
		}
		if tid == "" {
			c.JSON(200, gin.H{
				"code":    0,
				"message": "没有tid的参数!",
			})
			return
		}
		c.String(200, Go_CategoryContent(etd, tid, pg, true, extend, spider_file_path))
		return
	})
	r.GET("/detailContent", func(c *gin.Context) {
		spider_file_path := c.Query("spider_file_path")
		etd := c.Query("etd")
		ids := c.Query("ids")

		if spider_file_path == "" {
			c.JSON(200, gin.H{
				"code":    0,
				"message": "没有spider_file_path的参数!",
			})
			return
		}
		if ids == "" {
			c.JSON(200, gin.H{
				"code":    0,
				"message": "没有ids的参数!",
			})
			return
		}
		c.String(200, Go_DetailContent(etd, ids, spider_file_path))
		return
	})
	r.GET("/playerContent", func(c *gin.Context) {
		spider_file_path := c.Query("spider_file_path")
		etd := c.Query("etd")
		flag := c.Query("flag")
		id := Base64ToStr(c.Query("id"))

		if spider_file_path == "" {
			c.JSON(200, gin.H{
				"code":    0,
				"message": "没有spider_file_path的参数!",
			})
			return

		}
		if id == "" {
			c.JSON(200, gin.H{
				"code":    0,
				"message": "没有id的参数!或者Base64解密出错",
			})
			return

		}
		c.String(200, Go_PlayerContent(etd, flag, id, spider_file_path))
		return
	})
	r.GET("/searchContent", func(c *gin.Context) {
		spider_file_path := c.Query("spider_file_path")
		etd := c.Query("etd")
		key := c.Query("key")
		if spider_file_path == "" {
			c.JSON(200, gin.H{
				"code":    0,
				"message": "没有spider_file_path的参数!",
			})
			return

		}
		if key == "" {
			c.JSON(200, gin.H{
				"code":    0,
				"message": "没有key的参数!",
			})
			return

		}
		c.String(200, Go_SearchContent(etd, key, spider_file_path))
		return
	})
	r.GET("/proxy", func(c *gin.Context) {
		URL := Base64ToStr(c.Query("url"))
		HEADER := Base64ToStr(c.Query("header"))
		METHOD := Base64ToStr(c.Query("method"))
		DATA := Base64ToStr(c.Query("data"))
		resp := RequestClient(URL, METHOD, HEADER, DATA)
		if resp == nil {
			c.String(200, "请求URL:"+URL+"出错!")
			return
		}
		for k, v := range resp.Header {
			for i := 0; i < len(v); i++ {
				c.Header(k, v[i])
			}
		}
		body_bit, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		c.String(200, string(body_bit))
		return

		// c.String(200, res_str)
	})
	r.GET("/", func(c *gin.Context) {
		c.File("./html/" + "index.html")
	})
	r.GET("/detailContent.html", func(c *gin.Context) {
		uid := Base64ToStr(c.Query("uid"))
		p_url := ""
		if uid != "" {
			res := Go_PlayerContent("", "", uid, "/root/go/src/tvbox_web/python/NanGua.py")
			p_url = gjson.Get(res, "data.url").String()

		}
		fmt.Println(uid)
		h := `

<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>播放</title>
    <link href="/css/video-js.min.css" rel="stylesheet" />
    <link href="/css/detailContent.css" rel="stylesheet" />
    <script src="/js/video.min.js"></script>
  </head>
  <body>
    <video
      id="my-video"
      class="video-js vjs-default-skin"
      controls
      preload="auto"
      width="640"
      height="264"
      data-setup="{}"
    >
 <source src="` + p_url + `" type="application/x-mpegURL" />
    </video>
    <hr />
    <div id="play_from"></div>
    <div id="play_url"></div>
  </body>
  <script>

var player = videojs('my-video')
player.play()
  </script>
  <script src="/js/detailContent.js"></script>
  <script src="/js/play.js"></script>
</html>
 `
		c.Header("Content-Type", "text/html")
		c.String(200, h)
	})

	r.GET("/:html", func(c *gin.Context) {
		html := c.Param("html")
		c.File("./html/" + html)
	})
	r.GET("/js/:js", func(c *gin.Context) {
		js := c.Param("js")
		c.File("./html/js/" + js)
	})
	r.GET("/css/:css", func(c *gin.Context) {
		css := c.Param("css")
		c.File("./html/css/" + css)
	})
	r.GET("/m3u8/:m3u8", func(c *gin.Context) {
		m3u8 := c.Param("m3u8")
		c.File("./m3u8/" + m3u8)
	})

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run(":9987")
}
