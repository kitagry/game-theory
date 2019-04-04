# Prisoner's Dilemma

囚人のジレンマの実装

## Usage

```
$ go build .
$ ./prisoners-dilemma -h
Usage of ./prisoners-dilemma:
  -p1 int
        Player1's Algorithm, 0: random, 1: Grim Trigger, 2: Tip for tat
  -p2 int
        Player2's Algorithm, 0: random, 1: Grim Trigger, 2: Tip for tat (default 1)
```

## Algorithm

| | 説明 |
|:--:|:--:|
| Random | ランダムで裏切る |
| Grim Trigger | 相手が裏切るまで裏切らない.一度裏切れば、ずっと裏切る |
| Tip for tat | 相手が前回裏切っていたら、裏切る.それ以外では裏切らない. |

