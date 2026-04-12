CREATE TABLE messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    client_fingerprint TEXT NOT NULL,
    sender_type TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT (datetime('now'))
);

