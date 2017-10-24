## Google Vision (detect) API

```bash
$ go run detect.go main.go ../../doc/img/freddie.jpg
```
```bash
--- detectFaces
No faces found.
--- detectLabels
Labels:
performance
performing arts
entertainment
dancer
music
concert
event
performance art
stage
musician
percussion
music artist
public event
singing
pop music
profession
bassist
singer songwriter
--- detectLandmarks
No landmarks found.
--- detectText
No text found.
--- detectDocumentText
Text:
"17\n"
--- detectLogos
No logos found.
--- detectProperties
Dominant colors:
0.8% - #d49215
0.6% - #a76c16
2.0% - #e3bbbb
0.3% - #f8f223
0.3% - #2d1302
55.8% - #160f11
0.4% - #3c3012
0.5% - #f1ae28
0.7% - #e9bc33
0.4% - #c89c17
--- detectCropHints
Crop hints:
(22,0)
(583,0)
(583,899)
(22,899)
--- detectWeb
Web properties:
	Full image matches:
		https://lh3.googleusercontent.com/6NPqXWiv5dhiqN8_r7fL6xf8dixXZocbIDFwysEY7SavruoIvdK830i9jli7Iccimfcf1ZCLtZ25Uw=w1024-h1436-no
		https://pbs.twimg.com/media/CrlrvN8XgAAA3YC.jpg
		http://www.arcadia938.gr/images/articles/2015/11/15/22.jpg
		http://pm1.narvii.com/6276/f066c69676b933866f1c2febfca3aa68b0ea1fa7_hq.jpg
		https://pp.vk.me/c637725/v637725413/2e79/L_vZ23Wjbqo.jpg
		https://t-eska.cdn.smcloud.net/common/9/2/s/921828B3S3.jpg/ru-0-r-650,0-n-921828B3S3_freddie_mercury.jpg
		http://cdn2.thumbs.common.smcloud.net/common/2/2/s/2268441gEya.jpg/ru-1-r-640,0-n-2268441gEya.jpg
		http://ocdn.eu/pulscms-transforms/1/bEOktkpTURBXy8xOGYwNTFlM2NmNTAyYTI3MWI1ZmU1MmExMmU1Y2M3Mi5qcGeSlQLNA8AAwsOVAgDNA8DCww
		http://pm1.narvii.com/6457/5da388243982f998cb26a89480b339fa3dd60b88_hq.jpg
		https://brianmay.com/queen/queennews/newspix/15/Freddie-Mercury-in-1986-244006_600X915.jpg
	Pages with this image:
		http://www.express.co.uk/entertainment/music/552960/Queen-Adam-Lambert-singer-succeed-Freddie-Mercury
		https://www.pinterest.com/vierveijzer/freddie-mercury/
		https://brianmay.com/queen/queennews/queennewsjan15b.html
		https://www.pinterest.com/viedefun/freddie-mercury-queen/
		http://thegayguidenetwork.com/pop-culture/
		https://www.pinterest.co.uk/valerieperr0644/queen/
		http://www.thepicta.com/user/freddie_mercury_the_real_queen/4375218582
		https://www.pinterest.co.uk/lindaallison28/freddie-mercury-god-of-rock/
		http://www.washweb.net/queen-freddie-mercury
		http://www.throwbacks.com/theyve-announced-whos-playing-queen-in-the-new-movie-and-it-will-rock-you/
	Entities:
		/m/01jddz    Concert
		/m/0bk1p     Queen
		/m/01hdlq4   Queen
		/m/01hdkdr   Live at Wembley '86
		/m/04rlf     Music
		/m/011pyjjw  Queen + Adam Lambert
		/m/09l65     Singer
		/g/1s05j_49t Killer Queen
		/g/1q5j72r76 The Show Must Go On
		/m/0jg24     Image
--- detectSafeSearch
Safe Search properties:
Adult: VERY_UNLIKELY
Medical: VERY_UNLIKELY
Spoofed: UNLIKELY
Violence: VERY_UNLIKELY
```
