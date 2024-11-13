# Pokemon Service 🐾

Pokemon Service es una API RESTful construida con **Golang** que permite almacenar información de Pokémon consumida desde la [PokeAPI](https://pokeapi.co/). La base de datos utilizada es **SQLite**, y la aplicación está completamente dockerizada para facilitar su despliegue.

---

## 📂 Estructura del Proyecto

```plaintext
pokemon-service/
├── database/               # Capa de acceso a la base de datos
│   ├── connection.go       # Conexión y migraciones
├── models/                 # Definición de structs y datos
│   ├── pokemon.go          # Modelo de datos de Pokémon
├── routes/                 # Controladores y rutas HTTP
│   ├── pokemon_routes.go   # Rutas relacionadas con Pokémon
│   ├── router.go           # Configuración del router principal
├── services/               # Lógica de negocio e integraciones externas
│   ├── pokeapi.go          # Consumidor de la PokeAPI
├── Dockerfile              # Archivo de configuración para Docker
├── .dockerignore           # Archivos a ignorar durante la construcción de Docker
├── go.mod                  # Módulo de dependencias de Go
├── go.sum                  # Suma de verificación de dependencias de Go
├── main.go                 # Punto de entrada de la aplicación
```

---

## 🚀 Cómo Empezar

### **1. Requisitos Previos**
- **Golang** 1.23.1 o superior
- **Docker** y **Docker Compose**

### **2. Configuración del Proyecto**
1. Clona este repositorio:
   ```bash
   git clone https://github.com/tu-usuario/pokemon-service.git
   cd pokemon-service
   ```

2. Asegúrate de que las dependencias estén instaladas:
   ```bash
   go mod tidy
   ```

3. (Opcional) Prueba la aplicación localmente:
   ```bash
   go run main.go
   ```

### **3. Dockerizar la Aplicación**
1. Construye la imagen de Docker:
   ```bash
   docker build -t pokemon-service .
   ```

2. Ejecuta el contenedor:
   ```bash
   docker run -p 8080:8080 --name pokemon-service -v "$PWD/pokemon.db:/app/pokemon.db" pokemon-service
   ```

3. La API estará disponible en [http://localhost:8080](http://localhost:8080).

---

## 📜 Endpoints de la API

### **1. Crear un Pokémon**
- **URL:** `POST /pokemon`
- **Body (JSON):**
  ```json
  {
    "pokemon_id": 1
  }
  ```
- **Respuesta Exitosa:**
  ```json
  {
    "message": "Pokémon insertado con éxito",
    "data": {
        "name": "bulbasaur",
        "types": [{"type": {"name": "grass"}}, {"type": {"name": "poison"}}],
        "abilities": [{"ability": {"name": "overgrow"}}, {"ability": {"name": "chlorophyll"}}],
        "weight": 69
    }
  }
  ```

### **2. Obtener un Pokémon por ID**
- **URL:** `GET /pokemon/:id`
- **Respuesta Exitosa:**
  ```json
  {
    "id": 1,
    "name": "bulbasaur",
    "type": "grass",
    "ability": "overgrow",
    "weight": 69,
    "pokemon_id": 1
  }
  ```

---

## 📂 Persistencia de Datos

La base de datos `pokemon.db` se almacena en el contenedor Docker y se puede sincronizar con tu máquina local mediante un volumen. Para asegurarte de que los datos persisten:
- Usa el siguiente comando para montar la base de datos local:
  ```bash
  docker run -p 8080:8080 --name pokemon-service -v "$PWD/pokemon.db:/app/pokemon.db" pokemon-service
  ```

---

## 🛠 Herramientas y Tecnologías

- **Golang 1.23.1**
- **SQLite**
- **Docker**
- **Gin Framework**
- **PokeAPI**

---

## 🧹 Limpieza y Mantenimiento

### Detener el Contenedor
```bash
docker stop pokemon-service
```

### Eliminar el Contenedor
```bash
docker rm pokemon-service
```

### Reconstruir la Imagen
```bash
docker build -t pokemon-service .
```

---

## 🐛 Contribuciones y Problemas

Si encuentras algún problema, por favor abre un [issue](https://github.com/tu-usuario/pokemon-service/issues) en el repositorio. ¡Las contribuciones son siempre bienvenidas! 😊
