build:
	go build -o bin/godot-utils *.go

run: build
	./bin/godot-utils

create: build
	rm -rf foobar && ./bin/godot-utils -all -jobs 18 foobar gd421

