package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "strings"
)

type Producto struct {
    ID                    int     `json:"id"`
    Cantidad              int     `json:"cantidad"`
    Precio                float64 `json:"precio"`
    Descripcion           string  `json:"descripcion"`
    Marca                 string  `json:"marca"`
    FechaUltimoInventario string  `json:"fecha_ultimo_inventario"`
}

func main() {
    http.HandleFunc("/productos", productosHandler)
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func productosHandler(w http.ResponseWriter, r *http.Request) {
    // Ejemplo de arreglo de productos
    productos := []Producto{
        {
            ID:                    1,
            Cantidad:              10,
            Precio:                110.0,
            Descripcion:           "mayonesa",
            Marca:                 "KRAFT",
            FechaUltimoInventario: "2024-06-16",
        },
        {
            ID:                    2,
            Cantidad:              12,
            Precio:                112.0,
            Descripcion:           "salsa",
            Marca:                 "HEINZ",
            FechaUltimoInventario: "2024-06-18",
        },
        {
            ID:                    3,
            Cantidad:              24,
            Precio:                124.0,
            Descripcion:           "arroz",
            Marca:                 "MARY",
            FechaUltimoInventario: "2024-06-24",
        },

        {
            ID:                    4,
            Cantidad:              10,
            Precio:                110.0,
            Descripcion:           "Super 8",
            Marca:                 "NESTLE",
            FechaUltimoInventario: "2024-06-16",
        },
        {
            ID:                    5,
            Cantidad:              12,
            Precio:                112.0,
            Descripcion:           "Ken",
            Marca:                 "CERVECERIAS UNIDAS",
            FechaUltimoInventario: "2024-06-18",
        },
        {
            ID:                    6,
            Cantidad:              24,
            Precio:                124.0,
            Descripcion:           "Milo",
            Marca:                 "NESTLE",
            FechaUltimoInventario: "2024-06-24",
        },
    }

    // Obtener los query parameters
    queryParams := r.URL.Query()
    idsParam := queryParams.Get("ids")

    // Parsear los IDs recibidos
    ids := parseIDs(idsParam)

    // Filtrar los productos por los IDs recibidos
    productosFiltrados := filterProductosPorIDs(productos, ids)

    // Convertir a JSON y escribir la respuesta
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(productosFiltrados)
}

func parseIDs(idsParam string) []int {
    ids := []int{}

    // Remover los corchetes y separar por comas
    idsStr := strings.Trim(idsParam, "{}")
    idStrings := strings.Split(idsStr, ",")

    // Convertir los strings a ints
    for _, idStr := range idStrings {
        id, err := strconv.Atoi(strings.TrimSpace(idStr))
        if err == nil {
            ids = append(ids, id)
        }
    }

    return ids
}

func filterProductosPorIDs(productos []Producto, ids []int) []Producto {
    productosFiltrados := []Producto{}

    // Filtrar productos por IDs
    for _, producto := range productos {
        for _, id := range ids {
            if producto.ID == id {
                productosFiltrados = append(productosFiltrados, producto)
            }
        }
    }

    return productosFiltrados
}
