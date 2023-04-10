const urlParams = new URLSearchParams(window.location.search);
  const nome = urlParams.get('Nome');
  const email = urlParams.get('Email');
  const telefone = urlParams.get('Telefone');

  const nomeSpan = document.querySelector('#nome');
  const emailSpan = document.querySelector('#email');
  const telefoneSpan = document.querySelector('#telefone');

  nomeSpan.textContent = nome;
  emailSpan.textContent = email;
  telefoneSpan.textContent = telefone;