window.onload = function () {
  var filter = {}
  var homeContent = document.getElementById('homeContent')
  var homeContent_filter = document.getElementById('homeContent_filter')

  var xhr = new XMLHttpRequest()
  xhr.open(
    'GET',
    'http://localhost:9987/homeContent?spider_file_path=/root/go/src/TvBox/NanGua.py',
  )
  xhr.responseType = 'json'
  xhr.onload = function () {
    res = xhr.response
    if (res['code'] == 1) {
      // console.log('获取homeContent参数成功!!!')

      while (homeContent.firstChild) {
        homeContent.removeChild(homeContent.firstChild)
      }
      var home = document.createElement('span')
      home.innerText = '主页'
      home.setAttribute('class', 'home type_name')
      home.setAttribute('id', 'home')
      homeContent.appendChild(home)
      var classes = res['data']['class']
      for (var i = 0; i < classes.length; i++) {
        var span = document.createElement('span')
        span.innerText = classes[i]['type_name']
        span.setAttribute('class', 'type_name')
        span.setAttribute('id', classes[i]['type_id'])
        span.setAttribute('tid', classes[i]['type_id'])
        span.setAttribute('tname', classes[i]['type_name'])
        span.setAttribute(
          'onclick',
          'type_name_onclick("' + classes[i]['type_id'] + '")',
        )
        homeContent.appendChild(span)
      }
      filter = res['data']['filter']
    } else {
      console.log('获取homeContent参数失败!!!')
    }
  }
  xhr.send()
  //加载函数
  type_name_onclick = function (click_id) {
    var flr = filter[click_id]
    for (i = 0; i < flr.length; i++) {
      k = flr[i]['key']
      n = flr[i]['name']
      v = flr[i]['value']
      // console.log(k)
      var span = document.createElement('span')
      var div_null = document.createElement('div')
      span.setAttribute('id', k)
      span.setAttribute('key', k)
      span.setAttribute('class', 'filter_name')
      span.innerText = n
      div_null.appendChild(span)
      for (j = 0; j < v.length; j++) {
        // console.log(v)
        var span = document.createElement('span')
        span.innerText = v[j]['n']
        span.setAttribute('class', 'filter_value')

        div_null.appendChild(span)
      }

      homeContent_filter.appendChild(div_null)
    }
  }
}
