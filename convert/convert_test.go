package convert

// @Project: go-util
// @Author: houseme
// @Description:
// @File: convert_test
// @Version: 1.0.0
// @Date: 2021/4/6 02:10
// @Package goutil

import (
    "bytes"
    "reflect"
    "testing"
)

func TestStringToBytes(t *testing.T) {
    test1 := StringToBytes("abc")
    if !bytes.Equal(test1, []byte{97, 98, 99}) {
        t.Error("StringToBytes 'abc' must equal []byte{97, 98, 99}!")
    }
    
    test2 := StringToBytes("")
    if !bytes.Equal(test2, []byte{}) {
        t.Error("StringToBytes '' must equal []byte{}!")
    }
    
    test3 := StringToBytes("123.,/;'m")
    if !bytes.Equal(test3, []byte{49, 50, 51, 46, 44, 47, 59, 39, 109}) {
        t.Error("StringToBytes '' must equal []byte{49, 50, 51, 46, 44, 47, 59, 39, 109}!")
    }
}

func TestBytesToString(t *testing.T) {
    test1 := BytesToString([]byte{97, 98, 99})
    if test1 != "abc" {
        t.Error("BytesToString []byte{97, 98, 99} must equal 'abc'!")
    }
    
    test2 := BytesToString([]byte{})
    if test2 != "" {
        t.Error("BytesToString []byte{} must equal ''!")
    }
    
    test3 := BytesToString([]byte{49, 50, 51, 46, 44, 47, 59, 39, 109})
    if test3 != "123.,/;'m" {
        t.Error("BytesToString []byte{49, 50, 51, 46, 44, 47, 59, 39, 109} must equal '123.,/;'m'!")
    }
}

func TestJSONToMap(t *testing.T) {
    type args struct {
        jsonStr string
    }
    tests := []struct {
        name    string
        args    args
        want    map[string]string
        wantErr bool
    }{
        {
            name: "",
            args: args{
                jsonStr: `{"Port": "80","Namespace": "appboot"}`,
            },
            want: map[string]string{
                "Port":      "80",
                "Namespace": "appboot",
            },
            wantErr: false,
        },
        {
            name: "",
            args: args{
                jsonStr: `{"appboot"}`,
            },
            want:    map[string]string{},
            wantErr: true,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := JSONToMap(tt.args.jsonStr)
            if (err != nil) != tt.wantErr {
                t.Errorf("JSONToMap() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("JSONToMap() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestMapToJSON(t *testing.T) {
    type args struct {
        m map[string]string
    }
    tests := []struct {
        name    string
        args    args
        want    string
        wantErr bool
    }{
        {
            name: "",
            args: args{
                map[string]string{
                    "Port":      "80",
                    "Namespace": "appboot",
                },
            },
            want:    `{"Namespace":"appboot","Port":"80"}`,
            wantErr: false,
        },
        {
            name: "",
            args: args{
                map[string]string{},
            },
            want:    `{}`,
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := MapToJSON(tt.args.m)
            if (err != nil) != tt.wantErr {
                t.Errorf("MapToJSON() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("MapToJSON() = %v, want %v", got, tt.want)
            }
        })
    }
}
