package main

type FFmpegConfiguration struct {
	ffmpegBinaryPath string
}

type FFmpegOptions struct {
	stream_loop  int
	inputs       []string
	image_loop   uint
	preset       string
	framerate    uint
	sampleRate   uint
	gop          uint
	threads      uint
	maxRate      uint
	bufferSize   uint
	audioBitrate uint
	videoBitrate uint
	audioCodec   string
	videoCodec   string
	format       string
	output       string
}
