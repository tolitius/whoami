## whoami

context: Learning Tensorfow, OpenCV, Google Cloud Vision API

Initial goals:

- [x] Image recognition based on a pre trained TensorFlow model
- [x] Image recognition based on manually retrained TensorFlow model (Python)
- [ ] Image recognition based on manually retrained TensorFlow model (Go)
- [x] Image detection with Google Vision API
- [x] Face detection with OpenCV

A set of "vision" API examples from different sources

## Inception (TensorFlow)

### running it

```bash
$ cd tensor/inception
$ docker-compose up --build
```

### using it

```bash
$ ls -l doc/img
total 1064
-rw-r--r--@ 1 user group  459051 Oct 23 11:12 image.jpg
$
```

<img src="doc/img/image.jpg" width="400px">

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

## Google Cloud API

```bash
$ cd gcloud/vision
```

```bash
$ go run detect.go main.go ../../doc/img/freddie.jpg
```

## OpenCV

```bash
$ cd opencv
```

```bash
$ export PYTHONPATH=/usr/local/lib/python2.7/site-packages:$PYTHONPATH
```

```bash
$ python face_detect_cv3.py image-with-faces.jpg
faces found: 42
... exit?
```

## License

Copyright © 2017 tolitius

Distributed under the Eclipse Public License either version 1.0 or (at your option) any later version.
