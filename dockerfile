FROM golang:alpine AS dev

WORKDIR /work

RUN apk add git
RUN apk add ffmpeg

RUN pip install whisper
RUN pip install stable-ts
RUN pip install --upgrade git+https://github.com/huggingface/transformers.git accelerate datasets[audio]
RUN pip install optimum

