# Quick API Test Commands

## Working Endpoints (After Fix)

### 1. List Accounts (GET)
```bash
curl http://localhost:8080/accounts/
```

### 2. Create Account (POST)
```bash
curl -X POST http://localhost:8080/accounts/ \
  -H "Content-Type: application/json" \
  -d '{
    "username": "anurag",
    "email": "anurag@carret.com",
    "password": "secure"
  }'
```

### 3. Get My Account (Authenticated GET)
```bash
curl http://localhost:8080/accounts/me/ \
  -H "Authorization: Token demo-token-123"
```

---

## Posts Endpoints

### 1. List Posts with Pagination
```bash
curl "http://localhost:8080/posts/?page=1&page_size=10"
```

### 2. Create Post (Authenticated)
```bash
curl -X POST http://localhost:8080/posts/ \
  -H "Authorization: Token demo-token-456" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "My First Post",
    "content": "Hello World",
    "author_id": 1,
    "published": true
  }'
```

### 3. Get Post Detail (Authenticated)
```bash
curl http://localhost:8080/posts/1/ \
  -H "Authorization: Token demo-token-456"
```

### 4. Update Post (Authenticated PUT)
```bash
curl -X PUT http://localhost:8080/posts/1/ \
  -H "Authorization: Token demo-token-456" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Post",
    "content": "Updated content"
  }'
```

### 5. Delete Post (Authenticated DELETE)
```bash
curl -X DELETE http://localhost:8080/posts/1/ \
  -H "Authorization: Token demo-token-456"
```

### 6. Test Throttling (10 req/min limit)
```bash
for i in {1..12}; do
  echo "Request $i:"
  curl http://localhost:8080/posts/throttled/
  echo ""
  sleep 0.5
done
```

---

## URL Pattern Summary

| Endpoint | Full URL | Methods | Auth Required |
|----------|----------|---------|---------------|
| List Accounts | `/accounts/` | GET | No |
| Create Account | `/accounts/` | POST | No |
| My Account | `/accounts/me/` | GET | Yes |
| List Posts | `/posts/` | GET | No |
| Create Post | `/posts/` | POST | Yes |
| Post Detail | `/posts/{id}/` | GET, PUT, PATCH, DELETE | Yes |
| Throttled Posts | `/posts/throttled/` | GET | No (limited) |
| Search Posts | `/posts/published/` | GET | No |

---

## Important Notes

1. **Server runs on port 8080** (not 8000)
2. **Routes are now fixed** - use `/accounts/` not `/accounts/accounts/`
3. **Authentication tokens:**
   - Accounts: `demo-token-123`
   - Posts: `demo-token-456`
4. **Restart server** after route changes: `go run manage.go runserver`
