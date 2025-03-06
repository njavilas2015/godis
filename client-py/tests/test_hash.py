from godis.hash import HashStorage

if __name__ == "__main__":
    storage = HashStorage()

    response = storage.hset("food", country="ar", type="fruit", color="red")
    print(f"Respuesta del servidor (HSET): {response}")

    result = storage.hget("food", "country")
    print(f"Respuesta del servidor (HGET): {result}")