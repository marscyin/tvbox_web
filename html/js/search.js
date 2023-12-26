var spider_file_path = '/root/go/src/tvbox_web/python/NanGua.py'
var url = window.location.href
var key = ''

var search = function (spider_file_path, key) {
  var element_searchContent = document.getElementById('searchContent')
  // console.log(element_searchContent)

  url =
    'http://localhost:9987/searchContent?spider_file_path=' +
    spider_file_path +
    '&key=' +
    key
  var xhr = new XMLHttpRequest()
  xhr.open('GET', url)
  xhr.responseType = 'json'
  xhr.onload = function () {
    res = xhr.response
    if (res['code'] == 1) {
      // console.log('获取homeContent参数成功!!!')
      list = res['data']['list']
      // console.log(list

      while (element_searchContent.firstChild) {
        element_searchContent.removeChild(element_searchContent.firstChild)
      }
      for (i = 0; i < list.length; i++) {
        var a_table = document.createElement('a')
        a_table.setAttribute(
          'href',
          'http://localhost:9987/detailContent.html?ids=' + list[i]['vod_id'],
        )
        var pic_table = document.createElement('img')
        var name_table = document.createElement('span')
        var div_table = document.createElement('div')
        div_table.setAttribute('class', 'div_vod')
        pic_table.setAttribute('src', list[i]['vod_pic'])

        console.log(list[i]['vod_pic'])
        name_table.innerText = list[i]['vod_name']
        a_table.appendChild(pic_table)
        // a_table.appendChild(br_table)
        div_table.appendChild(a_table)
        div_table.appendChild(name_table)
        element_searchContent.appendChild(div_table)
      }
    }
  }
  xhr.send()
}

//====================================
if (url.split('=').length > 1) {
  key = url.split('=')[1]
}
if (key !== '') {
  search(spider_file_path, key)
}
s = function () {
  input_key = document.getElementById('input_key').value
  // console.log(input_key)
  search(spider_file_path, input_key)
}
