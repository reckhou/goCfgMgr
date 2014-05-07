package goCfgMgr

import (
  "bitbucket.org/kardianos/osext"
  js "github.com/bitly/go-simplejson"
  "log"
  "os"
)

var (
  cfgFile    = "config.json"
  cfgPath    = ""
  cfgContent string
  cfgMap     map[string]interface{}
)

func init() {
  cfgPath, _ = osext.ExecutableFolder()
  cfgPath += "config/" + cfgFile
  log.Println("cfgPath:", cfgPath)

  file, err := os.Open(cfgPath)
  if err != nil {
    log.Println("Config file read error:", err)
    os.Exit(1)
  }

  fileLen, _ := file.Seek(0, 2)
  data := make([]byte, fileLen)
  file.Seek(0, 0)
  file.Read(data)

  file.Close()
  cfgContent = string(data)
  //log.Println("cfgContent:", cfgContent)

  cfgJson, errJS := js.NewJson([]byte(cfgContent))
  if errJS != nil {
    log.Println("Config format error:", err)
    os.Exit(1)
  }

  cfgMap, errJS = cfgJson.Map()
  if errJS != nil {
    log.Println("Config map error:", err)
    os.Exit(1)
  }

  //log.Println(cfgMap)
}

func Get(field string, keys interface{}) interface{} {
  m, ok := cfgMap[field]
  if !ok {
    log.Println("Invalid field", field, "in config file")
    return nil
  }

  if keys == nil {
    return m
  }

  return m.(map[string]interface{})[keys.(string)]
}
