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
[GNU/Linux x64](https://mega.nz/file/p6RkWaDB#rH_d0moQCSDnQ4k_IaV5Pjzp1y9X-vPzNvbypkLkQJU)

[GNU/Linux ARM](https://mega.nz/file/svpBBCya#vgLTCWgdwhaqY6XKA6QOu7OBqsBzgApWMRSSDuyqZ24)

[Windows](https://mega.nz/file/dip31DAZ#KCh8qVafO0m_xIRd0DOHjBxatutJsPIktWhImQkwHxo)

[MacOS x64](https://mega.nz/file/06ZUHL7K#uGkwr7oW2XquK3F_6Bp706z8RB_7tKoPTf85Ns8c6y4)

[MacOS ARM](https://mega.nz/file/Q6IBSTYZ#woTQQL-dSbOB0bX4QfkhPO4QAmGpC-IPH9eCaO4Atqk)


## Version Changes
v1.1.2:
 - Fixed bug where search doesn't pull magnet links

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