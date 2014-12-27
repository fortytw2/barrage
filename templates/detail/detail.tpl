{{ define "title" }} Barrage | Eva {{ end }}


{{ define "body" }}

<video id="detail" class="video-js vjs-default-skin"
controls preload="auto"
poster="http://video-js.zencoder.com/oceans-clip.png"
data-setup='{"example_option":true}'>
  <source src="video/codeg.mp4" type="video/mp4" /><!-- WebKit -->
  <track kind="subtitles" src="video/codeg.ass" srclang="en-US" label="English"></track>
</video>

{{ end }}
