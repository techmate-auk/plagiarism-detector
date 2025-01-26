# Project Title

Plagiarim Detector will help you to detect copy content of your manuscript or document.

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Features
- Easy to use and generate report with source of copy content
- AI enabled Plagiarim Detector
- Remove Plagiarism with enabled Paraphrasing Tool

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/ausafumarkhan/plagiarism-detector.git

2. Go to project directory
    ```bash
    cd Plagiarism-detector

3. Run GO code
    ```bash
    go run main.go

4. Open browser
    ```text
    http://localhost:8080

## Using Docker file

### We used Podman as a container runtime platform
 
 1. Build image for plagiarism detector
    ```bash
    podman build -t plagiarim:v1 .

2. Run the image using podman
    ```bash
    podman run -d -p 8080:8080 --name plagiarism-app plagiarism:v1