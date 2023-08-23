# Go Load Tester

A simple yet effective tool to perform load testing on APIs or websites, developed in Go. This tool will help you understand the performance metrics like the average response time, successful requests, failed requests, memory, and CPU usage.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Features

- Send concurrent requests to an API or website.
- Gather performance metrics like:
    - Number of requests
    - Average response time
    - Min/Max response time
    - Total success and failed requests
    - Memory and CPU usage

## Installation

1. Ensure you have Go installed on your machine.
2. Clone the repository:

   ```bash
   git clone https://github.com/indranandjha1993/go-load-tester.git
   cd go-load-tester
    ```

3. Build the tool:

   ```bash
   go build -o loadtester
    ```

## Usage
Modify the main.go file to set the target URL, HTTP method, number of requests, and concurrency level.

1. Modify the main.go file to set the target URL, HTTP method, number of requests, and concurrency level.
2. Run the tester:

    ```bash
    ./loadtester
   ```
3. Modify the **main.go** file to set the target URL, HTTP method, number of requests, and concurrency level.
4. Run the tester:

    ```bash
   ./loadtester
   ```
5. Review the results on the console.

