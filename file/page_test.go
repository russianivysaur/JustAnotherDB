package file

import (
	assertPkg "github.com/stretchr/testify/assert"
	"testing"
)

func TestPage(t *testing.T) {
	assert := assertPkg.New(t)
	blockSize := 400
	// creating a page
	page := NewPage(blockSize)
	assert.Equalf(len(page.page), blockSize, "Page buffer allocation error")

	buffer := make([]byte, blockSize)
	page = NewPageWithBuffer(buffer)
	assert.Equalf(page.page, buffer, "Page not using the specified buffer")

	offset := 0
	// integer test
	intData := 1243
	page.setInt(offset, intData)
	assert.Equalf(page.getInt(offset), intData, "Integer data does not match in page at offset %d", offset)

	offset = 10
	// bytes test
	byteData := []byte("This is a test!")
	page.setBytes(offset, byteData)
	assert.Equalf(page.getBytes(offset), byteData, "Byte data does not match in page at offset %d", offset)

	// string test
	stringData := "This is another test!"
	err := page.setString(offset, stringData)
	assert.NoErrorf(err, "Could not get string in page buffer : %v", err)
	assert.Equalf(page.getString(offset), stringData, "String data does not match in page at offset %d", offset)
}