# EZ-Weeb (G)

This is a program to enable weebs through technology, but for terminal!

## Build

You can use Make:
`make build`

Otherwise, if you have Go installed on your machine:
Linux and Mac:
`go build -o ./bin/${BINARY_NAME}.exe ./cmd/ezweeb/ezweeb`

Windows
`go build -o ./bin/${BINARY_NAME}.exe ./cmd/ezweeb/ezweeb.go`

## Version Changes
v1.1.0:
 - Added a search function. It will now allow you to input a title and get results from Nyaa.
 - Added an option to change the safety level. Can use any or all:
    - Safe (Nyaa CSS ".success")
    - Potentially Dangerous (Nyaa CSS ".danger")
    - Default (Nyaa CSS ".default")

## Contributing

Submit a PR.

## Acknowledgements
* [Dennis Capone](https://github.com/dcap0)
