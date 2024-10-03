// Cargar las funciones que cargan los objetos a usar
window.onload = function () {
	cargarProductos();
	cargarMesoneros();
	cargarMesas();
};

function cargarProductos() {
	var xhr = new XMLHttpRequest();
	xhr.open('GET', 'http://localhost:3000/productos', true);
	xhr.send();

	xhr.onload = function () {
		if (xhr.status === 200) {
			var productos = JSON.parse(xhr.responseText);

			// Obtener el combo box
			var combo = document.getElementById('producto');

			// Agregar cada nombre de producto al combo box
			for (var i = 0; i < productos.length; i++) {
				var option = document.createElement('option');
				option.text = productos[i].nombre;
				combo.add(option);
			}
		}
		else {
			console.log('Error al cargar el archivo JSON');
		}
	};
}

function cargarMesoneros() {
	var xhr = new XMLHttpRequest();
	xhr.open('GET', 'http://localhost:3000/mesoneros', true);
	xhr.send();

	xhr.onload = function () {
		if (xhr.status === 200) {
			var mesoneros = JSON.parse(xhr.responseText);

			// Obtener el combo box
			var combo = document.getElementById('mesoneros');

			// Agregar cada nombre de mesonero al combo box
			for (var i = 0; i < mesoneros.length; i++) {
				var option = document.createElement('option');
				option.text = mesoneros[i].nombre;
				combo.add(option);
			}
		}
		else {
			console.log('Error al cargar el archivo JSON de Mesoneros');
		}
	};
}

function cargarMesas() {
	var xhr = new XMLHttpRequest();
	xhr.open('GET', 'http://localhost:3000/mesas', true);
	xhr.send();

	xhr.onload = function () {
		if (xhr.status === 200) {
			var mesas = JSON.parse(xhr.responseText);

			// Obtener el combo box
			var combo = document.getElementById('mesas');

			// Agregar cada nombre de mesonero al combo box
			for (var i = 0; i < mesas.length; i++) {
				var option = document.createElement('option');
				option.text = mesas[i].nombre;
				combo.add(option);
			}
		}
		else {
			console.log('Error al cargar el archivo JSON de Mesoneros');
		}
	};
}

function eliminarFila(boton) {
	// Obtener la fila actual
	var fila = boton.closest('tr');

	// Obtener el índice de la fila actual
	var indice = fila.rowIndex;

	// Eliminar la fila actual
	document.getElementById("tabla").deleteRow(indice);

	// Actualizar el total de la tabla
	sumarSubtotales();
}

function calcularSubtotal(input) {
	// Obtener el precio de la columna 3 del mismo renglón
	var precio = parseFloat(input.parentNode.parentNode.cells[2].innerHTML.replace("$", ""));

	// Obtener el valor del spinner
	var cantidad = parseFloat(input.value);

	// Calcular el subtotal
	var subtotal = precio * cantidad;

	// Colocar el subtotal en la columna 4 del mismo renglón
	input.parentNode.parentNode.cells[4].innerHTML = subtotal;
	sumarSubtotales();
}

function sumarSubtotales() {
	var tabla = document.getElementById("tabla");
	var total = 0;

	for (var i = 1; i < tabla.rows.length; i++) {
		// Obtener el valor de la columna 5 de la fila actual
		var subtotal = parseFloat(tabla.rows[i].cells[4].innerHTML);

		// Agregar el subtotal al total
		total += subtotal;
	}

	document.getElementById("resultado").value = total;
}

function agregarFila() {
	var xhr = new XMLHttpRequest();
	xhr.open('GET', 'http://localhost:3000/productos', true);
	xhr.send();

	xhr.onload = function () {
		if (xhr.status === 200) {
			var productos = JSON.parse(xhr.responseText);

			var combo = document.getElementById("producto");
			var nombreProducto = combo.options[combo.selectedIndex].text;
			var producto = productos.find(function (p) { return p.nombre === nombreProducto; });

			var fila = document.createElement("tr");
			var nombre = document.createElement("td");
			var cantidad = document.createElement("td");
			var precio = document.createElement("td");
			var nota = document.createElement("td");
			var subtotal = document.createElement("td");
			var accion = document.createElement("td");

			nombre.innerHTML = nombreProducto;
			precio.innerHTML = "$" + producto.precio.toFixed(2);
			nota.innerHTML = producto.nota;

			var prec = parseFloat(precio.innerHTML.replace("$", ""));
			cantidad.innerHTML = "<input type='number' value='1' min='1' max='10' onchange='calcularSubtotal( this , " + prec + ")'> </input>";
			subtotal.innerHTML = (prec).toFixed(2);
			document.getElementById("resultado").value = subtotal.innerHTML;

			accion.innerHTML = "<button id='eliminar' onclick='eliminarFila(this)'>X</button>";

			fila.appendChild(nombre);
			fila.appendChild(cantidad);
			fila.appendChild(precio);
			fila.appendChild(nota);
			fila.appendChild(subtotal);
			fila.appendChild(accion);

			document.getElementById("tabla").getElementsByTagName("tbody")[0].appendChild(fila);

			sumarSubtotales();
		}
		else {
			console.log('Error al cargar el archivo JSON');
		}
	};
}

// Función para obtener los datos de la tabla
function obtenerDatosDeLaTabla() {
	const filas = document.querySelectorAll('#tabla tbody tr');
	const datos = [];
  
	filas.forEach((fila) => {
	  const nombre = fila.querySelector('td:nth-child(1)').textContent;
	  const cantidad = parseInt(fila.querySelector('td:nth-child(2)').textContent, 10);
	  const precio = parseFloat(fila.querySelector('td:nth-child(3)').textContent);
	  const nota = fila.querySelector('td:nth-child(4)').textContent;
	  const subtotal = parseFloat(fila.querySelector('td:nth-child(5)').textContent);
  
	  // Agrega los datos a un objeto
	  const filaDatos = {
		nombre,
		cantidad,
		precio,
		nota,
		subtotal,
	  };
  
	  datos.push(filaDatos);
	});
  
	return datos;
  }
  
  function obtenerValoresSelectores() {
	const mesoneroSeleccionado = document.getElementById('mesoneros').value;
	const mesaSeleccionada = document.getElementById('mesas').value;
	console.log('Mesonero seleccionado:', mesoneroSeleccionado);
	console.log('Mesa seleccionada:', mesaSeleccionada);
  }  

//document.getElementById('btnprocesar').addEventListener('click', procesarPedido);

// Función para procesar el pedido
async function procesarPedido() {
  try {
	 const datosPedido = obtenerDatosDeLaTabla();

    // Ruta para realizar solicitud POST al servidor
    const response = await fetch('/procesar-pedido', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(datosPedido),
    });

    if (response.ok) {
      const data = await response.json();
      console.log('Pedido procesado:', data.message);
      // Puedes mostrar un mensaje de éxito al usuario si lo deseas
    } else {
      console.error('Error al procesar el pedido:', response.status);
      // Puedes mostrar un mensaje de error al usuario si lo deseas
    }
  } catch (error) {  
    console.error('Error al procesar el pedido:', error);
    // Puedes mostrar un mensaje de error al usuario si lo deseas

  }
}



