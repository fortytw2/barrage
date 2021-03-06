Barrage [![Build Status](https://travis-ci.org/fortytw2/barrage.svg)](https://travis-ci.org/fortytw2/barrage)
------

Self-hosted video streaming site, based around pre-transcoding of video files,
rather than live-transcoding. This makes a CPU-Disk Space tradeoff, using
around 50% more disk space in exchange for barely touching the CPU under max
load, meaning it works flawlessly on extremely under-performant devices, such as
a Raspberry Pi.

However, due to hard disk space limits, Barrage also supports live transcoding

Setup
------

Barrage depends on a particularly structured layout for the source
media. Essentially, you should have something like:
```
MEDIA -
     |- Group Name (Name of this folder can be anything)
     	|- Show 1
     		|- bunch of video files
     		|- series.toml
      |- Show 2
        |- bunch of video files
        |- series.toml
     |- Group Name
     	|- Movie 1
     		|- video file
     		|- movie.toml
```
Having a metadata file makes tracking everything accurately incredibly easy
compared to trying to guess at filenames and foldernames to figure out what's
there. This way, barrage knows exactly what files match up to what episodes.
This requires a bit more manual work than one would want, but maintains accuracy
in the end. I plan on adding a extra tool to assist in generating the metadata
files in the near future.

Once your media is set up as described above, you simply have to set up your
config file. Do this by copying config.ex.toml into config.toml, and editing
`sourcefolder` to be what MEDIA is in the above picture. Feel free to change the
other settings - they do what you'd expect them to do.

Now start barrage by launching the executable - on first run, it may take a long
time to start, as it re-encodes the existing video into 3 streams, low/med/high
quality (high is the source file, typically), instead of encoding it as you
stream the video. This is the biggest compromise barrage makes - sacrificing
disk space (each set of re-encodes shouldn't take up more space than the source)
for extremely low CPU usage - keep this in mind as you import files.

SSH Tunneling
------
`ssh -nNT remote_port:localhost:local_port user@remote_host`

Building
------
Compile the CSS and JS files with `go generate` - you need to have `lessc`,
`uglifyjs`, `cleancss`, and `browserify` installed and in your $PATH

Ensure you have go.rice installed with
`go install github.com/GeertJohan/go.rice/rice`

Use go.rice to embed static files in the source code, simply `rice embed-go`,
then `CGO_ENABLED=0 go build -a` to create the fully independent binary.

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
