package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-cmd/cmd"
)

func generateArgs(options FFmpegOptions) []string {
	args := []string{
		// "-y",
		"-loop", fmt.Sprintf("%d", options.loopAmount),
		// "-loop_input", "1",
	}
	for _, input := range options.inputs {
		args = append(args, "-i", input)
	}

	return append(args, []string{
		"-preset", options.preset,
		"-r", fmt.Sprintf("%d", options.framerate),
		"-g", fmt.Sprintf("%d", options.gop),
		"-ar", fmt.Sprintf("%d", options.sampleRate),
		"-threads", fmt.Sprintf("%d", options.threads),
		"-maxrate", fmt.Sprintf("%dk", options.maxRate),
		"-b:a", fmt.Sprintf("%d", options.audioBitrate),
		"-b:v", fmt.Sprintf("%dk", options.videoBitrate),
		"-bufsize", fmt.Sprintf("%dk", options.bufferSize),
		"-vcodec", options.videoCodec,
		"-acodec", options.audioCodec,
		"-f", options.format,
		options.output,
	}...)
}

func startLive(options FFmpegOptions) {
	config := FFmpegConfiguration{
		ffmpegBinaryPath: "/usr/bin/ffmpeg",
	}

	live := cmd.NewCmd(config.ffmpegBinaryPath, generateArgs(options)...)
	status := live.Start()

	tick := time.NewTicker(300 * time.Millisecond)

	go func() {
		s := live.Status()
		log.Printf("[CMD] %s", s.Cmd)
		log.Printf("[PID] %d", s.PID)
		initialLog := false

		for range tick.C {
			s := live.Status()
			cErr := len(s.Stderr) - 1
			if !initialLog {
				for _, v := range s.Stderr {
					log.Printf("%+v", v)
				}
				initialLog = true
			}
			if cErr > 0 {
				log.Printf("%+v", s.Stderr[cErr])
			}
		}
	}()

	finalStatus := <-status

	log.Println(finalStatus)
}
