# Shortform Wonder
![shortformwonder-logo](./assets/banner.png)

Shortform Wonder is a minimalist CLI short-form video editor. It aims to automate viral reel creation.
1. Merges given video with chosen <i> attention-seeking </i> video
2. Formats those stacked videos to the given format
3. Adds automatically generated subtitles using [Whisper](https://github.com/openai/whisper)

Along [Whisper](https://github.com/openai/whisper), <b> STW </b> uses [stable-ts](https://github.com/jianfch/stable-ts) to improve the quality of automatically generated subtitles

## Main premise
<b> STW </b> was created to allow users to develop reel-creation-pipelines upon. <b> STW </b> is a perfect solution to the problem allowing for bulk video processing with automatically generated subtitles.

## Docker
For user's comfort, it is recommended for <b> STW </b> to be launched using Docker environment using provided <i>.Dockerfile</i>
```bash
docker build --target dev . -t go
docker run -it -v ${PWD}:/work go sh
```
it launches a terminal of a Go container in the repo's directory.

## Local
If one desires to use <b> STW </b> on their local machine then the following libraries must be installed:
#### Go (Golang)
<b> STW </b> was made using [Golang](https://go.dev/). If run locally - Go environment must be installed and configured.
#### Video
[FFmpeg](https://github.com/FFmpeg/FFmpeg) is the backbone of this tool.
``` bash
# linux
apt-get install ffmpeg  

# windows
choco install ffmpeg
```
#### Subtitles
```bash
pip install whisper
pip install stable-ts
pip install --upgrade git+https://github.com/huggingface/transformers.git accelerate datasets[audio]
pip install optimum
```

## Usage
### Bulk run
```bash
./sfw bulk 
```
Bulk editor will edit all of the videos from ``` input ``` at once in series. Perfect for automization.

### Single run
```bash
./sfw run ... 
```
Single run allows users to select any video they want to edit.

### Options
There are several available options to achieve the desired results.

#### Duration
- ``` start ``` - allows to select the start timestamp of the video (optional) 
-  ``` end ``` - allows to select the end timestamp of the video (optional)  

if none of the above are specified, <b> STW </b> will use the whole duration of the given video.  
default: 00:00 ; 00:00

```bash
./sfw run -source="input\example.mp4" -start="00:00" -end="00:10"
```

#### Dimensions

- ``` height ``` - allows to specify the height of the video (optional) 
-  ``` width ``` - allows to specify the width of the video (optional) 

default: 1080 x 1920 (reels format)
```bash
./sfw run -source="input\example.mp4" -height=900 -width=900
```

#### Attention video
 - ``` additional ``` - allows specifying the source of attention video
-  ``` random ``` - allows to specify if random attention video should be chosen (optional) (default: True)

```bash
./sfw run -source="input\example.mp4" -additional="additional\a.mp4" -random=false
```

All of the above options can be used with ``` bulk ``` - but as for now, there is no way to specify different options for each processed video during such a run.


# License
This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details