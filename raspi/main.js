function test(){
    window.alert("Worked");
}
function getSubmitParams(){
    var input_value = document.getElementById('task-input-field').value;
    writeTaskToJSON(input_value);
}
function writeTaskToJSON(value){
    var text = value;
    const newTask = `{"tasks":["${text}","Abspühlen"]}`;
    const task = JSON.parse(newTask);
    console.log(task.tasks[0]);
}