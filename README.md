# mayalauncher
CLIのMayaランチャーです。
Mayaの言語やメジャーバージョンを指定して起動することができます。

# 使い方

## Goがインストールされている場合
下記コマンドでインストールしてください。

`$GOPATH/bin`(C:\Users\ユーザー名\go\bin)下に、
ビルドしたアプリがインストールされます。
```
go install github.com/Hum9183/mayalauncher@latest
```

下記コマンドでMayaを起動できます。
```sh
mayalauncher launch
```

## Goがインストールされていない場合
1. [releases](https://github.com/Hum9183/mayalauncher/releases)からビルド済のファイルをダウンロードする。
2. ダウンロードした.zipを解凍し、**mayalauncher.exe**を任意のディレクトリに配置する。
3. **mayalauncher.exe**を配置したディレクトリでターミナルを開く。

下記コマンドでMayaを起動できます。
```sh
./mayalauncher launch
```

# 言語フラグ
**-language**フラグをつけると、UI言語を指定して起動できます。
```sh
mayalauncher launch --language ja
```
```sh
mayalauncher launch -l en
```

# バージョンフラグ
**-version**フラグをつけると、メジャーバージョンを指定して起動できます。
```sh
mayalauncher launch --version 2022
```
```sh
mayalauncher launch -v 2024
```
