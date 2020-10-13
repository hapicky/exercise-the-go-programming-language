package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	_ "image/png"
	"io"
	"os"
)

type encoderFlag struct {
	encoder func(in io.Reader, out io.Writer) error
}

func (f *encoderFlag) String() string {
	// 怒られたから実装したけど、これ何だろう？
	return "[encoder]"
}

func (f *encoderFlag) Set(s string) error {
	switch s {
	case "jpeg":
		f.encoder = toJPEG
		return nil
	case "gif":
		f.encoder = toGIF
		return nil
	}
	return fmt.Errorf("invalid encoder %q", s)
}

func EncoderFlag(name string, value func(in io.Reader, out io.Writer) error, usage string) func(in io.Reader, out io.Writer) error {
	f := encoderFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return f.encoder
}

// TODO これだと引数指定しても参照が置き換わらない？
var encode = EncoderFlag("encoder", toJPEG, "the encoder('jpeg' OR 'png' OR 'gif'. default: jpeg)")

func main() {
	flag.Parse()
	if err := encode(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toGIF(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return gif.Encode(out, img, &gif.Options{})
}
