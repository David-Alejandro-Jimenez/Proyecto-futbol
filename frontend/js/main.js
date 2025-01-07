const imagenes = [
    './../assets/images/players/messi.png',
    './../assets/images/players/ronaldo.png',
    './../assets/images/players/Lewandowski.png',
];

function cargarImagenAleatoria() {
    const indiceAleatorio = Math.floor(Math.random() * imagenes.length);
    const imagenSeleccionada = imagenes[indiceAleatorio];

    const imagenElemento = document.getElementById("randomImage");
    imagenElemento.src = imagenSeleccionada;
    imagenElemento.alt = "Imagen aleatoria";
} 

window.onload = cargarImagenAleatoria;