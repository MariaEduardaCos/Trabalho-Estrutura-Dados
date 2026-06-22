package main

import ( "fmt")

type Grafo struct {
	ListaAdjacencia map[string][]string
}

func NovoGrafo() *Grafo {
	return &Grafo{
		ListaAdjacencia: make(map[string][]string),
	}
}

func (g *Grafo) AdicionarAresta(origem, destino string) {
	if g.ListaAdjacencia == nil {
		g.ListaAdjacencia = make(map[string][]string)
	}
	g.ListaAdjacencia[origem] = append(g.ListaAdjacencia[origem], destino)
	g.ListaAdjacencia[destino] = append(g.ListaAdjacencia[destino], origem)
}

