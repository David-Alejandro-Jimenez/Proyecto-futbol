const imagenes = [
    //Cuadrar esto más adelante ya que no es bueno esta ruta absoluta
    'http://127.0.0.1:5500/frontend/Imagenes/messi.png',
    'http://127.0.0.1:5500/frontend/Imagenes/ronaldo.png',
    'http://127.0.0.1:5500/frontend/Imagenes/Lewandowski.png',
];

function cargarImagenAleatoria() {
    const indiceAleatorio = Math.floor(Math.random() * imagenes.length);
    const imagenSeleccionada = imagenes[indiceAleatorio];

    // Establecemos la imagen en el elemento HTML con el id 'randomImage'
    const imagenElemento = document.getElementById("randomImage");
    imagenElemento.src = imagenSeleccionada;
    imagenElemento.alt = "Imagen aleatoria";
} 
  // Llamamos a la función para cargar la imagen aleatoria cuando cargue la página
window.onload = cargarImagenAleatoria;

