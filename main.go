// Made by GPT-4. See `README.md` for details.
package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

const (
	C = 261.63
	D = 293.66
	E = 329.63
	F = 349.23
	G = 392.00
	A = 440.00
	B = 493.88
)

type SineWave struct {
	freq  float64
	srate beep.SampleRate
	step  float64
}

func (sw *SineWave) Stream(samples [][2]float64) (n int, ok bool) {
	for i := range samples {
		samples[i][0] = math.Sin(2 * math.Pi * sw.freq * sw.step)
		samples[i][1] = samples[i][0]
		sw.step += 1.0 / float64(sw.srate)
	}
	return len(samples), true
}

func (sw *SineWave) Err() error {
	return nil
}

type tone struct {
	freq  float64
	delay time.Duration
}

func main() {
	jingleBells := []tone{
		{E, time.Millisecond * 500},
		{E, time.Millisecond * 500},
		{E, time.Millisecond * 1000},
		{E, time.Millisecond * 500},
		{E, time.Millisecond * 500},
		{E, time.Millisecond * 1000},
		{E, time.Millisecond * 500},
		{G, time.Millisecond * 500},
		{C, time.Millisecond * 500},
		{D, time.Millisecond * 500},
		{E, time.Millisecond * 1000},
		{F, time.Millisecond * 500},
		{F, time.Millisecond * 500},
		{F, time.Millisecond * 500},
		{F, time.Millisecond * 500},
		{F, time.Millisecond * 500},
		{E, time.Millisecond * 500},
		{E, time.Millisecond * 500},
		{E, time.Millisecond * 500},
		{E, time.Millisecond * 500},
		{D, time.Millisecond * 500},
		{D, time.Millisecond * 500},
		{E, time.Millisecond * 500},
		{D, time.Millisecond * 1000},
		{G, time.Millisecond * 1000},
	}

	err := speaker.Init(44100, 44100/10)
	if err != nil {
		panic(err)
	}

	exit := make(chan os.Signal, 1)

	go func() {
		fmt.Println("Press ENTER to stop playing")
		fmt.Scanln()
		exit <- os.Interrupt
	}()

playLoop:
	for {
		for _, t := range jingleBells {
			select {
			case <-exit:
				break playLoop
			default:
				samples := int(float64(44100) * t.delay.Seconds())
				sineWave := &SineWave{freq: t.freq, srate: beep.SampleRate(44100)}
				speaker.Play(beep.Take(samples, sineWave))
				time.Sleep(t.delay)
			}
		}
	}
}
