package images

import (
	"strings"

	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/daemon"
)

var privateRegistry string = "backupregistry"

func rename(name string) string {
	image := strings.Split(name, ":")
	img, version := image[0], image[1]
	newName := privateRegistry + "/" + img + ":" + version
	return newName
}

func retag(imgName string) (name.Tag, error) {
	tag, err := name.NewTag(imgName)
	if err != nil {
		return name.Tag{}, err
	}
	return tag, nil
}

func Process(imgName string) (string, error) {
	img, err := crane.Pull(imgName)
	if err != nil {
		return "", err
	}
	newName := rename(imgName)
	tag, err := retag(newName)
	if err != nil {
		return "", err
	}

	log, err := daemon.Write(tag, img)
	if err != nil {
		return "", err
	}
	if err := crane.Push(img, tag.String()); err != nil {
		return "", err
	}
	return log, nil
}
