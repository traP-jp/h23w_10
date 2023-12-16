package imggen

import (
	"context"
	"image"
	_ "image/gif"
	"image/png"
	"os"
	"testing"
)

var usernames = []string{
	"pirosiki",
	"hayatroid",
	"d_etteiu8383",
	"masky5859",
	"Ras",
	"toki",
	"ikura-hamu",
	"cp20",
	"hijiki51",
	"aya_se",
	"s9",
	"H1rono_K",
	"oribe",
	"mazrean",
	"Pugma",
}

var layerConfig = []LayerConfig{
	NewLayerConfig(0, 1, 220),
	NewLayerConfig(200, 8, 130),
	NewLayerConfig(330, 16, 115),
	NewLayerConfig(450, 21, 100),
}

func TestGenerateImage(t *testing.T) {
	svc := NewImggenService(layerConfig)
	baseURL := "https://q.trap.jp/api/v3/public/icon/"
	ch := make(chan image.Image)
	var result image.Image
	var err error
	done := make(chan struct{})
	go func() {
		result, err = svc.GenerateImage(context.Background(), len(usernames), ch)
		if err != nil {
			t.Error(err)
		}
		close(done)
	}()

	t.Log("start generating images")

	for _, username := range usernames {
		img, err := openImage(baseURL + username)
		if err != nil {
			t.Error(err)
		}
		t.Log("sending image")
		ch <- img
	}
	t.Log("finish sending images")
	close(ch)
	<-done

	if result == nil {
		t.Error("result is nil")
	}
	if err != nil {
		t.Error(err)
	}

	f, _ := os.Create("test.png")
	defer f.Close()
	png.Encode(f, result)
}
