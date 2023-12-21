const dp = new DPlayer({
  container: document.getElementById('dplayer'),
  video: {
    url: 'http://asters.tk:5244/d/%E9%98%BF%E9%87%8C%E4%BA%91%E7%9B%98/Alist3/%E6%B1%AA%E6%B1%AA%E9%98%9F%E7%AB%8B%E5%A4%A7%E5%8A%9F%E7%AC%AC8%E5%AD%A3/01.mp4',
    type: 'hls',
  },
})
dp.switchVideo({
  url: 'https://ydd.vsmql.com/m3u84/share/816355/963579/20231214/150306/1080/index.m3u8?sign=b2bd612a83a6142eba48bf68031ace95&t=1703169300',
})
