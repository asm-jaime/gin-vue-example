package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	gen "github.com/asm-jaime/gen"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// ========== server

type Config struct {
	Port         string
	StaticFolder string
	IndexFile    string
}

func (config *Config) SetDefault() { // {{{
	config.Port = "3000"
	config.StaticFolder = "../public"
	config.IndexFile = "../public/index.html"
} // }}}

// ========== data

type Data struct {
	Id   string `form:"id"`
	Data string `form:"data"`
}

// SetRnd set random data to a point// {{{
func (data *Data) SetRnd() {
	data.Id = gen.Str(6)
	data.Data = gen.Str(20)
} // }}}

// ========== dataState

type DataState struct {
	Dates map[string]Data
	sync.RWMutex
}

func NewDataState() (datast *DataState) { // {{{
	datast = &DataState{Dates: make(map[string]Data)}
	return datast
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

func (datast *DataState) Get(data *Data) (dates []Data) { // {{{
	datast.Lock()
	defer datast.Unlock()

	if data.Id == "" {
		for _, data := range datast.Dates {
			dates = append(dates, data)
		}
	} else {
		gdata, ok := datast.Dates[data.Id]
		if ok == true {
			dates = append(dates, gdata)
		}
	}
	return dates
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

// ========== interface for global vars, what should be set to context

type Vars struct { // {{{
	dataState DataState
} // }}}

// ========== controller

func GetData(c *gin.Context) { // {{{
	vars, ok := c.Keys["vars"].(*Vars)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "can't get vars from context", "body": nil})
		return
	}

	var req Data
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error(), "body": nil})
		return
	}

	dates := vars.dataState.Get(&req)
	if len(dates) < 1 {
		c.JSON(http.StatusNotFound, gin.H{"msg": "no data in state", "body": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "get data complete", "body": dates})
} // }}}

func PostData(c *gin.Context) { // {{{
	vars, ok := c.Keys["vars"].(*Vars)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "can't get vars from context", "body": nil})
		return
	}

	var req Data
	err := c.Bind(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error(), "body": nil})
		return
	}

	vars.dataState.Post(&req)
	c.JSON(http.StatusOK, gin.H{"msg": "get dates success", "body": req})
} // }}}

func PutData(c *gin.Context) { // {{{
	vars, ok := c.Keys["vars"].(*Vars)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "can't get vars from context", "body": nil})
		return
	}

	var req Data
	err := c.Bind(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error(), "body": nil})
		return
	}

	vars.dataState.Put(&req)
	c.JSON(http.StatusOK, gin.H{"msg": "put dates success", "body": req})
} // }}}

func DelData(c *gin.Context) { // {{{
	vars, ok := c.Keys["vars"].(*Vars)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "can't get vars from context", "body": nil})
		return
	}

	var req Data
	err := c.Bind(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error(), "body": nil})
		return
	}

	vars.dataState.Del(&req)

	c.JSON(http.StatusOK, gin.H{"msg": "get dates success", "body": req})
} // }}}

// ========== middleware

func middleVars(vars *Vars) gin.HandlerFunc { // {{{
	return func(c *gin.Context) {
		c.Set("vars", vars)
		c.Next()
	}
} // }}}

func NoRoute(c *gin.Context) { // {{{
	path := strings.Split(c.Request.URL.Path, "/")
	if (path[1] != "") && (path[1] == "api") {
		c.JSON(http.StatusNotFound, gin.H{"msg": "no route", "body": nil})
	} else {
		c.HTML(http.StatusOK, "index.html", "")
	}
} // }}}

// ========== router

func NewRouter(vars *Vars, config *Config) *gin.Engine { // {{{
	router := gin.Default()

	// middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleVars(vars))
	// no route, bad url
	router.NoRoute(NoRoute)

	// frontend
	router.Use(static.Serve("/", static.LocalFile(config.StaticFolder, true)))
	router.LoadHTMLGlob(config.IndexFile)

	// api
	api := router.Group("/api")
	{
		data := api.Group("/data")
		{
			data.GET("/", GetData)
			data.POST("/", PostData)
			data.PUT("/", PutData)
			data.DELETE("/", DelData)
		}
	}
	return router
} // }}}

func startServer(args []string) {
	// set config
	config := Config{}
	config.SetDefault()

	vars := Vars{dataState: *NewDataState()}
	vars.dataState.SetRnd(10)

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

	if _, err := os.Stat(config.IndexFile); os.IsNotExist(err) {
		fmt.Printf("\nfile: %v does not exist\n", config.IndexFile)
		return
	}

	// server config info
	fmt.Println("---------------")
	fmt.Println("Selected port: " + config.Port)
	fmt.Println("Selected static folder: " + config.StaticFolder)
	fmt.Println("Selected index.html: " + config.IndexFile)
	fmt.Println("---------------")

	// get and start router
	router := NewRouter(&vars, &config)
	router.Run(":" + config.Port)
}

func main() {
	args := os.Args
	startServer(args)
}
