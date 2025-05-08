package libpctl
 
import (
    "os"
    "io"
    "bufio"
    "errors"
)

const BUF_SIZE int = 32_768 // 2**15


func ReadLines(path string) ([]string, error) {
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
        lines = append(lines, fileScanner.Text())
    }

    return lines, nil
}


func FileCopy(srcpath, dstpath string) error {
    infile, err := os.Open(srcpath)
    if err != nil {
       return err
    }
    defer infile.Close()

    outfile, err := os.Create(dstpath)
    if err != nil {
       return err
    }
    defer outfile.Close()

    // Specifically do NOT use io.Copy() nor io.CopyBuf()
    //  as they seem to read all into memory using ReadFrom()
    //  and wierdly, if ReaderFrom is implemented on the reader, then
    //  the buffer is _ignored_ (wtf?)
    
    // var buff [BUF_SIZE]byte
    buff := make([]byte, BUF_SIZE)

    rcount, rerr := infile.Read(buff)
    for {
        got_genuine_error := (rerr != nil && !errors.Is(rerr, io.EOF))

        if rcount > 0 && !got_genuine_error {
            _, werr := outfile.Write(buff[:rcount])
            if werr != nil {
                return werr
            }
        }

        if rerr != nil {
            if got_genuine_error {
                return rerr
            } else {
                // rerr is now necessarily an io.EOF
                return nil
            }
        }

        rcount, rerr = infile.Read(buff)
    }
}
