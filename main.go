package main

import (
	"log"

	"github.com/fogleman/fauxgl"
	"github.com/hschendel/stl"
)

const (
	Width         = 1024
	Height        = 1024
	FOV           = 30   
	Near          = 1
	Far           = 10
	MetallicColor = 0.75 
)

func main() {
	filename := "input.stl" 
	reader, err := stl.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read STL file: %v", err)
	}

	mesh := fauxgl.NewEmptyMesh()

	for _, triangle := range reader.Triangles {
		v1 := fauxgl.Vector{float64(triangle.Vertices[0][0]), float64(triangle.Vertices[0][1]), float64(triangle.Vertices[0][2])}
		v2 := fauxgl.Vector{float64(triangle.Vertices[1][0]), float64(triangle.Vertices[1][1]), float64(triangle.Vertices[1][2])}
		v3 := fauxgl.Vector{float64(triangle.Vertices[2][0]), float64(triangle.Vertices[2][1]), float64(triangle.Vertices[2][2])}

		mesh.Triangles = append(mesh.Triangles, fauxgl.NewTriangleForPoints(v1, v2, v3))
	}

	mesh.BiUnitCube()

	context := fauxgl.NewContext(Width, Height)
	context.ClearColorBufferWith(fauxgl.HexColor("#ffffff")) // White background

	eye := fauxgl.Vector{3, 3, 3}     // Camera position
	center := fauxgl.Vector{0, 0, 0}  // Looking at origin
	up := fauxgl.Vector{0, 0, 1}      // Up direction
	matrix := fauxgl.LookAt(eye, center, up).Perspective(FOV, float64(Width)/float64(Height), Near, Far)

	light := fauxgl.Vector{1, 1, 1}.Normalize()

	shader := fauxgl.NewPhongShader(matrix, light, eye)
	shader.ObjectColor = fauxgl.Gray(MetallicColor) 
	shader.SpecularPower = 100                     
	shader.SpecularColor = fauxgl.Color{1, 1, 1, 1}

	context.Shader = shader
	context.DrawMesh(mesh)

	image := context.Image()
	err = fauxgl.SavePNG("output.png", image) 
	if err != nil {
		log.Fatalf("Failed to save PNG image: %v", err)
	}

	log.Println("Rendered image saved as output.png")
}

