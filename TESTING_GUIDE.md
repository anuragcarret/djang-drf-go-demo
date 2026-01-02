# DRF Features Demo - Testing Guide

This guide demonstrates how to test all the DRF features implemented in the `django-drf-go` framework using the demo application.

## Prerequisites

1. Start the demo application:
```bash
cd demo
go run manage.go runserver
```

2. Ensure PostgreSQL is running:
```bash
docker compose up -d
```

3. Run migrations:
```bash
go run manage.go migrate
```

---

## 1. Basic List View (Pagination)

### Test Default Pagination
```bash
curl -X GET http://localhost:8000/accounts/
```

**Expected Response:**
```json
{
  "count": 10,
  "next": "http://localhost:8000/accounts/?page=2",
  "previous": null,
  "results": [...]
}
```

### Test Custom Page Size
```bash
curl -X GET "http://localhost:8000/accounts/?page_size=5"
```

### Test Specific Page
```bash
curl -X GET "http://localhost:8000/accounts/?page=2"
```

---

## 2. Create Account (No Auth Required)

```bash
curl -X POST http://localhost:8000/accounts/ \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "securepass123",
    "first_name": "Test",
    "last_name": "User"
  }'
```

**Expected Response:** `201 Created`
```json
{
  "id": 1,
  "username": "testuser",
  "email": "test@example.com",
  ...
}
```

---

## 3. Authentication Required Endpoints

### Without Token (Should Fail)
```bash
curl -X GET http://localhost:8000/accounts/1/
```

**Expected Response:** `401 Unauthorized` or `403 Forbidden`

### With Valid Token (Should Succeed)
```bash
curl -X GET http://localhost:8000/accounts/1/ \
  -H "Authorization: Token demo-token-123"
```

**Expected Response:** `200 OK`
```json
{
  "id": 1,
  "username": "demo_user",
  "email": "demo@example.com"
}
```

### Update Account (Authenticated)
```bash
curl -X PUT http://localhost:8000/accounts/1/ \
  -H "Authorization: Token demo-token-123" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "updated_user",
    "email": "updated@example.com",
    "first_name": "Updated"
  }'
```

### Delete Account (Authenticated)
```bash
curl -X DELETE http://localhost:8000/accounts/1/ \
  -H "Authorization: Token demo-token-123"
```

**Expected Response:** `204 No Content`

---

## 4. Filtering

### Filter by Username
```bash
curl -X GET "http://localhost:8000/accounts/?username=testuser"
```

### Filter with Lookup (contains)
```bash
curl -X GET "http://localhost:8000/accounts/?username__contains=test"
```

### Filter with Lookup (gt - greater than)
```bash
curl -X GET "http://localhost:8000/accounts/?id__gt=5"
```

### Multiple Filters (AND logic)
```bash
curl -X GET "http://localhost:8000/accounts/?username__contains=user&email__contains=example"
```

---

## 5. Search & Ordering

### Search Across Multiple Fields
```bash
curl -X GET "http://localhost:8000/accounts/search/?search=john"
```

### Order by Username (Ascending)
```bash
curl -X GET "http://localhost:8000/accounts/search/?ordering=username"
```

### Order by Email (Descending)
```bash
curl -X GET "http://localhost:8000/accounts/search/?ordering=-email"
```

### Multi-field Ordering
```bash
curl -X GET "http://localhost:8000/accounts/search/?ordering=-created_at,username"
```

### Combined Search and Ordering
```bash
curl -X GET "http://localhost:8000/accounts/search/?search=test&ordering=-created_at"
```

---

## 6. Rate Limiting (Throttling)

### Test Throttle Endpoint (5 requests/minute for anonymous)

**Request 1-5:** Should succeed
```bash
for i in {1..5}; do
  echo "Request $i"
  curl -X GET http://localhost:8000/accounts/throttled/
  echo ""
done
```

**Request 6:** Should fail with 429
```bash
curl -i -X GET http://localhost:8000/accounts/throttled/
```

**Expected Response:**
```
HTTP/1.1 429 Too Many Requests
Retry-After: 60s
Content-Type: application/json

{"error": "Rate limit exceeded"}
```

### Test Authenticated Throttle (100/hour)
```bash
# First create many requests with token
for i in {1..6}; do
  echo "Authenticated Request $i"
  curl -X GET http://localhost:8000/accounts/throttled/ \
    -H "Authorization: Token demo-token-123"
  echo ""
done
```

---

## 7. Combined Features Test

### Pagination + Filtering + Ordering
```bash
curl -X GET "http://localhost:8000/accounts/?page=1&page_size=10&username__contains=user&ordering=-created_at"
```

### Search + Pagination
```bash
curl -X GET "http://localhost:8000/accounts/search/?search=test&page=1&page_size=5"
```

---

## 8. Permission Testing

### AllowAny (No Auth Needed)
```bash
# List accounts - allowed for anyone
curl -X GET http://localhost:8000/accounts/
```

### IsAuthenticated (Auth Required)
```bash
# Detail view - requires authentication
curl -X GET http://localhost:8000/accounts/1/ \
  -H "Authorization: Token demo-token-123"
```

---

## Expected Response Formats

### Success Responses

**200 OK - List with Pagination:**
```json
{
  "count": 25,
  "next": "http://localhost:8000/accounts/?page=2",
  "previous": null,
  "results": [
    {
      "id": 1,
      "username": "user1",
      "email": "user1@example.com"
    }
  ]
}
```

**201 Created:**
```json
{
  "id": 10,
  "username": "newuser",
  "email": "new@example.com",
  "created_at": "2026-01-02T19:30:00Z"
}
```

**204 No Content:**
```
(Empty body)
```

### Error Responses

**400 Bad Request:**
```json
{
  "error": "Invalid data provided"
}
```

**401 Unauthorized:**
```json
{
  "error": "Authentication credentials not provided"
}
```

**403 Forbidden:**
```json
{
  "error": "You do not have permission to perform this action"
}
```

**404 Not Found:**
```json
{
  "error": "Object not found"
}
```

**429 Too Many Requests:**
```json
{
  "error": "Rate limit exceeded"
}
```

---

## Testing Checklist

- [ ] Basic list pagination works
- [ ] Custom page sizes work
- [ ] Account creation (POST) works
- [ ] Authentication is enforced on detail endpoints
- [ ] Token authentication works correctly
- [ ] Filtering by field works
- [ ] Lookup expressions (contains, gt, etc.) work
- [ ] Search across multiple fields works
- [ ] Ordering (ascending and descending) works
- [ ] Throttling blocks after limit exceeded
- [ ] Retry-After header is present on 429
- [ ] Permissions allow/deny correctly

---

## Tips

1. **View Headers:** Add `-i` flag to curl to see response headers:
   ```bash
   curl -i -X GET http://localhost:8000/accounts/
   ```

2. **Pretty Print JSON:** Pipe to `jq` for formatted output:
   ```bash
   curl -X GET http://localhost:8000/accounts/ | jq .
   ```

3. **Test Parallel Requests:** Use `&` to test throttling:
   ```bash
   for i in {1..10}; do
     curl -X GET http://localhost:8000/accounts/throttled/ &
   done
   wait
   ```

4. **Debug Mode:** Check server logs for detailed error information.
