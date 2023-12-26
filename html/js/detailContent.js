var spider_file_path = '/root/go/src/TvBox/NanGua.py'
var play_from = document.getElementById('play_from')
var play_url = document.getElementById('play_url')
var url = window.location.href
var host = url.split('?')[0]
var pu = ''
var ids = ''
var uid = ''
if (url.split('?').length > 1) {
  // ids = url.split('=')[1]
  p = url.split('?')[1]
  p_list = p.split('&')
  for (i = 0; i < p_list.length; i++) {
    console.log(p_list[i])
    if (p_list[i].split('=')[0] === 'ids') {
      ids = p_list[i].split('=')[1]
    }
    if (p_list[i].split('=')[0] === 'uid') {
      uid = p_list[i].split('=')[1]
    }
  }
}
var click_pf = function (index) {
  while (play_url.firstChild) {
    play_url.removeChild(play_url.firstChild)
  }
  pu_list = pu.split('$$$')
  purl_list = pu_list[index].split('#')
  for (i = 0; i < purl_list.length; i++) {
    var a_table = document.createElement('a')
    a_table.setAttribute(
      'href',
      host + '?ids=' + ids + '&uid=' + btoa(purl_list[i].split('$')[1]),
    )
    // console.log(purl_list[i].split('$')[1])
    span = document.createElement('span')
    span.innerText = purl_list[i].split('$')[0]
    span.setAttribute(
      'onclick',
      'click_play("' + purl_list[i].split('$')[1] + '")',
    )
    span.setAttribute('class', 'purl_span')
    a_table.appendChild(span)

    play_url.appendChild(a_table)
  }
}
var detail = function (ids) {
  if (ids === '') {
    return
  }

  var url =
    'http://localhost:9987/detailContent?spider_file_path=' +
    spider_file_path +
    '&ids=' +
    ids
  var xhr = new XMLHttpRequest()
  xhr.open('GET', url)
  xhr.responseType = 'json'
  xhr.onload = function () {
    res = xhr.response
    if (res['code'] == 1) {
      vod = res['data']['list'][0]
      pf = vod['vod_play_from']
      pu = vod['vod_play_url']
      pf_list = pf.split('$$$')
      for (i = 0; i < pf_list.length; i++) {
        span = document.createElement('span')
        span.innerText = pf_list[i]
        span.setAttribute('class', 'pf_span')
        span.setAttribute('onclick', 'click_pf(' + i.toString() + ')')
        play_from.appendChild(span)
      }
      if (pf_list.length > 0) {
        click_pf(0)
      }
      // console.log(list)
    }
  }
  xhr.send()
}
detail(ids)
