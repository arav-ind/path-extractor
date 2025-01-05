## Path Extractor

Extract the required data from an SVG file and write it to a JavaScript file.

### Setup

1. Ensure the input file (`input.svg`) is located in the root directory of the project.

2. **Project structure**:

    ```
    myproject/
    ├── cmd/
    │   └── main.go
    ├── input.svg  # Your input SVG file
    ├── Makefile   # Makefile (optional)
    └── go.mod     # Go module file
    ```

### Running the Program

#### Option 1: Using `make`

If you have `make` installed, you can run the following command:

```bash
make run
```

#### Option 2: Using Go directly

If you don't have `make` installed, you can run the program directly with Go:

```bash
go run cmd/main.go
```

### Output

The program will read the `input.svg` file, extract the required data, and write the output to a JavaScript file.