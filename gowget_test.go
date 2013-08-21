package wget_test

import (
    "github.com/hahnicity/go-wget"
    "os"
    "path"
    "testing"
)

func TestWgetWithFileName(t *testing.T) {
    fileName := "foo"
    wget.Wget("https://google.com", fileName)
    defer os.Remove(fileName)   
    wd, _ := os.Getwd()
    _, err := os.Open(path.Join(wd, fileName))
    if err != nil {
        t.Errorf("A file with name %s was unable to be opened", fileName)    
    }
}

func TestWgetWithNoFileName(t *testing.T) {
    fileName := "google.com"
    wget.Wget("https://google.com", "")
    defer os.Remove(fileName)   
    wd, _ := os.Getwd()
    _, err := os.Open(path.Join(wd, fileName))
    if err != nil {
        t.Errorf("A file with name  was unable to be opened", fileName)    
    }
}
