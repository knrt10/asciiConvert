<p align="center">
  <img src="https://user-images.githubusercontent.com/24803604/45256385-603e1780-b3b3-11e8-83e5-a1f366de844c.png" />
</p>

> Convert images to ascii art using your command line. This was just me getting some start on writing some Go-Code.

## Installation

```
go get github.com/knrt10/asciiConvert
```

## Features

- Get images ascii Art
- GIF support

## Usage

`asciiConvert --help` 

```
Usage:
  asciiConvert [flags]

Flags:
  -h, --help          help for asciiArt
  -p, --path string   path of your file for which you want to convert ASCII Art
  -w, --width int     width of final file (default 100)
```

## Command

```
asciiConvert -p "path to file" // this will print with width 100

or

asciiConvert -p "path to file" -w 150
```

## Preview
![preview](https://user-images.githubusercontent.com/24803604/45258693-a4dca980-b3d9-11e8-9935-aa33646a16e6.gif)

## GIF support 

![preview](https://user-images.githubusercontent.com/22731013/67517637-a1c97900-f6c0-11e9-8da7-c5cf93261b44.gif)

## Contributors

- [@knrt10](https://github.com/knrt10)
- [@S Santosh Nagaraj](https://github.com/yolossn)
