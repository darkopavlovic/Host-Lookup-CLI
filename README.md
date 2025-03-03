# Host Lookup CLI [![Open Source Love](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/ellerbrock/open-source-badges/)

CLI tool to query DNS records. This tool uses the net package from Go's standard library to perform common DNS queries.

# Usage

- Prerequisite: Go 1.23.X
- Clone the repo and run with `go run cli.go command --host www.example.com`

###### Available commands

- `ns` ➡️ Look up host name servers
- `ip` ➡️ Look up host IP addresses
- `cname` ➡️ Look up host CNAME
- `mx` ➡️ Look up host MX records

# License

This project is licensed under the MIT License.
