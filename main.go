package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

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

func Baes64ToStr(b64 string) string {
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
		id := Baes64ToStr(c.Query("id"))

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
		url := c.Query("url")
		header := c.Query("header")
		c.String(200, url+"\n"+header)
	})
	r.GET("/", func(c *gin.Context) {
		c.File("./html/" + "index.html")
	})
	r.GET("/detailContent.html", func(c *gin.Context) {
		uid := Baes64ToStr(c.Query("uid"))
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
