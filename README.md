# IROIRO API

This is colorful IROIRO aplication.

## Usage

access with auth.

```
access_key: unkounko
```

### GET IroIro

```
$ curl -i localhost:8000/v1/iroiro\?access_key=unkounko
```

### GET Iro

```
$ curl -i localhost:8000/v1/iroiro/1\?access_key=unkounko
```

### POST Iro

```
$ curl -i -F "iro[content]=auth test" -F "iro[re_iro_id]=2" -F "iro[color_code]=#2196F3" -F "iro[color_name]=Blue500" localhost:8000/v1/iroiro\?access_key=unkounko
```

## Docs

- [Statement](https://github.com/funnythingz/IROIRO/wiki/Statement)
- [WBS](https://docs.google.com/a/nanapi.co.jp/spreadsheets/d/111eu2YoP1SF7jQuImFCmIVaM2p6fd0DPh40tLoDyaZc/edit#gid=0)
