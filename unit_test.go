package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertQaToZil(t *testing.T) {
	qa := 1000000000000.0
	ret := FromQa(qa, ZIL, false)
	assert.Equal(t, 1.0, ret)
}

func TestConvertQaToLi(t *testing.T) {
	qa := 1000000.0
	ret := FromQa(qa, LI, false)
	assert.Equal(t, 1.0, ret)
}

func TestConvertLiToQa(t *testing.T) {
	val := 1.0
	ret := ToQa(val, LI)
	assert.Equal(t, 1000000.0, ret)
}

func TestConvertZilToQa(t *testing.T) {
	val := 1.0
	ret := ToQa(val, ZIL)
	assert.Equal(t, 1000000000000.0, ret)
}

func TestFromQaNegativeNumber(t *testing.T) {
	val := -1000000000000.0
	ret := FromQa(val, ZIL, false)
	assert.Equal(t, -1.0, ret)
}

func TestFromQaWithPack(t *testing.T) {
	val := 1000000000001.0
	ret := FromQa(val, ZIL, true)
	assert.Equal(t, 1.0, ret)
}

func TestToQaNegativeNumber(t *testing.T) {
	val := -1.0
	ret := ToQa(val, ZIL)
	assert.Equal(t, -1000000000000.0, ret)
}
