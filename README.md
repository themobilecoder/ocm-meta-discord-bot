# ocm-meta-discord-bot
Unofficial Discord bot to check the metas of your OnChainMonkey.

## Permissions
Requires the following permissions to work on Discord
```
Scope: Bot
Bot Permissions: Use Slash Commands
```

## Traits
All OCM meta traits are obtained from [Metagood's Repo](https://github.com/metagood/OnChainMonkeyData), and are subject to change.

```
color match, mouth match, zeroes, nips, poker hands, twins
```
## Setup
```bash
$ export DISCORD_API_KEY="YOUR_API_KEY"
```
## Run in bash

```bash
$ task br
```

## Run in docker
```bash
$ task docker-build
$ task docker-run
```
## Usage
```bash
/meta id: ID_OF_OCM
```