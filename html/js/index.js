var spider_file_path = '/root/go/src/TvBox/NanGua.py'
window.onload = function () {
  var filter = {}
  var homeContent = document.getElementById('homeContent')
  var homeContent_filter = document.getElementById('homeContent_filter')
  var categoryContent = document.getElementById('categoryContent')

  var xhr = new XMLHttpRequest()
  xhr.open(
    'GET',
    'http://localhost:9987/homeContent?spider_file_path=' + spider_file_path,
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
    //筛选页面加载完成
    while (categoryContent.firstChild) {
      categoryContent.removeChild(categoryContent.firstChild)
    }
    var xhr_category = new XMLHttpRequest()
    var categoryContent_url =
      'http://localhost:9987/categoryContent?spider_file_path=' +
      spider_file_path +
      '&tid=' +
      click_id +
      '&pg=1'
    xhr_category.open('GET', categoryContent_url)
    xhr_category.responseType = 'json'
    xhr_category.onload = function () {
      res = xhr_category.response
      if (res['code'] == 1) {
        list = res['data']['list']
        // console.log(list)
        for (i = 0; i < list.length; i++) {
          var a_table = document.createElement('a')
          var pic_table = document.createElement('img')
          var name_table = document.createElement('span')
          var div_table = document.createElement('div')
          div_table.setAttribute('class', 'div_vod')
          div_table.setAttribute('url', categoryContent_url.slice(0, -1))
          div_table.setAttribute('pg', '1')

          pic_table.setAttribute('src', list[i]['vod_pic'])

          name_table.innerText = list[i]['vod_name']
          a_table.appendChild(pic_table)
          // a_table.appendChild(br_table)
          div_table.appendChild(a_table)
          div_table.appendChild(name_table)
          categoryContent.appendChild(div_table)
          // console.log(list[i]['vod_name'])
        }
      } else {
        console.log('获取categoryContent参数失败!!!')
        return
      }
    }
    xhr_category.send()
  }
  filter_onclick = function (parentId, fid) {
    // alert(parentId + ':' + value)
    element_pId = document.getElementById(parentId)
    element_click_filter_id = document.getElementById(fid)
    element_check_type_name_id = document
      .getElementsByClassName('check_type_name')[0]
      .getAttribute('tid')

    // console.log(element_click_filter_id)

    for (var i = 1; i < element_pId.children.length; i++) {
      var child = element_pId.children[i]
      // console.log(child) // 输出子元素
      child.setAttribute('class', 'filter_value')
    }

    element_click_filter_id.setAttribute('class', 'check_filter_value')
    elements_filters = document.getElementsByClassName('check_filter_value')
    flr_obj = {}
    for (i = 0; i < elements_filters.length; i++) {
      eid_list = elements_filters[i].id.split('$$$$$')
      flr_obj[eid_list[0]] = eid_list[1]
    }
    // console.log(flr_obj)

    while (categoryContent.firstChild) {
      categoryContent.removeChild(categoryContent.firstChild)
    }

    var xhr_category_filter = new XMLHttpRequest()
    var category_filter_url =
      'http://localhost:9987/categoryContent?spider_file_path=' +
      spider_file_path +
      '&tid=' +
      element_check_type_name_id +
      '&extend=' +
      JSON.stringify(flr_obj) +
      '&pg=1'
    xhr_category_filter.open('GET', category_filter_url)
    xhr_category_filter.responseType = 'json'
    xhr_category_filter.onload = function () {
      res = xhr_category_filter.response
      if (res['code'] == 1) {
        list = res['data']['list']
        for (i = 0; i < list.length; i++) {
          var a_table = document.createElement('a')
          var pic_table = document.createElement('img')
          var name_table = document.createElement('span')
          var div_table = document.createElement('div')
          div_table.setAttribute('class', 'div_vod')
          div_table.setAttribute('url', category_filter_url.slice(0, -1))
          div_table.setAttribute('pg', '1')

          pic_table.setAttribute('src', list[i]['vod_pic'])

          name_table.innerText = list[i]['vod_name']
          a_table.appendChild(pic_table)
          // a_table.appendChild(br_table)
          div_table.appendChild(a_table)
          div_table.appendChild(name_table)
          categoryContent.appendChild(div_table)
        }
      }
    }
    xhr_category_filter.send()
  }
  // document.body.addEventListener('scroll', function () {
  // console.log('aa')
  // })
  window.onscroll = function () {
    // console.log('aaa')
    var scrolltop = document.documentElement.scrollTop
    var dh = document.documentElement.clientHeight
    // console.log(scrolltop + dh)
    var body_offsetH = document.documentElement.offsetHeight
    // console.log(body_offsetH)
    // console.log(scrolltop + dh)
    if (scrolltop + dh >= body_offsetH - 10) {
      // console.log('到底部咯')
      // alert('到底了')
      let last_vod = categoryContent.lastElementChild
      let u = last_vod.getAttribute('url')
      let p = last_vod.getAttribute('pg')
      let new_pg = parseInt(p) + 1
      let url_cate = u + new_pg.toString()
      // console.log(url_cate)

      var xhr_categorys = new XMLHttpRequest()
      xhr_categorys.open('GET', url_cate)
      console.log('正在请求-->' + url_cate)
      xhr_categorys.responseType = 'json'
      xhr_categorys.onload = function () {
        let res = xhr_categorys.response
        if (res['code'] == 1) {
          let last_vod_new = categoryContent.lastElementChild
          let p_new = last_vod_new.getAttribute('pg')
          // console.log(p_new)
          // console.log(last_vod_new)
          if (p_new === p) {
            list = res['data']['list']

            for (i = 0; i < list.length; i++) {
              var a_table = document.createElement('a')
              var pic_table = document.createElement('img')
              var name_table = document.createElement('span')
              var div_table = document.createElement('div')
              div_table.setAttribute('class', 'div_vod')
              div_table.setAttribute('url', u)
              div_table.setAttribute('pg', new_pg)

              pic_table.setAttribute('src', list[i]['vod_pic'])

              name_table.innerText = list[i]['vod_name']
              a_table.appendChild(pic_table)
              // a_table.appendChild(br_table)
              div_table.appendChild(a_table)
              div_table.appendChild(name_table)
              categoryContent.appendChild(div_table)
              // console.log(list)
            }
          }
        }
      }
      xhr_categorys.send()
      // console.log(old_pg)
      // var categoryContent_urls =
    }
  }
}
