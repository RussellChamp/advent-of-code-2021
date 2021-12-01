default: build day=1
	# This is my 2021 Advent of Code project

init:
	mkdir day-$(day); sed 's/#DAY#/$(day)/g' main.go.template | sed 's/#TITLE#/$(title)/g' > day-$(day)/main.go; touch day-$(day)/input.txt

build:
	cd day-$(day); go run .

test:
	cd day-$(day); go test