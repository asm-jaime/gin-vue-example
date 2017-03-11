package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"math/rand"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// ========== addition methods

// random string {{{
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func rndStr(n int) string {
	rndStr := make([]rune, n)
	for i := range rndStr {
		rndStr[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(rndStr)
} // }}}

// ========== data

type Data struct {
	Id    string  `form:"id" binding:"required"`
	Name  string  `form:"name"`
	SomeF float64 `form:"somef"`
	SomeI int     `form:"somei"`
}

// SetRnd set random data to a point// {{{
func (data *Data) SetRnd() {
	data.Id = rndStr(6)
	data.Name = rndStr(8)
	data.SomeF = (rand.Float64() * 5) + 5
	data.SomeI = (rand.Int() * 5) + 5
} // }}}

// ========== datastate

type DataState struct {
	Dates map[string]Data
	sync.RWMutex
}

func NewDataState() *DataState { // {{{
	return &DataState{
		Dates: make(map[string]Data),
	}
} // }}}

func (datast *DataState) Get(id string) (data Data, ok bool) { // {{{
	datast.Lock()
	defer datast.Unlock()

	data, ok = datast.Dates[id]
	return data, ok
} // }}}

func (datast *DataState) Post(data *Data) { // {{{
	datast.Lock()
	defer datast.Unlock()
	datast.Dates[data.Id] = *data
} // }}}

func (datast *DataState) Put(data *Data) { // {{{
	datast.Lock()
	defer datast.Unlock()

	datast.Dates[data.Id] = *data
} // }}}

func (datast *DataState) Del(data *Data) { // {{{
	datast.Lock()
	defer datast.Unlock()

	delete(datast.Dates, data.Id)
} // }}}

func (datast *DataState) SetRnd(num int) { // {{{
	datast.Lock()
	defer datast.Unlock()

	data := new(Data)
	for i := 0; i < num; i++ {
		data.SetRnd()
		datast.Dates[data.Id] = *data
	}
} // }}}

// ========== server

type Server struct{}

type Config struct {
	Port         string
	StaticFolder string
	IndexFile    string
}

func (config *Config) SetDefault() { // {{{
	config.Port = "3000"
	config.StaticFolder = "./public"
	config.IndexFile = "./public/index.html"
} // }}}

func (config *Config) CheckFiles() (err error) { // {{{
	if _, err := os.Stat(config.IndexFile); os.IsNotExist(err) {
		fmt.Printf("\nfile: %v does not exist\n", config.IndexFile)
		// make dir if it does't exist
		os.MkdirAll(config.StaticFolder, os.ModePerm)

		var file, err = os.Create(config.IndexFile)
		if err != nil {
			return err
		}

		file, err = os.OpenFile(config.IndexFile, os.O_RDWR, 0644)
		defer file.Close()
		if err != nil {
			return err
		}

		// write default index.html
		_, err = file.WriteString(`
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>some project</title>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/css/bootstrap.min.css" integrity="sha384-rwoIResjU2yc3z8GV/NPeZWAv56rSmLldC3R/AZzGRnGxQQKnKkoFVhFQhNUwEyJ" crossorigin="anonymous">
</head>

<body>
  <div id="app"></div>
  <script src="build.js"></script>
</body>
</html>`)
		// save changes
		err = file.Sync()
		if err != nil {
			fmt.Printf("\ncan't write to index.html\n")
			return err
		}
	}
	return
} // }}}

// ========== vars

var dataState *DataState

// ========== controller

func GetDates(c *gin.Context) { // {{{
	var url_request string
	url_request = c.Request.URL.Query().Get("id")

	if url_request != "" {
		fmt.Printf("\nget data id: %v\n", url_request)
		if data, ok := dataState.Get(url_request); ok == true {
			c.JSON(http.StatusOK, gin.H{"message": "get data on id success", "body": data})
		} else {
			c.JSON(400, gin.H{"message": "no data on id in data state", "body": nil})
		}
		return
	}

	if len(dataState.Dates) > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "get dates success", "body": dataState})
	} else {
		c.JSON(400, gin.H{"message": "no data on server", "body": nil})
	}
} // }}}

func PostData(c *gin.Context) { // {{{
	var request Data
	err := c.Bind(&request)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "Incorrect data", "body": nil})
		return
	}

	dataState.Post(&request)

	c.JSON(http.StatusOK, gin.H{"message": "get dates success", "body": request})
} // }}}

func PutData(c *gin.Context) { // {{{
	var request Data
	err := c.Bind(&request)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "Incorrect data", "body": nil})
		return
	}

	dataState.Put(&request)

	c.JSON(http.StatusOK, gin.H{"message": "put dates success", "body": request})
} // }}}

func DelData(c *gin.Context) { // {{{
	var request Data
	err := c.Bind(&request)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "Incorrect data", "body": nil})
		return
	}

	dataState.Del(&request)

	c.JSON(http.StatusOK, gin.H{"message": "get dates success", "body": request})
} // }}}

// ========== middleware

func noRoute(c *gin.Context) { // {{{
	path := strings.Split(c.Request.URL.Path, "/")
	if (path[1] != "") && (path[1] == "api") {
		c.JSON(http.StatusNotFound, gin.H{"message": "no route"})
	} else {
		c.HTML(http.StatusOK, "index.html", "")
	}
} // }}}

func (server *Server) NewEngine(config *Config) { // {{{
	router := gin.Default()

	// router
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// all frontend
	fmt.Printf("\nstatic folder: %v\n", config.StaticFolder)

	router.Use(static.Serve("/", static.LocalFile(config.StaticFolder, true)))
	router.LoadHTMLGlob(config.IndexFile)

	// api
	api := router.Group("/api")
	{
		data := api.Group("/data")
		{
			data.GET("/", GetDates)
			data.POST("/", PostData)
			data.PUT("/", PutData)
			data.DELETE("/", DelData)
		}
	}

	// server config info
	fmt.Println("---------------")
	fmt.Println("Selected port: " + config.Port)
	fmt.Println("Selected static folder: " + config.StaticFolder)
	fmt.Println("Selected index.html: " + config.IndexFile)
	fmt.Println("---------------")

	router.Run(":" + config.Port)
} // }}}

func startServer(args []string) {
	// set config
	config := Config{}
	config.SetDefault()
	dataState = NewDataState()
	dataState.SetRnd(10)

	// custom params
	if len(args) > 1 { // set custom port
		config.Port = args[1]
	}
	if len(args) > 2 { // set custom static folder
		config.StaticFolder = args[2]
	}
	if len(args) > 3 { // set custom index file
		config.IndexFile = args[3]
	}
	err := config.CheckFiles()
	if err != nil {
		fmt.Printf("\nerror index.html file: %v\n", err)
		return
	}

	// start server
	server := Server{}
	server.NewEngine(&config)
}

func main() {
	args := os.Args
	startServer(args)
}
