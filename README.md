# r3c -- a wrapper of rsync

## Motivation

I always have problem to remember the weired semantic difference between `rsync -r dir1/ ...` and `rsync -r dir ...`.

Thus I wrote this wrapper.

## Usage

    r3c dir1 dir2

will sync `dir1` and `dir2`. That's it.

In some cases, if you do not care about unix features like permissions, links, etc, you can add an `--simple` option:

    r3c -simple dir1 dir2

For rsync users:

- Without `-simple`, `-a --partial --delete`.
- With `-simple`, same as above, but using `-r` instead of `-a`.
- With `-compress`, passing the `-z` option.
- With `-progress`, passing the `--progress` option.

## License

0BSD



