# Tinder Fist

## Tinder をとにかく速く全部いいねしたい

[python で tinder を自動化](https://github.com/Dragon-taro/tinder)したが、100 人いいねするのに 30 秒弱かかる。待てない。そうだ goroutine を使おう。

## 使い方

まずは、[このサイト](https://gist.github.com/taseppa/66fc7239c66ef285ecb28b400b556938)から facebook の access_token と id を用意（実装も含め実はここが一番だるいかも）

あとは、それを`functions/api.go`の中に埋め込んで`main.go`を実行！

## よくわからないところ

- 並行処理の実装の仕方全般（これでいいのかな？）
- channel の受信（かなり雑に実装してしまった）
- そもそも POST を goroutine で書いてもいいのか（こっちはいいけどサーバーに負担をかけてしまいそう）
- エラー処理
- ファイル構成

## 課題

- トークンとかをハードコーディングしてるところをなんとかしたい → yml にして git 管理から外すのが綺麗？
- update とかを実装したい
- てか sdk にしてイキりたい
