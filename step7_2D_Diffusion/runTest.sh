#!/bin/bash

# main go filename
filename="step7_2D_Diffusion"

# Create output dir if it doesn't already exist
mkdir output

# Build step7
go build

# Run step7
./$filename

# Generate video from output
bash make_video.sh output/1d_sim400x400_%03dS.png $filename.mp4

# Open video
open $filename.mp4