# Installation

## Nuclei CLI

Installation of nuclei can be done using various options depending on the platform.

**Using Go**:

The following go command can be used to install the latest version of nuclei.

```
go install -v github.com/projectdiscovery/nuclei/v2/cmd/nuclei@latest
```

**Building from Code**:

To clone and build the project tip from Github, the follow command can be used.

```
git clone https://github.com/projectdiscovery/nuclei.git; \
cd nuclei/v2/cmd/nuclei; \
go build; \
mv nuclei /usr/local/bin/; \
nuclei -version;
```

**For MacOS (Brew)**:

```
brew install nuclei
```

**Using Docker**:

Tagged Docker Images are uploaded to Dockerhub with every Nuclei release, which can be downloaded using the below command.

```
docker pull projectdiscovery/nuclei:latest
```

**Using Binary**:

Pre-built binaries for various platforms generated using Github Actions can be downloaded from [Releases](https://github.com/projectdiscovery/nuclei/releases) page. Download the latest binary for your OS and unzip it to get a nuclei installation without go installation.


## Nuclei Templates

Nuclei has built-in support for automatic update/download templates since version v2.4.0. [Nuclei-Templates](https://github.com/projectdiscovery/nuclei-templatess) project provides a community-contributed list of ready-to-use templates that is constantly updated.

You may still use the update-templates flag to update the nuclei templates at any time; automatic updates happen every 24 hours.

The templates are stored in `${USER_HOME}/nuclei-templates` directory.