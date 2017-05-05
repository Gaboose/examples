#!/usr/bin/bash
[ -z "$1" ] && AMP=10 || AMP=$1 # Amplitude
[ -z "$2" ] && FREQ=1 || FREQ=$2 # Frequency
echo $A
while true; do
    sleep 0.01
    TS=$(date +"%s%N")
    DX=$(echo "s($FREQ*$TS/1000000000)*$AMP" | bc -l)
    xdotool mousemove_relative -- $DX 0
done