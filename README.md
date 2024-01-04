# r3c -- a wrapper of rsync

## Motivation

I always have problem to remember the weired semantic difference between `rsync -r dir1/ ...` and `rsync -r dir ...`.

Thus, I wrote this wrapper.

## Usage

    r3c dir1 dir2

will sync `dir1` and `dir2`. That's it.

In some cases, if you do not care about unix features like permissions, links, etc, you can add a `-simple` option:

    r3c -simple dir1 dir2

For rsync users:

- Without `-simple`, `-a --partial --delete`.
- With `-simple`, same as above, but using `-r` instead of `-a`.
- With `-nodel`, not passing the `--delete` option.
- With `-compress`, passing the `-z` option.
- With `-progress`, passing the `--progress` option.
- With `-dry`, just print out the generated rsync command line and exit.

## Bugs

Since r3c uses the `--sparse` option of rsync for better support of sparse files, it will corrupt the files on a Solaris "tmpfs" destination.
Please do not use r3c when the destination is on a Solaris "tmpfs" file system.
If you do want to do so, please use r3c `v0.0.0`, which does not use the `--sparse` option.

## Install

Compile from the source and install to `/usr/local/bin`:

```sh
make
make install
```

Depending on your file system permission configuration, you may need to prefix the `make install` command with `sudo`.
Change `PREFIX` to install r3c to another directory.
For example:

```sh
make PREFIX=~/.local install
```

The Makefile is compatible with both GNU and BSD make.

## Uninstall

Just run:

```
make uninstall
```

If you have changed the `PREFIX` variable when installing, you need to use the same value when uninstalling.
For example:

```sh
make PREFIX=~/.local uninstall
```

## License

0BSD



