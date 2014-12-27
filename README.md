Barrage
------

Self-hosted video streaming site, based around pre-transcoding of video files,
rather than live-transcoding. This makes a CPU-Disk Space tradeoff, using
around 50% more disk space in exchange for barely touching the CPU under max
load, meaning it works flawlessly on extremely under-performant devices, such as
a Raspberry Pi.

Building
------

Compile the CSS and JS files with `go generate` - you need to have `lessc` and `cleancss` 
installed and in your $PATH (both are avaliable from npm).

Ensure you have go.rice installed with
`go install github.com/GeertJohan/go.rice/rice`

Using go.rice to embed static files in the source code, simply `rice embed-go`,
then `CGO_ENABLED=0 go build -a` to create the fully independent binary.
If `static.rice-box.go` exists next to `main.go`, barrage should be able to be
cross-compiled for any platform that Go runs on.

Setup a `config.toml` as the example `config.ex.toml` shows - default values
should be more than fine.

Runtime Dependencies
------

Barrage uses HandbrakeCLI internally for subtitle/audio track extraction and
creation of varying bitrate/resolution video tracks. Eventually, I want to
migrate away from this, but for now, ensure the HandbrakeCLI binary exists in
your $PATH (or set 'handbrake' in your config.toml)

Bugs
------
Without a doubt. Report them [here](https://github.com/fortytw2/barrage/issues)


LICENSE
------

MIT
