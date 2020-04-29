var socket = io('http://localhost:1323');
socket.on('reply', function(msg){
  console.log(msg)
    // $('#messages').append($('<li>').text(msg));
  });

const form = document.forms[0]
console.log(form)
form.onsubmit = function () {
  console.log('i will submit')
  console.log(form.m.value)
  return false
}
// $('form').submit(function() {
//   s2.emit('msg', $('#m').val(), function(data){
//     $('#messages').append($('<li>').text('ACK CALLBACK: ' + data));
//   });
//   socket.emit('notice', $('#m').val());
//   $('#m').val('');
//   return false;
// });
