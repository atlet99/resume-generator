# Resume Generator [![Docker Image Version](https://img.shields.io/docker/v/zetfolder17/resume-generator?label=Docker%20Image&sort=semver)](https://hub.docker.com/r/zetfolder17/resume-generator) [![CI](https://github.com/atlet99/resume-generator/actions/workflows/ci-builds-and-releases.yml/badge.svg)](https://github.com/atlet99/resume-generator/actions/workflows/ci-builds-and-releases.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/atlet99/resume-generator)](https://goreportcard.com/report/github.com/atlet99/resume-generator) [![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/atlet99/resume-generator/badge)](https://securityscorecards.dev/viewer/?uri=github.com/atlet99/resume-generator)

This project is a simple resume generator built using Golang. It allows you to create resumes in DOCX and PDF formats based on predefined templates.

## Features
- Generate resumes in DOCX and PDF formats.
- Environment-based configuration using `.env` file.
- Easy to customize templates for different resume styles.

## Installation
### Prerequisites
- [Go](https://golang.org/doc/install)
- Docker (optional, for containerized usage)

### Steps
1. Clone the repository:
```shell
git clone https://github.com/atlet99/resume-generator.git;
cd resume-generator
```
2. Install dependencies:
```shell
go mod tidy
```
3. Build the application:
```shell
go build -o resume-generator
```
4. Run the application:
```shell
./resume-generator
```
5. Alternatively, run with Docker:
```shell
docker-compose up --build
```

### Usage

1. Configure the `.env` file to specify your environment variables.
2. Run the application with the following command:
```shell
./resume-generator
```

### Contribution

We welcome contributions. Please fork the repository and submit pull requests.

### License

```shell
© MIT License
```