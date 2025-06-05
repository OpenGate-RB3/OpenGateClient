# Open Gate Frontend client

## About
***THIS PROJECT REQUIRES FFMPEG***
This repository provides a web-like interface for interacting with the RB3, specifically for fetching the video stream and sending/receiving audio. 


## Platforms
Currently, macOS and Linux are supported. Although Windows can fetch the stream, sending and receiving audio will be inop.

## Build instructions
The project leverages wails.io; more information about wails can be found on their website [wails](https://wails.io).
Below are the default build instructions for running and building the app.


### Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

### Building

To build a redistributable, production-mode package, use `wails build`.
