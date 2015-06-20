# IROIRO API

This is colorful IROIRO aplication API.

## Usage

access with auth.

```
access_key: unkounko
```

### Iro

#### GET IroIro

```
$ curl -i iroiro.space/v1/iroiro
```

#### GET Iro

```
$ curl -i iroiro.space/v1/iroiro/:id
```

#### POST Iro

```
$ curl -i -F "iro[content]=auth test" -F "iro[re_iro_id]=2" -F "iro[color_id]=1" iroiro.space/v1/iroiro
```

### Color

#### GET ColorList

```
$ curl -i iroiro.space/v1/colors
```

#### GET Color

```
$ curl -i iroiro.space/v1/colors/:id
```

#### POST Color

```
$ curl -i -F "color[name]=Blue500" -F "color[code]=#2196F3" -F "color[text_code]=#FFFFFF" iroiro.space/v1/colors
```
