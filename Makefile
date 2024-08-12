FIND_TARGET=./bin/myFind
ROTATE_TARGET=./bin/myRotate
WC_TARGET=./bin/myWc
XARGS_TARGET=./bin/myXargs

BUILD=go build -o

all: $(FIND_TARGET)

$(FIND_TARGET):
	$(BUILD) $(FIND_TARGET) cmd/cli/find/main.go

$(ROTATE_TARGET):
	$(BUILD) $(ROTATE_TARGET) cmd/cli/rotate/main.go

$(WC_TARGET):
	$(BUILD) $(WC_TARGET) cmd/cli/wc/main.go

$(XARGS_TARGET)):
	$(BUILD) $(XARGS_TARGET) cmd/cli/xargs/main.go

clean:
	rm -f $(FIND_TARGET)
	rm -f $(ROTATE_TARGET)
	rm -f $(WC_TARGET)
	rm -f $(XARGS_TARGET)



