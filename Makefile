build: clean
	go build -o bin/gdextension main.go

clean:
	rm -f bin/gdextension

clean-mock:
	rm -rf foobar

run: build
	./bin/gdextension

mock1: build clean-mock
	./bin/gdextension -no-clean -all -jobs 18 foobar gd421

mock2: build clean-mock
	./bin/gdextension -no-clean foobar gd421

mock3: build clean-mock
	./bin/gdextension -no-clean -all -jobs 18 foobar gd43

mock4: build clean-mock
	./bin/gdextension -no-clean foobar gd43

