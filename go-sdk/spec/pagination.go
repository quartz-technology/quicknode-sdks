package spec

type PageInfo struct {
	HasNextPage bool
	EndCursor   string
}

type PaginationArgs struct {
	First int
	After string
}
