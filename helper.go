package main

import "fmt"

const primaryStreamServerYoutube string = "rtmp://a.rtmp.youtube.com/live2"

func rtmpURL(key string) string {
	return fmt.Sprintf("%s/%s", primaryStreamServerYoutube, key)
}
