var player = videojs('my-video')
var play_src = document.getElementById('play_src')
var url = window.location.href
// alert('videojs')
player.on('ended', function () {
  next_play = document.getElementById('next_play')
  u = next_play.getAttribute('href')
  window.location.href = u
})
