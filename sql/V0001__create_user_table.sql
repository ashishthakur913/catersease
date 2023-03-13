CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  city TEXT NOT NULL,
  country TEXT NOT NULL,
  google_url TEXT NOT NULL,
  postal_code TEXT NOT NULL,
  state TEXT NOT NULL,
  street TEXT NOT NULL,
  unit TEXT,
  name TEXT NOT NULL,
  notes TEXT,
  phone_number TEXT NOT NULL
);
