present2html
-----

Convert a presentation from go present format to a statically servable HTML page.

A community day project for GopherCon 2024.

## Usage

```
present2html presentation.slide > presentation.html
```

The resulting HTML has the JavaScript and CSS embedded directly into it (as opposed to referenced as a `<link>`).

Files that are referenced within command invocations (e.g. `.code some_file.go`) will be read relative to the input file.

## Feature Support

This currently supports rendering `.slide` presentations to a single HTML page with all the relevant inline markdown. Support for the command invocations is limited. I have tested both `.code` and `.link` and can confirm they work. I have tested `.image` and confirm is does _not_ (TODO: I want to load the image as base64 right into the rendered page). I have not tested any of the other command invocations.

This library does not support playground features or presenter notes and I have no real plans to add support for them.

I have only tested this with the "new" markdown format and not the legacy present format. I have no reason to suspect it wouldn't work though.

## License

This work is licensed under the terms found in `LICENSE.txt` (MIT License).

Large portions of this have been adapted from code found in [golang.org/x/tools](https://cs.opensource.google/go/x/tools). Per the conditions of their license, the original terms are included in `GOLANG_X_TOOLS_LICENSE.txt`, and I have added comments to the files that were adapted directly.
