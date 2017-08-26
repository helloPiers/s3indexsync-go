package sync // import "hellopiers.io/s3indexsync/internal/sync"

import (
	"fmt"
	"mime"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var s3re = regexp.MustCompile(`s3://([^/]+)(/.*)?`)

func Do(dir, s3uri string) error {

	sm := s3re.FindStringSubmatch(s3uri)

	if len(sm) != 3 { // whole match in sm[0]
		return fmt.Errorf("couldn't parse <s3uri> %s", s3uri)
	}

	sess := session.New()

	s := &syncer{
		s3:     s3.New(sess),
		dir:    dir,
		bucket: sm[1],
		prefix: strings.Trim(sm[2], "/"),
	}

	fmt.Println(s)

	return filepath.Walk(dir, s.walkFn)
}

type syncer struct {
	s3     *s3.S3
	dir    string
	bucket string
	prefix string
}

func (s *syncer) walkFn(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	relp, err := filepath.Rel(s.dir, p)
	if err != nil {
		return err
	}

	mimeType := mime.TypeByExtension(filepath.Ext(p))
	if mimeType == "" {
		mimeType = "text/html" // suits me
	}

	key := path.Join(s.prefix, filepath.ToSlash(relp))
	err = s.upload(p, key, info.Size(), mimeType)
	if err != nil {
		return err
	}

	if path.Base(key) == "index.html" && key != "index.html" {
		key := filepath.Dir(key) + "/"
		err := s.upload(p, key, info.Size(), mimeType)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *syncer) upload(p string, key string, length int64, ctype string) error {

	fmt.Printf("%s => bucket:%s key:%s (%s)\n", p, s.bucket, key, ctype)

	f, err := os.Open(p)
	if err != nil {
		return err
	}
	defer f.Close()

	i := &s3.PutObjectInput{
		Bucket:        &s.bucket,
		Key:           &key,
		ContentLength: &length,
		ContentType:   &ctype,
		Body:          f,
	}

	_, err = s.s3.PutObject(i)
	return err
}
