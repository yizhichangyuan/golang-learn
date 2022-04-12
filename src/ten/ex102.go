package main

import (
	"archive/tar"
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var compressType = flag.String("compressType", "zip", "choose which compress format（example: zip/tar/unzip)")
var cprPath = flag.String("filePath", "", "the path you want to compress or decompress")

func main() {
	flag.Parse()

	// 创建压缩文件，取出文件名
	name := (*cprPath)[strings.LastIndex(*cprPath, string(filepath.Separator))+1:]

	switch *compressType {
	case "tar":
		f, err := os.Create(name + ".tar")
		defer f.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tw := tar.NewWriter(f)
		defer tw.Close()
		compressTar(*cprPath, tw)
	case "zip":
		f, err := os.Create(name + ".zip")
		defer f.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tw := zip.NewWriter(f)
		defer tw.Close()
		compressZip(*cprPath, tw)
	case "unzip":
		err := unZip(*cprPath)
		fmt.Println(err)
	}
}

func compressTar(path string, w *tar.Writer) {
	filepath.Walk(path, func(curPath string, info fs.FileInfo, errBack error) (err error) {
		h, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return
		}

		// 修改文件头为相对名称
		h.Name = strings.TrimPrefix(curPath, path)
		// 文件夹需要添加/，表示该为一个文件夹否则无法解压表示该为文件夹否则打开不是目录
		if info.IsDir() {
			h.Name += string(filepath.Separator)
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		// 写入文件头
		w.WriteHeader(h)
		// 写入文件内容
		f, err := os.Open(path)
		if err != nil {
			return
		}
		_, err = io.Copy(w, f)
		return
	})
}

func compressZip(path string, w *zip.Writer) {
	filepath.Walk(path, func(curPath string, info fs.FileInfo, errBack error) (err error) {
		h, err := zip.FileInfoHeader(info)
		if err != nil {
			return
		}

		// 修改文件头中的文件名为该zip的绝对路径，此步很重要
		h.Name = strings.TrimPrefix(curPath, path)
		if info.IsDir() {
			h.Name += "/"
		}

		// 写入文件头
		wr, err := w.CreateHeader(h)
		if err != nil {
			return
		}

		// 检测，如果不是标准文件就只写入头信息，不写入文件数据到 w
		// 如目录，也没有数据需要写
		if !info.Mode().IsRegular() {
			return nil
		}

		// 打开要压缩的文件内
		fr, err := os.Open(curPath)
		if fr != nil {
			return
		}
		defer fr.Close()

		// 写入文件内容
		io.Copy(wr, fr)
		return nil
	})
}

func unZip(zipPath string) (err error) {
	if !strings.HasSuffix(zipPath, ".zip") {
		fmt.Println("must be zip format")
		return
	}

	//fInfo, err := os.Stat(zipPath)
	//if err != nil {
	//	return
	//}
	//
	//fi, err := os.Open(zipPath)
	//if err != nil {
	//	return
	//}
	//
	//reader, err := zip.NewReader(fi, fInfo.Size())
	//if err != nil {
	//	return
	//}
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return
	}
	defer reader.Close()

	for _, file := range reader.File {
		fp := file.Name

		if file.FileInfo().IsDir() {
			if err = os.MkdirAll(fp, 0755); err != nil {
				return
			}
			continue
		}

		fd, err := os.Create(fp)
		defer fd.Close()
		if err != nil {
			continue
		}

		rc, err := file.Open()
		defer rc.Close()
		if err != nil {
			continue
		}

		n, err := io.Copy(fd, rc)
		if err != nil {
			continue
		}

		fmt.Printf("成功解压%s， 共写入%d个字符的数据\n", fp, n)
	}
	return nil
}
