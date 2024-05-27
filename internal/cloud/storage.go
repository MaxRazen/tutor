package cloud

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/MaxRazen/tutor/internal/config"
	"google.golang.org/api/option"
)

var client *storage.Client

func PrepareClient(credDir embed.FS, name string) {
	ctx := context.TODO()
	b, err := credDir.ReadFile(name)

	if err != nil {
		log.Fatalln(err.Error())
	}

	opt := option.WithCredentialsJSON(b)
	c, err := storage.NewClient(ctx, opt)

	if err != nil {
		log.Fatalln(err.Error())
	}

	client = c
}

// Generates a signed URL that public accessable
// Specify object name on the bucket and expires in seconds
func SignUrl(objectName string, expires int) string {
	bucket := config.GetEnv(config.STORAGE_BUCKET_NAME, "")
	options := storage.SignedURLOptions{
		Method:  http.MethodGet,
		Expires: time.Now().Add(time.Second * time.Duration(expires)),
	}

	url, err := client.
		Bucket(bucket).
		SignedURL(objectName, &options)

	if err != nil {
		log.Println(err.Error())
	}

	return url
}

func Upload(objectName string, content []byte) error {
	bucket := config.GetEnv(config.STORAGE_BUCKET_NAME, "")
	r := bytes.NewReader(content)

	ctx := context.Background()
	wc := client.Bucket(bucket).Object(objectName).NewWriter(ctx)
	wc.ObjectAttrs.ContentType = resolveContentType(objectName)

	if _, err := io.Copy(wc, r); err != nil {
		log.Println(err.Error())
		return fmt.Errorf("cloud/storage:copy %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("cloud/storage:write %v", err)
	}
	return nil
}

func resolveContentType(filename string) string {
	if strings.HasSuffix(filename, ".ogg") {
		return "audio/ogg"
	}
	return ""
}
