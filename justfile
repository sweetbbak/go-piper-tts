build:
	bash -c ./rebuild.sh
	\cat squashfs-start.sh squashfuse.tar piper.squashfs > piper.sh 
	chmod +x piper.sh
	./piper.sh "Hello" | aplay -r 22050 -c 1 -f S16_LE -t raw   
