var end = 0;
var time = 0;
var interval = null;
var audio = new Audio("ring.wav")

function display_time() {
  var minutes = Math.floor(time / 60);
  var seconds = time % 60;

  var fMinutes = minutes.toString().padStart(2, '0');
  var fSeconds = seconds.toString().padStart(2, '0');

  $(".clock").text(fMinutes + ":" + fSeconds);
  $(document).attr('title', fMinutes + ":" + fSeconds);
}

$(".pom90").click(function() {
  time = 90 * 60;
  display_time();
});

$(".pom60").click(function() {
  time = 60 * 60;
  display_time();
});

$(".pom30").click(function() {
  time = 30 * 60;
  display_time();
});

$(".start").click(function() {
  console.log($(this).text());

  if ($(this).text() == "start") {
    console.log("start");
    $(this).text("pause")
    end = Math.floor(Date.now() / 1000) + time;
    interval = setInterval(function () {
      if (time != 0) {
        time = end - Math.floor(Date.now() / 1000);
      } else {
        audio.play();
        clearInterval(interval);
        $(".start").text("start");
      }
      display_time();
    }, 1000);
  } else if ($(this).text() == "pause") {
    clearInterval(interval);
    $(".start").text("start");
  }
});
