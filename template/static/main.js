var start = false
var timerId
var UID = ""


function setUID(UID) {
this.UID = UID;
}

function makeStart() {  
  this.start = true;
  $('#status').html("Процесс обработки... Ожидайте результат");
  
}

function makeFinish(data) {
  this.start = false;
  $('#status').html("Обработка завершена");
  $('#full').html(data.meta.nums.join( ", " ));
  $('#last').html( data.meta.last );
}

function executeQuery() { 
  if (this.start === true) {
  $.ajax({
    url: '/status',
    data:{UID: this.UID},
    success: function(data) {
      if (data.done) {
        makeFinish(data);
      } else {
        $('#last').html(data.meta.last);
      }

    }
  });
  this.timerId  = setTimeout(executeQuery, 300); 
 } else {
  clearTimeout(timerId);
 }
}

