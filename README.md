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

## Binaries
[GNU/Linux x64](https://mega.nz/file/gnAjBYgK#8GNHQUGlkHV9O-9Gj6nFOloj9EJJvqLfcrBUZJ0VGkk)

[GNU/Linux ARM](https://mega.nz/file/QrRVgZYZ#9nYZmVbkSxgyY1-53yq_4kMqRkvGBxwjwIc7-MArAho)

[Windows](https://mega.nz/file/BuZRDQYZ#SLtb2EhZkX0zXUwy8X5W2cjuRH1OfoWnH8DplfQi0qE)

[MacOS x64](https://mega.nz/file/p2YWnbgC#1etKGLsgdHLrlJ4_-L3vZReKQLIT6PHfU338-C8pjgU)

[MacOS ARM](https://mega.nz/file/puJnnKIQ#GxHPyyvbub5eEu24-b8tFvGw0sPOF3aKce6pgFo2S_0)


## Version Changes
v1.1.1:
 - Fixed imports

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

## License
This software is licensed under the MIT License. See the [LICENSE.md](LICENSE.md) file for more info.