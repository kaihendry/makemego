INFILES = $(shell shopt -s globstar; for i in *.JPG; do echo $$i; done)
OUTPUT = out
IMG = $(patsubst %.JPG,$(OUTPUT)/%.JPG,$(INFILES))

all: $(OUTPUT) $(IMG)

$(OUTPUT):
	mkdir $(OUTPUT)

$(OUTPUT)/%.JPG: %.JPG
	convert "$<" -gravity SouthEast -font Nimbus-Sans -pointsize 88 -fill white -stroke black -annotate +30+30  "http://prazefarm.co.uk" $@

clean:
	rm -rf $(OUTPUT)