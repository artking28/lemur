.SILENT:

OS := $(shell uname -s)
ifeq ($(OS),Windows_NT)
	BIN_DIR := $(USERPROFILE)\bin
	EXE := lemur.exe
	MKDIR := cmd /C "if not exist \"$(BIN_DIR)\" mkdir \"$(BIN_DIR)\""
	COPY := cmd /C "copy /Y $(EXE) \"$(BIN_DIR)\""
	REMOVE := cmd /C "del /Q \"$(BIN_DIR)\\$(EXE)\""
	CLEAN_RM := cmd /C "del /Q $(EXE)"
	ECHO := cmd /C echo
else
	BIN_DIR := /usr/local/bin
	EXE := lemur
	MKDIR := mkdir -p $(BIN_DIR)
	COPY := cp $(EXE) $(BIN_DIR)/
	REMOVE := rm -f $(BIN_DIR)/$(EXE)
	CLEAN_RM := rm -f $(EXE)
	CHMOD := chmod +x $(BIN_DIR)/$(EXE)
	ECHO := echo
endif

all: install

install:
	$(ECHO) "Installing lemur..."
	go build -o $(EXE) cmd/main.go cmd/actions.go
	$(MKDIR)
	$(COPY)
	$(if $(CHMOD),$(CHMOD))
	$(ECHO) "ok! Binary is at \"$(BIN_DIR)\""

uninstall:
	$(ECHO) "Uninstalling lemur..."
	$(REMOVE)
	$(ECHO) "ok!"

clean:
	$(CLEAN_RM)
