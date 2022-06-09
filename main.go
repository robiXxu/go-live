package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func exec(c *cli.Context) error {
	startLive(FFmpegOptions{
		stream_loop: -1,
		// always mp3 second
		inputs:       c.StringSlice("resource"),
		framerate:    15,
		gop:          30,
		sampleRate:   44100,
		threads:      1,
		image_loop:   1,
		preset:       "medium",
		maxRate:      2500,
		audioBitrate: 2500,
		videoBitrate: 2500,
		bufferSize:   512,
		audioCodec:   "copy",
		videoCodec:   "h264",
		format:       "flv",
		output:       rtmpURL(c.String("streamKey")),
	})

	return nil
}

func main() {
	app := &cli.App{
		Name: "live",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "streamKey",
				Aliases:  []string{"k", "key"},
				Usage:    "Provide your streaming key for Youtube*",
				Required: true,
			},
			&cli.StringSliceFlag{
				Name:     "resource",
				Aliases:  []string{"r"},
				Usage:    "Provide your media resources (image & mp3 - in this order)",
				Required: true,
			},
		},
		Action: exec,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("hiroshima %+v", err)
	}
}
