# music-junk-cleaner

It scans the music directory provided in input, rename them by removing the junk data. If it finds a new junk data pattern, it will ask you to mark it, and it will clean(rename) such kind of patterns.

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
