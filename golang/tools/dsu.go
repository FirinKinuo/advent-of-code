package tools

type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}

	return &DSU{parent, size}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}

	return d.parent[x]
}

func (d *DSU) Union(x, y int) bool {
	rootX := d.Find(x)
	rootY := d.Find(y)

	if rootX == rootY {
		return false
	}

	if d.size[rootX] < d.size[rootY] {
		rootX, rootY = rootY, rootX
	}

	d.parent[rootY] = rootX
	d.size[rootX] += d.size[rootY]

	return true
}

func (d *DSU) GetComponentSizes() []int {
	n := len(d.parent)

	sizes := make([]int, 0, n)
	visited := UniqueAny{}

	for i := 0; i < n; i++ {
		root := d.Find(i)
		if !visited.Exists(root) {
			visited.Add(root)
			sizes = append(sizes, d.size[root])
		}
	}

	return sizes
}

func (d *DSU) CountComponents() int {
	n := len(d.parent)
	visited := UniqueAny{}

	for i := 0; i < n; i++ {
		root := d.Find(i)
		visited.Add(root)
	}

	return len(visited)
}
