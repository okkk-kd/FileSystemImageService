package diskStorage

const (
	queryUploadFile = `
	insert into store.images(client_id, created_at, updated_at, name)
	values($1, current_timestamp, current_timestamp, $2)
	returning id;
`
	queryDownloadFile = `
	select name, created_at, updated_at, id from store.images where client_id = $1 and name = $2
`
	queryGetList = `
	select name, created_at, updated_at, id from store.images where client_id = $1;
`
)
