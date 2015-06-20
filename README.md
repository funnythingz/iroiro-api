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
$ curl -i iroiro.space/v1/iroiro\?access_key=unkounko
```

#### GET Iro

```
$ curl -i iroiro.space/v1/iroiro/:id\?access_key=unkounko
```

#### POST Iro

```
$ curl -i -F "iro[content]=auth test" -F "iro[re_iro_id]=2" -F "iro[color_id]=1" iroiro.space/v1/iroiro\?access_key=unkounko
```

### Color

#### GET ColorList

```
$ curl -i iroiro.space/v1/colors\?access_key=unkounko
```

#### GET Color

```
$ curl -i iroiro.space/v1/colors/:id\?access_key=unkounko
```

#### POST Color

```
$ curl -i -F "color[name]=Blue500" -F "color[code]=#2196F3" -F "color[text_code]=#FFFFFF" iroiro.space/v1/colors\?access_key=unkounko
```

&copy; funnythingz
