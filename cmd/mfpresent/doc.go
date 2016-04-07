// mfpresenter project doc.go

/*
mfpresenter plays media files that are found on inserted usb sticks (at first)
and plays them full screen. Rule is that if the stick contains more than one
media file, then the newest is played.

Media files are determined by their extension (windows style). Per type a
player can be specified.

When a new usb stick is inserted, and a newer file is found, and if a file
is being played, then the current file is stopped and the new file is played.

The program expects that inserted usb sticks are automatically mounted. Use
usbmount for this. The program just monitors a directory and when a change
occurs it scans the directory for media files. The newest file is copied to
the cache.

Environment variables control the player.

- MFP_LOCAL_CACHE: file directory where the file to be played is cached.
- MFP_WATCH_DIR_0: direcgtory that is watched for changes. This can be more
than one, increase the number, they must be sequential.
- MFP_PLAYER_BIN_EXT0: the executable to play the files with extension EXT0.
If a player can play more than one file, then this line must be repeated for
each extension.

*/
package main

// TODO: create sub packages for config, player, scanner and watcher.
// TODO: refactor.
