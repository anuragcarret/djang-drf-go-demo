# DRF Features Testing Guide - Posts API

Complete guide for testing all DRF features in the Posts API.

## Prerequisites

```bash
cd demo
go run manage.go runserver
# Server runs on http://localhost:8000
```

---

## 1. List Posts with Pagination

### Default Pagination (10 items per page)
```bash
curl http://localhost:8000/posts/
```

**Response:**
```json
{
  "count": 50,
  "next": "http://localhost:8000/posts/?page=2",
  "previous": null,
  "results": [...]
}
```

### Custom Page Size
```bash
curl "http://localhost:8000/posts/?page_size=5"
```

### Navigate Pages
```bash
curl "http://localhost:8000/posts/?page=3&page_size=10"
```

---

## 2. Filtering Posts

### Filter by Title (Exact Match)
```bash
curl "http://localhost:8000/posts/?title=My%20First%20Post"
```

### Filter with Contains
```bash
curl "http://localhost:8000/posts/?title__contains=tutorial"
```

### Filter by Author ID
```bash
curl "http://localhost:8000/posts/?author_id=1"
```

### Filter Published Posts Only
```bash
curl "http://localhost:8000/posts/?published=true"
```

### Multiple Filters (AND logic)
```bash
curl "http://localhost:8000/posts/?author_id=1&published=true&title__contains=guide"
```

---

## 3. Search Posts

### Search in Title and Content
```bash
curl "http://localhost:8000/posts/published/?search=python"
```

### Search with Pagination
```bash
curl "http://localhost:8000/posts/published/?search=tutorial&page=1&page_size=5"
```

---

## 4. Ordering Posts

### Order by Creation Date (Ascending)
```bash
curl "http://localhost:8000/posts/?ordering=created_at"
```

### Order by Creation Date (Descending - newest first)
```bash
curl "http://localhost:8000/posts/?ordering=-created_at"
```

### Order by Title
```bash
curl "http://localhost:8000/posts/?ordering=title"
```

### Multi-field Ordering
```bash
curl "http://localhost:8000/posts/?ordering=-created_at,title"
```

---

## 5. Combined Features

### Pagination + Filtering + Ordering
```bash
curl "http://localhost:8000/posts/?page=1&page_size=10&author_id=1&ordering=-created_at"
```

### Search + Ordering
```bash
curl "http://localhost:8000/posts/published/?search=tutorial&ordering=-created_at"
```

### Full Stack: Filter + Search + Order + Paginate
```bash
curl "http://localhost:8000/posts/?title__contains=guide&search=python&ordering=-created_at&page=1&page_size=5"
```

---

## 6. Create Post (Authentication Required)

### Without Auth (Should Fail - 401/403)
```bash
curl -X POST http://localhost:8000/posts/ \
  -H "Content-Type: application/json" \
  -d '{
    "title": "New Post",
    "content": "This is my post content"
  }'
```

### With Valid Token (Should Succeed - 201)
```bash
curl -X POST http://localhost:8000/posts/ \
  -H "Authorization: Token demo-token-456" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Authenticated Post",
    "content": "Created with authentication",
    "author_id": 1,
    "published": true
  }'
```

---

## 7. Update Post (Authentication Required)

### Update Full Post (PUT)
```bash
curl -X PUT http://localhost:8000/posts/1/ \
  -H "Authorization: Token demo-token-456" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Title",
    "content": "Updated content",
    "author_id": 1,
    "published": true
  }'
```

### Partial Update (PATCH)
```bash
curl -X PATCH http://localhost:8000/posts/1/ \
  -H "Authorization: Token demo-token-456" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Just updating the title"
  }'
```

---

## 8. Delete Post (Authentication Required)

```bash
curl -X DELETE http://localhost:8000/posts/1/ \
  -H "Authorization: Token demo-token-456"
```

**Expected:** `204 No Content`

---

## 9. Rate Limiting (Throttling)

### Test Anonymous Rate Limit (10 requests/minute)

Run 11 requests quickly:
```bash
for i in {1..11}; do
  echo "Request $i:"
  curl -i http://localhost:8000/posts/throttled/
  sleep 0.5
done
```

**Requests 1-10:** Success (200)
**Request 11:** Throttled (429)

**Expected Response for 11th:**
```
HTTP/1.1 429 Too Many Requests
Retry-After: 60s

{"error": "Rate limit exceeded. Try again later."}
```

### Test Authenticated Rate Limit (100 requests/minute)

Authenticated users get higher limit:
```bash
for i in {1..15}; do
  echo "Authenticated Request $i:"
  curl http://localhost:8000/posts/throttled/ \
    -H "Authorization: Token demo-token-456"
  sleep 0.2
done
```

Should allow more requests before throttling.

---

## 10. Complete CRUD Workflow

```bash
# 1. List all posts
curl http://localhost:8000/posts/

# 2. Create a new post (with auth)
curl -X POST http://localhost:8000/posts/ \
  -H "Authorization: Token demo-token-456" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "My Complete Guide",
    "content": "Full tutorial content here",
    "author_id": 1,
    "published": true
  }'

# 3. Get specific post (with auth)
curl http://localhost:8000/posts/1/ \
  -H "Authorization: Token demo-token-456"

# 4. Update post (with auth)
curl -X PUT http://localhost:8000/posts/1/ \
  -H "Authorization: Token demo-token-456" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Guide",
    "content": "Updated content",
    "author_id": 1,
    "published": true
  }'

# 5. Delete post (with auth)
curl -X DELETE http://localhost:8000/posts/1/ \
  -H "Authorization: Token demo-token-456"
```

---

## Feature Checklist

### Pagination ✅
- [x] Page number pagination
- [x] Custom page size
- [x] Next/previous links
- [x] Total count

### Filtering ✅
- [x] Exact match filtering
- [x] Lookup expressions (contains, gt, lt, etc.)
- [x] Multiple field filtering
- [x] AND logic for combined filters

### Search ✅
- [x] Multi-field search
- [x] Case-insensitive search
- [x] Search + pagination
- [x] Search + ordering

### Ordering ✅
- [x] Single field ordering
- [x] Descending order (-)
- [x] Multi-field ordering
- [x] Field whitelist validation

### Authentication ✅
- [x] Token authentication
- [x] Auth middleware
- [x] Protected endpoints
- [x] Public endpoints

### Permissions ✅
- [x] AllowAny (public access)
- [x] IsAuthenticated (require login)
- [x] Permission checking

### Throttling ✅
- [x] Anonymous rate limiting
- [x] User rate limiting
- [x] Retry-After header
- [x] 429 status code

---

## Pro Tips

### 1. View Full Response Headers
```bash
curl -i http://localhost:8000/posts/
```

### 2. Pretty Print JSON
```bash
curl http://localhost:8000/posts/ | jq .
```

### 3. Save Response to File
```bash
curl http://localhost:8000/posts/ > posts.json
```

### 4. Test Multiple Requests in Parallel
```bash  
for i in {1..5}; do
  curl http://localhost:8000/posts/throttled/ &
done
wait
```

### 5. Benchmark Performance
```bash
time curl http://localhost:8000/posts/
```

---

## Troubleshooting

**401 Unauthorized:**
- Check token in Authorization header
- Ensure format: `Authorization: Token <token>`

**404 Not Found:**
- Verify endpoint URL
- Check if resource exists

**429 Too Many Requests:**
- Wait for rate limit window to reset
- Use authenticated requests for higher limits

**500 Internal Server Error:**
- Check server logs
- Verify database connection
- Ensure migrations are run
