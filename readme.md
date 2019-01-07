# Azerite Goal

Query your guilds Azerite level and check if they achieved set goals for the guild :)

Download precompiled bin, [see releases page](https://github.com/BilboTheGreedy/Azerite/releases). 

## Usage

edit config.json with your region, realm, guild and ranks. Ranks can be bit tricky to figure out at first. But usually 0 equals guild master etc.

Azerite.exe -goal 40 -above -below

Will show characters above and below the goal. Use only -above / -below to show that result only.

#### Example output

```
  GOAL |  CHARACTER   | LEVEL |
+------+--------------+-------+----+
  Yes  | Scrype       |    40
  Yes  | Pottermonk   |    40
  Yes  | Nnogga       |    41
  Yes  | Wõrstplâÿêr  |    42
  Yes  | Narcolies    |    41
  Yes  | Joshpriest   |    40
  Yes  | Deepshades   |    41
  Yes  | Xerwo        |    41
  Yes  | Philwestside |    41
  Yes  | Gingì        |    41
  Yes  | Meeres       |    41
  Yes  | Fleksshades  |    40
  Yes  | Flêks        |    41
  Yes  | Flêkspriest  |    40
  Yes  | Swagfist     |    41
  Yes  | Doitpssy     |    41
+------+--------------+-------+----+
                        GOAL  | 40
                      +-------+----+
```

#### Azerite -h

```
Usage Azerite.exe:
  -above
        show characters above goal
  -below
        show characters below goal
  -debug
        debug
  -goal int
        Azerite neck level goal (default 38)

```
