#!/bin/bash

/bin/ollama serve &
pid=$!

sleep 5

ollama pull tinyllama

wait $pid