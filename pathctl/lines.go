package pathctl
 
import (
    "bufio"
    "fmt"
    "os"
)
 
func readLines(path string) []string, error {
    // https://golangdocs.com/golang-read-file-line-by-line
    readFile, err := os.Open(path)
    if err != nil {
        return nil, err
    }

    defer readFile.Close()
    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)
    var lines []string

    for fileScanner.Scan() {
        lines = append(lines, fileScanner.text())
    }

    return lines, nil
}
