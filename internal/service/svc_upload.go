package service

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"

	"github.com/disintegration/imaging"
	"github.com/gen2brain/avif"
	"github.com/google/uuid"
	_ "github.com/gookit/goutil/dump"
	"github.com/pkg/errors"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"

	"github.com/haierkeys/obsidian-image-api-gateway/global"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/cloud_storage"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/cloud_storage/aws_s3"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/cloud_storage/cloudflare_r2"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/cloud_storage/local_fs"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/cloud_storage/oss"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/fsutil"
)

type FileInfo struct {
	ImageTitle string `json:"imageTitle"`
	ImageUrl   string `json:"imageUrl"`
}

type ClientUploadParams struct {
	Key    string `form:"key"`
	Type   string `form:"type"`
	Width  int    `form:"width"`
	Height int    `form:"height"`
}

// UploadFile 上传文件
func (svc *Service) UploadFile(fileType fsutil.FileType, file multipart.File, fileHeader *multipart.FileHeader, form *ClientUploadParams) (*FileInfo, error) {

	var fileName string

	// dump.P(fileHeader)

	// 通过剪切板上传的附件 都是一个默认名字
	if fileHeader.Filename == "image.png" {
		fileName = fsutil.GetFileName(uuid.New().String() + fileHeader.Filename)
	} else {
		fileName = fsutil.GetFileName(fileHeader.Filename)
	}

	cType := fileHeader.Header.Get("Content-Type")

	if !fsutil.CheckContainExt(fileType, fileName, global.Config.App.UploadAllowExts) {
		return nil, errors.New("file suffix is not supported.")
	}
	if fsutil.CheckMaxSize(fileType, file, global.Config.App.UploadMaxSize) {
		return nil, errors.New("exceeded maximum file limit.")
	}

	fileKey := fsutil.GetSavePreDirPath() + fileName

	var up = make(map[string]cloud_storage.CloudStorage)
	var dstFileKey string

	writer := &bytes.Buffer{}

	// 压缩
	_, err := file.Seek(0, 0)

	img, filetype, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	size := img.Bounds().Size()

	// 默认裁剪 | 居中裁剪 | 固定尺寸拉伸 | 固定尺寸等比缩放不裁切 | 不处理
	// type: "fill-topleft" | "fill-center" | "resize" | "fit" | "none";

	// 服务器强制限制图片的宽度和高度
	var imageMaxWidth = global.Config.App.ImageMaxSizeWidth
	var imageMaxHeight = global.Config.App.ImageMaxSizeHeight
	var newWidth, newHeight int
	var newImage image.Image
	var isNewImage bool

	if form.Type == "none" || form.Type == "" {

		newWidth = imageMaxWidth
		newHeight = imageMaxHeight

		if (size.X != newWidth || size.Y != newHeight) && (newWidth != 0 || newHeight != 0) {

			if newWidth == 0 || newHeight == 0 {
				newImage = imaging.Resize(img, newWidth, newHeight, imaging.Lanczos)
			} else {
				newImage = imaging.Fit(img, newWidth, newHeight, imaging.Lanczos)
			}

			isNewImage = true
		}
	} else if form.Type == "fill-topleft" {
		if form.Width < imageMaxWidth || imageMaxWidth == 0 {
			newWidth = form.Width
		} else {
			newWidth = imageMaxWidth
		}
		if form.Height < imageMaxHeight || imageMaxHeight == 0 {
			newHeight = form.Height
		} else {
			newHeight = imageMaxHeight
		}

		newImage = imaging.Fill(img, newWidth, newHeight, imaging.TopLeft, imaging.Lanczos)
		isNewImage = true
	} else if form.Type == "fill-center" {
		if form.Width < imageMaxWidth || imageMaxWidth == 0 {
			newWidth = form.Width
		} else {
			newWidth = imageMaxWidth
		}
		if form.Height < imageMaxHeight || imageMaxHeight == 0 {
			newHeight = form.Height
		} else {
			newHeight = imageMaxHeight
		}
		// newImage = imaging.Fit(img, newWidth, newHeight, imaging.Lanczos)
		newImage = imaging.Fill(img, newWidth, newHeight, imaging.Center, imaging.Lanczos)
		isNewImage = true
	} else if form.Type == "resize" {

		if form.Width < imageMaxWidth || imageMaxWidth == 0 {
			newWidth = form.Width
		} else {
			newWidth = imageMaxWidth
		}
		if form.Height < imageMaxHeight || imageMaxHeight == 0 {
			newHeight = form.Height
		} else {
			newHeight = imageMaxHeight
		}

		if form.Width != 0 && form.Height != 0 && (size.X != newWidth || size.Y != newHeight) {
			newImage = imaging.Resize(img, newWidth, newHeight, imaging.Lanczos)
			isNewImage = true
		}
	} else if form.Type == "fit" {

		if form.Width < imageMaxWidth || imageMaxWidth == 0 {
			newWidth = form.Width
		} else {
			newWidth = imageMaxWidth
		}
		if form.Height < imageMaxHeight || imageMaxHeight == 0 {
			newHeight = form.Height
		} else {
			newHeight = imageMaxHeight
		}

		if (size.X != newWidth || size.Y != newHeight) && (newWidth != 0 || newHeight != 0) {

			if newWidth == 0 || newHeight == 0 {
				newImage = imaging.Resize(img, newWidth, newHeight, imaging.Lanczos)
			} else {
				newImage = imaging.Fit(img, newWidth, newHeight, imaging.Lanczos)
			}

			isNewImage = true
		}
	}

	if isNewImage {

		// 调整图片大小

		switch filetype {
		case "png":
			err = png.Encode(writer, newImage)
		case "gif":
			err = gif.Encode(writer, newImage, &gif.Options{NumColors: 256})
		case "jpeg", "jpg":
			err = jpeg.Encode(writer, newImage, &jpeg.Options{Quality: global.Config.App.ImageQuality})
		case "bmp":
			err = bmp.Encode(writer, newImage)
		case "tif", "tiff":
			err = tiff.Encode(writer, newImage, nil)
		case "webp":
			cType = "image/jpg"
			ext := fsutil.GetFileExt(fileKey)
			fileKey = fileKey[0:len(fileKey)-len(ext)] + ".jpg"

			err = jpeg.Encode(writer, newImage, &jpeg.Options{Quality: global.Config.App.ImageQuality})
		case "avif":
			err = avif.Encode(writer, newImage, avif.Options{Quality: global.Config.App.ImageQuality})

		default:
			return nil, errors.New("Unknown image type:" + filetype)
		}

		if err != nil {
			return nil, err
		}

	} else {
		file.Seek(0, 0)
		io.Copy(writer, file)
	}

	reader := bytes.NewReader(writer.Bytes())

	for _, v := range []string{"local_fs", "oss", "cloudflare_r2", "aws_s3"} {

		if v == "local_fs" && global.Config.LocalFS.Enable {

			up[v] = new(local_fs.LocalFS)

		} else if v == "oss" && global.Config.OSS.Enable {

			c, _ := oss.NewClient()
			up[v] = &oss.OSS{
				Client: c,
			}
		} else if v == "cloudflare_r2" && global.Config.CloudfluR2.Enable {

			c, _ := cloudflare_r2.NewClient()

			up[v] = &cloudflare_r2.R2{
				S3Client: c,
			}
		} else if v == "aws_s3" && global.Config.AWSS3.Enable {

			c, _ := aws_s3.NewClient()

			up[v] = &aws_s3.S3{
				S3Client: c,
			}

		} else {
			continue
		}

		reader.Seek(0, 0)

		var err error
		dstFileKey, err = up[v].SendFile(fileKey, reader, cType)
		if err != nil {
			return nil, err
		}

	}

	accessUrl := fsutil.PathSuffixCheckAdd(global.Config.App.UploadUrlPre, "/") + fsutil.UrlEscape(dstFileKey)

	return &FileInfo{ImageTitle: fileHeader.Filename, ImageUrl: accessUrl}, nil
}
func MemDupReader(r io.Reader) func() io.Reader {
	b := bytes.NewBuffer(nil)
	t := io.TeeReader(r, b)

	return func() io.Reader {
		br := bytes.NewReader(b.Bytes())
		return io.MultiReader(br, t)
	}
}
