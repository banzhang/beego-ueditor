package controllers

import (
	"blog/util"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/astaxie/beego"
)

type UeditorController struct {
	beego.Controller
}

type FileInfo struct {
	State    string `json:"state"`
	Url      string `json:"url"`
	Title    string `json:"title"`
	Original string `json:"original"`
	Type     string `json:"type"`
	Size     string `json:"size"`
}

func (c *UeditorController) URLMapping() {
	c.Mapping("Config", c.Config)
	c.Mapping("Uploadimage", c.Uploadimage)
	c.Mapping("Uploadscrawl", c.Uploadscrawl)
	c.Mapping("Uploadvideo", c.Uploadvideo)
	c.Mapping("Uploadfile", c.Uploadfile)
	c.Mapping("Listimage", c.Listimage)
	c.Mapping("Listfile", c.Listfile)
	c.Mapping("Catchimage", c.Catchimage)
	c.Mapping("Index", c.Index)
}

// Config ...
// @Title Get Config
// @Description get config with json type
// @router /config [get]
func (c *UeditorController) Config() {
	c.Redirect("/admin/ueditor/config.json", http.StatusMovedPermanently)
}

func (c *UeditorController) getConfig() (map[string]interface{}, error) {
	cs := "admin/ueditor/config.json"
	var err error
	err = nil
	if fi, err := os.Stat(cs); fi != nil {
		ct, err := ioutil.ReadFile(cs)
		if err == nil {
			cts := string(ct[:])
			beego.Debug(cts)
			re := regexp.MustCompile(`\/\*[\s\S]+?\*\/`)
			var ar interface{}
			cts = re.ReplaceAllString(cts, "")
			beego.Debug(cts)
			err := json.Unmarshal([]byte(cts), &ar)
			if err != nil {
				beego.Error(err)
			}
			beego.Debug(ar)
			arm := ar.(map[string]interface{})
			return arm, nil
		} else {
			beego.Error(err)
		}
	} else {
		beego.Error(err)
	}
	return nil, err
}

// @router /uploadimage [post]
func (c *UeditorController) Uploadimage() {
	arm, err := c.getConfig()
	if err == nil {
		base64 := "upload"

		var config = map[string]interface{}{
			"pathFormat": arm["imagePathFormat"],
			"maxSize":    arm["imageMaxSize"],
			"allowFiles": arm["imageAllowFiles"],
		}
		fieldName := (arm["imageFieldName"]).(string)

		beego.Info(config)
		beego.Info(fieldName)
		beego.Info(base64)

		u := util.NewUpload(config, fieldName, base64)
		fi, err := u.Save2Local(c.Ctx.Request)
		if err == nil {
			c.Data["json"] = fi
			c.ServeJSON()
		} else {
			beego.Error(err)
			c.Data["json"] = map[string]string{"state": err.Error()}
			c.ServeJSON()
		}
	} else {
		beego.Error(err)
		c.Data["json"] = map[string]string{"state": err.Error()}
		c.ServeJSON()
	}
}

// @router /uploadscrawl [post]
func (c *UeditorController) Uploadscrawl() {
	arm, err := c.getConfig()
	if err == nil {
		base64 := "upload"

		var config = map[string]interface{}{
			"pathFormat": arm["scrawlPathFormat"],
			"maxSize":    arm["scrawlMaxSize"],
			"allowFiles": arm["scrawlAllowFiles"],
		}
		fieldName := (arm["scrawlFieldName"]).(string)

		beego.Info(config)
		beego.Info(fieldName)
		beego.Info(base64)

		u := util.NewUpload(config, fieldName, base64)
		fi, err := u.Save2Local(c.Ctx.Request)
		if err == nil {
			c.Data["json"] = fi
			c.ServeJSON()
		} else {
			beego.Error(err)
			c.Data["json"] = map[string]string{"state": err.Error()}
			c.ServeJSON()
		}
	} else {
		beego.Error(err)
		c.Data["json"] = map[string]string{"state": err.Error()}
		c.ServeJSON()
	}
}

// @router /uploadvideo [post]
func (c *UeditorController) Uploadvideo() {
	arm, err := c.getConfig()
	if err == nil {
		base64 := "upload"

		var config = map[string]interface{}{
			"pathFormat": arm["videoPathFormat"],
			"maxSize":    arm["videoMaxSize"],
			"allowFiles": arm["videoAllowFiles"],
		}
		fieldName := (arm["videoFieldName"]).(string)

		beego.Info(config)
		beego.Info(fieldName)
		beego.Info(base64)

		u := util.NewUpload(config, fieldName, base64)
		fi, err := u.Save2Local(c.Ctx.Request)
		if err == nil {
			c.Data["json"] = fi
			c.ServeJSON()
		} else {
			beego.Error(err)
			c.Data["json"] = map[string]string{"state": err.Error()}
			c.ServeJSON()
		}
	} else {
		beego.Error(err)
		c.Data["json"] = map[string]string{"state": err.Error()}
		c.ServeJSON()
	}
}

// @router /uploadfile [post]
func (c *UeditorController) Uploadfile() {
	arm, err := c.getConfig()
	if err == nil {
		base64 := "upload"
		var config = map[string]interface{}{
			"pathFormat": arm["filePathFormat"],
			"maxSize":    arm["fileMaxSize"],
			"allowFiles": arm["fileAllowFiles"],
		}
		fieldName := (arm["fileFieldName"]).(string)
		beego.Info(config)
		beego.Info(fieldName)
		beego.Info(base64)

		u := util.NewUpload(config, fieldName, base64)
		fi, err := u.Save2Local(c.Ctx.Request)
		if err == nil {
			c.Data["json"] = fi
			c.ServeJSON()
		} else {
			beego.Error(err)
			c.Data["json"] = map[string]string{"state": err.Error()}
			c.ServeJSON()
		}
	} else {
		beego.Error(err)
		c.Data["json"] = map[string]string{"state": err.Error()}
		c.ServeJSON()
	}
}

