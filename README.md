# go-yaml-sample
golangでyamlを読み込むサンプルです
http://tweeeety.hateblo.jp/entry/2017/06/04/231043

# setup

```bash
$ git clone https://github.com/tweeeety/go-yaml-sample.git
$ cd go-yaml-sample/src/yaml_sample/
direnv: loading ../../.envrc
direnv: export ~GOPATH

$ glide install
$ glide update

$ go run yaml_sample.go 
buf: users:
    -
        name: ほげ ほげ男
        full_name:
            first_name: ほげ
            last_name: ほげ男
        sex: male
        birthday: 1990-12-12
        self_introduction:
            long: ほげほげ男です。よろしくおねがいします！
            short: しゃす！
        image_urls:
            - /my/photo/01.png
            - /my/photo/02.png
            - /my/photo/03.png
        shemale: false
    -
        name: ふが ふが子
        full_name:
            first_name: ふが
            last_name: ふが子
        sex: female
        birthday: 1994-03-03
        self_introduction:
            long: ふが子です！(≧∀≦*)
        image_urls:
            - /my/photo/01.png
        shemale: true


d: {Users:[{Name: FullName:{FirstName:ほげ LastName:ほげ男} Sex:male SelfIntroduction:{Long:ほげほげ男です。よろしくおねがいします！ Short:しゃす！} ImageURLs:[/my/photo/01.png /my/photo/02.png /my/photo/03.png] Shemale:false} {Name: FullName:{FirstName:ふが LastName:ふが子} Sex:female SelfIntroduction:{Long:ふが子です！(≧∀≦*) Short:} ImageURLs:[/my/photo/01.png] Shemale:true}]}


```

