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
[GNU/Linux x64](https://mega.nz/file/w3BGTJAZ#oWeToZJTKwNaGhVNd-YSCX97cVoyIQHLur8Af2w23DE)

[GNU/Linux ARM](https://mega.nz/file/N3J3SIgD#n9DiCbOCP_hwKV1Cz14WNKQKShDclll8AfCnVYIeiME)

[Windows](https://mega.nz/file/Br5wHKAb#AoWBbMCfeN66Zuo1umWyGRg_8M7ZVlS_dVrTFi_niok)

[MacOS x64](https://mega.nz/file/RnInmATT#FELKFoZ6um6l1jcUbJZzk0N-4vKms3HELB5hEi2fW00)

[MacOS ARM](https://mega.nz/file/dqQVHKAB#Bc2f0ZdQVh1R01ls6y0hgP11voIqVmjLyrPr04TlakE)


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

## License
This software is licensed under the MIT License. See the [LICENSE.md](LICENSE.md) file for more info.