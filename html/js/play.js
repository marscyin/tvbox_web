var player = videojs('my-video')
var play_src = document.getElementById('play_src')
var url = window.location.href
var spider_file_path = '/root/go/src/tvbox_web/python/NanGua.py'

var click_play = function (url) {
  u =
    'http://localhost:9987/playerContent?spider_file_path=' +
    spider_file_path +
    '&id=' +
    btoa(url)
  var xhr = new XMLHttpRequest()
  xhr.open('GET', u)
  xhr.responseType = 'json'
  xhr.onload = function () {
    res = xhr.response
    if (res['code'] == 1) {
    }
  }
  xhr.send()
}
