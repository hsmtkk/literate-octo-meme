package zip

import (
	"fmt"

	"github.com/mholt/archiver/v3"
)

type Unzipper interface {
	Unzip(src, dst string) error
}

type unzipperImpl struct{}

func New() Unzipper {
	return &unzipperImpl{}
}

func (u *unzipperImpl) Unzip(src, dst string) error {
	if err := archiver.Unarchive(src, dst); err != nil {
		return fmt.Errorf("failed to extract zip from %s to %s; %w", src, dst, err)
	}
	return nil
}
