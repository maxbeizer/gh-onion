# gh-onion 🧅

A [GitHub CLI](https://cli.github.com/) extension that brings The Onion to your terminal. Bundled dataset, zero API dependencies, pure fun.

## Install

```bash
gh extension install maxbeizer/gh-onion
```

## Commands

```bash
gh onion                  # Random headline (Onion or real news)
gh onion motd             # Headline of the day — add to your .zshrc
gh onion search <term>    # Search headlines by keyword
gh onion fresh            # Fetch a fresh headline from The Onion's RSS feed
```

### Random Headline

```
$ gh onion

  ╔══════════════════════════════════════════════════════════╗
  ║ Study Finds Every Style Of Parenting Produces            ║
  ║ Disturbed, Miserable Adults                              ║
  ╚══════════════════════════════════════════════════════════╝

  — The Onion
```

### Message of the Day

Add to your `.zshrc` or `.bashrc`:

```bash
gh onion motd
```

Output:
```
📰 Today's Headline from America's Finest News Source:
   "Study Finds Every Style Of Parenting Produces Disturbed, Miserable Adults"
```

### Search

```
$ gh onion search florida

Found 4 headline(s) matching "florida":

  📰 Florida Man Arrested For Trying To Open Airplane Emergency Exit Mid-Flight
  📰 Florida Man Uses Alligator To Shotgun A Beer
  📰 69-Year-Old Florida Man Hits Nephew With Pizza Over Prior Pizza Incident
  📰 Florida Man Steals Elderly Woman's Purse, Returns It After Finding Bible Inside
```

### Fresh (RSS)

```bash
gh onion fresh   # Requires internet
```

### JSON Output

All commands support `--json` and `--jq` flags:

```bash
gh onion --json
gh onion search dog --jq .text
gh onion motd --json
```

## Development

```bash
make help          # see all targets
make build         # build binary
make test          # run tests
make ci            # build + vet + test-race
make install-local # install extension from checkout
make relink-local  # reinstall after changes
```

## Dataset

This extension bundles the [Onion or Not](https://www.kaggle.com/datasets/chrisfilo/onion-or-not) dataset (~24k headlines) by [Chris Filo Finan](https://www.kaggle.com/chrisfilo), licensed under [CC BY-SA 4.0](https://creativecommons.org/licenses/by-sa/4.0/). The dataset contains headlines from The Onion and r/NotTheOnion.

## Releasing

Tag a version to trigger a release build:

```bash
git tag v0.1.0
git push origin v0.1.0
```
