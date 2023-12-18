window.onload = function () {
  var homeContent = document.getElementById('homeContent')

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
      home.setAttribute('class', 'home title')
      homeContent.appendChild(home)
      var classes = res['data']['class']
      for (var i = 0; i < classes.length; i++) {
        console.log(classes[i])
      }
    } else {
      console.log('获取homeContent参数失败!!!')
    }
  }
  xhr.send()
}
