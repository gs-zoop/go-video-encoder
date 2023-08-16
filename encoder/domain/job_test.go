package domain_test

import (
	"testing"
	"time"

	"github.com/gs-zoop/encoder-go/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNewJob(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "/"
	video.CreatedAt = time.Now()
	job, err := domain.NewJob("/", "Converted", video)
	require.NotNil(t, job)
	require.Nil(t, err)

}
