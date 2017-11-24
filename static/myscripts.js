//adapted from https://bootsnipp.com/snippets/featured/simple-chat
var me = {};
me.avatar = "http://wfarm2.dataknet.com/static/resources/icons/set108/b5cdab07.png";

var you = {};
you.avatar = "http://downloadicons.net/sites/default/files/women-business-user-icon-44928.png";

//the function to get the time
function formatAMPM(date) {
  var hours = date.getHours();
  var minutes = date.getMinutes();
  var ampm = hours >= 12 ? 'PM' : 'AM';
  hours = hours % 12;
  hours = hours ? hours : 12; 
  minutes = minutes < 10 ? '0' + minutes : minutes;
  var strTime = hours + ':' + minutes + ' ' + ampm;
  return strTime;
}

//this function displays the image and formats the text 
function insertChat(who, text, time = 0) {
  var control = "";
  var date = formatAMPM(new Date());
  //the text user sends gets formatted 
  if (who == "me") {
    control = '<li style="width:100%">' +
      '<div class="msj macro">' +
      '<div class="avatar"><img class="img-circle" style="width:100%;" src="' + me.avatar + '" /></div>' +
      '<div class="text text-l">' +
      '<p style = "margin-top: -200">' + text + '</p>' +
      '<p><small>'+ date + '</small></p>' +
      '</div>' +
      '</div>' +
      '</li>';
  } else {//
    control = '<li style="width:100%;">' +
      '<div class="msj-rta macro">' +
      '<div class="text text-r">' +
      '<p>' + text + '</p>' +
      '<p><small>' + date + '</small></p>' +
      '</div>' +
      '<div class="avatar" style="padding:0px 0px 0px 10px !important"><img class="img-circle" style="width:100%;" src="' + you.avatar + '" /></div>' +
      '</li>';
  }
  setTimeout(
    function () {
      $("ul").append(control);
      $("ul").animate({scrollTop: $("ul")[0].scrollHeight},500);

    }, time);

}

function resetChat() {
  $("ul").empty();
}

$("#mytext").on("keyup", function (e) {
  if (e.which == 13) {
    var text = $(this).val();
    if (text !== "") {
      insertChat("me", text, 100);
      $.get('/user-input', { value: $('#mytext').val() })
        .done(function (data) {
          insertChat("you", data, 500)
        })
      $(this).val('');
    }
  }
});

//-- Clear Chat
resetChat();

$("#user-input-form").submit(
function (event) {
  event.preventDefault();

});