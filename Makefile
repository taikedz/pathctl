build: bin/pathctl
	bash ./build.sh


# An example of hot to use `pathctl` in a makefile
install: build
	dest=$(bin/pathctl)
	cp bin/pathctl $dest/pathctl
