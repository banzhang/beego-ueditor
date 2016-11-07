package util

import (
	"bytes"
	"errors"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beeku"
)

type upload struct {
	Config    map[string]interface{}
	FieldName string
	Type      string
	//	fileInfo  FileInfo
}

var State = map[string]string{
	"SUCCESS":                  "SUCCESS",
	"ERROR_TMP_FILE":           "临时文件错误",
	"ERROR_TMP_FILE_NOT_FOUND": "找不到临时文件",
	"ERROR_SIZE_EXCEED":        "文件大小超出网站限制",
	"ERROR_TYPE_NOT_ALLOWED":   "文件类型不允许",
	"ERROR_CREATE_DIR":         "目录创建失败",
	"ERROR_DIR_NOT_WRITEABLE":  "目录没有写权限",
	"ERROR_FILE_MOVE":          "文件保存时出错",
	"ERROR_FILE_NOT_FOUND":     "找不到上传文件",
	"ERROR_WRITE_CONTENT":      "写入文件内容错误",
	"ERROR_UNKNOWN":            "未知错误",
	"ERROR_DEAD_LINK":          "链接不可用",
	"ERROR_HTTP_LINK":          "链接不是http链接",
	"ERROR_HTTP_CONTENTTYPE":   "链接contentType不正确",
	"INVALID_URL":              "非法 URL",
	"INVALID_IP":               "非法 IP",
}

type FileInfo struct {
	State    string `json:"state"`
	Url      string `json:"url"`      //fullName
	Title    string `json:"title"`    //fileName
	Original string `json:"original"` //original filename
	Type     string `json:"type"`
	Size     string `json:"size"`
	FilePath string `json:"filepath"`
	Source   string `json:"source, omitempty"`
}

func NewUpload(c map[string]interface{}, f string, t string) *upload {
	var u upload = upload{c, f, t}
	return &u
}

func (u *upload) Remote2Local() (*FileInfo, error) {
	imgUrl := html.EscapeString(u.FieldName)
	re := regexp.MustCompile("$amp")
	imgUrl = re.ReplaceAllString(imgUrl, "&")
	fi := &FileInfo{"", "", "", "", "", "", "", ""}
	if !strings.HasPrefix("http://", "http://") {
		fi.State = "ERROR_HTTP_LINK"
		return fi, errors.New(fi.State)
	}
	re = regexp.MustCompile(`(^https*:\/\/[^:\/]+)`)
	imgUrls := re.FindAllString(imgUrl, -1)
	hostWithProtocol := ""
	if len(imgUrls) > 0 {
		hostWithProtocol = imgUrls[0]
	}

	_, err := url.Parse(hostWithProtocol)
	if err != nil {
		fi.State = "INVALID_URL"
		return fi, errors.New(fi.State)
	}
	re = regexp.MustCompile(`^https*:\/\/(.+)`)
	imgUrlss := re.FindAllStringSubmatch(hostWithProtocol, -1)
	hostWithoutProtocol := ""
	if len(imgUrls) > 0 {
		hostWithoutProtocol = string(imgUrlss[0][1])
	} else {
		fi.State = "获取主机名失败"
		return fi, errors.New(fi.State)
	}
	ips, err := net.LookupIP(hostWithoutProtocol)
	if err != nil {
		fi.State = err.Error()
		return fi, err
	}
	if len(ips) < 1 {
		fi.State = "找不到IP"
		return fi, errors.New(fi.State)
	}
	for _, ip := range ips {
		if IsPrivate(ip) {
			fi.State = "INVALID_IP" + ip.String()
			return fi, errors.New(fi.State)
		}
	}

	resp, err := http.Head(imgUrl)
	if !(resp.StatusCode == 200 && strings.Contains(resp.Status, "OK")) {
		fi.State = "ERROR_DEAD_LINK"
		return fi, errors.New(fi.State)
	}
	ext := strings.ToLower(resp.Header.Get("Content-Type"))
	if len(ext) > 0 && !strings.Contains(ext, "image") {
		fi.State = "ERROR_HTTP_CONTENTTYPE"
		beego.Info(ext)
		return fi, errors.New(fi.State)
	}

	ext = path.Ext(imgUrl)
	ext = strings.ToLower(ext)
	if !u.CheckAllowMimeType(ext) {
		fi.State = "ERROR_HTTP_CONTENTTYPE"
		return fi, errors.New(fi.State)
	}

	resp, err = http.Get(imgUrl)
	binImg, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	fi.Size = strconv.FormatInt(resp.ContentLength, 10)
	if resp.ContentLength > int64(u.Config["maxSize"].(float64)) {
		fi.State = "ERROR_SIZE_EXCEED"
		return fi, errors.New(fi.State)
	}
	fi.Original = path.Base(imgUrl)
	fi.Type = strings.TrimLeft(ext, ".")
	u.GetFullName(fi)
	u.GetFilePath(fi)
	os.MkdirAll(fi.FilePath, 0666)
	err = ioutil.WriteFile(fi.FilePath+"/"+fi.Title, binImg, 0666)
	if err != nil {
		beego.Error(err)
		fi.State = fmt.Sprintf("%s", err)
		return fi, err
	}
	fi.State = "SUCCESS"
	fi.Source = imgUrl
	return u.GetFileInfo(fi), nil
}

