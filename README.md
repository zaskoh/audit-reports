# go-starter template

<p>
    <a href="https://pkg.go.dev/github.com/zaskoh/go-starter">
        <img alt="Go reference" src="https://img.shields.io/badge/reference-grey?style=flat-square&logo=Go">
    </a>
    <a href="https://github.com/zaskoh/go-starter/actions/workflows/test.yml">
        <img alt="GitHub Workflow Status" src="https://github.com/zaskoh/go-starter/workflows/Test/badge.svg?style=flat-square">
    </a>
    <a href="https://goreportcard.com/report/github.com/zaskoh/go-starter">
        <img alt="Go Report Card" src="https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=flat-square">
    </a>
    <a href="https://github.com/zaskoh/go-starter/blob/main/go.mod">
        <img alt="go version" src="https://img.shields.io/github/go-mod/go-version/zaskoh/go-starter?style=flat-square&logo=Go">
    </a>
    <a href="https://github.com/zaskoh/go-starter/blob/main/LICENSE">
        <img alt="license" src="https://img.shields.io/github/license/zaskoh/go-starter?style=flat-square">
    </a>
    <a href="https://github.com/zaskoh/go-starter/releases">
        <img alt="GitHub Release" src="https://img.shields.io/github/v/release/zaskoh/go-starter?style=flat-square&include_prereleases&sort=semver">
    </a>
</p>

### Starting template for new go projects
The go template comes with:
- a [Makefile](Makefile)
    - build
    - lint
    - test 
    - benchmark 
- configured [main](main.go)
    - clean shutdown behaviour
    - logging setup with go.uber.org/zap
    - configuration via yml file and env variables

---

### todos after installing the template
- rename **__name_your_project__** in Makefile to your project name.
- remove .github/workflows if you don't need them / want them
- checkout [config](/config/config.go) and extend baseConfig or add / repalce the anotherConfig to add more possible configurations you can use in the config.yml
---