import psycopg2
import random
import string

dsn = "host=db port=5432 user=user password=password dbname=mydb sslmode=disable"

def insert_users():
    conn = psycopg2.connect(dsn)
    cur = conn.cursor()
    for i in range(1000000):
        user = generate_user()
        cur.execute("INSERT INTO users (city, country, email, google_url, name, notes, phone_number, postal_code, state, street, unit) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", (user['city'], user['country'], user['email'], user['google_url'], user['name'], user['notes'], user['phone_number'], user['postal_code'], user['state'], user['street'], user['unit']))
    if i % 100 == 0:
        print(f"Inserted {i} users")
    conn.commit()
    cur.close()
    conn.close()

def generate_user():
    city = ''.join(random.choices(string.ascii_uppercase, k=10))
    country = ''.join(random.choices(string.ascii_uppercase, k=10))
    email = ''.join(random.choices(string.ascii_lowercase, k=10)) + "@example.com"
    google_url = "https://google.com/" + ''.join(random.choices(string.ascii_lowercase, k=10))
    name = ''.join(random.choices(string.ascii_uppercase, k=10))
    notes = None if random.random() < 0.5 else ''.join(random.choices(string.ascii_uppercase, k=10))
    phone_number = ''.join(random.choices(string.digits, k=10))
    postal_code = ''.join(random.choices(string.digits, k=6))
    state = ''.join(random.choices(string.ascii_uppercase, k=10))
    street = ''.join(random.choices(string.ascii_uppercase, k=10))
    unit = None if random.random() < 0.5 else ''.join(random.choices(string.ascii_uppercase, k=5))
    return {"city": city, "country": country, "email": email, "google_url": google_url, "name": name, "notes": notes, "phone_number": phone_number, "postal_code": postal_code, "state": state, "street": street, "unit": unit}

if __name__ == "__main__":
    insert_users()
