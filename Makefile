# Set the default goal
.DEFAULT_GOAL=all

# A phony target is one that is not really the name of a file; rather it is just a name for a recipe to be executed when you make an explicit request. There are two reasons to use a phony target: to avoid a conflict with a file of the same name, and to improve performance.
.PHONY: all protocol fmt clean mostlyclean distclean realclean clobber install print tar shar dist TAGS check test

# Make all the top-level targets the makefile knows about.


all: protocol fmt

# Generate all protocol
PROTO_PATH=pb
PROTO_FILES=$(wildcard $(PROTO_PATH)/*.proto)

PB_GO_PATH=pb
PB_GO_FILES=$(patsubst $(PROTO_PATH)/%.proto,$(PB_GO_PATH)/%.pb.go,$(PROTO_FILES))

TEST_PATH=test
PB_PY_FILES=$(patsubst $(PROTO_PATH)/%.proto,$(TEST_PATH)/%_pb2.py,$(PROTO_FILES))

protocol: $(PB_GO_FILES) $(PB_PY_FILES)

# *.proto to *.pb.go
$(PB_GO_PATH)/%.pb.go: $(PROTO_PATH)/%.proto
	@echo generating $@
	@protoc --proto_path=$(PROTO_PATH)/ --gofast_out=plugins=grpc:$(PB_GO_PATH)/ $(patsubst $(PB_GO_PATH)/%.pb.go,$(PROTO_PATH)/%.proto,$@)

$(TEST_PATH)/%_pb2.py: $(PROTO_PATH)/%.proto
	@echo generating $@
	@python -m grpc_tools.protoc -I$(PROTO_PATH)/ --python_out=$(TEST_PATH)/ --grpc_python_out=$(TEST_PATH)/ $(patsubst $(TEST_PATH)/%_pb2.py,$(PROTO_PATH)/%.proto,$@)

# Format all sources
fmt: GO_FILES=$(wildcard *.go) $(wildcard */*.go) $(wildcard */*/*.go)
fmt:
	goimports -w $(GO_FILES)
	clang-format -i $(PROTO_FILES)

# Delete all files that are normally created by running make.
clean:
	rm -f $(PB_GO_PATH)/*.pb.go $(TEST_PATH)/*_pb2*.py

# Like ‘clean’, but may refrain from deleting a few files that people normally don’t want to recompile. For example, the ‘mostlyclean’ target for GCC does not delete libgcc.a, because recompiling it is rarely necessary and takes a lot of time.
mostlyclean:

# Any of these targets might be defined to delete more files than ‘clean’ does. For example, this would delete configuration files or links that you would normally create as preparation for compilation, even if the makefile itself cannot create these files.
distclean:
realclean:
clobber:
	git clean -dxf

# Copy the executable file into a directory that users typically search for commands; copy any auxiliary files that the executable uses into the directories where it will look for them.
install:

# Print listings of the source files that have changed.
print:
	@echo $(PROTO_FILES)
	@echo $(PB_GO_FILES)
	@echo $(PB_PY_FILES)

# Create a tar file of the source files.
tar:

# Create a shell archive (shar file) of the source files.
shar:

# Create a distribution file of the source files. This might be a tar file, or a shar file, or a compressed version of one of the above, or even more than one of the above.
dist:

# Update a tags table for this program.
TAGS:

# Perform self tests on the program this makefile builds.
check:
test: protocol
	@cd $(TEST_PATH) && python client.py