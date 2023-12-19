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
    for (var i = 1; i < homeContent.children.length; i++) {
      var child = homeContent.children[i]
      // console.log(child) // 输出子元素
      child.setAttribute('class', 'type_name')
    }
    element_click_id = document.getElementById(click_id)
    element_click_id.setAttribute('class', 'check_type_name')

    while (homeContent_filter.firstChild) {
      homeContent_filter.removeChild(homeContent_filter.firstChild)
    }
    var flr = filter[click_id]
    if (typeof flr === 'undefined') {
      return
    }

    for (i = 0; i < flr.length; i++) {
      k = flr[i]['key']
      n = flr[i]['name']
      v = flr[i]['value']
      // console.log(k)
      var span = document.createElement('span')
      var div_null = document.createElement('div')
      div_null.setAttribute('filter_key', k)
      div_null.setAttribute('filter_name', n)
      div_null.setAttribute('id', 'div_' + k)
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
        span.setAttribute('id', k + '$$$$$' + v[j]['v'])
        // span.setAttribute('id', 'a$' + j)
        // span.setAttribute('v', v[j]['v'])
        div_null.appendChild(span)
        span.setAttribute(
          'onclick',
          'filter_onclick("div_' + k + '","' + k + '$$$$$' + v[j]['v'] + '")',
        )
      }

      homeContent_filter.appendChild(div_null)
    }
    var hr = document.createElement('hr')
    homeContent_filter.appendChild(hr)
  }
  filter_onclick = function (parentId, fid) {
    // alert(parentId + ':' + value)
    element_pId = document.getElementById(parentId)
    element_click_filter_id = document.getElementById(fid)

    console.log(element_click_filter_id)

    for (var i = 1; i < element_pId.children.length; i++) {
      var child = element_pId.children[i]
      // console.log(child) // 输出子元素
      child.setAttribute('class', 'filter_value')
    }

    element_click_filter_id.setAttribute('class', 'check_filter_value')
  }
}
