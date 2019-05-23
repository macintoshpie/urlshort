## tldr;
Toy project for learning web services with Golang; 

## Setup
```bash
go get github.com/macintoshpie/urlshort
go build -o urlshort $GOHOME/src/github.com/macintoshpie/urlshort
```

## Usage
After getting the repo building the binary, you can run the program like so:
```bash
./urlshort [-dbname <name>] [-json <path>] [-port <port>]
  -dbname string
        Name of database where shortened urls are stored
  -json string
        Path to json file for preconfiguring redirects (zero time for fetching)
  -port uint
        Port on which to expose the API (default 8080)
```
