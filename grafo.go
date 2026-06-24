package main

import "fmt"

// Grafo representa um grafo usando lista de adjacência
type Grafo struct {
	vertices  int
	adjacencia map[int][]int
}

// NovoGrafo cria um novo grafo vazio
func NovoGrafo() *Grafo {
	return &Grafo{
		adjacencia: make(map[int][]int),
	}
}

// AdicionarVertice insere um vértice no grafo
func (g *Grafo) AdicionarVertice(v int) {
	if _, existe := g.adjacencia[v]; !existe {
		g.adjacencia[v] = []int{}
		g.vertices++
		fmt.Printf("Vértice %d adicionado.\n", v)
	} else {
		fmt.Printf("Vértice %d já existe.\n", v)
	}
}

// AdicionarAresta insere uma aresta entre dois vértices (não direcionado)
func (g *Grafo) AdicionarAresta(origem, destino int) {
	// Verifica se os vértices existem
	if _, existe := g.adjacencia[origem]; !existe {
		fmt.Printf("Vértice %d não existe. Adicione-o primeiro.\n", origem)
		return
	}
	if _, existe := g.adjacencia[destino]; !existe {
		fmt.Printf("Vértice %d não existe. Adicione-o primeiro.\n", destino)
		return
	}

	g.adjacencia[origem] = append(g.adjacencia[origem], destino)
	g.adjacencia[destino] = append(g.adjacencia[destino], origem) // grafo não direcionado
	fmt.Printf("Aresta %d <-> %d adicionada.\n", origem, destino)
}

// ExibirGrafo imprime a lista de adjacência do grafo
func (g *Grafo) ExibirGrafo() {
	fmt.Println("\n--- Lista de Adjacência ---")
	for vertice, vizinhos := range g.adjacencia {
		fmt.Printf("  %d -> %v\n", vertice, vizinhos)
	}
	fmt.Println("---------------------------")
}

// BFS realiza a busca em largura entre dois nós e retorna o caminho encontrado
func (g *Grafo) BFS(inicio, destino int) {
	fmt.Printf("\n=== BFS: de %d até %d ===\n", inicio, destino)

	// Verifica se os vértices existem
	if _, existe := g.adjacencia[inicio]; !existe {
		fmt.Printf("Vértice de início %d não existe.\n", inicio)
		return
	}
	if _, existe := g.adjacencia[destino]; !existe {
		fmt.Printf("Vértice de destino %d não existe.\n", destino)
		return
	}

	visitado := make(map[int]bool)   // controla nós já visitados
	anterior := make(map[int]int)    // guarda o caminho percorrido
	fila := []int{inicio}            // fila do BFS
	visitado[inicio] = true
	anterior[inicio] = -1           // início não tem antecessor

	encontrado := false
	ordem := []int{}                // ordem de visita dos nós

	for len(fila) > 0 {
		// Remove o primeiro elemento da fila
		atual := fila[0]
		fila = fila[1:]
		ordem = append(ordem, atual)

		fmt.Printf("Visitando nó: %d\n", atual)

		if atual == destino {
			encontrado = true
			break
		}

		// Adiciona os vizinhos não visitados na fila
		for _, vizinho := range g.adjacencia[atual] {
			if !visitado[vizinho] {
				visitado[vizinho] = true
				anterior[vizinho] = atual
				fila = append(fila, vizinho)
			}
		}
	}

	// Exibe a ordem de visita
	fmt.Printf("\nOrdem de visita: %v\n", ordem)

	if encontrado {
		// Reconstrói o caminho do destino até a origem
		caminho := []int{}
		for no := destino; no != -1; no = anterior[no] {
			caminho = append([]int{no}, caminho...)
		}
		fmt.Printf("Caminho encontrado: %v\n", caminho)
	} else {
		fmt.Printf("Não existe caminho entre %d e %d.\n", inicio, destino)
	}
}

func main() {
	// Cria o grafo
	g := NovoGrafo()

	fmt.Println("=============================")
	fmt.Println("   GRAFO - BUSCA EM LARGURA  ")
	fmt.Println("=============================")

	// Adiciona vértices
	fmt.Println("\n-- Inserindo vértices --")
	g.AdicionarVertice(1)
	g.AdicionarVertice(2)
	g.AdicionarVertice(3)
	g.AdicionarVertice(4)
	g.AdicionarVertice(5)
	g.AdicionarVertice(6)

	// Adiciona arestas
	fmt.Println("\n-- Inserindo arestas --")
	g.AdicionarAresta(1, 2)
	g.AdicionarAresta(1, 3)
	g.AdicionarAresta(2, 4)
	g.AdicionarAresta(2, 5)
	g.AdicionarAresta(3, 6)
	g.AdicionarAresta(4, 6)

	// Exibe o grafo
	g.ExibirGrafo()

	// Realiza BFS entre dois nós
	g.BFS(1, 6)
	g.BFS(5, 3)
}