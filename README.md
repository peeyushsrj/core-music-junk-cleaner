# music-junk-cleaner

It scans the music directory provided in input, rename music files by removing the junk data. If it finds a new junk data pattern, it will prompt you to mark and then it will clean/rename such patterns.

### Build instructions

- Golang must be installed.
- Installing package `go get github.com/gorilla/websocket`.
- Clone it `git clone https://github.com/peeyushsrj/music-junk-cleaner/`
- Changed directory & build it `go build`.
- Run it! `./music-junk-cleaner /home/user/some_music_directory/`

### Future TODO

- [ ] Launching a browser on start of this pgm.
- [ ] X platform Binaries.
- [ ] Demo gif on readme.
- [ ] Add support to other music formats.
- [ ] Cutting out core from UI.
- [ ] Advertizing for fun & usability.

### License

The MIT License (MIT) Copyright (c) 2017 Peeyush Singh
