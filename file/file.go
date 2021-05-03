package file

import (
    "bufio"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
    
    "github.com/housemecn/goutil/convert"
)

// BytesToFile 字节数组写入到文件
func BytesToFile(data []byte, filepath string) error {
    file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        return err
    }
    defer file.Close()
    
    writer := bufio.NewWriter(file)
    _, err = writer.Write(data)
    if err != nil {
        return err
    }
    _ = writer.Flush()
    return nil
}

// Exists reports whether the file or directory exists.
func Exists(path string) (existed bool) {
    existed, _ = Exist(path)
    return
}

// Exist reports whether the file or directory exists.
func Exist(path string) (existed bool, isDir bool) {
    info, err := os.Stat(path)
    if err != nil {
        return !os.IsNotExist(err), false
    }
    return true, info.IsDir()
}

// Mode get file mode
func Mode(path string) os.FileMode {
    fileInfo, err := os.Stat(path)
    if err != nil {
        return 0755
    }
    return fileInfo.Mode()
}

// WriteStringToFile write string to file
func WriteStringToFile(content, path string, mode os.FileMode) (err error) {
    bytes := convert.StringToBytes(content)
    return ioutil.WriteFile(path, bytes, mode)
}

// AppendStringToFile append string to file
func AppendStringToFile(content string, filePath string) error {
    f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer f.Close()
    if _, err := f.WriteString(content); err != nil {
        return err
    }
    return nil
}

// GetDirList get directory list
func GetDirList(path string) ([]string, error) {
    return GetDirListWithFilter(path, nil)
}

// Filter file filter
type Filter func(os.FileInfo) bool

// GetDirListWithFilter get directory list with filter
func GetDirListWithFilter(path string, filter Filter) ([]string, error) {
    var dirList []string
    
    paths, err := filepath.Glob(filepath.Join(path, "*"))
    
    log.Printf("paths: %v", paths)
    
    for _, value := range paths {
        f, err := os.Stat(value)
        if err != nil {
            return dirList, err
        }
        if filter != nil && !filter(f) {
            continue
        }
        if f.IsDir() {
            dir := strings.Replace(value, path, "", 1)
            if strings.HasPrefix(dir, "/") {
                dir = strings.Replace(dir, "/", "", 1)
            }
            dirList = append(dirList, dir)
        }
    }
    
    return dirList, err
}

// RecreateDir recreate dir
func RecreateDir(dir string) error {
    mode := Mode(dir)
    _ = os.RemoveAll(dir)
    return os.MkdirAll(dir, mode)
}

// GetFilepaths get all filepaths in a directory tree
func GetFilepaths(dir string) ([]string, error) {
    var paths []string
    err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            paths = append(paths, path)
        }
        return nil
    })
    return paths, err
}

// File template file
type File struct {
    Path    string
    Content string
}

// GetFiles get files
func GetFiles(dir string) ([]*File, error) {
    var files []*File
    
    paths, err := GetFilepaths(dir)
    if err != nil {
        return files, err
    }
    
    for _, path := range paths {
        bytes, err := ioutil.ReadFile(path)
        if err != nil {
            return files, err
        }
        content := string(bytes)
        file := &File{
            Path:    path,
            Content: content,
        }
        files = append(files, file)
    }
    
    return files, nil
}
