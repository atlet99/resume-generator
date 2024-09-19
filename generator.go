package main

import (
    "fmt"
    "os"
    "text/template"
)

func generateResume(resume Resume) error {
    tmpl, err := template.ParseFiles("templates/london-template.txt")
    if err != nil {
        return err
    }

    f, err := os.Create("resume.txt")
    if err != nil {
        return err
    }
    defer f.Close()

    return tmpl.Execute(f, resume)
}