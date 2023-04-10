const formulario = document.getElementById("formulario");

formulario.addEventListener("submit", function(event) {
	event.preventDefault();//deixa o formulario como default para envio

	const nome = document.getElementById("nome").value;//pega os dados enviados por ID
	const email = document.getElementById("email").value;
	const telefone = document.getElementById("telefone").value;

	const dados = { nome, email, telefone };
	const dadosJson = JSON.stringify(dados);

	localStorage.setItem("dados", dadosJson);

	window.location.href = "exibir.html";
});
