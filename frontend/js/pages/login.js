document.getElementById("login").addEventListener("submit", function(event) {
    event.preventDefault();

    const userName = document.getElementById("UserName").value;
    const password = document.getElementById("Password").value;

    fetch('/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            userName: userName,
            password: password
        }),
        credentials: 'include',
    })
    .then(response => {
        if (!response.ok) {
            throw new Error("Credenciales inválidas");
        }
        return response.json();
    })
    .then(data => {
        if (data.redirect) {
            console.log('redireccionando a:', data.redirect);
            window.location.href = data.redirect;
        } else {
            alert('Error: ' + data.message);
        }
    })
    .catch(error => {
        console.error('Error al iniciar sesión:', error);
        alert('Hubo un error al procesar la solicitud.');
    });
});

function accederRutaProtegida() {
    fetch('/home', {
        method: 'GET',
        credentials: 'include',
    })
    .then(response => {
        if (!response.ok) {
            throw new Error("No autorizado");
        }
        return response.text();
    })
    .then(data => {
        console.log("Respuesta de la ruta protegida:", data);
    })
    .catch(error => {
        console.error('Error al acceder a la ruta protegida:', error);
        alert('No tienes acceso a esta ruta.');
    });
};