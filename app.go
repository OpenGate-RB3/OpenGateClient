package main

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os/exec"
	goRuntime "runtime"
)

// App struct
type App struct {
	ctx          context.Context
	audioOutProc *exec.Cmd
	audioInProc  *exec.Cmd
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) shutdown(context.Context) {
	if a.audioOutProc != nil {
		a.audioOutProc.Process.Kill()
	}
	if a.audioInProc != nil {
		a.audioInProc.Process.Kill()
	}
}

// StartAudioReceive starts the audio in Pipeline
func (a *App) StartAudioReceive(ip string) bool {
	fmt.Println("Starting audio receive")
	a.audioInProc = exec.Command("ffplay",
		"-nodisp",
		"-autoexit",
		"-loglevel",
		"quiet",
		"rtsp://"+ip+":8554/audiostream",
	)
	err := a.audioInProc.Start()
	if err != nil {
		return false
	}
	go func(cmd *exec.Cmd, cxt context.Context) {
		err := cmd.Wait()
		runtime.EventsEmit(cxt, "audio_receive_proc", "stopped")
		if err != nil {
			if a.audioInProc != nil {
				a.audioInProc = nil
			}
		}
		return
	}(a.audioInProc, a.ctx)
	runtime.EventsEmit(a.ctx, "audio_receive_proc", "spawned")
	return true
}

func (a *App) StartAudioSend(ip string) bool {
	if a.audioOutProc != nil {
		return false
	}
	switch goRuntime.GOOS {
	case "darwin":
		a.audioOutProc = exec.Command("ffmpeg",
			"-f", "avfoundation",
			"-i", ":0",
			"-af", "volume=2.0",
			"-acodec", "libmp3lame",
			"-ab", "128k",
			"-ac", "1",
			"-f", "mpegts",
			"-pkt_size", "188",
			"udp://"+ip+":9000")
		break
	case "linux":
		a.audioOutProc = exec.Command("ffmpeg",
			"-f", "pulse",
			"-i", "default",
			"-af", "volume=2.0",
			"-acodec", "libmp3lame",
			"-ab", "128k",
			"-ac", "1",
			"-f", "mpegts",
			"-pkt_size", "188",
			"udp://"+ip+":9000")
		break
	default:
		return false
	}

	err := a.audioOutProc.Start()
	if err != nil {
		a.audioOutProc = nil
		return false
	}
	go func(cmd *exec.Cmd, cxt context.Context) {
		err := cmd.Wait()
		runtime.EventsEmit(cxt, "audio_send_proc", "stopped")
		if err != nil {
			if a.audioOutProc != nil {
				a.audioOutProc = nil
			}
			return
		}
	}(a.audioOutProc, a.ctx)
	runtime.EventsEmit(a.ctx, "audio_send_proc", "spawned")
	return true
}

func (a *App) StopAudioReceive() bool {
	if a.audioInProc == nil {
		return false
	}
	err := a.audioInProc.Process.Kill()
	if err != nil {
		return false
	}
	a.audioInProc = nil
	return true
}

func (a *App) StopAudioSend() bool {
	if a.audioOutProc == nil {
		return false
	}
	err := a.audioOutProc.Process.Kill()
	if err != nil {
		return false
	}
	a.audioOutProc = nil
	return true
}
