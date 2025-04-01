CREATE TABLE IF NOT EXISTS transactions (
    id INTEGER PRIMARY KEY AUTOINCREMENT, -- Unique transaction ID
    ticker TEXT NOT NULL, -- Stock symbol (e.g., "AAPL", "TSLA")
    price_per_unit REAL NOT NULL, -- Price per unit
    currency TEXT NOT NULL, -- Currency as string
    amount REAL NOT NULL, -- Amount of stocks bought/sold
    date DATETIME NOT NULL, -- Date and time of transaction
    is_buy BOOLEAN NOT NULL -- True for buy, False for sell
);

CREATE TABLE IF NOT EXISTS tracked_stocks (
    id INTEGER PRIMARY KEY AUTOINCREMENT, -- Unique ID for tracked stock
    ticker TEXT NOT NULL, -- Stock symbol (e.g., "AAPL", "TSLA")
    date_added DATETIME NOT NULL -- Date when stock was added to tracking list
);
