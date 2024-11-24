telnet localhost 6379

```telnet

user@user:$ SET foo bar


```

# Comandos `LeftPush`, `RightPush` y `LeftPop` en Godis

En **Godis** (una implementación ligera de Redis escrita en Go), los comandos `LeftPush`, `RightPush` y `LeftPop` permiten manipular listas de forma similar a Redis.

## Explicación de los comandos

### `LeftPush` (Left Push)
- **Significado:** Inserta uno o más elementos al **inicio** de una lista.
- Si la lista no existe, se crea una nueva lista y se insertan los elementos.

**Ejemplo:**
```plaintext
LeftPush mylist "a" "b" "c"
```

Esto crea o modifica la lista mylist, colocando "c", "b" y "a" al inicio de la lista.
Resultado:
```plaintext
["c", "b", "a"]
```

### `RightPush` (Right Push)
- **Significado:**  Inserta uno o más elementos al **final** de una lista.
- Si la lista no existe, se crea una nueva lista y se insertan los elementos.

**Ejemplo:**
```plaintext
RightPush mylist "a" "b" "c"
```

Esto crea o modifica la lista mylist, colocando "a", "b" y "c" al final de la lista.
Resultado:

```plaintext
["a", "b", "c"]
```

### `LeftPop` (Left Pop)
- **Significado:**  Remueve y retorna el elemento del inicio de la lista.
- Si la lista está vacía o no existe, retorna un valor nulo o un mensaje de error.

**Ejemplo:**
Supongamos que mylist contiene:

```plaintext
["a", "b", "c"]
```

El comando:

```plaintext
LeftPop mylist
```

Remueve y retorna "a".
Resultado:

```plaintext
["b", "c"]
```