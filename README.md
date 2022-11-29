# c4c-reports

<p>
    <a href="https://pkg.go.dev/github.com/zaskoh/c4c-reports">
        <img alt="Go reference" src="https://img.shields.io/badge/reference-grey?style=flat-square&logo=Go">
    </a>
    <a href="https://github.com/zaskoh/c4c-reports/actions/workflows/test.yml">
        <img alt="GitHub Workflow Status" src="https://github.com/zaskoh/c4c-reports/workflows/Test/badge.svg?style=flat-square">
    </a>
    <a href="https://goreportcard.com/report/github.com/zaskoh/c4c-reports">
        <img alt="Go Report Card" src="https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=flat-square">
    </a>
    <a href="https://github.com/zaskoh/c4c-reports/blob/main/go.mod">
        <img alt="go version" src="https://img.shields.io/github/go-mod/go-version/zaskoh/c4c-reports?style=flat-square&logo=Go">
    </a>
    <a href="https://github.com/zaskoh/c4c-reports/blob/main/LICENSE">
        <img alt="license" src="https://img.shields.io/github/license/zaskoh/c4c-reports?style=flat-square">
    </a>
    <a href="https://github.com/zaskoh/c4c-reports/releases">
        <img alt="GitHub Release" src="https://img.shields.io/github/v/release/zaskoh/c4c-reports?style=flat-square&include_prereleases&sort=semver">
    </a>
</p>

**c4c-reports** will check https://code4rena.com/reports/ for new reports and inform you via discord / per log.

![discord message](/img/discord-info.png)

## Getting started

### Installing and building
```bash
git clone git@github.com:zaskoh/c4c-reports.git

make build
```

### Run
With the flag --config you can define the path to the config file.  

If you want to get informed via discord you need the token and channel (add them to your config or set as environment variables DISCORD_TOKEN + DISCORD_CHANNEL).

If you just want to get the log in your terminal, use --no-discord flag.

Examples
```bash
go run main.go --config=config.yml --no-discord

DISCORD_TOKEN=xyz DISCORD_CHANNEL=zzz ./c4c-reports

DISCORD_TOKEN=xyz DISCORD_CHANNEL=zzz ./build/c4c-reports --config=/path/to/config.yml
```

### Note
You need a discord token and a channel. Checkout https://github.com/zaskoh/discordbooter how to get the token.