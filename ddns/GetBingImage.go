package ddns

import (
	"net/http"
	"io/ioutil"
	"regexp"
	"os"
	"time"
	"github.com/sirupsen/logrus"
	logUtil "com.dingjinlin.getBing/util/log"
	fileUtil "com.dingjinlin.getBing/util/file"
)

const URL = "https://cn.bing.com"

var logger *logrus.Logger

func init() {
	logger = logUtil.GetLoggerInstance()
}

/*func RunTase(unit time.Duration, count time.Duration) {
	ticker := time.NewTicker(unit * count)
	go generateWallpaper("images")
	for range ticker.C {
		go generateWallpaper("images")
	}

}*/

func RunTime(directName string, hour, min, sec int) {
	go generateWallpaper(directName)

	now := time.Now()
	var next time.Time

	next = time.Date(now.Year(), now.Month(), now.Day(), hour, min, sec, 0, now.Location())

	var value time.Duration
	if !now.Before(next) {
		value = next.Add(time.Hour * 24).Sub(now)
	} else {
		value = next.Sub(now)
	}

	ticker := time.NewTicker(value)
	<-ticker.C
	go generateWallpaper(directName)

	ticker = time.NewTicker(time.Hour * 24)
	//ticker = time.NewTicker(time.Second * 5)
	for range ticker.C {
		go generateWallpaper(directName)
	}
}
func generateWallpaper(directName string) {
	absPath, err := os.Getwd()
	if err != nil {
		logger.Errorln(err)
	}

	filePath, _ := GetBingImg(absPath, directName)

	if filePath != "" {
		direct := absPath + "/" + "wallpaper/"
		if !fileUtil.CheckFileIsExist(direct) {
			err := os.Mkdir(direct, os.ModePerm)
			if err != nil {
				logger.Warnln(err)
			}
		}

		symlinkfile := direct + "wallpaper.jpg"
		if fileUtil.CheckFileIsExist(symlinkfile) {
			err := os.Remove(symlinkfile)
			if err != nil {
				logger.Warnln(err)
			}
		}

		err := os.Symlink(filePath, symlinkfile)
		if err != nil {
			logger.Warnln(err)
		}
	}

}

func GetBingImg(absPath, directName string) (filePath string, filename string) {
	var resp *http.Response
	var err error
	resp, err = http.Get(URL)

	if err != nil {
		logger.Info(err)
		return "", ""
	}

	defer resp.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	pattern := "(?i)g_img={URL: \"(.+\\.jpg)"
	//"[\w-]+\.jpg"

	flysnowRegexp := regexp.MustCompile(pattern)
	value := flysnowRegexp.FindStringSubmatch(string(body))

	pattern = "[\\w-]+\\.jpg"
	flysnowRegexp = regexp.MustCompile(pattern)
	fileNames := flysnowRegexp.FindStringSubmatch(string(body))

	direct := absPath + "/" + directName + "/"
	if !fileUtil.CheckFileIsExist(direct) {
		err := os.Mkdir(direct, os.ModePerm)
		if err != nil {
			logger.Warnln(err)
		}
	}

	fileName := fileNames[0]
	file := direct + fileName
	if !fileUtil.CheckFileIsExist(file) {
		imgUrl := URL + value[1]
		resp, err = http.Get(imgUrl)

		if err != nil {
			logger.Errorln(err)
		}

		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)

		fi, err := os.Create(file)
		if nil != err {
			panic(err)
		}
		defer fi.Close()
		err = ioutil.WriteFile(file, body, 0666)
		if nil != err {
			panic(err)
		}
		logger.Infoln("download file success!")
		return file, fileName
	} else {
		logger.Infoln("have same file!")
		return "", ""
	}

}


