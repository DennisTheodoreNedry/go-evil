CC := g++
LIB := lib/
SRC := src/
CC_FLAGS := -Wpedantic
BIN := ml

compile:
	$(CC) $(CC_FLAGS) -o $(BIN) -I$(LIB) $(SRC)*.cpp

clean:
	rm $(BIN)	
