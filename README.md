# Piper-TTS Portable

Piper tts is a faster than real time text to speech program. It requires you to
download the binary, the shared object files and the espeak data. It is generally
run like this:

```sh

echo "Hello World" | piper-tts --model /path/to/voice.onnx --output_raw | aplay -r 22050 -c 1 -f S16_LE -t raw

```

This is generally fine but I wanted to experiment with Squashfs and squashfuse to embed the binary and dependencies
into a single file. Then allow you to run that file and listen to the speech that is created on the fly.
To do this, I've concatenated a squashfs archive, with the squashfuse dependency, and an executable ascii file (shell script)
that controls the execution. It can self extract, self-mount, and run the app all by itself.

## Create the binary

**Insert build instructions here **

```sh
  ./piper.sh "Hello world"
```
