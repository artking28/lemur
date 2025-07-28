.SILENT:

OS := $(shell uname -s)
ifeq ($(OS),Windows_NT)
	BIN_DIR := $(USERPROFILE)\bin
	EXE := lemur.exe
	MKDIR := if not exist "$(BIN_DIR)" mkdir "$(BIN_DIR)"
	COPY := copy lemur.exe "$(BIN_DIR)\"
	REMOVE := del "$(BIN_DIR)\lemur.exe"
	CLEAN_RM := del lemur.exe
	# CLEAR_CONSOLE := cls
else
	BIN_DIR := /usr/local/bin
	EXE := lemur
	MKDIR := mkdir -p $(BIN_DIR)
	COPY := cp lemur $(BIN_DIR)/
	REMOVE := rm -f $(BIN_DIR)/lemur
	CLEAN_RM := rm -f lemur
	CHMOD := chmod +x $(BIN_DIR)/$(EXE)
	# CLEAR_CONSOLE := clear
endif

all: install

install:
	go build -o $(EXE) cmd/main.go
	$(MKDIR)
	$(COPY)
	$(if $(CHMOD),$(CHMOD))
	# $(CLEAR_CONSOLE)

uninstall:
	$(REMOVE)

clean:
	$(CLEAN_RM)