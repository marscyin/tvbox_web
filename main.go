package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func Go_HomeContent(etd string, filter bool, file_name string) string {
	M := make(map[string]interface{})
	flr := "False"
	if filter {
		flr = "True"
	}
	head, file_name := filepath.Split(file_name)
	File_Name_Remove_py := strings.TrimSuffix(file_name, ".py")
	cmd := exec.Command("python3", "-c", "import sys;sys.path.append(\""+head+"\");from "+File_Name_Remove_py+" import homeContent,init;init(\""+etd+"\");homeContent("+flr+")")
	fmt.Println(cmd)
	content, err := cmd.Output()
	if err != nil {
		M["code"] = 0
		M["message"] = "homeContent运行出错，请检查!!!"
		M["data"] = err.Error()
		fmt.Printf("%v\n", cmd)

	} else {
		M["code"] = 1
		M["message"] = "success"
		M["data"] = strings.TrimSpace(string(content))
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
	flr := "False"
	if filter {
		flr = "True"
	}
	head, file_name := filepath.Split(file_name)
	File_Name_Remove_py := strings.TrimSuffix(file_name, ".py")
	cmd := exec.Command("python3", "-c", "import sys;sys.path.append(\""+head+"\");from "+File_Name_Remove_py+" import categoryContent,init;init(\""+etd+"\");categoryContent(\""+tid+"\",\""+pg+"\","+flr+",\""+extend+"\")")
	content, err := cmd.Output()
	if err != nil {
		M["code"] = 0
		M["message"] = "categoryContent运行出错，请检查!!!"
		M["data"] = err.Error()
		fmt.Printf("%v\n", cmd)

	} else {
		M["code"] = 0
		M["message"] = "success"
		M["data"] = strings.TrimSpace(string(content))
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
	head, file_name := filepath.Split(file_name)
	File_Name_Remove_py := strings.TrimSuffix(file_name, ".py")
	cmd := exec.Command("python3", "-c", "import sys;sys.path.append(\""+head+"\");from "+File_Name_Remove_py+" import detailContent,init;init(\""+etd+"\");detailContent(\""+ids+"\")")
	content, err := cmd.Output()
	if err != nil {
		M["code"] = 0
		M["message"] = "detailContent运行出错，请检查!!!"
		M["data"] = err.Error()
		fmt.Printf("%v\n", cmd)

	} else {
		M["code"] = 0
		M["message"] = "success"
		M["data"] = strings.TrimSpace(string(content))
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
	head, file_name := filepath.Split(file_name)
	File_Name_Remove_py := strings.TrimSuffix(file_name, ".py")
	cmd := exec.Command("python3", "-c", "import sys;sys.path.append(\""+head+"\");from "+File_Name_Remove_py+" import playerContent,init;init(\""+etd+"\");playerContent(\""+flag+"\",\""+id+"\")")
	content, err := cmd.Output()
	if err != nil {
		M["code"] = 0
		M["message"] = "playerContent运行出错，请检查!!!"
		M["data"] = err.Error()
		fmt.Printf("%v\n", cmd)

	} else {
		M["code"] = 0
		M["message"] = "success"
		M["data"] = strings.TrimSpace(string(content))
	}
	jstr, err := json.Marshal(M)
	if err != nil {
		return "{\"code\":0,\"message\":\"格式化json出错，请检查!!!函数名为Go_PlayerContent!!\",\"data\":\"\"}"
	} else {
		return string(jstr)
	}
}

func Go_SearchContent(etd string, key string, file_name string) string {
	M := make(map[string]interface{})
	head, file_name := filepath.Split(file_name)
	File_Name_Remove_py := strings.TrimSuffix(file_name, ".py")
	cmd := exec.Command("python3", "-c", "import sys;sys.path.append(\""+head+"\");from "+File_Name_Remove_py+" import searchContent,init;init(\""+etd+"\");searchContent(\""+key+"\")")
	content, err := cmd.Output()
	if err != nil {
		M["code"] = 0
		M["message"] = "searchContent运行出错，请检查!!!"
		M["data"] = err.Error()
		fmt.Printf("%v\n", cmd)

	} else {
		M["code"] = 0
		M["message"] = "success"
		M["data"] = strings.TrimSpace(string(content))
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
		id := c.Query("id")
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
				"message": "没有id的参数!",
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
	r.GET("/", func(c *gin.Context) {
		c.File("./html/" + "index.html")
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

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run(":9987")
}
