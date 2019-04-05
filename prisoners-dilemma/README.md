# Prisoner's Dilemma

囚人のジレンマの実装

## Game

プレイヤーは二人.
プレイヤーは裏切ると裏切らないの2つの行動が出来る.
それぞれの行動を行ったときに与えられる得点は以下の通り.

| 自分の行動 \ 相手の行動 | 裏切る | 裏切らない |
|:--:|:--:|:--:|
| 裏切る     | 1 | 5 |
| 裏切らない | 0 | 4 |

表を見る限りは、裏切るという行為が最も合理的である.
しかし、相手も合理的ならば必ず裏切るのでポイントは1点しか手に入れられない.
2人が合計で手に入れられる点数は以下の通り.

| 自分の行動 \ 相手の行動 | 裏切る | 裏切らない |
|:--:|:--:|:--:|
| 裏切る     | 2 | 5 |
| 裏切らない | 5 | 8 |

つまり、2人が自分のために合理的に動くことで得られる点数は合計して最も低くなる.
これを「囚人のジレンマ」という.


## Usage

```
$ go build .
$ ./prisoners-dilemma -h
Usage of ./prisoners-dilemma:
  -n int
        Number of game attempts (default 10)
  -p1 int
        Player1's Algorithm, 0: betray, 1: random, 2: Grim Trigger, 3: Tip for tat
  -p2 int
        Player2's Algorithm, 0: betray, 1: random, 2: Grim Trigger, 3: Tip for tat (default 1)
```

## Algorithm

| | 説明 |
|:--:|:--:|
| Betray | 常に裏切る |
| Random | ランダムで裏切る |
| Grim Trigger | 相手が裏切るまで裏切らない.一度裏切れば、ずっと裏切る |
| Tip for tat | 相手が前回裏切っていたら、裏切る.それ以外では裏切らない. |

## Result

|| Betray | Random | Grim trigger | Tip for tat |
|:--:|:--:|:--:|:--:|:--:|
|    Betray   | 10 | 34 | 14 | 14 |
|    Random   | 8 | 18 | 16 | 25 |
| Grim trigger| 9 | 28 | 40 | 40 |
| Tip for tat | 9 | 28 | 40 | 40 |
