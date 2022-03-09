var times = 0
function lol(){
    check()
    times++
}
function check(){
    if (times > 0) {
        window.alert("Du hast diesen Knopf zu oft gedrückt")
        document.getElementById("button").disabled = true
        document.getElementById("disabled").hidden = false

    }else{
        document.getElementById("text").innerHTML += " sagte Ich."
    }
}
function boden(){
    if(document.getElementById("button").disabled){
        document.getElementById("button").disabled = false
        document.getElementById("disabled").hidden = true
        window.alert("You just enabled a button")
    }
}