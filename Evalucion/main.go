package main

import (
	"fmt"
	"time"
)

type Medicamento struct {
	nombre           string
	fabricante       string
	fechaFabricacion time.Time
	vidaUtil         int
}

type Tableta struct {
	Medicamento
	dosisPorTableta float64
	requiereReceta  bool
}

type Jarabe struct {
	Medicamento
	volumen int
	sabor   string
}

func NuevoMedicamento(nombre, fabricante, fecha string, vidaUtil int) (Medicamento, bool) {
	fechaFabricacion, errorFecha := time.Parse("02/01/2006", fecha)

	if nombre == "" || fabricante == "" || vidaUtil <= 0 || errorFecha != nil {
		return Medicamento{}, false
	}

	return Medicamento{nombre, fabricante, fechaFabricacion, vidaUtil}, true
}

func NuevaTableta(nombre, fabricante, fecha string, vidaUtil int, dosis float64, receta bool) (Tableta, bool) {
	medicamento, valido := NuevoMedicamento(nombre, fabricante, fecha, vidaUtil)

	if !valido || dosis <= 0 {
		return Tableta{}, false
	}

	return Tableta{medicamento, dosis, receta}, true
}

func NuevoJarabe(nombre, fabricante, fecha string, vidaUtil, volumen int, sabor string) (Jarabe, bool) {
	medicamento, valido := NuevoMedicamento(nombre, fabricante, fecha, vidaUtil)

	if !valido || volumen <= 0 || sabor == "" {
		return Jarabe{}, false
	}

	return Jarabe{medicamento, volumen, sabor}, true
}

func (m Medicamento) ObtenerNombre() string {
	return m.nombre
}

func (m Medicamento) ObtenerFabricante() string {
	return m.fabricante
}

func (m Medicamento) ObtenerFechaFabricacion() time.Time {
	return m.fechaFabricacion
}

func (m Medicamento) ObtenerVidaUtil() int {
	return m.vidaUtil
}

func (m *Medicamento) EstablecerNombre(nombre string) bool {
	if nombre == "" {
		return false
	}

	m.nombre = nombre
	return true
}

func (m *Medicamento) EstablecerFabricante(fabricante string) bool {
	if fabricante == "" {
		return false
	}

	m.fabricante = fabricante
	return true
}

func (m *Medicamento) EstablecerFechaFabricacion(fecha string) bool {
	fechaFabricacion, errorFecha := time.Parse("02/01/2006", fecha)

	if errorFecha != nil {
		return false
	}

	m.fechaFabricacion = fechaFabricacion
	return true
}

func (m *Medicamento) EstablecerVidaUtil(vidaUtil int) bool {
	if vidaUtil <= 0 {
		return false
	}

	m.vidaUtil = vidaUtil
	return true
}

func (m Medicamento) CalcularCaducidad() time.Time {
	return m.fechaFabricacion.AddDate(0, m.vidaUtil, 0)
}

func (t Tableta) MostrarDetalles() {
	fmt.Println("Tipo: Tableta")
	fmt.Println("Nombre:", t.ObtenerNombre())
	fmt.Println("Fabricante:", t.ObtenerFabricante())
	fmt.Println("Fecha de fabricacion:", t.ObtenerFechaFabricacion().Format("02/01/2006"))
	fmt.Println("Vida util:", t.ObtenerVidaUtil(), "meses")
	fmt.Println("Dosis:", t.dosisPorTableta, "mg")
	fmt.Println("Requiere receta:", t.requiereReceta)
}
