## whoami

context: Learning Go and Tensorfow

Initial goals:

- [x] Image recognition based on a pre trained model
- [ ] Image recognition based on manually retrained model

By default `whoami` uses a pre-trained TensorFlow [Inception-V3](https://arxiv.org/abs/1512.00567) model
and is initially based on [this](https://outcrawl.com/image-recognition-api-go-tensorflow/) great example.

## running it

```bash
docker-compose up --build
```

## using it

```bash
$ ls -l doc/img
total 1064
-rw-r--r--@ 1 user group  459051 Oct 23 11:12 image.jpg
$
```

```bash
$ curl -s http://localhost:4242/whoami -F 'image=@doc/img/image.jpg'
```
```json
{
  "filename": "image.jpg",
  "labels": [
    {
      "label": "tiger",
      "probability": 0.66044205
    },
    {
      "label": "tiger cat",
      "probability": 0.33603966
    },
    {
      "label": "lynx",
      "probability": 0.0015672832
    },
    {
      "label": "tabby",
      "probability": 0.0009888832
    },
    {
      "label": "lion",
      "probability": 0.00036086622
    }
  ]
}
```

## License

Copyright © 2017 tolitius

Distributed under the Eclipse Public License either version 1.0 or (at your option) any later version.
