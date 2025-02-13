package postgres

const (
	CreateQuery string = `INSERT INTO urls (short_url, long_url)
	VALUES ($1, $2)
	RETURNING id, short_url, long_url, created_at, updated_at`

	GetShortQuery string = `SELECT id, short_url, long_url, created_at, updated_at FROM urls WHERE long_url = $1`

	GetLongQuery string = `SELECT id, short_url, long_url, created_at, updated_at FROM urls WHERE short_url = $1`
)
