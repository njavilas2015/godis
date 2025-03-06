# Cliente de python para Godis

**Godis** es una implementación ligera de Redis escrita en Go

## Uso
Para usar `hash`, puedes ejecutar el siguiente comando en la terminal:

```python
from godis.hash import HashStorage

if __name__ == "__main__":
    storage = HashStorage()

    response = storage.hset("food", country="ar", type="fruit", color="red")
    print(f"Respuesta del servidor (HSET): {response}")

    result = storage.hget("food", "country")
    print(f"Respuesta del servidor (HGET): {result}")
```

## Contacto
Si tienes alguna pregunta o necesitas soporte, no dudes en contactarme:

Nombre: Javier Avila
Email: [njavilas2015@gmail.com]
GitHub: njavilas2015

## Apóyame con un café ☕️

Si te gusta mi trabajo y quieres apoyarme, ¡puedes invitarme a un café! 😊

[![Buy Me a Coffee](https://img.buymeacoffee.com/button-api/?text=Buy%20Me%20a%20Coffee&emoji=coffee&slug=tu_nombre&button_colour=FF5F5F&font_colour=ffffff&font_family=Cookie)](https://buymeacoffee.com/njavilas
)
