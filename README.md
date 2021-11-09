# URL Shortener

The simplest URL shortening service possible.

## Get Started

```bash
make run
```

### Create a short link

```bash
curl http://localhost?url=http://averylongurl.com
```

**Response**

```json
{
  "id": "<id_sequence>",
  "url": "http://averylongurl.com"
}
```

To test your short url link open your browser and navigate to `http://localhost/<id_sequence>`

### Test
```bash
make test
```