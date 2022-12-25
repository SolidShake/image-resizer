package watermark

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"go.uber.org/zap"
)

func AddWatermark(logger *zap.SugaredLogger, dir string, selection []string) bool {
	var success int
	for _, file := range selection {
		f, err := os.Open(file)
		if err != nil {
			logger.Errorf("file open error: %s, %w", file, err)
			continue
		}
		defer f.Close()
		img, err := jpeg.Decode(f)
		if err != nil {
			logger.Errorf("jpeg decode: %w", err)
			continue
		}

		wmb, err := os.Open("watermark.png")
		if err != nil {
			logger.Errorf("open watermark: %w", err)
			continue
		}
		watermark, err := png.Decode(wmb)
		if err != nil {
			logger.Errorf("decode watermark: %w", err)
			continue
		}
		defer wmb.Close()

		offset := image.Pt(200, 200)
		b := img.Bounds() // ? missing SOI marker
		m := image.NewRGBA(b)
		draw.Draw(m, b, img, image.ZP, draw.Src)
		draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

		filePath := strings.Split(file, "/")
		fileName := filePath[len(filePath)-1]

		imgw, err := os.Create(dir + "/" + fileName)
		if err != nil {
			logger.Errorf("create file result: %w", err)
			return false
		}
		jpeg.Encode(imgw, m, &jpeg.Options{jpeg.DefaultQuality})
		defer imgw.Close()

		success++

		// err = ioutil.WriteFile(a.dir+"/"+fileName, f, 0644)
		// if err != nil {
		// 	fmt.Println(err)
		// }
	}

	return success > 0
}
