package main

import (
	"encoding/binary"
	"fmt"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func escape_string(input string) string {
	input = strings.Replace(input, "'", "", -1)
	input = strings.Replace(input, "|", "", -1)
	input = strings.Replace(input, "\\", "", -1)
	input = strings.Replace(input, "\"", "", -1)
	input = strings.Replace(input, "\n", " ", -1)
	input = strings.Replace(input, "  ", " ", -1)
	input = strings.TrimSpace(input)

	return input
}

func piper(text string) (io.Reader, error) {
	text = escape_string(text)
	// fileName := hashString(input) + ".wav"
	voice := "/home/sweet/ssd/pipertts/ivona-8-23/amy.onnx"
	cmd := exec.Command("piper-tts", "--model", voice, "--output_file", "-")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	_, err = io.WriteString(stdin, text)
	if err != nil {
		return nil, err
	}

	stdin.Close()
	return stdoutPipe, nil

}
func convert_audio() {
	// Read raw PCM data from input file.
	in, err := os.Open("out.pcm")
	if err != nil {
		log.Fatal(err)
	}

	// Output file.
	out, err := os.Create("output.wav")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// 8 kHz, 16 bit, 1 channel, WAV.
	e := wav.NewEncoder(out, 22050, 16, 1, 1)

	// Create new audio.IntBuffer.
	audioBuf, err := newAudioIntBuffer(in)
	if err != nil {
		log.Fatal(err)
	}
	// Write buffer to output file. This writes a RIFF header and the PCM chunks from the audio.IntBuffer.
	if err := e.Write(audioBuf); err != nil {
		log.Fatal(err)
	}
	if err := e.Close(); err != nil {
		log.Fatal(err)
	}
}

func newAudioIntBuffer(r io.Reader) (*audio.IntBuffer, error) {
	buf := audio.IntBuffer{
		Format: &audio.Format{
			NumChannels: 1,
			SampleRate:  8000,
		},
	}
	for {
		var sample int16
		err := binary.Read(r, binary.LittleEndian, &sample)
		switch {
		case err == io.EOF:
			return &buf, nil
		case err != nil:
			return nil, err
		}
		buf.Data = append(buf.Data, int(sample))
	}
}
func main() {
	fmt.Println("cunt")
	text := "This is a test"
	x, err := piper(text)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("return type from piper: %T\n", x)
	filename := "out.pcm"
	outFile, err := os.Create(filename)
	// handle err
	defer outFile.Close()
	_, err = io.Copy(outFile, x)
	if err != nil {
		fmt.Println(err)
	}

	convert_audio()
}
