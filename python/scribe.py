import stable_whisper
AUDIO='./temp/audio.mp3'
SRT='./temp/audio.srt'
model = stable_whisper.load_hf_whisper('large')
result = model.transcribe(AUDIO)
# model.refine(AUDIO, result, precision=0.5)
result.to_srt_vtt(SRT, tag=('<font color="red">', '</font>'))

