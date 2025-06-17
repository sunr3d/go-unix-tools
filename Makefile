.PHONY: all clean myFind myWc myXargs myRotate

BINARY_DIR = bin
SRC_DIR = ./cmd

all: clean myFind myWc myXargs myRotate

myFind:
	mkdir -p $(BINARY_DIR)
	go build -o $(BINARY_DIR)/myFind $(SRC_DIR)/myFind

myWc:
	mkdir -p $(BINARY_DIR)
	go build -o $(BINARY_DIR)/myWc $(SRC_DIR)/myWc

myXargs:
	mkdir -p $(BINARY_DIR)
	go build -o $(BINARY_DIR)/myXargs $(SRC_DIR)/myXargs

myRotate:
	mkdir -p $(BINARY_DIR)
	go build -o $(BINARY_DIR)/myRotate $(SRC_DIR)/myRotate

clean:
	rm -rf $(BINARY_DIR)