func (u *upload) Save2Local(r *http.Request) (*FileInfo, error) {
	mf, mfh, err := r.FormFile(u.FieldName)
	if err != nil {
		return nil, err
	}
	defer mf.Close()
	beego.Info(mfh)
	fi := &FileInfo{"", "", "", "", "", "", "", ""}
	o := mfh.Filename
	fi.Original = o

	ext := strings.ToLower(path.Ext(o))
	if len(ext) < 1 {
		ext = u.GetMimeType(mfh)
	}
	if !u.CheckAllowMimeType(ext) {
		msg := fmt.Sprintf("not allow type:%s", ext)
		fi.State = msg
		return fi, errors.New(msg)
	}
	fi.Type = strings.TrimLeft(ext, ".")
	fn := u.GetFullName(fi)
	beego.Info(fn)
	u.GetFilePath(fi)
	os.MkdirAll(fi.FilePath, 0666)
	uf, err := os.OpenFile(fi.FilePath+"/"+fi.Title, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		beego.Error(err)
		fi.State = fmt.Sprintf("%s", err)
		return fi, err
	}
	defer uf.Close()
	io.Copy(uf, mf)
	ufst, _ := uf.Stat()
	fi.Size = strconv.FormatInt(ufst.Size(), 10)
	fi.State = "SUCCESS"
	return u.GetFileInfo(fi), nil
}

//获取扩展类型 png, jpg
func (u *upload) GetMimeType(mfh *multipart.FileHeader) (s string) {
	beego.Info(mfh.Header)
	ty := mfh.Header.Get("Content-Type")
	beego.Info(ty)
	tys := strings.Split(ty, "/")
	beego.Info(tys)
	ext := "." + tys[len(tys)-1]
	return ext
}

//检测是否允许的上传类型
func (u *upload) CheckAllowMimeType(ext string) (b bool) {
	allow := u.Config["allowFiles"].([]interface{})
	if !beeku.In_slice(ext, allow) {
		msg := fmt.Sprintf("not allow type:%s", ext)
		beego.Error(msg)
		return false
	}
	return true
}

func (u *upload) GetFullName(fi *FileInfo) string {
	format := u.Config["pathFormat"].(string)
	t := time.Now()
	layout := "2006-01-02-15-04-05"
	s := t.Format(layout)
	sa := strings.Split(s, "-")

	y := sa[0]
	m := sa[1]
	d := sa[2]
	h := sa[3]
	min := sa[4]
	sec := sa[5]

	//replace {yyyy}
	re := regexp.MustCompile("{yyyy}")
	format = re.ReplaceAllString(format, y)

	//replace {yy}
	re = regexp.MustCompile("{yy}")
	z := []byte(y)
	my := (bytes.NewBuffer(z[2:])).String()
	format = re.ReplaceAllString(format, my)

	re = regexp.MustCompile("{mm}")
	format = re.ReplaceAllString(format, m)

	re = regexp.MustCompile("{dd}")
	format = re.ReplaceAllString(format, d)

	re = regexp.MustCompile("{hh}")
	format = re.ReplaceAllString(format, h)

	re = regexp.MustCompile("{ii}")
	format = re.ReplaceAllString(format, min)

	re = regexp.MustCompile("{ss}")
	format = re.ReplaceAllString(format, sec)

	re = regexp.MustCompile("{time}")
	format = re.ReplaceAllString(format, strconv.FormatInt(t.Unix(), 10))

	//过滤文件名的非法自负,并替换文件名
	re = regexp.MustCompile(`[\|\?\"\<\>\/\*\\\\]+`)
	ons := strings.Split(fi.Original, ".")
	on := ons[0]
	on = re.ReplaceAllString(on, "")
	re = regexp.MustCompile("{filename}")
	format = re.ReplaceAllString(format, on)

	//替换随机字符串
	rand.Seed(time.Now().UnixNano())
	rn := strconv.FormatInt(rand.Int63(), 10)
	re = regexp.MustCompile(`\{rand\:([\d]*)\}`)

	mt := re.FindAllStringSubmatch(format, -1)
	sb := []byte(rn)
	if len(mt) > 0 && len(mt[0]) > 1 {
		l, _ := strconv.Atoi(mt[0][1])
		rn = (bytes.NewBuffer(sb[0:l])).String()
	}
	beego.Info(format)
	beego.Info(rn)
	format = re.ReplaceAllString(format, rn)
	beego.Info(format)
	//拼接扩展名
	fi.Url = format + "." + fi.Type
	fi.Title = path.Base(fi.Url)
	beego.Info(fi.Url)
	return fi.Url
}

func (u *upload) GetFilePath(fi *FileInfo) (s string) {
	rootPath := "admin/upload"
	fi.FilePath = path.Dir(rootPath + fi.Url)
	return fi.FilePath
}

func (u *upload) GetFileInfo(fi *FileInfo) *FileInfo {
	fi.Url = "/img" + fi.Url
	return fi
}
