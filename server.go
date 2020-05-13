package main

import (
    "fmt"
    "net/http"
    "os/exec"
    "strings"
    "gopkg.in/alessio/shellescape.v1"
    "bytes"
    "math/rand"
)

func runcommand(bg bool, args ...string) string {
    cmd := exec.Command(args[0], args[1:]...)
    var out bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &stderr

    runc := cmd.Run

    if bg {
        runc = cmd.Start
    }

    err := runc()

    if err != nil {
        fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
        return ""
    }

    return out.String()
}

func randint(min int, max int) int {
    r := rand.Intn(max - min) + min

    switch r {
    case 8000, 33099: // skip some system ports
        return randint(min, max)
    }

    return r
}

func scan(w http.ResponseWriter, req *http.Request) {
    req.ParseForm()
    url := strings.TrimSpace(req.Form.Get("url"))

    if len(url) <= 0 {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    port := randint(9222, 65535)

    runcommand(true, "nohup", "sh", "-c", fmt.Sprintf("./browser.sh %d", port), "&")

    scan := runcommand(false, "CookieScanner", "cli", "--headless", "--port", fmt.Sprintf("%d",port), shellescape.Quote(url))

    w.Header().Set("Content-Type", "application/json")

    w.Write([]byte(scan))
}

func main() {
    http.HandleFunc("/scan-website", scan)
    http.ListenAndServe(":8000", nil)
}