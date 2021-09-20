package main

import (
    "io/ioutil"
    "fmt"
    "os"
    "regexp"
    "strings"

)

const READ_LEN = 20

func showLine(lines []rune, totalLen, startLen int) (bool) {
    for i := 0; i < 30; i++ {
        fmt.Println("")
    }

    end := startLen + READ_LEN
    isGoOn := true

    if end > totalLen {
        end = totalLen
        isGoOn = false
    }

    line := lines[startLen:end]
    fmt.Println(string(line))
    return isGoOn
}

func getText(s string) (string) {
    text := ""
    for {
        begin := strings.Index(s, "<")
        if begin != -1 {
            text += s[0:begin]

            end := strings.Index(s, ">")
            s = s[end+1:]
        } else {
            text += s
            return text
        }
    }
}

func readBook(book string) {
    b, err := ioutil.ReadFile(book)
    if err != nil {
        fmt.Println(err)
        return
    }
    content := string(b)

    // <taga>text<tagb>text</tagb></taga> => text<tagb>text</tagb>
    // <taga>text</taga> => text
    re := regexp.MustCompile(`>(.*)<`)
    matches := re.FindAllStringSubmatch(content, -1)

    if len(matches) == 0 {
        fmt.Println("No Text")
        return
    }

    allText := ""
    for _, match := range matches {
        allText += getText(match[1])
    }

    lines := []rune(allText)

    totalLen := len(lines)
    startLen := 0

    var input string
    for {
        if !showLine(lines, totalLen, startLen) {
            return
        }

        fmt.Scanln(&input)
        if input == "q" {
            return
        }

        startLen += READ_LEN
    }
}

func notifyBook(book string) {
    var input string
    fmt.Println(book)
    fmt.Scanln(&input)
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: ParseHTML.exe xxx.html")
        return
    }

    book := os.Args[1]
    readBook(book)
    notifyBook(book)
}