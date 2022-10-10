# Kali CLI

Small CLI tool to install tools from the kali repository and not destroy your installation because you forgot to remove the repos again before updating or upgrading. 

<div align="center">
    <img src="screenshot.gif" alt="Screenshot" />
</div>

## Installation

### Binary Releases

You can download the latest compiled binary release [here](https://github.com/schufeli/kalicli/releases). 

After the download, you must add it to the `usr/local/bin/` directory so that the command can be found and executed.

In the download directory execute the following command:

```bash
sudo cp ./kalicli /usr/local/bin/
 ```

### Using `go install`

If you have a [Go](https://go.dev/) environment ready to go (at least go 1.19), it's as easy as:

```bash
go install github.com/schufeli/kalicli@latest
```
PS: You need at least go 1.19 to compile kalicli.

## Usage

```
Usage: kalicli COMMAND

Description: Easily install tools from the official kali repository

Commands:
	repo   - Manage kali repository (key, file)
	tools  - Manage kali tools and packages
```
### `repo` subcommand

```
Usage: kalicli repo COMMAND

Description: Manage kali repository (key, file)

Commands:
	add     - Add repository-key and file
	remove  - Remove repository-key and file
```

### `tools` subcommand

```
Usage: kalicli tools COMMAND

Description: Manage kali tools and packages

Commands:
	install    - Install tools
	uninstall  - Uninstall tools
	search     - Search tool in repository
```