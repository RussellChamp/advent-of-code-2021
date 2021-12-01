default: 1
	# This is my 2021 Advent of Code project

init:
	mkdir day-$(day); sed 's/#DAY#/$(day)/g' main.go.template | sed 's/#TITLE#/$(title)/g' > day-$(day)/main.go; touch day-$(day)/input.txt

0 1:
	cd day-$@; go run .