// @router /listimage [get]
func (c *UeditorController) Listimage() {
	arm, err := c.getConfig()
	allow := make([]string, 5)
	allowFiles := (arm["imageManagerAllowFiles"]).([]interface{})
	for _, it := range allowFiles {
		allow = append(allow, it.(string))
	}
	listSize := arm["imageManagerListSize"]
	path := strings.Trim((arm["imageManagerListPath"]).(string), "/") + "/"
	ls := listSize.(float64)
	lz := int64(ls)
	beego.Info("ls, lz", ls, lz)
	size, _ := c.GetInt64("size", lz)
	start, _ := c.GetInt64("start", 0)
	end := size + start
	beego.Info("size, start, end", size, start, end)
	var res []map[string]interface{}
	res = make([]map[string]interface{}, 0)
	err = util.GetFiles("admin/upload/"+path, &res, allow)
	if err != nil {
		beego.Error(err)
	}
	l := int64(len(res))
	beego.Info("len:[%d],info[%s]", l, res)
	if l < 1 {
		c.Data["json"] = map[string]interface{}{
			"state": "no match file",
			"list":  make([]string, 0),
			"start": start,
			"total": 0,
		}
		c.ServeJSON()
	}
	ress := make([]map[string]interface{}, 0)
	i := l
	if end < l {
		i = end
	}
	for i = i - 1; i < l && i >= 0 && i >= start; i-- {
		itu := ((res[i])["url"]).(string)
		re := regexp.MustCompile(`admin\/upload\/`)
		itu = re.ReplaceAllString(itu, "/img/")
		(res[i])["url"] = itu
		beego.Info(itu, res[i])
		ress = append(ress, res[i])
	}
	c.Data["json"] = map[string]interface{}{
		"state": "SUCCESS",
		"list":  ress,
		"start": start,
		"total": len(ress),
	}
	c.ServeJSON()
}

// @router /listfile [get]
func (c *UeditorController) Listfile() {
	arm, err := c.getConfig()
	if err != nil {
		c.Data["json"] = map[string]string{
			"state": err.Error(),
		}
		c.ServeJSON()
	}
	allow := make([]string, 5)
	allowFiles := (arm["fileManagerAllowFiles"]).([]interface{})
	for _, it := range allowFiles {
		allow = append(allow, it.(string))
	}
	listSize := arm["fileManagerListSize"]
	path := strings.Trim((arm["fileManagerListPath"]).(string), "/") + "/"
	ls := listSize.(float64)
	lz := int64(ls)
	beego.Info("ls, lz", ls, lz)
	size, _ := c.GetInt64("size", lz)
	start, _ := c.GetInt64("start", 0)
	end := size + start
	beego.Info("size, start, end", size, start, end)
	var res []map[string]interface{}
	res = make([]map[string]interface{}, 0)
	err = util.GetFiles("admin/upload/"+path, &res, allow)
	if err != nil {
		beego.Error(err)
	}
	l := int64(len(res))
	beego.Info("len:[%d],info[%s]", l, res)
	if l < 1 {
		c.Data["json"] = map[string]interface{}{
			"state": "no match file",
			"list":  make([]string, 0),
			"start": start,
			"total": 0,
		}
		c.ServeJSON()
	}
	ress := make([]map[string]interface{}, 0)
	i := l
	if end < l {
		i = end
	}
	for i = i - 1; i < l && i >= 0 && i >= start; i-- {
		itu := ((res[i])["url"]).(string)
		re := regexp.MustCompile(`admin\/upload\/`)
		itu = re.ReplaceAllString(itu, "/img/")
		(res[i])["url"] = itu
		beego.Info(itu, res[i])
		ress = append(ress, res[i])
	}
	c.Data["json"] = map[string]interface{}{
		"state": "SUCCESS",
		"list":  ress,
		"start": start,
		"total": len(ress),
	}
	c.ServeJSON()
}

// @router /catchimage [post]
// @router /catchimage [get]
func (c *UeditorController) Catchimage() {
	arm, err := c.getConfig()
	if err != nil {
		c.Data["json"] = map[string]string{
			"state": err.Error(),
		}
		c.ServeJSON()
	}
	var config = map[string]interface{}{
		"pathFormat": arm["catcherPathFormat"],
		"maxSize":    arm["catcherMaxSize"],
		"allowFiles": arm["catcherAllowFiles"],
		"oriName":    "remote.png",
	}
	fieldName := ((arm["catcherFieldName"]).(string)) + "[]"
	base64 := "remote"
	beego.Info(config)
	beego.Info(fieldName)
	beego.Info(base64)
	source := c.GetStrings(fieldName)
	beego.Info(source)
	for _, imgUrl := range source {
		u := util.NewUpload(config, imgUrl, base64)
		fi, err := u.Remote2Local()
		if err == nil {
			c.Data["json"] = fi
			c.ServeJSON()
		} else {
			beego.Error(err)
		}
	}
}

// @Title 编辑器的入口URL
// @Description 编辑器的入口URL
// @router / [get,post]
func (u *UeditorController) Index() {
	action := u.GetString("action", "config")
	hmap := map[string]func(){
		"config":       u.Config,
		"uploadfile":   u.Uploadfile,
		"uploadimage":  u.Uploadimage,
		"uploadscrawl": u.Uploadscrawl,
		"uploadvideo":  u.Uploadvideo,
		"listfile":     u.Listfile,
		"listimage":    u.Listimage,
		"catchimage":   u.Catchimage,
	}
	hf, ok := hmap[action]
	if ok {
		hf()
	}
}
