use actix_web::{web, App, HttpServer, Responder, HttpResponse};
use redis::{AsyncCommands, Client};
use serde::{Deserialize, Serialize};
use std::error::Error;

#[derive(Debug, Serialize, Deserialize)]
struct Asignacion {
    curso: String,
    facultad: String,
    carrera: String,
    region: String,
}

//Handler para recibir datos y almacenarlos en Redis
async fn insertar(asignacion: web::Json<Asignacion>,) -> impl Responder {
    //Conectar a Redis
    let redis_url = "redis://127.0.0.1/";
    let client = match Client::open(redis_url) {
        Ok(c) => c,
        Err(e) => return HttpResponse::InternalServerError().body(format!("Error al conectar a Redis: {}", e)),
    };

    //Conectar a la base de datos
    let mut con = match client.get_multiplexed_async_connection().await {
        Ok(c) => c,
        Err(e) => return HttpResponse::InternalServerError().body(format!("Error al conectar a la base de datos: {}", e)),
    };

    //serializar la estructura a JSON
    let asignacion_json = match serde_json::to_string(&asignacion.into_inner()) {
        Ok(a) => a,
        Err(e) => return HttpResponse::InternalServerError().body(format!("Error al serializar la estructura: {}", e)),
    };

    //Almacenar la asignación en Redis
    if let Err(e) = con.set::<&str, String, ()>("asignacion:1", asignacion_json).await {
        return HttpResponse::InternalServerError().body(format!("Error al almacenar la asignación: {}", e));
    }

    HttpResponse::Ok().body("Asignación almacenada correctamente")
}

//Handler para mostrar un mensaje en la raíz del servidor
async fn index() -> impl Responder {
    HttpResponse::Ok().body("Servidor de almacenamiento en Redis")
}


#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    //Print para indicar que el servidor está corriendo
    println!("Servidor corriendo en el puerto 8080");
    //Iniciar el servidor HTTP
    HttpServer::new(|| {
        App::new()
            .route("/", web::get().to(index))
            .route("/insertar", web::post().to(insertar))
    })
    .bind("127.0.0.1:8080")? //Cambiar la dirección IP si se desea y el puerto
    .run()
    .await?;

    Ok(())
}