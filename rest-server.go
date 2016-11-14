package main

import (
    "github.com/ant0ine/go-json-rest/rest"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "log"
    "net/http"
    "errors"
)

func get_weather() (map[string]string, error) {
    
    weather := make(map[string]string)
    
    db, err := sql.Open("mysql", "weewx:weewx@/weewx")
    if err != nil {
        return weather, errors.New("Can`t connect to database")
        }
    
    err = db.Ping()
    if err != nil {
        return weather, errors.New("Can`t connect to database")
    }
    
    rows, err := db.Query("SELECT * FROM raw LIMIT 1")
    if err != nil {
        return weather, errors.New("Can`t get data from database")
    }
    
    columns, err := rows.Columns()
    if err != nil {
        return weather, log.Fatal(err)
    }
    
    values := make([]sql.RawBytes, len(columns))
    scanArgs := make([]interface{}, len(values))
    
    for i := range values {
        scanArgs[i] = &values[i]
    }

    for rows.Next() {
        err = rows.Scan(scanArgs...)
        if err != nil {
            return weather, log.Fatal(err)
        }
        
        for i, col := range values {
            if col == nil {
                weather[columns[i]] = "NULL"
            } else {
                weather[columns[i]] = string(col)
            }
        }
    }
    return weather, nil
}

func run_restserver() {
    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)
    
    router, err := rest.MakeRouter(
        rest.Get("/api/get", func(w rest.ResponseWriter, req *rest.Request) {
            data, err := get_weather()
            if err != nil {
                rest.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            w.WriteJson(data)
        }),
    )
    if err != nil {
        log.Fatal(err)
    }
    api.SetApp(router)
    log.Fatal(http.ListenAndServe("0.0.0.0:8080", api.MakeHandler()))
}

func main() {
    run_restserver()
}
