package convert

// @Project: go-util
// @Author: houseme
// @Description:
// @File: convert
// @Version: 1.0.0
// @Date: 2021/4/6 02:09
// @Package goutil

import "encoding/json"

// StringToBytes convert string to bytes
func StringToBytes(s string) []byte {
    return []byte(s)
}

// BytesToString convert bytes to string
func BytesToString(bytes []byte) string {
    return string(bytes)
}

// JSONToMap json string to map
func JSONToMap(jsonStr string) (map[string]string, error) {
    var dat = map[string]string{}
    if err := json.Unmarshal([]byte(jsonStr), &dat); err != nil {
        return dat, err
    }
    return dat, nil
}

// MapToJSON map to json string
func MapToJSON(m map[string]string) (string, error) {
    jsonByte, err := json.Marshal(m)
    if err != nil {
        return "", err
    }
    jsonStr := string(jsonByte)
    return jsonStr, nil
}