INFILES = $(wildcard *.JPG)
OUTPUT = out
IMG = $(patsubst %.JPG,$(OUTPUT)/%.JPG,$(INFILES))

all: $(OUTPUT) $(IMG)

$(OUTPUT):
	@mkdir $(OUTPUT)

$(OUTPUT)/%.JPG: %.JPG
	@convert "$<" -resize 25% $@
	@echo "$< -> $@"

clean:
	rm -rf $(OUTPUT)
