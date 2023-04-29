package graph

import "os"

const GRAPH_SERVICE_ENDPOINT = "GRAPH_SERVICE_ENDPOINT"

// Constants for routes
var BASE_URL = os.Getenv(GRAPH_SERVICE_ENDPOINT)

const GET_GRAPH = "/graphSearch/graph"
const GET_GRAPH_WITH_FILTER = "/graphSearch/filteredGraph"
