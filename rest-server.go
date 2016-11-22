package main

import (
    "database/sql"
    "log"
    "net/http"
    "fmt"
    "os"
    
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
    
    rows, err := db.Query("SELECT * FROM raw LIMIT 1")
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
        log.Print(err)
        os.Exit(1)
    }
}


func main() {    
    
    var config tomlConfig
    
    if len(os.Args) > 1 {
        config = readConfig(os.Args[1])
    } else {
        fmt.Println("Usage: rest-server <config_file>")
        os.Exit(0)
    }
    
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
    log.Fatal(http.ListenAndServe("0.0.0.0:8080", api.MakeHandler()))
}
