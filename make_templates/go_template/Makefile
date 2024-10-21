# Makefile for Go project

GOCMD = go
GOBUILD = $(GOCMD) build
GORUN = $(GOCMD) run

TARGET = program

.PHONY: run runfile clean

run:
	$(GOBUILD) -o $(TARGET) .
	./$(TARGET)

runfile:
	@if [ -z "$(FILE)" ]; then \
		echo "Usage: make runfile FILE=<filename>"; \
	else \
		$(GORUN) $(FILE); \
	fi

clean:
	rm -f $(TARGET)
