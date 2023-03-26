# traveler

## 注意点

- Dockerfileの一番最後に`RUN chown -R vscode:vscode /home/go`している箇所は、  
コンテナ起動後、go run main.goした際に、パッケージのインストールがpermission deniedになり、コードにジャンプできなかったため。(ビルドは成功するのに...)

- コンテナに入ったあとbashとすれば、bashが起動する。なおvimは入っていない。

## travelerAPI

### 起動方法
1. VSCodeにDev Containers拡張をインストール。
2. VSCodeで`/traveler`を開く
3. 左メニューのリモートエクスプローラで`travelerAPI`ディレクトリを指定。（２回め以降はそのまま選択）
4. コンテナ内に入るので、`/var/traveler/travelerAPI/cmd`内で`go run main.go`
