# tempura
## はじめに
毎回、同じコード書いてるとしんどくなってきたので、テンプレートと設定ファイルから生成できるようにしたい。
`text/template`を使ってためしてみる。
2021/07/10現在、まだできてないけども。

## インストール
たぶん、こんな雰囲気でインストールする

```
go install xxx
```

## 使い方

### 設定ファイル（json）を作る
`sample.json`を参考に設定ファイルを作る

### 実行する
こんな感じでできる予定
```
./tempura inputs/hoge.json
```

### ファイルを移動する
適切な場所にファイルを移動する

### 修正する
`TODO`が書かれている部分を修正する。

## そのうち、やりたいこと
- [ ] 既存ファイルに関数・メソッドを追加
- [ ] コマンドとクエリの両方に対応（jsonファイルの中身みて変更？）
