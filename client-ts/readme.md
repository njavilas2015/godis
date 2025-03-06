# Cliente de Typescript (Browser) para Godis

**Godis** es una implementación ligera de Redis escrita en Go

## Uso
Para usar `hash`, puedes ejecutar el siguiente comando en la terminal:

```bash
python3 test_hash.py
```

```typescript
const { HashStorage } = require('./dist')

const storage = new HashStorage("http://localhost:8080")

async function main() {

    await storage.hset("food", { country: "ar", price: "5" })

    const result = await storage.hget("food", "country")

    console.log(result)
}

main()
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
