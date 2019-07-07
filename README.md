# komodo-ecosysboard

## Prerequisites

Below is the list of prerequisites to compile `komodo-ecosysboard` on your machine:

- Golang 1.12.6 minimum

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.
See deployment for notes on how to deploy the project on a live system.

### Build

To build the project please follow the instructions below:

```bash
go build -o komodo-ecosysboard ecosysboard/ecosysboard.go
```

## Installing

:construction:

## Running the tests

To run the tests type the following instructions below:

```
chmod 000 ecosysboard/config/samples/not_good_rights_config.json
go test ./...
chmod 644 ecosysboard/config/samples/not_good_rights_config.json
```

### Coverage

To run the tests under coverage type the following instructions below:

```
chmod 000 ecosysboard/config/samples/not_good_rights_config.json
bash scripts/coverage.sh 
chmod 644 ecosysboard/config/samples/not_good_rights_config.json
```

## Deployment

:construction:

## Authors

-  **Roman Sztergbaum** - Creator - [Milerius](https://github.com/Milerius)
- Retrieve the list of authors here: [AUTHORS](LEGAL/AUTHORS)

## Legal

All information about the legal part of the project is available in the `LEGAL` folder at the root of the repository.

- COPYING: GPL V2.0: see the [COPYING](LEGAL/COPYING) file for details about copying the software.
- DEVELOPER-AGREEMENT: see the [DEVELOPER-AGREEMENT](LEGAL/DEVELOPER-AGREEMENT) file for details about authors.
- LICENSE: GPL V2.0: see the [LICENSE](LEGAL/LICENSE) file for details about the license.

***

| Badges     |                                                                                                                                                              |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Unit tests | [![badge](https://report.ci/status/Milerius/komodo-ecosysboard/badge.svg?branch=master)](https://report.ci/status/Milerius/komodo-ecosysboard?branch=master) |
| Coverage   | [![codecov](https://codecov.io/gh/Milerius/komodo-ecosysboard/branch/master/graph/badge.svg)](https://codecov.io/gh/Milerius/komodo-ecosysboard)                                                                                                                                                              |
| Build Status | [![Build Status](https://travis-ci.com/Milerius/komodo-ecosysboard.svg?branch=master)](https://travis-ci.com/Milerius/komodo-ecosysboard)                                                                                                                                                              |