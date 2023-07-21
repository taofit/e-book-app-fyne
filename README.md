# eBook-App-Fyne

## Description

This project is a cross platform ebook application made in Fyne toolkit. It can be run in mobile and desktop. The app has some books pre-installed, now it contains non-fiction and fiction books, of course book categories can be changed and user can add or remove books by modifying `internal/articles/assets/articles_index.json` file and add text files that contains the book text to `internal/articles/assets` folder. A later step in the tutorial will explain how to add more books.

## Run the program

Some instructions are in Fyne's web link: https://developer.fyne.io/started/

Besides Go is required in the system, you need to download the Fyne module and helper tool, and the following commands are needed

```
$ go get fyne.io/fyne/v2@latest
$ go install fyne.io/fyne/v2/cmd/fyne@latest
$ go run .
or
$ go run -tags mobile main.go
```

`go run .` is for desktop version <br>
`go run -tags mobile main.go` is to simulate a mobile application

## Package the desktop version

Instruction is in this link: https://developer.fyne.io/started/packaging

go install fyne.io/fyne/v2/cmd/fyne@latest

- for macOS

```
fyne package -os darwin -icon myapp.png
```

- for linux and window version

```
fyne package -os linux -icon myapp.png
fyne package -os windows -icon myapp.png
```

### Package the mobile version

Instruction link: https://developer.fyne.io/started/mobile
For Android builds, you must have the Android SDK and NDK installed
For iOS build, you will need Xcode installed on your macOS computer as well as the command line tools optional package.

build commands for android and ios is as follows:

```
fyne package -os android -appID com.example.myapp -icon mobileIcon.png
fyne package -os ios -appID com.example.myapp -icon mobileIcon.png
```

### How to add more books to app

A book consists of chapters, and each chapter may contain some parts

For example:

```
 "key": "bookKeyValue",
 "title": "book Title",
 "tableOfContents": [
    {
        "key":"chapter1Key",
        "title": "chapter 1 title"
    },
    {
        "key":"chapter2Key",
        "title": "chapter 2 title"
    },
 ]
```

key must be a unique in the `articles_index.json` file, if there is no key `tableOfContents` in a json object {key, title} then, this json object depicts a file whose name is key's value, such as if key's value is `chapter2Key`, the file `chapter2Key.txt` should be saved under `internal/articles/assets` folder.
