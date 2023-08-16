package domain_test

import (
	"testing"
	"time"

	"github.com/gs-zoop/encoder-go/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()
	require.Error(t, err)
}

func TestVideoIdIsNotAUUID(t *testing.T) {
	video := domain.NewVideo()
	video.ID = "abc"
	video.ResourceId = "a"
	video.FilePath = "/"
	video.CreatedAt = time.Now()
	err := video.Validate()
	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceId = "a"
	video.FilePath = "/"
	video.CreatedAt = time.Now()
	err := video.Validate()
	require.Nil(t, err)
}
