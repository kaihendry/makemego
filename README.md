# Make versus GO

	X1C3:~/tmp/makeme$ make clean
	rm -rf out
	X1C3:~/tmp/makeme$ time make -j8 | grep -v IMG

	real    0m7.349s
	user    0m26.333s
	sys     0m2.180s
	X1C3:~/tmp/makeme$ make clean
	rm -rf out
	X1C3:~/tmp/makeme$ time ./makeme | grep -v IMG

	real    0m7.450s
	user    0m26.577s
	sys     0m2.483s
	X1C3:~/tmp/makeme$

