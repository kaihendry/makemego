# Make versus GO

Single process, run one after another:

	X1C3:~/tmp/makeme$ time make
	mkdir out
	convert "IMG_8131.JPG" -gravity SouthEast -font Nimbus-Sans -pointsize 88 -fill white -stroke black -annotate +30+30  "blah blah blah blah" out/IMG_8131.JPG
	convert "IMG_8132.JPG" -gravity SouthEast -font Nimbus-Sans -pointsize 88 -fill white -stroke black -annotate +30+30  "blah blah blah blah" out/IMG_8132.JPG
	convert "IMG_8130.JPG" -gravity SouthEast -font Nimbus-Sans -pointsize 88 -fill white -stroke black -annotate +30+30  "blah blah blah blah" out/IMG_8130.JPG

	real    0m1.290s
	user    0m1.280s
	sys     0m0.160s
	X1C3:~/tmp/makeme$ make clean
	rm -rf out

Run in parallel:

	X1C3:~/tmp/makeme$ time make -j4
	mkdir out
	convert "IMG_8131.JPG" -gravity SouthEast -font Nimbus-Sans -pointsize 88 -fill white -stroke black -annotate +30+30  "blah blah blah blah" out/IMG_8131.JPG
	convert "IMG_8132.JPG" -gravity SouthEast -font Nimbus-Sans -pointsize 88 -fill white -stroke black -annotate +30+30  "blah blah blah blah" out/IMG_8132.JPG
	convert "IMG_8130.JPG" -gravity SouthEast -font Nimbus-Sans -pointsize 88 -fill white -stroke black -annotate +30+30  "blah blah blah blah" out/IMG_8130.JPG

	real    0m0.599s
	user    0m1.497s
	sys     0m0.170s

As you can see, `make -j4` runs much faster. But what's the _simplest example_
to emulate this parallel behaviour in go?
