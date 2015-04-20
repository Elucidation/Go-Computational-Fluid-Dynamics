#!/bin/bash

# $1 = outputs/template
# $2 = out.mp4 filename
ffmpeg -framerate 20 -i $1 -c:v libx264 -r 20 -pix_fmt yuv420p $2
