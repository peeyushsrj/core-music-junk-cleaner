# music-junk-cleaner

It scans the music directory provided in input, rename music files by removing the junk data. If it finds a new junk data pattern, it will prompt you to mark and then it will clean/rename such patterns.

[![Screenshot](https://s20.postimg.org/lrrw1qf65/Screenshot_from_2017-06-23_19-27-01.png)](https://postimg.org/image/66akhs37t/)

### Build instructions

- Golang must be installed.
- Installing package `go get github.com/gorilla/websocket`.
- Clone it `git clone https://github.com/peeyushsrj/music-junk-cleaner/`
- Changed directory & build it `go build`.
- Run the binary file ! `./music-junk-cleaner /home/user/some_music_directory/`
- Open `http://localhost:7899`.


### Future TODO

- [ ] Launching a browser on start of this pgm.
- [ ] X platform Binaries.
- [ ] Demo gif on readme.
- [ ] Add support to other music formats.
- [ ] Cutting out core from UI.
- [ ] Advertizing for fun & usability.

### License

The MIT License (MIT) Copyright (c) 2017 Peeyush Singh
