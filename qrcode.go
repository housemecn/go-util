package goutil

import (
    "fmt"
    "io"
    
    "github.com/tuotoo/qrcode"
    qrCodeGo "github.com/yeqown/go-qrcode"
)

// @Project: go-util
// @Author: houseme
// @Description:
// @File: qrcode
// @Version: 1.0.0
// @Date: 2021/4/6 00:12
// @Package goutil

// QRCodeParse 二维码图片解析
func QRCodeParse(fi io.Reader) (string, error) {
    qrMatrix, err := qrcode.Decode(fi)
    if err != nil {
        return "", err
    }
    return qrMatrix.Content, nil
}

// QRCreate .
func QRCreate(text, saveToPath string, opts ...qrCodeGo.ImageOption) error {
    qrc, err := qrCodeGo.New(text, opts...)
    if err != nil {
        fmt.Printf("could not generate QRCode: %v", err)
        return err
    }
    
    // save file
    if err := qrc.Save(saveToPath); err != nil {
        fmt.Printf("could not save image: %v", err)
        return err
    }
    return nil
}
