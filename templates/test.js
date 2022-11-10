document.getElementById("res").innerHTML = "New text!";
function enviarpost(buttonid){
	let Data ={Test:document.getElementById(buttonid).value}
        fetch("/test",{method: 'post', body: JSON.stringify(Data)})
};
