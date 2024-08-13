FIND_TARGET=./bin/myFind
ROTATE_TARGET=./bin/myRotate
WC_TARGET=./bin/myWc
XARGS_TARGET=./bin/myXargs

BUILD=go build -o

all: clean $(FIND_TARGET) $(WC_TARGET) $(XARGS_TARGET) $(ROTATE_TARGET)

$(FIND_TARGET):
	$(BUILD) $(FIND_TARGET) cmd/cli/find/main.go

$(ROTATE_TARGET):
	$(BUILD) $(ROTATE_TARGET) cmd/cli/rotate/main.go

$(WC_TARGET):
	$(BUILD) $(WC_TARGET) cmd/cli/wc/main.go

$(XARGS_TARGET):
	$(BUILD) $(XARGS_TARGET) cmd/cli/xargs/main.go

test_task1_1: $(FIND_TARGET)
	@echo "====./myFind text===="
	@echo "----find all .go files at ./----"
	$(FIND_TARGET) -f -ext 'go' ./

test_task1_2: $(FIND_TARGET)
	@echo "----find all files/dirs/symblinks .go files at ./----"
	$(FIND_TARGET) ./

test_task1_3: $(FIND_TARGET)
	@echo "----find all symblinks .go files at ./----"
	$(FIND_TARGET) -ls ./

test_task1_4: $(FIND_TARGET)
	@echo "----find all dirs .go files at ./----"
	$(FIND_TARGET) -d ./

test_task3_1: $(XARGS_TARGET)
	$(FIND_TARGET) -f -ext 'go' . | $(XARGS_TARGET) $(WC_TARGET) -l

test_task4_1: $(ROTATE_TARGET)
	$(ROTATE_TARGET) file.log

test_task4_2: $(ROTATE_TARGET)
	$(ROTATE_TARGET) -a ./archive file.log file1.log

clean:
	rm -f $(FIND_TARGET)
	rm -f $(ROTATE_TARGET)
	rm -f $(WC_TARGET)
	rm -f $(XARGS_TARGET)
	find . -name "*.tar.gz" | xargs rm -f



