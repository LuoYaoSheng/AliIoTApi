package util

import (
	"math/rand"
	"time"
)

func ProductNameGenerator() string {
	chars := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	r := len(chars)

	name := "product_"
	length := 4
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<length;i++  {
		name = name + string(chars[rand.Intn(r)])
	}
	return name
}

func DeviceNameGenertor(length int64) map[int64] string  {
	return DeviceNameForPrefixGenertor("",length)
}

func DeviceNameForPrefixGenertor(prefix string,length int64) map[int64] string  {
	chars := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	r := len(chars)

	names := make(map[int64] string, length)

	rand.Seed(time.Now().UnixNano())
	for i := int64(0); i < length;i++  {
		l := int64(5)

		name :=  "dev_"
		if len(prefix)>0 {
			name = name +  prefix + "_"
		}

		for j := int64(0); j < l;j++  {
			name = name +  string(chars[rand.Intn(r)])
		}
		names[i] = name
	}
	return names
}

