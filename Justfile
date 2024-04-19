build:
		go build -o permbits && chmod +x permbits 

install:
		just build
		cp permbits ~/.local/bin/

demo:
		vhs demo.tape
