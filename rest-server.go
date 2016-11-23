package main

import (
    "database/sql"
    "log"
    "net/http"
    "fmt"
    "os"
    
    "github.com/simonleung8/flags"
    "github.com/ant0ine/go-json-rest/rest"
    _ "github.com/go-sql-driver/mysql"
    "github.com/BurntSushi/toml"
)

    
func readConfig(config_file string) tomlConfig {
    
    var config tomlConfig
    _, err := toml.DecodeFile(config_file, &config)
    checkError(err)
    return config
}


type tomlConfig struct {
    User string
    Passwd string
    Database string
    ListenIP string
    ListenPort int
    LogFile string
}


func getWeather(User, Passwd, Database string) (map[string]string, error) {
    
    weather := make(map[string]string)
    db, err := sql.Open("mysql", User+":"+Passwd+"@/"+Database)
    if err != nil {
        return weather, err
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        return weather, err
    }
    
    rows, err := db.Query("SELECT * FROM raw ORDER BY dateTime DESC LIMIT 1")
    if err != nil {
        return weather, err
    }
    
    columns, err := rows.Columns()
    if err != nil {
        return weather, err
    }
    
    values := make([]sql.RawBytes, len(columns))
    scanArgs := make([]interface{}, len(values))
    
    for i := range values {
        scanArgs[i] = &values[i]
    }

    for rows.Next() {
        err = rows.Scan(scanArgs...)
        if err != nil {
            return weather, err
        }
        
        for i, col := range values {
            if col == nil {
                weather[columns[i]] = "NULL"
            } else {
                weather[columns[i]] = string(col)
            }
        }
    }
    db.Close()
    return weather, nil
}


func checkError(err error) {
    if err != nil {
        log.Println(err)
        os.Exit(1)
    }
}


func main() {    
    
    var config tomlConfig
    
    args := flags.New()
    args.NewStringFlag("config", "c", "configuration_file")
    args.Parse(os.Args...)
    
    if (args.IsSet("c") == true) {
        config = readConfig(args.String("c"))
        logger, err := os.OpenFile(config.LogFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0664)
        checkError(err)
        defer logger.Close()
        log.SetOutput(logger)
    } else {
        fmt.Println("Usage: rest-server -c <configuration_file> -l <log_file>")
        os.Exit(1)
    }
    
    log.Println("Starting server...")
    
    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)
    
    router, err := rest.MakeRouter(
        rest.Get("/api/get", func(w rest.ResponseWriter, req *rest.Request) {
            data, err := getWeather(config.User, config.Passwd, config.Database)
            if err != nil {
                rest.Error(w, err.Error(), http.StatusInternalServerError)
                return  
            }
            w.WriteJson(data)
        }),
    )
    checkError(err)
    api.SetApp(router)
    log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d",config.ListenIP, config.ListenPort), api.MakeHandler()))
    
}
