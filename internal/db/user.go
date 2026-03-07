package db

type UpsertUserRequest struct {
	fingerprint string
	clientIP string
}

func (db *DB) UpsertUser(r UpsertUserRequest) {
}

