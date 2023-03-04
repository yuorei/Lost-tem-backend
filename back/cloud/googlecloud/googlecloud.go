package googlecloud

import (
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/storage"
	vision "cloud.google.com/go/vision/apiv1"
)

type GCloud struct {
	vision_cli  *vision.ImageAnnotatorClient
	storage_cli *storage.Client
	ctx         context.Context
}

func NewGoogleCloud(ctx context.Context) (*GCloud, error) {
	storage_cli, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	vision_cli, err := vision.NewImageAnnotatorClient(ctx)
	return &GCloud{
		vision_cli:  vision_cli,
		storage_cli: storage_cli,
		ctx:         ctx,
	}, err
}

func (gc GCloud) UploadImage(data []byte, filename string) error {
	writer := gc.storage_cli.Bucket("lost-item").Object(filename).NewWriter(gc.ctx)
	if n, err := writer.Write(data); err != nil {
		return err
	} else if n != len(data) {
		return errors.New("クラウドのデータとファイルのデータの長さが一致しません")
	}

	if err := writer.Close(); err != nil {
		return err
	}
	return nil
}

func (gc GCloud) ObjectRecognition(filename string) ([]string, error) {
	image := vision.NewImageFromURI(fmt.Sprintf("gs://lost-item/%s", filename))

	objects, err := gc.vision_cli.LocalizeObjects(gc.ctx, image, nil)
	if err != nil {
		return []string{}, err
	}

	obj := []string{}

	for _, o := range objects {
		obj = append(obj, o.Name)
	}
	return obj, nil
}

func (gc GCloud) GetURL(filename string) (string, error) {
	return fmt.Sprintf("https://storage.googleapis.com/lost-item/%s", filename), nil
}

func (gc GCloud) Close() {
	gc.vision_cli.Close()
	gc.storage_cli.Close()
}
