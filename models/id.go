package models

import (
	"strconv"
	"strings"
)

// ID id
type ID int64

// Int64 转换为int64
func (i ID) Int64() int64 {
	return int64(i)
}

// String 转换为string
func (i ID) String() string {
	return strconv.FormatInt(int64(i), 10)
}

// Base36 转换为base36
func (i ID) Base36() string {
	return strings.ToUpper(strconv.FormatInt(int64(i), 36))
}
