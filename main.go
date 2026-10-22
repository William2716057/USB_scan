package main

import (
        "fmt"
        "os"
        "path/filepath"
        "strings"
)

var suspiciousExtensions = map[string]bool{
        ".exe": true,
        ".bat": true,
        ".cmd": true,
        ".ps1": true,
        ".vbs": true,
        ".js":  true,
        ".scr": true,
        ".lnk": true,
        ".msi": true,
}

func main() {
        if len(os.Args) < 2 {
                fmt.Println("Usage: go run script.go <USB-path>")
                fmt.Println(`Example: go run script.go E:\`)
                return
        }

        root := os.Args[1]

        err := filepath.WalkDir(root, func(
                path string,
                entry os.DirEntry,
                err error,
        ) error {
                if err != nil {
                        fmt.Printf("[ERROR] %s: %v\n", path, err)
                        return nil
                }

                if entry.IsDir() {
                        return nil
                }

                ext := strings.ToLower(filepath.Ext(entry.Name()))

                if suspiciousExtensions[ext] {
                        fmt.Printf("Flagged: %s\n", path)
                }

                return nil
        })

        if err != nil {
                fmt.Println("Scan error:", err)
                return
        }

        fmt.Println("Scan complete. Review flagged files before opening them.")
}
