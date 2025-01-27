package main

import (
    "fmt"
    "net"
    "net/url"
)

func main() {
    rawUrl := "postgres://user:pass@host.com:5432/path?k=v#f"

    parsedUrl, err := url.Parse(rawUrl)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Scheme: %s\n", parsedUrl.Scheme)

    userinfo := parsedUrl.User
    fmt.Printf("Userinfo: %v\n", userinfo)

    username := userinfo.Username()
    fmt.Printf("Username: %s\n", username)

    password, _ := userinfo.Password()
    fmt.Printf("Password: %s\n", password)

    host, port, _ := net.SplitHostPort(parsedUrl.Host)
    fmt.Printf("Host: %s\n", host)
    fmt.Printf("Port: %s\n", port)

    path := parsedUrl.Path
    fmt.Printf("Path: %s\n", path)

    fragment := parsedUrl.Fragment
    fmt.Printf("Fragment: %s\n", fragment)

    rawQuery := parsedUrl.RawQuery
    fmt.Printf("RawQuery: %s\n", rawQuery)

    queries, _ := url.ParseQuery(parsedUrl.RawQuery)
    fmt.Printf("Queries: %v\n", queries)
}
