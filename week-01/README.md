# Week 1 Execution Plan

This document defines my **weekly execution system**

Focus areas:

1. **Problem Solving (LeetCode)**
2. **Backend Engineering in Go**
3. **System Design Learning**
4. **Consistency**

Schedule constraint:

```
Deep Work: 9 PM – 11 PM
Light Work: During office hours
```

---

# Weekly Goals

By the end of the week:

* Solve **21 LeetCode problems**
* Build **a production-style Go backend service**
* Study **7 system design videos**
* Read **2 deep engineering blogs**
* Maintain **daily consistency**

---

# System Design Reading (Office Hours)

This week’s deep dives:

### 1. Discord Message Indexing

[https://discord.com/blog/how-discord-indexes-billions-of-messages](https://discord.com/blog/how-discord-indexes-billions-of-messages)

Key takeaway:

Discord built a search system capable of handling billions of messages using **Elasticsearch**, where messages are indexed as documents and distributed across shards to allow scalable search queries. ([Medium][1])

Important concepts to observe:

* Distributed search
* Index sharding
* Message indexing pipeline
* Search latency optimization

---

### 2. Instagram Video Upload Latency Improvements

[https://instagram-engineering.com/video-upload-latency-improvements-at-instagram-bcf4b4c5520a](https://instagram-engineering.com/video-upload-latency-improvements-at-instagram-bcf4b4c5520a)

Key takeaway:

Instagram reduced upload latency by performing **only minimal processing before making a video publishable**, allowing uploads to complete faster while background processing continues afterward. ([Instagram Engineering][2])

Important concepts:

* async processing
* upload pipelines
* background jobs
* latency optimization

---

### Daily System Design Task

Each day during office hours:

```
Watch 1 Arpit Bhayani system design video
Take short notes
```

Suggested note format:

```
Problem:
Architecture:
Key tradeoffs:
Scaling technique used:
```

---

# Go Backend Project (Main Weekly Project)

### Project: Distributed URL Shortener API

Goal:

Build a **production-style backend service in Go** similar to a simplified version of Bitly.

Example:

```
POST /shorten
GET /{short_id}
```

Why this project?

This teaches **real backend engineering fundamentals**:

* API design
* database interaction
* caching
* concurrency
* production project structure

Skills learned:

* Go HTTP servers
* routing
* database usage
* caching
* backend architecture

---

## Project Architecture

```
url-shortener/

cmd/
    server/

internal/
    handlers/
    services/
    repository/
    models/

pkg/
    utils/

configs/
```

Tech stack:

* Go
* PostgreSQL / SQLite
* Redis (optional later)
* REST API

---

# Daily Deep Work Plan (9–11 PM)

Structure every day:

```
9:00 – 10:00   LeetCode
10:00 – 11:00  Go Backend Project
```

---

# LeetCode Problem Plan (Striver Style)

Starting with **Array fundamentals**.

Difficulty progression:
Easy → Medium → Interview classics.

---

# Monday

### LeetCode

1. Two Sum
2. Largest Element in Array
3. Second Largest Element in Array

Focus:

```
Hashing
Linear scans
Array basics
```

### Go

Project Setup

Tasks:

```
Initialize Go module
Create folder structure
Create basic HTTP server
Add health check endpoint
```

Endpoint:

```
GET /health
```

---

# Tuesday

### LeetCode

4. Check if Array is Sorted
5. Remove Duplicates from Sorted Array
6. Left Rotate Array by One

Concepts:

```
Two pointers
In-place array manipulation
```

### Go

Short URL Creation

Tasks:

```
POST /shorten
Generate random short id
Store mapping in memory map
```

Example:

```
POST /shorten
{
 "url":"https://google.com"
}
```

Response:

```
{
 "short_url":"abc123"
}
```

---

# Wednesday

### LeetCode

7. Move Zeroes
8. Union of Two Sorted Arrays
9. Missing Number

Concepts:

```
Two pointer technique
Hash set
```

### Go

Redirect Service

Tasks:

```
GET /{shortID}
Redirect to original URL
```

Add:

```
HTTP 302 redirect
```

---

# Thursday

### LeetCode

10. Maximum Consecutive Ones
11. Single Number
12. Longest Subarray With Sum K

Concepts:

```
Prefix sum
Hashmap
```

### Go

Database Integration

Replace in-memory map with database.

Tasks:

```
Integrate SQLite
Create URL mapping table
Store shortened links
```

---

# Friday

### LeetCode

13. Kadane's Algorithm
14. Rearrange Array by Sign
15. Best Time to Buy and Sell Stock

Concepts:

```
Dynamic programming intuition
Greedy patterns
```

### Go

Service Layer Refactor

Tasks:

```
Separate layers
Handlers
Services
Repository
```

Introduce:

```
clean architecture
```

---

# Saturday

### LeetCode

16. Next Permutation
17. Majority Element
18. Set Matrix Zeroes

Concepts:

```
Interview classics
Edge case handling
```

### Go

Add Analytics

Track usage.

New endpoint:

```
GET /stats/{shortID}
```

Return:

```
{
 clicks: 15
}
```

Add click counting middleware.

---

# Sunday

### LeetCode

19. Pascal's Triangle
20. Rotate Matrix
21. Merge Intervals

Concepts:

```
Matrix manipulation
Interval problems
```

---

# Sunday Reflection

Fill this out:

```
Week: 1

Problems solved: __ / 21

Hardest problem:
Most interesting Go concept learned:
System design insight this week:
```

---

# Weekly Success Criteria

Week successful if:

```
✔ 21 problems solved
✔ Go URL shortener working
✔ 7 system design videos watched
✔ 2 engineering blogs read
```

---

# Important Rule

Never break the chain.

Even on bad days:

```
Solve at least 1 problem.
```

Consistency > intensity.