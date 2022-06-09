package main

type FFmpegConfiguration struct {
	ffmpegBinaryPath string
}

type FFmpegOptions struct {
	inputs       []string
	output       string
	format       string
	loopAmount   uint
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
}
