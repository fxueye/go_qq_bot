package main

import (
	"common/utils"
	"log"
)

type Config struct {
	Port int
}

func main() {
	// config := Config{}
	// bytes, err := ioutil.ReadFile("./config.json")
	// err = json.Unmarshal(bytes, &config)
	// wd, err := os.Getwd()
	// http.Handle("/", http.FileServer(http.Dir(wd)))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("server start on:%d", config.Port)
	// http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
	params := make(map[string]interface{})
	ret, err := utils.HttpGet("http://baidu.com", params)
	if err == nil {
		log.Printf("ret : %v", ret)
	}
}
