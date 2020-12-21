package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	err error
	db  *sql.DB
	tab = "membresia"
)

func main() {
	conexionBD()
	agregarDatos()
	//eliminarDatos(4)
	registroExiste()
	mostrarTabla(tab)
	mostrarID(tab)

	cerrarBD()
}

func conexionBD() {
	db, err = sql.Open("mysql", string("bienhechor:Bienhechor_1234;@tcp(74.208.31.248:3306)/bienhechor"))
	revisarError(err)

	err = db.Ping()
	revisarError(err)
	fmt.Println("Conexión establecida")
}

func agregarDatos() {
	tipoMembresia := "Oro"
	queryTexto := "insert into membresia (tipo_membresia)"
	agregar, err := db.Exec(queryTexto+" values(?)", tipoMembresia)
	revisarError(err)
	status, err := agregar.LastInsertId()
	revisarError(err)

	//mostrarTabla(tab)
	fmt.Println(status)
}

func mostrarID(tabla string) {
	query, _ := db.Query("SELECT * FROM membresia where id_membresia=8")

	for query.Next() {
		var idMembresia, tipoMembresia string
		err = query.Scan(&idMembresia, &tipoMembresia)
		revisarError(err)
		fmt.Println("ID de Registro: " + idMembresia)
	}
}

func mostrarTabla(tabla string) {
	query, _ := db.Query("SELECT * FROM membresia")

	for query.Next() {
		var idMembresia, tipoMembresia string
		err = query.Scan(&idMembresia, &tipoMembresia)
		revisarError(err)
		fmt.Println("ID: "+idMembresia, tipoMembresia)
	}
}

func registroExiste() {
	//var datoExiste = false
	query, _ := db.Query("SELECT * FROM membresia where tipo_membresia LIKE '%Lizeth%Rodriguez%Flores%'")
	for query.Next() {
		var idMembresia, tipoMembresia string
		err = query.Scan(&idMembresia, &tipoMembresia)
		revisarError(err)
		fmt.Println("El ID es: " + idMembresia)
	}
}

/*
func eliminarDatos(lastID int64) {
	eliminar, err := db.Exec("delete from membresia where id_membresia=?", lastID)
	revisarError(err)
	status, err := eliminar.RowsAffected()
	revisarError(err)
	fmt.Println(status)

	//mostrarTabla(tab)
}*/

func revisarError(err error) {
	if err != nil {
		fmt.Println("Error")
		panic(err)
	}
}

func cerrarBD() {
	defer db.Close()
	fmt.Printf("Se cerró la conexión")
}
