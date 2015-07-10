INFILES = $(wildcard *.JPG)
OUTPUT = out
IMG = $(patsubst %.JPG,$(OUTPUT)/%.JPG,$(INFILES))

all: $(OUTPUT) $(IMG)

$(OUTPUT):
	mkdir $(OUTPUT)

$(OUTPUT)/%.JPG: %.JPG
	convert "$<" -gravity SouthEast -font Nimbus-Sans -pointsize 88 -fill white -stroke black -annotate +30+30  "blah blah blah blah" $@

clean:
	rm -rf $(OUTPUT)
