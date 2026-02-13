# goCSVTimeAdd

A simple Go program that reads a CSV file with time entries and outputs the total time.

## Usage

```bash
go run main.go [-file <filename>]
```

Or build and run:

```bash
go build
./goCSVTimeAdd [-file <filename>]
```

The default filename is `time.csv`.

## CSV Format

The CSV file should have the format: `<min>,<sec>,<desc>`

Example:
```csv
5,30,Task 1
10,45,Task 2
2,15,Task 3
```

The program will sum all the minutes and seconds and output the result in `H:MM:SS` format (hours, zero-padded minutes, zero-padded seconds):
```
H:MM:SS
```

## Example

With the sample `time.csv` provided:
```bash
$ ./goCSVTimeAdd
1:06:06
```