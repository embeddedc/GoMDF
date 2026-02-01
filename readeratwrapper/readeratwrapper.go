package readeratwrapper

import (
	"fmt"
	"io"
)

type ReaderAtWrapper struct {
	reader io.ReaderAt
	offset int64
}

func NewReaderAtWrapper(reader io.ReaderAt, offset int64) *ReaderAtWrapper {
	return &ReaderAtWrapper{
		reader: reader,
		offset: offset,
	}
}

func (wrapper *ReaderAtWrapper) Read(p []byte) (int, error) {
	n, err := wrapper.reader.ReadAt(p, wrapper.offset)
	wrapper.offset += int64(n)
	return n, err
}

func (wrapper *ReaderAtWrapper) Seek(offset int64, whence int) (int64, error) {
	var newOffset int64

	switch whence {
	case io.SeekStart:
		newOffset = offset
	case io.SeekCurrent:
		newOffset = wrapper.offset + offset
	default:
		return 0, fmt.Errorf("invalid whence value: %d", whence)
	}

	wrapper.offset = newOffset
	return newOffset, nil
}
