# Chi router

## Content

- [Tools](#tools)
- [Installation](#installation)
- [Testing](#testing)
- [Author](#author)

## Tools

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![GitHub Actions](https://img.shields.io/badge/github%20actions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=white)
![Neovim](https://img.shields.io/badge/NeoVim-%2357A143.svg?&style=for-the-badge&logo=neovim&logoColor=white)

- [chi](https://go-chi.io)
- [chi/middleware](https://github.com/go-chi/chi/v5/middleware)

## Installation

In order to install the application please follow the instructions bellow:

- Make sure you have [Go](https://golang.org) and [Postgres](https://www.postgresql.org/) installed. 
In both sites you will find installation instructions for your platform
- Clone this repository

```bash
git clone https://github.com/alcb/chi-learn.git
cd chi-learn
```

- Install all dependencies

```bash
go mod tidy
```

- Build the application

```bash
make build
```

- Run the application

```bash
./bin/server
```

## Testing

The following will run all unit tests that the application has

```bash
make unit-test
```

## Author

Andres Court

[![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)](https://github.com/alcb1310)
[![LinkedIn](https://img.shields.io/badge/linkedin-%230077B5.svg?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/andres-court/)
[![X](https://img.shields.io/badge/X-%23000000.svg?style=for-the-badge&logo=X&logoColor=white)](https://twitter.com/alcb1310)
[![Dev.to blog](https://img.shields.io/badge/dev.to-0A0A0A?style=for-the-badge&logo=dev.to&logoColor=white)](https://dev.to/alcb1310)
